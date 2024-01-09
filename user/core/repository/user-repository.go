package repository

import "github.com/nanwp/jajan-yuk/user/core/entity"

type UserRepository interface {
	AddUser(user entity.User) (response entity.User, err error)
	CheckUsername(username string) (bool, error)
	CheckEmail(email string) (bool, error)
	StoredTokenToRedis(key, value string, ttl int) error
	GetTokenFromRedis(key string) (value string, err error)
	GetRoleByID(id string) (role entity.Role, err error)
}
