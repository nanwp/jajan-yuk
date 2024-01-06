package service

import (
	"fmt"
	"github.com/nanwp/jajan-yuk/email/config"
	"github.com/nanwp/jajan-yuk/email/core/entity"
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendEmail(email entity.Email) error
}

type emailService struct {
	cfg config.Config
}

func NewEmailService(cfg config.Config) EmailService {
	return &emailService{
		cfg: cfg,
	}
}

func (s *emailService) SendEmail(email entity.Email) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", fmt.Sprintf("%s<%s>", email.Title, s.cfg.SMTP_USER))
	mail.SetHeader("To", email.Receiver)
	mail.SetHeader("Subject", email.Subject)
	mail.SetBody("text/html", email.Body)

	dialer := gomail.NewDialer(s.cfg.SMTP_HOST, s.cfg.SMTP_PORT, s.cfg.SMTP_USER, s.cfg.SMTP_PASS)
	if err := dialer.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
