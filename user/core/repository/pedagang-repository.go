package repository

import "github.com/nanwp/jajan-yuk/user/core/entity"

type PedagangRepository interface {
	CreatePedagang(pedagang entity.Pedagang) (response entity.Pedagang, err error)
}
