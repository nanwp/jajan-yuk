package pedagang_repository

import (
	"github.com/nanwp/jajan-yuk/user/client/pedagang"
	"github.com/nanwp/jajan-yuk/user/config"
	"github.com/nanwp/jajan-yuk/user/core/entity"
	repository_intf "github.com/nanwp/jajan-yuk/user/core/repository"
)

type repository struct {
	cfg            config.Config
	pedagangClient pedagang.PedagangClient
}

func New(cfg config.Config, pedagangClient pedagang.PedagangClient) repository_intf.PedagangRepository {
	return &repository{
		cfg:            cfg,
		pedagangClient: pedagangClient,
	}
}

func (r repository) CreatePedagang(pedagang entity.Pedagang) (response entity.Pedagang, err error) {
	pedagang, err = r.pedagangClient.CreatePedagang(pedagang)
	if err != nil {
		return response, err
	}

	return pedagang, nil
}
