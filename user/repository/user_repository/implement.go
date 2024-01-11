package user_repository

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/nanwp/jajan-yuk/user/core/entity"
	repository_intf "github.com/nanwp/jajan-yuk/user/core/repository"
	"gorm.io/gorm"
	"time"
)

type repository struct {
	db    *gorm.DB
	redis *redis.Pool
}

func (r repository) GetUserByID(id string) (response entity.User, err error) {
	db := r.db.Model(&User{}).Where("id = ?", id)

	result := User{}

	if err := db.First(&result).Error; err != nil {
		return response, err
	}

	return result.ToEntity(), nil
}

func (r repository) ActivateUser(id string) error {
	db := r.db.Model(&User{}).Where("id = ?", id)

	if err := db.Update("activated_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) GetRoleByID(id string) (role entity.Role, err error) {
	db := r.db.Model(&Role{}).Where("id = ?", id)

	result := Role{}

	if err := db.First(&result).Error; err != nil {
		return role, err
	}

	role = result.ToEntity()

	return role, err
}

func (r repository) GetTokenFromRedis(key string) (value string, err error) {
	conn := r.redis.Get()
	defer conn.Close()

	value, err = redis.String(conn.Do("GET", value))
	if err != nil {
		return value, entity.ErrorTokenExpired
	}

	return
}

func (r repository) StoredTokenToRedis(key, value string, ttl int) error {
	conn := r.redis.Get()
	defer conn.Close()

	expTime := time.Minute * time.Duration(ttl)

	_, err := conn.Do("SET", key, value, "EX", expTime.Seconds())
	if err != nil {
		return err
	}

	return nil
}

func (r repository) CheckUsername(username string) (bool, error) {
	db := r.db.Model(&User{}).Where("username = ?", username)

	result := User{}
	if err := db.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r repository) CheckEmail(email string) (bool, error) {
	db := r.db.Model(&User{}).Where("email = ?", email)

	result := User{}
	if err := db.First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r repository) AddUser(user entity.User) (response entity.User, err error) {
	db := r.db.Model(&User{})
	value := User{}
	value.FromEntity(user)
	value.CreatedAt = time.Now()
	value.UpdatedAt = time.Now()

	if err := db.Create(&value).Error; err != nil {
		return response, err
	}

	return value.ToEntity(), nil
}

func New(db *gorm.DB, redis *redis.Pool) repository_intf.UserRepository {
	return repository{
		db:    db,
		redis: redis,
	}
}
