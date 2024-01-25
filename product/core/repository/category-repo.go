package repository

import "github.com/nanwp/jajan-yuk/product/core/entity"

type CategoryRepo interface {
	CreateCategory(category entity.Category, userID string) (record entity.Category, err error)
	GetCategoryByID(id int64) (record entity.Category, err error)
	GetCategoryByIDs(ids []int64) (records []entity.Category, err error)
	GetCategoryByUserCreated(userID string) (records []entity.Category, err error)
}
