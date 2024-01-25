package product_repository

import (
	"github.com/nanwp/jajan-yuk/product/config"
	"github.com/nanwp/jajan-yuk/product/core/entity"
	repository_intf "github.com/nanwp/jajan-yuk/product/core/repository"
	"gorm.io/gorm"
)

type repository struct {
	cfg config.Config
	db  *gorm.DB
}

func New(cfg config.Config, db *gorm.DB) repository_intf.ProductRepo {
	return &repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *repository) CreateProduct(product entity.Product, userID string) (record entity.Product, err error) {
	db := r.db.Model(&Product{})
	val := Product{}
	val.FromEntity(product)
	val.CreatedBy = userID
	val.UpdatedBy = userID

	if err := db.Create(&val).Error; err != nil {
		return entity.Product{}, err
	}

	return val.ToEntity(), nil
}

func (r *repository) GetProductByID(id int64) (record entity.Product, err error) {
	db := r.db.Model(&Product{}).Where("id = ?", id)

	rec := Product{}

	if err := db.First(&rec).Error; err != nil {
		return entity.Product{}, err
	}

	return rec.ToEntity(), nil
}

func (r *repository) GetProductByIDs(ids []int64) (records []entity.Product, err error) {
	db := r.db.Model(&Product{})

	products := []Product{}

	if err := db.Find(&products, "id IN ?", ids).Error; err != nil {
		return records, err
	}

	for _, product := range products {
		records = append(records, product.ToEntity())
	}

	return records, nil
}

func (r *repository) GetProductByPedagangID(pedagangID string) (records []entity.Product, err error) {
	db := r.db.Model(&Product{})

	products := []Product{}

	if err := db.Find(&products, "pedagang_id = ?", pedagangID).Error; err != nil {
		return records, err
	}

	for _, product := range products {
		records = append(records, product.ToEntity())
	}

	return records, nil
}

func (r *repository) GetProductByUserCreated(userID string) (records []entity.Product, err error) {
	db := r.db.Model(&Product{})

	products := []Product{}

	if err := db.Find(&products, "created_by = ?", userID).Error; err != nil {
		return records, err
	}

	for _, product := range products {
		records = append(records, product.ToEntity())
	}

	return records, nil
}
