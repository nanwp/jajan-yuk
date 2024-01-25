package repository

import "github.com/nanwp/jajan-yuk/auth/core/entity"

type AuthRepository interface {
	Login(params entity.LoginRequest) (response entity.LoginResponse, err error)
	GetUserByID(id string) (response entity.User, err error)
	StoredAccessTokenInRedis(token string, user entity.User) (err error)
	StoredRefreshTokenInRedis(token string, userID string) (err error)
	GetAccessTokenFromRedis(token string) (resp entity.User, err error)
	GetRefreshTokenFromRedis(token string) (resp string, err error)
	DeleteAccessTokenFromRedis(token, userID string) (err error)
	DeleteRefreshTokenFromRedis(token, userID string) (err error)
	ValidateSecretKey(key string) (secretKey entity.SecretKey, err error)
	GetRoleByID(id string) (role entity.Role, err error)

	StoredAccessTokenInRedisV2(token, userID string) (err error)
	GetAccessTokenFromRedisV2(token string) (resp string, err error)
	DeleteAccessTokenFromRedisV2(token, userID string) (err error)
}
