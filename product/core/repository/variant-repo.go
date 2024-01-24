package repository

import "github.com/nanwp/jajan-yuk/product/core/entity"

type VariantRepo interface {
	CreateVariant(variant entity.Variant, userID string) (record entity.Variant, err error)
	GetVariantByID(id int64) (record entity.Variant, err error)
	GetVariantByIDs(ids []int64) (records []entity.Variant, err error)
	// GetVariantByProductID(productID string) (records []entity.Variant, err error)
	GetVariantByUserCreated(userID string) (records []entity.Variant, err error)

	//VariantType
	CreateVariantType(variantType entity.VariantType, userID string) (record entity.VariantType, err error)
	GetVariantTypeByVariantIDs(variantIDs []int64) (records []entity.VariantType, err error)
	// GetVariantTypeByID(id int64) (record entity.VariantType, err error)
	// GetVariantTypeByIDs(ids []int64) (records []entity.VariantType, err error)
	// GetVariantTypeByUserCreated(userID string) (records []entity.VariantType, err error)
}
