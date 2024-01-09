package publisher

import (
	"context"
	"github.com/nanwp/jajan-yuk/user/core/entity"
)

type EmailPublisher interface {
	SendEmail(ctx context.Context, email entity.Email) error
}
