package entity

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	ErrorUserNotFound         = errors.New("user not found")
	ErrorPasswordNotMatch     = errors.New("password not match")
	ErrorSecretKeyNotFound    = errors.New("secret key not found")
	ErrorUserNotVerified      = errors.New("user not verified")
	ErrorSigningMethodInvalid = errors.New("signing method invalid")
	ErrorTokenInvalid         = errors.New("token invalid")
	ErrorTokenExpired         = errors.New("token expired")
	ErrorUserNotActivated     = errors.New("user not activated")

	DefaultSucessCode     = 200
	DefaultSuccessMessage = "success"

	APPLICATION_NAME            = "JAJAN-YUK"
	LOGIN_EXPIRATION_DURATION   = time.Duration(24) * time.Hour
	REFRESH_EXPIRATION_DURATION = time.Duration(240) * time.Hour
	JWT_SIGNING_METHOD          = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY           = []byte("jajanankelilingkuy")
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *LoginRequest) Validate() error {
	var msg []string

	if u.Username == "" {
		msg = append(msg, "Username required")
	}

	if u.Password == "" {
		msg = append(msg, "Password required")
	}

	if len(msg) > 0 {
		return fmt.Errorf("Validate Error: %v", strings.Join(msg, ", "))
	}

	return nil
}

type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	ActivatedAt time.Time `json:"activated_at"`
	Image       string    `json:"image"`
	Address     string    `json:"address"`
	Role        Role      `json:"role"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type GetCurrentUserResponse struct {
	User User `json:"user"`
}

type MyClaims struct {
	jwt.RegisteredClaims
	ID       string `json:"id"`
	Username string `json:"username"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type SecretKey struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}
