package auth_repository

import (
	"database/sql"
	"github.com/nanwp/jajan-yuk/auth/core/entity"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          string         `gorm:"column:id;primary_key" json:"id"`
	Name        string         `gorm:"column:name" json:"name"`
	Username    string         `gorm:"column:username" json:"username"`
	Password    string         `gorm:"column:password" json:"password"`
	Email       string         `gorm:"column:email" json:"email"`
	ActivatedAt sql.NullTime   `gorm:"column:activated_at" json:"activated_at"`
	Image       string         `gorm:"column:image" json:"image"`
	RoleID      string         `gorm:"column:role_id" json:"role_id"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type Role struct {
	ID   string `gorm:"column:id;primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type SecretKey struct {
	ID  string `gorm:"column:id;primary_key" json:"id"`
	Key string `gorm:"column:key" json:"key"`
}

func (u *User) ToEntity() entity.User {
	return entity.User{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Image:    u.Image,
		Role: entity.Role{
			ID: u.RoleID,
		},
	}
}

func (r *Role) ToEntity() entity.Role {
	return entity.Role{
		ID:   r.ID,
		Name: r.Name,
	}
}

func (s *SecretKey) ToEntity() entity.SecretKey {
	return entity.SecretKey{
		ID:  s.ID,
		Key: s.Key,
	}
}
