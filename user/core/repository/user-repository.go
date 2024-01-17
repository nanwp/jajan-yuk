package repository

import "github.com/nanwp/jajan-yuk/user/core/entity"

type UserRepository interface {
	AddUser(user entity.User) (response entity.User, err error)
	GetUserByID(id string) (response entity.User, err error)
	GetUserByEmail(email string) (response entity.User, err error)
	CheckUsername(username string) (bool, error)
	CheckEmail(email string) (bool, error)
	UpdateUser(user entity.User) error
	StoredTokenToRedis(key, value string, ttl int) error
	GetTokenFromRedis(key string) (value string, err error)
	DeleteTokenFromRedis(key string) error
	GetRoleByID(id string) (role entity.Role, err error)
	AddPedagangInfo(field entity.Pedagang) (response entity.Pedagang, err error)
}
