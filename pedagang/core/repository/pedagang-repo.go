package repository

import "github.com/nanwp/jajan-yuk/pedagang/core/entity"

type PedagangRepository interface {
	GetPedagangByID(id string) (entity.Pedagang, error)
	GetAllPedagangNearby(params entity.GetAllPedagangNearbyRequest) ([]entity.Pedagang, error)
	CreatePedagang(pedagang entity.Pedagang) (entity.Pedagang, error)
	UpdateLocation(params entity.UpdateLocationRequest) error
}
