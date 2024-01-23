package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/user/client/pedagang"
	"github.com/nanwp/jajan-yuk/user/config"
	"github.com/nanwp/jajan-yuk/user/core/module"
	"github.com/nanwp/jajan-yuk/user/handler/api"
	"github.com/nanwp/jajan-yuk/user/pkg/conn"
	"github.com/nanwp/jajan-yuk/user/publisher/email_publisher"
	"github.com/nanwp/jajan-yuk/user/repository/user_repository"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (http.Handler, conn.CacheService) {
	coreRedis, redis := conn.InitRedis(cfg)

	userRepository := user_repository.New(db, redis)
	emailPublisher := email_publisher.New(cfg)
	pedagangClient := pedagang.NewPedagangClient(cfg)
	userUsecase := module.NewUserRepository(cfg, userRepository, emailPublisher, pedagangClient)
	httpHandler := api.NewHttpHandler(userUsecase)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://jajan-yuk.pegelinux.my.id"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/register/{role}", httpHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/verification", httpHandler.Verification).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/request-reset-password", httpHandler.RequestResetPassword).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/reset-password", httpHandler.ResetPassword).Methods(http.MethodPost)

	handler := c.Handler(router)

	return handler, coreRedis
}
