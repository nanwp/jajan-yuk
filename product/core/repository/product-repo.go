package repository

import "github.com/nanwp/jajan-yuk/product/core/entity"

type ProductRepo interface {
	CreateProduct(product entity.Product, userID string) (record entity.Product, err error)
	GetProductByID(id int64) (record entity.Product, err error)
	GetProductByIDs(ids []int64) (records []entity.Product, err error)
	GetProductByPedagangID(pedagangID string) (records []entity.Product, err error)
	GetProductByUserCreated(userID string) (records []entity.Product, err error)
}
