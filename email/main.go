package main

import (
	"context"
	"github.com/nanwp/jajan-yuk/email/config"
	"github.com/nanwp/jajan-yuk/email/core/service"
	"github.com/nanwp/jajan-yuk/email/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	projectID = "jajan-yuk-409318"
	subID     = "email-sub"
)

type email struct {
	to   string
	body string
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Get()
	emailService := service.NewEmailService(cfg)
	handler := handler.NewPubsub(cfg, emailService)

	go func() {
		err := handler.Subcriber(ctx)
		if err != nil {
			log.Fatalf("failed to pull message: %v\n", err)
		}
	}()

	log.Println("subscriber is starting ....")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	log.Printf("signal %d received, shutting down gracefully...", <-quit)
	cancel()

	log.Println("finished graceful shutdown")
}
