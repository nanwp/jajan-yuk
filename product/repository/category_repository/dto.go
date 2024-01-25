package category_repository

import (
	"time"

	"github.com/nanwp/jajan-yuk/product/core/entity"
	"gorm.io/gorm"
)

type Category struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255)" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

func (c *Category) ToEntity() entity.Category {
	category := entity.Category{
		ID:   c.ID,
		Name: c.Name,
	}
	return category
}

func (c *Category) FromEntity(entity entity.Category) {
	c.ID = entity.ID
	c.Name = entity.Name
}
