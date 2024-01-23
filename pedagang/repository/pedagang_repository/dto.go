package pedagang_repository

import (
	"time"

	"github.com/nanwp/jajan-yuk/pedagang/core/entity"
	"gorm.io/gorm"
)

type Pedagang struct {
	ID           string         `gorm:"column:id;primary_key;type:uuid;default:null" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy    string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy    string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy    string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
	UserID       string         `gorm:"column:user_id;type:uuid;default:null" json:"user_id"`
	NameMerchant string         `gorm:"column:name_merchant;type:varchar(255);default:null" json:"name_merchant"`
	Phone        string         `gorm:"column:phone;type:varchar(255);default:null" json:"phone"`
	Latitude     float64        `gorm:"column:latitude;type:double precision;default:null" json:"latitude"`
	Longitude    float64        `gorm:"column:longitude;type:double precision;default:null" json:"longitude"`
	Image        string         `gorm:"column:image;type:varchar(255);default:null" json:"image"`
	IsActive     bool           `gorm:"column:is_active;type:boolean;default:true" json:"is_active"`
	Distance     float64        `gorm:"-" json:"distance"`
}

func (p *Pedagang) ToEntity() entity.Pedagang {
	return entity.Pedagang{
		ID:           p.ID,
		UserID:       p.UserID,
		NameMerchant: p.NameMerchant,
		Phone:        p.Phone,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		Image:        p.Image,
		IsActive:     p.IsActive,
	}
}

func (p *Pedagang) FromEntity(e entity.Pedagang) {
	p.ID = e.ID
	p.UserID = e.UserID
	p.NameMerchant = e.NameMerchant
	p.Phone = e.Phone
	p.Latitude = e.Latitude
	p.Longitude = e.Longitude
	p.Image = e.Image
	p.IsActive = e.IsActive
}
