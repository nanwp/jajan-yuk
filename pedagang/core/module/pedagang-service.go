package module

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/nanwp/jajan-yuk/pedagang/config"
	"github.com/nanwp/jajan-yuk/pedagang/core/entity"
	"github.com/nanwp/jajan-yuk/pedagang/core/repository"
	"github.com/nanwp/jajan-yuk/pedagang/pkg/helper"
)

type PedagangService interface {
	GetPedagangByID(id string) (entity.Pedagang, error)
	GetAllPedagangNearby(params entity.GetAllPedagangNearbyRequest) ([]entity.Pedagang, error)
	CreatePedagang(pedagang entity.Pedagang) (entity.Pedagang, error)
	UpdateLocation(params entity.UpdateLocationRequest) error
	SaveImage(image multipart.File, handler *multipart.FileHeader) (string, error)
}

type pedagangService struct {
	cfg          config.Config
	pedagangRepo repository.PedagangRepository
}

func NewPedagangService(cfg config.Config, pedagangRepo repository.PedagangRepository) PedagangService {
	return &pedagangService{
		cfg:          cfg,
		pedagangRepo: pedagangRepo,
	}
}

func (s *pedagangService) SaveImage(image multipart.File, handler *multipart.FileHeader) (string, error) {
	if image == nil {
		errMsg := errors.New("image is required")
		return "", errMsg
	}

	if handler == nil {
		errMsg := errors.New("handler is required")
		return "", errMsg
	}

	if handler.Size > 1000000 {
		errMsg := errors.New("image size is too large")
		return "", errMsg
	}

	if handler.Header.Get("Content-Type") != "image/jpeg" && handler.Header.Get("Content-Type") != "image/png" {
		log.Println(handler.Header.Get("Content-Type"))
		// errMsg := errors.New("image type must be jpeg or png")
		// return "", errMsg
	}

	imagePath, err := helper.UploadImage(image, handler)
	if err != nil {
		errMsg := fmt.Errorf("[PedagangService.SaveImage] error when upload image: %w", err)
		return "", errMsg
	}

	return imagePath, nil
}

func (s *pedagangService) GetPedagangByID(id string) (entity.Pedagang, error) {
	if id == "" {
		errMsg := errors.New("id is required")
		return entity.Pedagang{}, errMsg
	}

	response, err := s.pedagangRepo.GetPedagangByID(id)
	if err != nil {
		errMsg := fmt.Errorf("[PedagangService.GetPedagangByID] error when get pedagang by id: %w", err)
		return entity.Pedagang{}, errMsg
	}

	return response, nil
}

func (s *pedagangService) GetAllPedagangNearby(params entity.GetAllPedagangNearbyRequest) ([]entity.Pedagang, error) {
	if params.MaxDistance == 0 {
		params.MaxDistance = 1
	}

	response, err := s.pedagangRepo.GetAllPedagangNearby(params)
	if err != nil {
		errMsg := fmt.Errorf("[PedagangService.GetAllPedagangNearby] error when get all pedagang nearby: %w", err)
		return nil, errMsg
	}

	for i, v := range response {
		response[i].Distance = helper.Distance(params.Latitude, params.Longitude, v.Latitude, v.Longitude, "m")
		if response[i].Image != "" {
			response[i].Image = fmt.Sprintf("%s/%s", s.cfg.BaseURL, v.Image)
		}
	}

	return response, nil
}

func (s *pedagangService) CreatePedagang(pedagang entity.Pedagang) (entity.Pedagang, error) {
	if pedagang.UserID == "" {
		errMsg := errors.New("user id is required")
		return entity.Pedagang{}, errMsg
	}

	if pedagang.Image == "" {
		pedagang.Image = "/images/default.png"
	}

	response, err := s.pedagangRepo.CreatePedagang(pedagang)
	if err != nil {
		errMsg := fmt.Errorf("[PedagangService.CreatePedagang] error when create pedagang: %w", err)
		return entity.Pedagang{}, errMsg
	}

	return response, nil
}

func (s *pedagangService) UpdateLocation(params entity.UpdateLocationRequest) error {
	if params.UserID == "" {
		errMsg := errors.New("user id is required")
		return errMsg
	}

	err := s.pedagangRepo.UpdateLocation(params)
	if err != nil {
		errMsg := fmt.Errorf("[PedagangService.UpdateLocation] error when update location: %w", err)
		return errMsg
	}

	return nil
}
