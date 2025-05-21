package main

import (
	"fmt"
	"log"
	"sistem-manajemen-armada/api/router"
	"sistem-manajemen-armada/config"
	"sistem-manajemen-armada/database"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validatorInstance = validator.New()

func main() {
	database.InitializeDB()

	app := fiber.New()
	port := strconv.Itoa(config.Env.PORT)

	router.VehicleLocationRouter(app, validatorInstance)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Printf("Server running on port %s", port)

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Printf("Server failed to run: %v", err)
	}
}
