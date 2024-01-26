package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/product/config"
	"github.com/nanwp/jajan-yuk/product/core/module"
	"github.com/nanwp/jajan-yuk/product/handler/api"
	"github.com/nanwp/jajan-yuk/product/pkg/conn"
	"github.com/nanwp/jajan-yuk/product/repository/category_repository"
	"github.com/nanwp/jajan-yuk/product/repository/product_repository"
	"github.com/nanwp/jajan-yuk/product/repository/variant_repository"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (http.Handler, conn.CacheService) {
	coreRedis, _ := conn.InitRedis(cfg)

	categoryRepo := category_repository.New(cfg, db)
	variantRepo := variant_repository.New(cfg, db)
	productRepo := product_repository.New(cfg, db)

	categoryService := module.NewCategoryService(categoryRepo)
	variantService := module.NewVariantService(variantRepo)
	productService := module.NewProductService(cfg, productRepo, variantService, categoryService)

	httpHandler := api.NewHTTPHandler(variantService, categoryService, productService)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/variant/ids", httpHandler.GetVariantByIDs).Methods("GET")
	router.HandleFunc("/api/v1/variant/id", httpHandler.GetVariantByID).Methods("GET")
	router.HandleFunc("/api/v1/variant", httpHandler.GetVariantByUserCreated).Methods("GET")
	router.HandleFunc("/api/v1/variant", httpHandler.CreateVariant).Methods("POST")

	router.HandleFunc("/api/v1/category", httpHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/api/v1/category", httpHandler.GetCategoryByUserCreated).Methods("GET")
	router.HandleFunc("/api/v1/category/ids", httpHandler.GetCategoryByIDs).Methods("GET")

	router.HandleFunc("/api/v1/product", httpHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/api/v1/product", httpHandler.GetProductByUserCreated).Methods("GET")
	// router.HandleFunc("/api/v1/product/ids", httpHandler.GetProductByIDs).Methods("GET")
	router.HandleFunc("/api/v1/product/pedagang", httpHandler.GetProductByUserID).Methods("GET")

	router.HandleFunc("/images/{image_name}", httpHandler.GetImage).Methods("GET")

	handler := c.Handler(router)

	return handler, coreRedis
}
