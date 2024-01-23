package pedagang_repository

import (
	"fmt"
	"log"

	"github.com/nanwp/jajan-yuk/pedagang/config"
	"github.com/nanwp/jajan-yuk/pedagang/core/entity"
	repository_intf "github.com/nanwp/jajan-yuk/pedagang/core/repository"
	"gorm.io/gorm"
)

type repository struct {
	cfg config.Config
	db  *gorm.DB
}

func NewPedagangRepository(cfg config.Config, db *gorm.DB) repository_intf.PedagangRepository {
	return &repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *repository) GetPedagangByID(id string) (entity.Pedagang, error) {
	db := r.db.Model(&Pedagang{}).Where("id = ?", id)
	var pedagang Pedagang

	if err := db.First(&pedagang).Error; err != nil {
		return entity.Pedagang{}, err
	}

	return pedagang.ToEntity(), nil
}

func (r *repository) GetAllPedagangNearby(params entity.GetAllPedagangNearbyRequest) ([]entity.Pedagang, error) {
	db := r.db.Model(&Pedagang{}).Where("is_active = ?", true)

	if params.Keyword != "" {
		keywordStr := fmt.Sprintf("%%%s%%", params.Keyword)
		db = db.Where("name_merchant ILIKE ?", keywordStr)
	}

	if params.Latitude != 0 && params.Longitude != 0 {
		latitude := params.Latitude
		longitude := params.Longitude
		maxDistance := params.MaxDistance

		// Haversine Formula
		db = db.Select("*").
			Select(
				"*, "+"6371 * acos(cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) + sin(radians(?)) * sin(radians(latitude))) AS distance", latitude, longitude, latitude).
			Where("6371 * acos(cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) + sin(radians(?)) * sin(radians(latitude))) < ?", latitude, longitude, latitude, maxDistance).
			Order("distance ASC")
	}

	db = db.Order("name_merchant ASC")

	var pedagangs []Pedagang
	if err := db.Find(&pedagangs).Error; err != nil {
		return nil, err
	}

	var result []entity.Pedagang
	for _, pedagang := range pedagangs {
		result = append(result, pedagang.ToEntity())
	}

	return result, nil
}

func (r *repository) CreatePedagang(pedagang entity.Pedagang) (entity.Pedagang, error) {
	db := r.db.Model(&Pedagang{})
	pedagangModel := Pedagang{}
	pedagangModel.FromEntity(pedagang)
	pedagangModel.CreatedBy = pedagang.UserID
	pedagangModel.UpdatedBy = pedagang.UserID

	if err := db.Create(&pedagangModel).Error; err != nil {
		return entity.Pedagang{}, err
	}

	return pedagangModel.ToEntity(), nil
}

func (r *repository) UpdatePedagang(params entity.Pedagang) error {
	db := r.db.Model(&Pedagang{}).Where("user_id = ?", params.UserID)
	pedagang := Pedagang{}
	pedagang.FromEntity(params)
	pedagang.UpdatedBy = params.UserID
	log.Println(pedagang)

	if err := db.Updates(&pedagang).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetPedagangByUserID(userID string) (entity.Pedagang, error) {
	db := r.db.Model(&Pedagang{}).Where("user_id = ?", userID)
	var pedagang Pedagang

	if err := db.First(&pedagang).Error; err != nil {
		return entity.Pedagang{}, err
	}

	return pedagang.ToEntity(), nil
}

func (r *repository) SwitchActiveStatus(id string, status bool) error {
	db := r.db.Model(&Pedagang{}).Where("id = ?", id)
	pedagang := Pedagang{}
	pedagang.IsActive = status

	updatedFields := map[string]interface{}{
		"is_active": status,
	}

	if err := db.Updates(updatedFields).Error; err != nil {
		return err
	}

	return nil
}
