package router

import (
	"sistem-manajemen-armada/api/controller"
	"sistem-manajemen-armada/api/repository"
	"sistem-manajemen-armada/api/service"
	"sistem-manajemen-armada/database"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func VehicleLocationRouter(router fiber.Router, validator *validator.Validate) {
	repository := repository.NewRepository(database.DB)
	service := service.NewService(repository)
	controller := controller.NewController(service, validator)

	vehicleLocationRouter := router.Group("/vehicle-locations")

	vehicleLocationRouter.Post("", controller.Create)
	vehicleLocationRouter.Get("", controller.FindAll)
	vehicleLocationRouter.Get("/:vehicleId/history", controller.FindHistory)
	vehicleLocationRouter.Get("/:vehicleId/location", controller.FindLatestLocationById)
	vehicleLocationRouter.Put("/:vehicleId", controller.Update)
	vehicleLocationRouter.Delete("/:vehicleId", controller.Delete)
}
