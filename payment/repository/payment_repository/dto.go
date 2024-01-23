package payment_repository

import (
	"time"

	"github.com/nanwp/jajan-yuk/payment/core/entity"
	"gorm.io/gorm"
)

type Payment struct {
	ID        string         `gorm:"column:id;primary_key;type:uuid;default:null" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
	UserID    string         `gorm:"column:user_id;type:uuid;default:null" json:"user_id"`
	Status    string         `gorm:"column:status;type:varchar(255);default:null" json:"status"`
	Ammount   int64          `gorm:"column:ammount;type:integer;default:null" json:"ammount"`
	SnapURL   string         `gorm:"column:snap_url;type:varchar(255);default:null" json:"snap_url"`
}

func (p *Payment) ToEntoty() entity.Payment {
	return entity.Payment{
		ID:      p.ID,
		UserID:  p.UserID,
		Status:  p.Status,
		Ammount: p.Ammount,
		SnapURL: p.SnapURL,
	}
}

func (p *Payment) FromEntoty(payment entity.Payment) {
	p.ID = payment.ID
	p.UserID = payment.UserID
	p.Status = payment.Status
	p.Ammount = payment.Ammount
	p.SnapURL = payment.SnapURL
}
