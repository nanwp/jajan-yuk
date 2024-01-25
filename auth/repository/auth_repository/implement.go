package auth_repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/nanwp/jajan-yuk/auth/core/entity"
	repositoryintf "github.com/nanwp/jajan-yuk/auth/core/repository"
	"github.com/nanwp/jajan-yuk/auth/pkg/helper"
	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	redis *redis.Pool
}

func NewAuthRepository(db *gorm.DB, redis *redis.Pool) repositoryintf.AuthRepository {
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Role{})
	//db.AutoMigrate(&SecretKey{})
	return &repository{db, redis}
}

func (r *repository) Login(params entity.LoginRequest) (response entity.LoginResponse, err error) {
	db := r.db.Model(&User{})
	result := User{}

	if err := db.Where("username = ?", params.Username).First(&result).Error; err != nil {
		return response, entity.ErrorUserNotFound
	}

	if result.ActivatedAt.Valid == false {
		return response, entity.ErrorUserNotActivated
	}

	password, _ := helper.GeneratePasswordString(params.Password)

	if result.Password != password {
		return response, entity.ErrorPasswordNotMatch
	}

	response.User = result.ToEntity()

	return response, err
}

func (r *repository) GetUserByID(id string) (response entity.User, err error) {
	db := r.db.Model(&User{})
	result := User{}

	if err := db.Where("id = ?", id).First(&result).Error; err != nil {
		return response, err
	}

	response = result.ToEntity()

	return response, err
}

func (r *repository) StoredAccessTokenInRedis(token string, user entity.User) (err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("access_token:%s", token)
	expired := time.Now().Add(entity.LOGIN_EXPIRATION_DURATION).Unix()
	value, _ := json.Marshal(user)

	_, err = conn.Do("SET", redisKey, value, "EX", expired)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) StoredRefreshTokenInRedis(token string, userID string) (err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("refresh_token:%s", token)
	expired := time.Now().Add(entity.REFRESH_EXPIRATION_DURATION).Unix()

	_, err = conn.Do("SET", redisKey, token, "EX", expired)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAccessTokenFromRedis(token string) (resp entity.User, err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("access_token:%s", token)

	respStr, err := redis.String(conn.Do("GET", redisKey))
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (r *repository) GetRefreshTokenFromRedis(token string) (resp string, err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("refresh_token:%s", token)

	resp, err = redis.String(conn.Do("GET", redisKey))
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *repository) DeleteAccessTokenFromRedis(token, userID string) (err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("access_token:%s", token)

	_, err = conn.Do("DEL", redisKey)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteRefreshTokenFromRedis(token, userID string) (err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("refresh_token:%s", token)

	_, err = conn.Do("DEL", redisKey)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ValidateSecretKey(key string) (secretKey entity.SecretKey, err error) {
	db := r.db.Model(&SecretKey{})
	result := SecretKey{}

	if err := db.Where("key = ?", key).First(&result).Error; err != nil {
		return secretKey, err
	}

	secretKey = result.ToEntity()
	return secretKey, nil
}

func (r *repository) GetRoleByID(id string) (role entity.Role, err error) {
	db := r.db.Model(&Role{}).Where("id = ?", id)

	result := Role{}

	if err := db.First(&result).Error; err != nil {
		return role, err
	}

	return result.ToEntity(), nil
}

func (r *repository) StoredAccessTokenInRedisV2(token, userID string) (err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("access_token:%s", token)
	expired := time.Now().Add(entity.LOGIN_EXPIRATION_DURATION).Unix()

	_, err = conn.Do("SET", redisKey, userID, "EX", expired)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAccessTokenFromRedisV2(token string) (resp string, err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("access_token:%s", token)

	resp, err = redis.String(conn.Do("GET", redisKey))
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *repository) DeleteAccessTokenFromRedisV2(token, userID string) (err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("access_token:%s", token)

	_, err = conn.Do("DEL", redisKey)
	if err != nil {
		return err
	}
	return nil
}
