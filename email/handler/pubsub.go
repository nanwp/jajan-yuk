package handler

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/nanwp/jajan-yuk/email/config"
	"github.com/nanwp/jajan-yuk/email/core/entity"
	"github.com/nanwp/jajan-yuk/email/core/service"
	"log"
)

type Handler interface {
	Subcriber(ctx context.Context) error
}

type handler struct {
	cfg          config.Config
	emailService service.EmailService
}

func NewPubsub(cfg config.Config, emailService service.EmailService) Handler {
	return &handler{cfg: cfg, emailService: emailService}
}

func (ps *handler) Subcriber(ctx context.Context) error {
	client, err := pubsub.NewClient(ctx, ps.cfg.ProjectID)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	subs := client.Subscription(ps.cfg.SubcriberID)

	err = subs.Receive(ctx, func(_ context.Context, message *pubsub.Message) {
		msg := entity.Email{
			Title:    string(message.Data),
			Receiver: message.Attributes["receiver"],
			Subject:  message.Attributes["subject"],
			Body:     message.Attributes["body"],
		}

		err := ps.emailService.SendEmail(msg)
		if err != nil {
			log.Printf("failed send email to : %v, error : %v", msg.Receiver, err)
			return
		}

		log.Printf("success send email to: %v, message: %v", msg.Receiver, msg.Body)

		message.Ack()
	})

	if err != nil {
		return err
	}

	return nil
}
