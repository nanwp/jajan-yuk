package email_publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/nanwp/jajan-yuk/user/config"
	"github.com/nanwp/jajan-yuk/user/core/entity"
	"github.com/nanwp/jajan-yuk/user/core/publisher"
	"log"
)

type emailPublisher struct {
	cfg config.Config
}

func (e emailPublisher) SendEmail(ctx context.Context, email entity.Email) error {
	client, err := pubsub.NewClient(ctx, e.cfg.ProjectID)
	if err != nil {
		return err
	}

	defer client.Close()

	topic := client.Topic(e.cfg.TopicID)

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(email.Title),
		Attributes: map[string]string{
			"receiver": email.Receiver,
			"subject":  email.Subject,
			"body":     email.Body,
		},
	})

	_, err = result.Get(ctx)
	if err != nil {
		log.Printf("error to send email: %v", err.Error())
	}

	return nil
}

func New(cfg config.Config) publisher.EmailPublisher {
	return emailPublisher{
		cfg: cfg,
	}
}
