package middleware

import (
	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/auth/config"
	"github.com/nanwp/jajan-yuk/auth/core/module"
	"github.com/nanwp/jajan-yuk/auth/handler/api"
	"github.com/nanwp/jajan-yuk/auth/pkg/conn"
	"github.com/nanwp/jajan-yuk/auth/repository/auth_repository"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (mux.Router, conn.CacheService) {
	coreRedis, redis := conn.InitRedis(cfg)

	authRepository := auth_repository.NewAuthRepository(db, redis)
	authUsecase := module.NewAuthUsecase(authRepository)

	router := mux.NewRouter()
	apiHttp := api.NewHttpHandler(authUsecase)

	router.HandleFunc("/login", apiHttp.Login)
	router.HandleFunc("/current-user", apiHttp.GetCurrentUser)
	router.HandleFunc("/refresh", apiHttp.RefreshToken)

	return *router, coreRedis
}
