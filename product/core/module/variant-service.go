package module

import (
	"fmt"

	"github.com/nanwp/jajan-yuk/product/core/entity"
	"github.com/nanwp/jajan-yuk/product/core/repository"
)

type VariantService interface {
	GetVariantByIDs(ids []int64) (records []entity.Variant, err error)
	GetVariantByUserCreated(userID string) (records []entity.Variant, err error)
	CreateVariant(variant entity.Variant, userID string) (record entity.Variant, err error)
}

type variantService struct {
	variantRepo repository.VariantRepo
}

func NewVariantService(variantRepo repository.VariantRepo) VariantService {
	return &variantService{
		variantRepo: variantRepo,
	}
}

func (s *variantService) CreateVariant(variant entity.Variant, userID string) (record entity.Variant, err error) {
	record, err = s.variantRepo.CreateVariant(variant, userID)
	if err != nil {
		errMsg := fmt.Errorf("[VariantService.CreateVariant] error when create variant, err: %w", err)
		return entity.Variant{}, errMsg
	}

	for _, variantType := range variant.VariantTypes {
		variantType.VariantID = record.ID
		variantType, err = s.variantRepo.CreateVariantType(variantType, userID)
		if err != nil {
			errMsg := fmt.Errorf("[VariantService.CreateVariant] error when create variant type, err: %w", err)
			return entity.Variant{}, errMsg
		}
		record.VariantTypes = append(record.VariantTypes, variantType)
	}

	return record, nil
}

func (s *variantService) GetVariantByUserCreated(userID string) (records []entity.Variant, err error) {
	records, err = s.variantRepo.GetVariantByUserCreated(userID)
	if err != nil {
		errMsg := fmt.Errorf("[VariantService.GetVariantByUserCreated] error when get variant by user created, err: %w", err)
		return []entity.Variant{}, errMsg
	}

	variantIDs := []int64{}
	for _, record := range records {
		variantIDs = append(variantIDs, record.ID)
	}

	variantTypes, err := s.variantRepo.GetVariantTypeByVariantIDs(variantIDs)
	if err != nil {
		errMsg := fmt.Errorf("[VariantService.GetVariantByUserCreated] error when get variant type by variant ids, err: %w", err)
		return []entity.Variant{}, errMsg
	}

	mapVariantTypes := map[int64][]entity.VariantType{}
	for _, variantType := range variantTypes {
		mapVariantTypes[variantType.VariantID] = append(mapVariantTypes[variantType.VariantID], variantType)
	}

	for i, record := range records {
		records[i].VariantTypes = mapVariantTypes[record.ID]
	}

	return records, nil
}

func (s *variantService) GetVariantByIDs(ids []int64) (records []entity.Variant, err error) {
	records, err = s.variantRepo.GetVariantByIDs(ids)
	if err != nil {
		errMsg := fmt.Errorf("[VariantService.GetVariantByIDs] error when get variant by ids, err: %w", err)
		return []entity.Variant{}, errMsg
	}

	variantIDs := []int64{}
	for _, record := range records {
		variantIDs = append(variantIDs, record.ID)
	}

	variantTypes, err := s.variantRepo.GetVariantTypeByVariantIDs(variantIDs)
	if err != nil {
		errMsg := fmt.Errorf("[VariantService.GetVariantByIDs] error when get variant type by variant ids, err: %w", err)
		return []entity.Variant{}, errMsg
	}

	mapVariantTypes := map[int64][]entity.VariantType{}
	for _, variantType := range variantTypes {
		mapVariantTypes[variantType.VariantID] = append(mapVariantTypes[variantType.VariantID], variantType)
	}

	for i, record := range records {
		records[i].VariantTypes = mapVariantTypes[record.ID]
	}

	return records, nil
}
