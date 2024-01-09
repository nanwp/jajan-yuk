package middleware

import (
	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/user/config"
	"github.com/nanwp/jajan-yuk/user/core/module"
	"github.com/nanwp/jajan-yuk/user/handler/api"
	"github.com/nanwp/jajan-yuk/user/pkg/conn"
	"github.com/nanwp/jajan-yuk/user/publisher/email_publisher"
	"github.com/nanwp/jajan-yuk/user/repository/user_repository"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (mux.Router, conn.CacheService) {
	coreRedis, redis := conn.InitRedis(cfg)

	userRepository := user_repository.New(db, redis)
	emailPublisher := email_publisher.New(cfg)
	userUsecase := module.NewUserRepository(cfg, userRepository, emailPublisher)

	httpHandler := api.NewHttpHandler(userUsecase)

	router := mux.NewRouter()
	router.HandleFunc("/register/{role}", httpHandler.Register)

	return *router, coreRedis
}
