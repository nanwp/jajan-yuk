package module

import (
	"fmt"

	"github.com/nanwp/jajan-yuk/product/core/entity"
	"github.com/nanwp/jajan-yuk/product/core/repository"
)

type CategoryService interface {
	GetCategoryByIDs(ids []int64) (records []entity.Category, err error)
	GetCategoryByUserCreated(userID string) (records []entity.Category, err error)
	CreateCategory(category entity.Category, userID string) (record entity.Category, err error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepo
}

func NewCategoryService(categoryRepo repository.CategoryRepo) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) CreateCategory(category entity.Category, userID string) (record entity.Category, err error) {
	if err := category.Validate(); err != nil {
		errMsg := fmt.Errorf("[CategoryService.CreateCategory] error when validate category, err: %w", err)
		return entity.Category{}, errMsg
	}

	record, err = s.categoryRepo.CreateCategory(category, userID)
	if err != nil {
		errMsg := fmt.Errorf("[CategoryService.CreateCategory] error when create category, err: %w", err)
		return entity.Category{}, errMsg
	}

	return record, nil
}

func (s *categoryService) GetCategoryByUserCreated(userID string) (records []entity.Category, err error) {
	records, err = s.categoryRepo.GetCategoryByUserCreated(userID)
	if err != nil {
		errMsg := fmt.Errorf("[CategoryService.GetCategoryByUserCreated] error when get category by user created, err: %w", err)
		return []entity.Category{}, errMsg
	}

	return records, nil
}

func (s *categoryService) GetCategoryByIDs(ids []int64) (records []entity.Category, err error) {
	records, err = s.categoryRepo.GetCategoryByIDs(ids)
	if err != nil {
		errMsg := fmt.Errorf("[CategoryService.GetCategoryByIDs] error when get category by ids, err: %w", err)
		return []entity.Category{}, errMsg
	}

	return records, nil
}
