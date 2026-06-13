package config

import (
	"fmt"

	core_logger "github.com/Lesnekkk/golang-todo-app/internal/core/logger"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPAddr string `envconfig:"HTTP_ADDR" required:"true"`

	PostgresUser     string `envconfig:"POSTGRES_USER"     required:"true"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDB       string `envconfig:"POSTGRES_DB"       required:"true"`
	PostgresHost     string `envconfig:"POSTGRES_HOST"     required:"true"`
	PostgresPort     int    `envconfig:"POSTGRES_PORT"     required:"true"`

	RedisHost string `envconfig:"REDIS_HOST" required:"true"`
	RedisPort int    `envconfig:"REDIS_PORT" required:"true"`

	Logger core_logger.LoggerConfig
}

func New() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("process config: %w", err)
	}
	return cfg, nil
}

func (c Config) RedisAddr() string {
	return fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort)
}

func (c Config) PostgresDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDB,
	)
}
