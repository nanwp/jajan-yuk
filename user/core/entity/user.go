package entity

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type RoleID string

var (
	UserID     RoleID = "3e76048d-f9f2-4974-845f-60137f9e2f4b"
	PedagangID RoleID = "ea8e1e87-ae6e-44b1-9854-4dbb0c70a330"
)

var (
	ErrorUsernameUsed = errors.New("username already exist")
	ErrorEmailUsed    = errors.New("email already exist")
	ErrorTokenExpired = errors.New("token expired")
)

type ActivateAccount struct {
	Token string `json:"token"`
}

type ResetPassword struct {
	Token              string `json:"token"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

func (r ResetPassword) Validate() error {
	var errMsg []string
	if r.Token == "" {
		errMsg = append(errMsg, "token required")
	}

	if len(r.NewPassword) < 8 {
		errMsg = append(errMsg, "password min 8")
	}

	if r.ConfirmNewPassword != r.NewPassword {
		errMsg = append(errMsg, "password not match")
	}

	if len(errMsg) > 0 {
		return fmt.Errorf("Validate error: %v", strings.Join(errMsg, ", "))
	}
	return nil
}

type RequestResetPassword struct {
	Email string `json:"email"`
}

type User struct {
	ID             string       `json:"id,omitempty"`
	Name           string       `json:"name,omitempty"`
	Username       string       `json:"username,omitempty"`
	Email          string       `json:"email,omitempty"`
	Password       string       `json:"password,omitempty"`
	Address        string       `json:"address,omitempty"`
	DateOfBirthday time.Time    `json:"date_of_birthday,omitempty"`
	Image          string       `json:"image"`
	Role           Role         `json:"role"`
	ActivatedAt    sql.NullTime `json:"activated_at"`
}

func (u *User) Validate() error {
	var errMsg []string
	if u.Username == "" {
		errMsg = append(errMsg, "username is required")
	}
	if len(u.Username) < 5 {
		errMsg = append(errMsg, "username minimal 8")
	}
	if u.Name == "" {
		errMsg = append(errMsg, "name is required")
	}
	if u.Email == "" {
		errMsg = append(errMsg, "email is required")
	}
	if u.Password == "" {
		errMsg = append(errMsg, "password is required")
	}
	if len(u.Password) < 8 {
		errMsg = append(errMsg, "password minimal 8")
	}

	if len(errMsg) > 0 {
		return fmt.Errorf("Validate Error: %v", strings.Join(errMsg, ", "))
	}
	return nil
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}
