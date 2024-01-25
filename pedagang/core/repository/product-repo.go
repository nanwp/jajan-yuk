package repository

import "github.com/nanwp/jajan-yuk/pedagang/core/entity"

type ProductRepository interface {
	GetProductByUserID(userID string) ([]entity.Product, error)
}
