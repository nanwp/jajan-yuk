package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/auth/config"
	"github.com/nanwp/jajan-yuk/auth/core/module"
	"github.com/nanwp/jajan-yuk/auth/handler/api"
	"github.com/nanwp/jajan-yuk/auth/pkg/conn"
	"github.com/nanwp/jajan-yuk/auth/repository/auth_repository"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (http.Handler, conn.CacheService) {
	coreRedis, redis := conn.InitRedis(cfg)

	authRepository := auth_repository.NewAuthRepository(db, redis)
	authUsecase := module.NewAuthUsecase(authRepository)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://jajan-yuk.pegelinux.my.id"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	apiHttp := api.NewHttpHandler(authUsecase)

	router.HandleFunc("/api/v1/login", apiHttp.Login).Methods("POST")
	router.HandleFunc("/api/v1/current-user", apiHttp.GetCurrentUser).Methods("GET")
	router.HandleFunc("/api/v1/refresh", apiHttp.RefreshToken).Methods("POST")
	router.HandleFunc("/api/v1/validate-secret-key", apiHttp.ValidateSecretKey).Methods("POST")

	handler := c.Handler(router)

	return handler, coreRedis
}
