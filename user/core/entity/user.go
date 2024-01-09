package entity

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrorUsernameUsed = errors.New("username already exist")
	ErrorEmailUsed    = errors.New("email already exist")
	ErrorTokenExpired = errors.New("token expired")
)

type User struct {
	ID             string    `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Username       string    `json:"username,omitempty"`
	Email          string    `json:"email,omitempty"`
	Password       string    `json:"password,omitempty"`
	Address        string    `json:"address,omitempty"`
	DateOfBirthday time.Time `json:"date_of_birthday,omitempty"`
	Image          string    `json:"image"`
	Role           Role      `json:"role"`
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
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
