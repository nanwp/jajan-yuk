package auth_repository

import (
	"database/sql"
	"time"

	"github.com/nanwp/jajan-yuk/auth/core/entity"
	"gorm.io/gorm"
)

type User struct {
	ID          string         `gorm:"column:id;primary_key;type:uuid" json:"id"`
	Name        string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Username    string         `gorm:"column:username;type:varchar(70)" json:"username"`
	Password    string         `gorm:"column:password;type:varchar(200)" json:"password"`
	Email       string         `gorm:"column:email;type:varchar(150)" json:"email"`
	ActivatedAt sql.NullTime   `gorm:"column:activated_at" json:"activated_at"`
	Image       string         `gorm:"column:image;type:varchar(255)" json:"image"`
	RoleID      string         `gorm:"column:role_id;type:uuid" json:"role_id"`
	Address     string         `gorm:"column:address;type:varchar(255)" json:"address"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy   string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy   string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy   string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

type Role struct {
	ID        string         `gorm:"column:id;primary_key;type:uuid" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255)" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

type SecretKey struct {
	ID        string         `gorm:"column:id;primary_key;type:uuid" json:"id"`
	Key       string         `gorm:"column:key;type:varchar(255)" json:"key"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

func (u *User) ToEntity() entity.User {
	return entity.User{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
		Image:    u.Image,
		Address:  u.Address,
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
