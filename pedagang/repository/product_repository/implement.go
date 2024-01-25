package product_repository

import (
	"github.com/nanwp/jajan-yuk/pedagang/client/product_client"
	"github.com/nanwp/jajan-yuk/pedagang/config"
	"github.com/nanwp/jajan-yuk/pedagang/core/entity"
	repository_intf "github.com/nanwp/jajan-yuk/pedagang/core/repository"
)

type repository struct {
	cfg    config.Config
	client product_client.ProductClient
}

func New(cfg config.Config, client product_client.ProductClient) repository_intf.ProductRepository {
	return &repository{
		cfg:    cfg,
		client: client,
	}
}

func (r *repository) GetProductByUserID(userID string) ([]entity.Product, error) {
	record := make([]entity.Product, 0)

	response, err := r.client.GetProductByUserID(userID)
	if err != nil {
		return []entity.Product{}, err
	}

	for _, v := range response.Data {
		record = append(record, v.ToEntity())
	}

	return record, nil
}
