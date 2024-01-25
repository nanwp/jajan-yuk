package module

import (
	"fmt"
	"log"

	"github.com/nanwp/jajan-yuk/product/config"
	"github.com/nanwp/jajan-yuk/product/core/entity"
	"github.com/nanwp/jajan-yuk/product/core/repository"
)

type ProductService interface {
	GetProductByIDs(ids []int64) (records []entity.Product, err error)
	GetProductByUserCreated(userID string) (records []entity.Product, err error)
	GetProductByPedagangID(pedagangID string) (records []entity.Product, err error)
	CreateProduct(product entity.Product, userID string) (record entity.Product, err error)
}

type productService struct {
	cfg             config.Config
	productRepo     repository.ProductRepo
	variantService  VariantService
	categoryService CategoryService
}

func NewProductService(cfg config.Config, productRepo repository.ProductRepo, variantService VariantService, categoryService CategoryService) ProductService {
	return &productService{
		cfg:             cfg,
		productRepo:     productRepo,
		variantService:  variantService,
		categoryService: categoryService,
	}
}

func (s *productService) GetProductByPedagangID(pedagangID string) (records []entity.Product, err error) {
	records, err = s.productRepo.GetProductByPedagangID(pedagangID)
	if err != nil {
		return []entity.Product{}, err
	}

	categoryIDs := []int64{}
	for _, record := range records {
		categoryIDs = append(categoryIDs, record.Category.ID)
	}

	categories, err := s.categoryService.GetCategoryByIDs(categoryIDs)
	if err != nil {
		return []entity.Product{}, err
	}

	mapCategories := map[int64]entity.Category{}
	for _, category := range categories {
		mapCategories[category.ID] = category
	}

	variantIDs := []int64{}
	for _, record := range records {
		for _, variant := range record.Variant {
			variantIDs = append(variantIDs, variant.ID)
		}
	}

	variants, err := s.variantService.GetVariantByIDs(variantIDs)
	if err != nil {
		return []entity.Product{}, err
	}

	mapVariants := map[int64][]entity.Variant{}
	for _, variant := range variants {
		mapVariants[variant.ID] = append(mapVariants[variant.ID], variant)
	}

	for i, record := range records {
		records[i].Category = mapCategories[record.Category.ID]
		records[i].Variant = mapVariants[record.ID]
		records[i].Image = fmt.Sprintf("%s/%s", s.cfg.BaseURL, record.Image)
	}

	return records, nil
}

func (s *productService) GetProductByUserCreated(userID string) (records []entity.Product, err error) {
	records, err = s.productRepo.GetProductByUserCreated(userID)
	if err != nil {
		return []entity.Product{}, err
	}

	categoryIDs := []int64{}
	for _, record := range records {
		categoryIDs = append(categoryIDs, record.Category.ID)
	}

	categories, err := s.categoryService.GetCategoryByIDs(categoryIDs)
	if err != nil {
		return []entity.Product{}, err
	}

	mapCategories := map[int64]entity.Category{}
	for _, category := range categories {
		mapCategories[category.ID] = category
	}

	variantIDs := []int64{}
	for _, record := range records {
		for _, variant := range record.Variant {
			variantIDs = append(variantIDs, variant.ID)
		}
	}

	variants, err := s.variantService.GetVariantByIDs(variantIDs)
	if err != nil {
		return []entity.Product{}, err
	}
	log.Println("variants", variants)

	// mapVariants := map[int64][]entity.Variant{}
	// for _, variant := range variants {
	// 	mapVariants[variant.ID] = append(mapVariants[variant.ID], variant)
	// }

	for i, record := range records {
		records[i].Category = mapCategories[record.Category.ID]
		records[i].Variant = variants
		records[i].Image = fmt.Sprintf("%s/%s", s.cfg.BaseURL, record.Image)
	}

	return records, nil
}

func (s *productService) CreateProduct(product entity.Product, userID string) (record entity.Product, err error) {
	if err := product.Validate(); err != nil {
		return entity.Product{}, err
	}

	if product.Image == "" {
		product.Image = "images/default.png"
	}

	record, err = s.productRepo.CreateProduct(product, userID)
	if err != nil {
		return entity.Product{}, err
	}

	return record, nil
}

func (s *productService) GetProductByIDs(ids []int64) (records []entity.Product, err error) {
	records, err = s.productRepo.GetProductByIDs(ids)
	if err != nil {
		return []entity.Product{}, err
	}

	categoryIDs := []int64{}
	for _, record := range records {
		categoryIDs = append(categoryIDs, record.Category.ID)
	}

	categories, err := s.categoryService.GetCategoryByIDs(categoryIDs)
	if err != nil {
		return []entity.Product{}, err
	}

	mapCategories := map[int64]entity.Category{}
	for _, category := range categories {
		mapCategories[category.ID] = category
	}

	variantIDs := []int64{}
	for _, record := range records {
		for _, variant := range record.Variant {
			variantIDs = append(variantIDs, variant.ID)
		}
	}

	variants, err := s.variantService.GetVariantByIDs(variantIDs)
	if err != nil {
		return []entity.Product{}, err
	}

	mapVariants := map[int64][]entity.Variant{}
	for _, variant := range variants {
		mapVariants[variant.ID] = append(mapVariants[variant.ID], variant)
	}

	for i, record := range records {
		records[i].Category = mapCategories[record.Category.ID]
		records[i].Variant = mapVariants[record.ID]
		records[i].Image = fmt.Sprintf("%s/%s", s.cfg.BaseURL, record.Image)
	}

	return records, nil
}
