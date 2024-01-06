package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBType            string `envconfig:"DB_TYPE" default:"postgres"`
	DBHost            string `envconfig:"DB_HOST" default:"103.171.182.206"`
	DBPort            string `envconfig:"DB_PORT" default:"5432"`
	DBUsername        string `envconfig:"DB_USERNAME" default:"postgres"`
	DBPassword        string `envconfig:"DB_PASSWORD" default:"Latihan"`
	DBName            string `envconfig:"DB_NAME" default:"jajan_yuk_user"`
	LogMode           bool   `envconfig:"LOG_MODE" default:"true"`
	DBMaxIdleConns    int    `envconfig:"DB_MAX_IDLE_CONNS" default:"5"`
	DBMaxOpenConns    int    `envconfig:"DB_MAX_OPEN_CONNS" default:"10"`
	DBConnMaxLifeTime int    `envconfig:"DB_CONN_MAX_LIFETIME" default:"10"`

	RedisHost     string `envconfig:"REDIS_HOST" default:"127.0.0.1"`
	RedisPort     string `envconfig:"REDIS_PORT" default:"6379"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`
	RedisMaxIdle  int    `envconfig:"REDIS_MX_IDLE" default:"10"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}