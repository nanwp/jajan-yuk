package category_repository

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

func New(cfg config.Config, db *gorm.DB) repository_intf.CategoryRepo {
	return &repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *repository) CreateCategory(category entity.Category, userID string) (record entity.Category, err error) {
	db := r.db.Model(&Category{})
	val := Category{}
	val.FromEntity(category)
	val.CreatedBy = userID
	val.UpdatedBy = userID

	if err := db.Create(&val).Error; err != nil {
		return entity.Category{}, err
	}

	return val.ToEntity(), nil
}

func (r *repository) GetCategoryByID(id int64) (record entity.Category, err error) {
	db := r.db.Model(&Category{}).Where("id = ?", id)

	rec := Category{}

	if err := db.First(&rec).Error; err != nil {
		return entity.Category{}, err
	}

	return rec.ToEntity(), nil
}

func (r *repository) GetCategoryByIDs(ids []int64) (records []entity.Category, err error) {
	db := r.db.Model(&Category{})

	categories := []Category{}

	if err := db.Find(&categories, "id IN ?", ids).Error; err != nil {
		return records, err
	}

	for _, category := range categories {
		records = append(records, category.ToEntity())
	}

	return records, nil
}

func (r *repository) GetCategoryByUserCreated(userID string) (records []entity.Category, err error) {
	db := r.db.Model(&Category{})

	categories := []Category{}

	if err := db.Find(&categories, "created_by = ?", userID).Error; err != nil {
		return records, err
	}

	for _, category := range categories {
		records = append(records, category.ToEntity())
	}

	return records, nil
}
