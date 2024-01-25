package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBType            string `envconfig:"DB_TYPE" default:"postgres"`
	DBHost            string `envconfig:"DB_HOST" default:"localhost"`
	DBPort            string `envconfig:"DB_PORT" default:"5432"`
	DBUsername        string `envconfig:"DB_USERNAME" default:"postgres"`
	DBPassword        string `envconfig:"DB_PASSWORD" default:"password"`
	DBName            string `envconfig:"DB_NAME" default:"jajan_yuk_product"`
	LogMode           bool   `envconfig:"LOG_MODE" default:"true"`
	DBMaxIdleConns    int    `envconfig:"DB_MAX_IDLE_CONNS" default:"5"`
	DBMaxOpenConns    int    `envconfig:"DB_MAX_OPEN_CONNS" default:"10"`
	DBConnMaxLifeTime int    `envconfig:"DB_CONN_MAX_LIFETIME" default:"10"`

	ProjectID string `envconfig:"PROJECT_ID" default:"jajan-yuk-409318"`
	TopicID   string `envconfig:"TOPIC_ID" default:"email"`

	HttpPort string `envconfig:"HTTP_PORT" default:"8083"`
	WebURL   string `envconfig:"WEB_URL" default:"http://localhost:3000"`

	RedisHost     string `envconfig:"REDIS_HOST" default:"127.0.0.1"`
	RedisPort     string `envconfig:"REDIS_PORT" default:"6379"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`
	RedisMaxIdle  int    `envconfig:"REDIS_MX_IDLE" default:"10"`

	BaseURL  string `envconfig:"BASE_URL" default:"http://localhost:8083"`
	AUTH_URL string `envconfig:"AUTH_URL" default:"http://localhost:8080"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
