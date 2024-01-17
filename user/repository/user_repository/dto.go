package user_repository

import (
	"database/sql"
	"time"

	"github.com/nanwp/jajan-yuk/user/core/entity"
	"gorm.io/gorm"
)

type User struct {
	ID             string         `gorm:"column:id;primary_key;type:uuid;default:null" json:"id"`
	Name           string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Username       string         `gorm:"column:username;type:varchar(70)" json:"username"`
	Password       string         `gorm:"column:password;type:varchar(200)" json:"password"`
	Email          string         `gorm:"column:email;type:varchar(150)" json:"email"`
	ActivatedAt    sql.NullTime   `gorm:"column:activated_at" json:"activated_at"`
	Image          string         `gorm:"column:image;type:varchar(255)" json:"image"`
	RoleID         string         `gorm:"column:role_id;type:uuid" json:"role_id"`
	DateOfBirthday time.Time      `gorm:"column:date_of_birthday" json:"date_of_birthday"`
	Address        string         `gorm:"column:address" json:"address"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy      string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy      string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy      string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

func (u *User) ToEntity() entity.User {
	return entity.User{
		ID:             u.ID,
		Name:           u.Name,
		Username:       u.Username,
		Email:          u.Email,
		Password:       u.Password,
		Address:        u.Address,
		Image:          u.Image,
		DateOfBirthday: u.DateOfBirthday,
		Role: entity.Role{
			ID: u.RoleID,
		},
		ActivatedAt: u.ActivatedAt,
	}
}

func (u *User) FromEntity(user entity.User) {
	u.ID = user.ID
	u.Name = user.Name
	u.Username = user.Username
	u.Password = user.Password
	u.Email = user.Email
	u.Image = user.Image
	u.Address = user.Address
	u.RoleID = user.Role.ID
	u.DateOfBirthday = user.DateOfBirthday
	u.ActivatedAt = user.ActivatedAt
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

func (r *Role) ToEntity() entity.Role {
	return entity.Role{
		ID:   r.ID,
		Name: r.Name,
	}
}

type Pedagang struct {
	ID        string         `gorm:"column:id;primary_key;type:uuid;default:null" json:"id"`
	UserID    string         `gorm:"column:user_id" json:"user_id"`
	Name      string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Image     string         `gorm:"column:image" json:"image"`
	Telephone string         `gorm:"column:telephone" json:"telephone"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

func (p *Pedagang) ToEntity() entity.Pedagang {
	return entity.Pedagang{
		ID:        p.ID,
		UserID:    p.UserID,
		Name:      p.Name,
		Image:     p.Image,
		Telephone: p.Telephone,
	}
}

func (p *Pedagang) FromEntity(pedagang entity.Pedagang) {
	p.ID = pedagang.ID
	p.UserID = pedagang.UserID
	p.Name = pedagang.Name
	p.Image = pedagang.Image
	p.Telephone = pedagang.Telephone
}
