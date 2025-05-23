package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	MQTT_PROTOCOL        string `env:"MQTT_PROTOCOL" validate:"required"`
	MQTT_BROKER_URL      string `env:"MQTT_BROKER_URL" validate:"required"`
	MQTT_BROKER_PORT     int    `env:"MQTT_BROKER_PORT" validate:"required"`
	MQTT_BROKER_USERNAME string `env:"MQTT_BROKER_USERNAME" validate:"required"`
	MQTT_BROKER_PASSWORD string `env:"MQTT_BROKER_PASSWORD" validate:"required"`
	MQTT_CLIENT_ID       string `env:"MQTT_CLIENT_ID" validate:"required"`
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
	_ = godotenv.Load()

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	Env = validateEnv()
}
