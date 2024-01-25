package product_repository

import (
	"strconv"
	"strings"
	"time"

	"github.com/nanwp/jajan-yuk/product/core/entity"
	"gorm.io/gorm"
)

type Product struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"column:name;type:varchar(255)" json:"name"`
	Description string         `gorm:"column:description;type:varchar(255)" json:"description"`
	Image       string         `gorm:"column:image;type:varchar(255)" json:"image"`
	Price       int64          `gorm:"column:price;type:bigint" json:"price"`
	CategoryID  int64          `gorm:"column:category_id;type:bigint" json:"category_id"`
	VarianIDs   string         `gorm:"column:varian_ids;type:varchar(255)" json:"varian_ids"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy   string         `gorm:"column:created_by;type:varchar(255)" json:"created_by"`
	UpdatedBy   string         `gorm:"column:updated_by;type:varchar(255)" json:"updated_by"`
	DeletedBy   string         `gorm:"column:deleted_by;type:varchar(255)" json:"deleted_by"`
}

func (p *Product) ToEntity() entity.Product {
	variantIDs := strings.Split(p.VarianIDs, ",")
	variant := []entity.Variant{}
	for _, id := range variantIDs {
		idInt, _ := strconv.ParseInt(id, 10, 64)
		variant = append(variant, entity.Variant{
			ID: idInt,
		})
	}

	return entity.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
		Price:       p.Price,
		Category: entity.Category{
			ID: p.CategoryID,
		},
		Variant: variant,
	}
}

func (p *Product) FromEntity(entity entity.Product) {
	variantIDs := []string{}
	for _, variant := range entity.Variant {
		IDstr := strconv.FormatInt(variant.ID, 10)
		variantIDs = append(variantIDs, IDstr)
	}

	p.ID = entity.ID
	p.Name = entity.Name
	p.Image = entity.Image
	p.Description = entity.Description
	p.Price = entity.Price
	p.CategoryID = entity.Category.ID
	p.VarianIDs = strings.Join(variantIDs, ",")
}
