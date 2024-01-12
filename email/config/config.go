package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ProjectID   string `envconfig:"PROJECT_ID" default:"jajan-yuk-409318"`
	SubcriberID string `envconfig:"SUBCRIBER_ID"  default:"email-sub"`

	SMTP_HOST string `envconfig:"SMTP_HOST" default:"mx2.mailspace.id"`
	SMTP_PORT int    `envconfig:"SMTP_PORT" default:"465"`
	SMTP_USER string `envconfig:"SMTP_USER" default:"noreply@hiline.my.id"`
	SMTP_PASS string `envconfig:"SMTP_PASS" default:"password"`
	SMTP_FROM string `envconfig:"SMTP_FROM" default:"noreply<noreply@hiline.my.id>"`
}

func Get() Config {
	var c Config
	envconfig.MustProcess("", &c)
	return c
}
