package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	PORT         int    `env:"PORT" validate:"required"`
	DB_HOST      string `env:"DB_HOST" validate:"required"`
	DB_PORT      string `env:"DB_PORT" validate:"required"`
	DB_USER      string `env:"DB_USER" validate:"required"`
	DB_PASSWORD  string `env:"DB_PASSWORD" validate:"required"`
	DB_NAME      string `env:"DB_NAME" validate:"required"`
	DB_SSL_MODE  string `env:"DB_SSL_MODE" validate:"required"`
	DB_TIME_ZONE string `env:"DB_TIME_ZONE" validate:"required"`
}

var Env *EnvConfig

func validateEnv() *EnvConfig {
	var cfg EnvConfig

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("Failed to load env vars: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		log.Fatalf("Invalid environment config: %v", err)
	}

	return &cfg
}

func init() {
	err := godotenv.Load()
	
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	Env = validateEnv()
}
