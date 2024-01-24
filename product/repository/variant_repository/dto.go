package variant_repository

import (
	"time"

	"github.com/nanwp/jajan-yuk/product/core/entity"
	"gorm.io/gorm"
)

type Variant struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255)" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

type VariantType struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Price     int64          `gorm:"column:price;type:bigint" json:"price"`
	VariantID int64          `gorm:"column:variant_id;type:bigint" json:"variant_id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

func (v *Variant) ToEntity() entity.Variant {
	variant := entity.Variant{
		ID:   v.ID,
		Name: v.Name,
	}
	return variant
}

func (v *Variant) FromEntity(entity entity.Variant) {
	v.ID = entity.ID
	v.Name = entity.Name
}

func (vt *VariantType) ToEntity() entity.VariantType {
	return entity.VariantType{
		ID:        vt.ID,
		Name:      vt.Name,
		Price:     vt.Price,
		VariantID: vt.VariantID,
	}
}

func (vt *VariantType) FromEntity(entity entity.VariantType) {
	vt.ID = entity.ID
	vt.Name = entity.Name
	vt.Price = entity.Price
	vt.VariantID = entity.VariantID
}
