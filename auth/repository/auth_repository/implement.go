package auth_repository

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/nanwp/jajan-yuk/auth/core/entity"
	repositoryintf "github.com/nanwp/jajan-yuk/auth/core/repository"
	"github.com/nanwp/jajan-yuk/auth/pkg/helper"
	"gorm.io/gorm"
	"time"
)

type repository struct {
	db    *gorm.DB
	redis *redis.Pool
}

func NewAuthRepository(db *gorm.DB, redis *redis.Pool) repositoryintf.AuthRepository {
	return &repository{db, redis}
}

func (r *repository) Login(params entity.LoginRequest) (response entity.LoginResponse, err error) {
	db := r.db.Model(&User{})
	result := User{}

	if err := db.Where("username = ?", params.Username).First(&result).Error; err != nil {
		return response, err
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

func (r *repository) StoredAccessTokenInRedis(token string, userID string) (err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("access_token:%s", token)
	expired := time.Now().Add(entity.LOGIN_EXPIRATION_DURATION).Unix()

	_, err = conn.Do("SET", redisKey, token, "EX", expired)
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

func (r *repository) GetAccessTokenFromRedis(token string) (resp string, err error) {
	conn := r.redis.Get()
	defer conn.Close()

	redisKey := fmt.Sprintf("access_token:%s", token)

	resp, err = redis.String(conn.Do("GET", redisKey))
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
