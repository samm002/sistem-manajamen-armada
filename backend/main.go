package main

import (
	"fmt"
	"log"
	"sistem-manajemen-armada/api/router"
	"sistem-manajemen-armada/config"
	"sistem-manajemen-armada/database"
	"sistem-manajemen-armada/pkg/mqtt_client"
	"sistem-manajemen-armada/pkg/rabbitmq"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validatorInstance = validator.New()

func main() {
	database.InitializeDB()
	mqtt_client.InitializeMqtt(validatorInstance)
	rabbitmq.InitializeRabbitMq()

	app := fiber.New()
	port := strconv.Itoa(config.Env.PORT)

	router.VehicleLocationRouter(app, validatorInstance)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Fleet Management System API")
	})

	log.Printf("Server running on port %s", port)

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Printf("Server failed to run: %v", err)
	}
}
