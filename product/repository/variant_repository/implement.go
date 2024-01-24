package variant_repository

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

func New(cfg config.Config, db *gorm.DB) repository_intf.VariantRepo {
	return &repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *repository) CreateVariant(variant entity.Variant, userID string) (record entity.Variant, err error) {
	db := r.db.Model(&Variant{})
	val := Variant{}
	val.FromEntity(variant)
	val.CreatedBy = userID
	val.UpdatedBy = userID

	if err := db.Create(&val).Error; err != nil {
		return entity.Variant{}, err
	}

	return val.ToEntity(), nil
}

func (r *repository) GetVariantByID(id int64) (record entity.Variant, err error) {
	db := r.db.Model(&Variant{}).Where("id = ?", id)
	db = db.Preload("VariantType")

	rec := Variant{}

	if err := db.First(&rec).Error; err != nil {
		return entity.Variant{}, err
	}

	return rec.ToEntity(), nil
}

func (r *repository) GetVariantByIDs(ids []int64) (records []entity.Variant, err error) {
	db := r.db.Model(&Variant{})

	variants := []Variant{}

	if err := db.Find(&variants, "id IN ?", ids).Error; err != nil {
		return records, err
	}

	for _, variant := range variants {
		records = append(records, variant.ToEntity())
	}

	return records, nil
}

func (r *repository) GetVariantTypeByVariantIDs(ids []int64) (records []entity.VariantType, err error) {
	db := r.db.Model(&VariantType{})

	variantTypes := []VariantType{}

	if err := db.Find(&variantTypes, "variant_id IN ?", ids).Error; err != nil {
		return records, err
	}

	for _, variantType := range variantTypes {
		records = append(records, variantType.ToEntity())
	}

	return records, nil
}

func (r *repository) GetVariantByUserCreated(userID string) (records []entity.Variant, err error) {
	db := r.db.Model(&Variant{})

	variants := []Variant{}

	if err := db.Find(&variants, "created_by = ?", userID).Error; err != nil {
		return records, err
	}

	for _, variant := range variants {
		records = append(records, variant.ToEntity())
	}
	return records, nil
}

func (r *repository) CreateVariantType(variantType entity.VariantType, userID string) (record entity.VariantType, err error) {
	db := r.db.Model(&VariantType{})
	val := VariantType{}
	val.FromEntity(variantType)
	val.CreatedBy = userID
	val.UpdatedBy = userID

	if err := db.Create(&val).Error; err != nil {
		return entity.VariantType{}, err
	}

	return val.ToEntity(), nil
}
