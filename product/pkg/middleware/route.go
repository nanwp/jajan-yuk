package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/product/config"
	"github.com/nanwp/jajan-yuk/product/core/module"
	"github.com/nanwp/jajan-yuk/product/handler/api"
	"github.com/nanwp/jajan-yuk/product/pkg/conn"
	"github.com/nanwp/jajan-yuk/product/repository/variant_repository"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (http.Handler, conn.CacheService) {
	coreRedis, _ := conn.InitRedis(cfg)

	variantRepo := variant_repository.New(cfg, db)
	variantService := module.NewVariantService(variantRepo)
	httpHandler := api.NewHTTPHandler(variantService)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/variant/ids", httpHandler.GetVariantByIDs).Methods("GET")
	router.HandleFunc("/api/v1/variant", httpHandler.GetVariantByUserCreated).Methods("GET")
	router.HandleFunc("/api/v1/variant", httpHandler.CreateVariant).Methods("POST")

	handler := c.Handler(router)

	return handler, coreRedis
}
