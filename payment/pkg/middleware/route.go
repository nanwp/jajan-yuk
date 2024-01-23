package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/payment/config"
	"github.com/nanwp/jajan-yuk/payment/core/service"
	"github.com/nanwp/jajan-yuk/payment/handler/api"
	"github.com/nanwp/jajan-yuk/payment/pkg/conn"
	"github.com/nanwp/jajan-yuk/payment/repository/payment_repository"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func InitRouter(cfg config.Config, db *gorm.DB) (http.Handler, conn.CacheService) {
	coreRedis, _ := conn.InitRedis(cfg)

	paymentRepo := payment_repository.New(cfg, db)
	midtransService := service.NewMidtransService(cfg)
	paymentService := service.NewPaymentService(paymentRepo, midtransService)
	httpHandler := api.NewHttpHandler(paymentService, midtransService)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://jajan-yuk.pegelinux.my.id"},
		AllowCredentials: true,
	})

	router := mux.NewRouter()
	router.HandleFunc("/payment/init", httpHandler.InitializePayment).Methods(http.MethodPost)
	router.HandleFunc("/payment/callback", httpHandler.PaymentNotification).Methods(http.MethodPost)

	handler := c.Handler(router)

	return handler, coreRedis
}
