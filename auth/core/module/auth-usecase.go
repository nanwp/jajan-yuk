package module

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nanwp/jajan-yuk/auth/core/entity"
	"github.com/nanwp/jajan-yuk/auth/core/repository"
)

type AuthUsecase interface {
	Login(params entity.LoginRequest) (response entity.LoginResponse, err error)
	RefreshToken(params entity.RefreshTokenRequest) (response entity.LoginResponse, err error)
	GetCurrentUser(token string) (response entity.GetCurrentUserResponse, err error)
	Logout(token string) (err error)
	ValidateSecretKey(secretKey string) (response entity.SecretKey, err error)
}

type authUsecase struct {
	authRepo repository.AuthRepository
}

func NewAuthUsecase(authRepo repository.AuthRepository) AuthUsecase {
	return &authUsecase{authRepo}
}

func (a *authUsecase) ValidateSecretKey(secretKey string) (response entity.SecretKey, err error) {
	response, err = a.authRepo.ValidateSecretKey(secretKey)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (a *authUsecase) Login(params entity.LoginRequest) (response entity.LoginResponse, err error) {
	if err := params.Validate(); err != nil {
		return entity.LoginResponse{}, err
	}
	response, err = a.authRepo.Login(params)
	if err != nil {
		return response, err
	}
	accessToken, err := a.generateAccessToken(response.User.ID, response.User.Username)
	if err != nil {
		return response, err
	}

	refreshToken, err := a.generateRefreshToken(response.User.ID, response.User.Username)
	if err != nil {
		return response, err
	}

	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	if err := a.authRepo.StoredAccessTokenInRedisV2(accessToken, response.User.ID); err != nil {
		return response, err
	}

	if err := a.authRepo.StoredRefreshTokenInRedis(refreshToken, response.User.ID); err != nil {
		return response, err
	}

	response.User.Password = ""
	return response, err
}

func (a *authUsecase) RefreshToken(params entity.RefreshTokenRequest) (response entity.LoginResponse, err error) {
	_, err = a.authRepo.GetRefreshTokenFromRedis(params.RefreshToken)
	if err != nil {
		return response, err
	}

	token, err := jwt.Parse(params.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entity.ErrorSigningMethodInvalid
		} else if method != entity.JWT_SIGNING_METHOD {
			return nil, entity.ErrorSigningMethodInvalid
		}
		return entity.JWT_SIGNATURE_KEY, nil
	})
	if err != nil {
		return response, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return response, entity.ErrorTokenInvalid
	}

	userID := claims["id"].(string)
	username := claims["username"].(string)

	user, err := a.authRepo.GetUserByID(userID)
	if err != nil {
		return entity.LoginResponse{}, err
	}

	role, err := a.authRepo.GetRoleByID(user.Role.ID)
	if err != nil {
		return entity.LoginResponse{}, err
	}

	user.Role = role

	accessToken, err := a.generateAccessToken(userID, username)
	if err != nil {
		return response, err
	}

	refreshToken, err := a.generateRefreshToken(userID, username)
	if err != nil {
		return response, err
	}

	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	response.User = user

	if err := a.authRepo.StoredAccessTokenInRedisV2(accessToken, userID); err != nil {
		return response, err
	}

	if err := a.authRepo.StoredRefreshTokenInRedis(refreshToken, userID); err != nil {
		return response, err
	}

	response.User.Password = ""
	return response, nil
}

func (a *authUsecase) GetCurrentUser(accessToken string) (response entity.GetCurrentUserResponse, err error) {
	userID, err := a.authRepo.GetAccessTokenFromRedisV2(accessToken)
	if err != nil {
		return response, err
	}

	user, err := a.authRepo.GetUserByID(userID)
	if err != nil {
		return response, err
	}

	role, err := a.authRepo.GetRoleByID(user.Role.ID)
	if err != nil {
		return response, err
	}

	user.Role = role

	response.User = user

	response.User.Password = ""
	return response, nil
}

func (a *authUsecase) Logout(token string) (err error) {
	response, err := a.GetCurrentUser(token)
	if err != nil {
		return err
	}

	if err := a.authRepo.DeleteAccessTokenFromRedisV2(token, response.User.ID); err != nil {
		return err
	}

	if err := a.authRepo.DeleteRefreshTokenFromRedis(token, response.User.ID); err != nil {
		return err
	}

	return nil
}

func (a *authUsecase) generateAccessToken(id, username string) (token string, err error) {
	claims := entity.MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: entity.APPLICATION_NAME,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(entity.LOGIN_EXPIRATION_DURATION),
			},
		},
		ID:       id,
		Username: username,
	}
	signetToken := jwt.NewWithClaims(
		entity.JWT_SIGNING_METHOD,
		claims,
	)
	tokenString, err := signetToken.SignedString(entity.JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *authUsecase) generateRefreshToken(id, username string) (token string, err error) {
	claims := entity.MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: entity.APPLICATION_NAME,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(entity.REFRESH_EXPIRATION_DURATION),
			},
		},
		ID:       id,
		Username: username,
	}
	signetToken := jwt.NewWithClaims(
		entity.JWT_SIGNING_METHOD,
		claims,
	)
	tokenString, err := signetToken.SignedString(entity.JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
