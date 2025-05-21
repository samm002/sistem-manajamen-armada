package controller

import (
	"errors"
	"log"
	"net/http"
	"sistem-manajemen-armada/api/common/util"
	responseUtil "sistem-manajemen-armada/api/common/util/response"
	validatorUtil "sistem-manajemen-armada/api/common/util/validator"
	"sistem-manajemen-armada/api/dto"
	"sistem-manajemen-armada/api/service"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Create(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type controller struct {
	service   service.Service
	validator *validator.Validate
}

func NewController(service service.Service, validator *validator.Validate) Controller {
	return &controller{service, validator}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	var payload dto.CreateVehicleLocationDto

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("invalid request body", err))
	}

	if payload.VehicleId == "" {
		payload.VehicleId = util.GenerateRandomVehicleId()
	}

	if !validatorUtil.IsValidVehicleId(payload.VehicleId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("invalid vehicleId format", nil))
	}

	if err := c.validator.Struct(&payload); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(responseUtil.GenerateFailedResponse("validation failed", err))
	}

	log.Println("payload :", payload)
	log.Println("&payload :", &payload)
	vehicleLocation, err := c.service.Create(&payload)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("create vehicle location failed", err))
	}

	return ctx.Status(http.StatusCreated).JSON(responseUtil.GenerateSuccessResponse("create vehicle location success", vehicleLocation))
}

func (c *controller) FindAll(ctx *fiber.Ctx) error {
	vehicleLocations, err := c.service.FindAll()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(responseUtil.GenerateFailedResponse("failed to get vehicle locations", err))
	}

	return ctx.JSON(responseUtil.GenerateSuccessResponse("get vehicle locations success", vehicleLocations))
}

func (c *controller) FindById(ctx *fiber.Ctx) error {
	vehicleId := ctx.Params("vehicleId")

	if !validatorUtil.IsValidVehicleId(vehicleId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("get vehicle location failed", errors.New("invalid vehicleId format")))
	}

	vehicleLocation, err := c.service.FindById(vehicleId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.Status(http.StatusNotFound).JSON(responseUtil.GenerateFailedResponse("get vehicle location failed", err))
		} else {
			return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("get vehicle location failed", err))
		}
	}

	return ctx.JSON(responseUtil.GenerateSuccessResponse("get vehicle location success", vehicleLocation))
}

func (c *controller) Update(ctx *fiber.Ctx) error {
	vehicleId := ctx.Params("vehicleId")

	if !validatorUtil.IsValidVehicleId(vehicleId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("update vehicle location failed", errors.New("invalid vehicleId format")))
	}

	var payload dto.UpdateVehicleLocationDto

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("invalid request body", err))
	}

	if err := c.validator.Struct(&payload); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(responseUtil.GenerateFailedResponse("validation failed", err))
	}

	vehicleLocation, err := c.service.Update(vehicleId, &payload)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.Status(http.StatusNotFound).JSON(responseUtil.GenerateFailedResponse("update vehicle location failed", err))
		} else {
			return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("update vehicle location failed", err))
		}
	}

	return ctx.Status(http.StatusCreated).JSON(responseUtil.GenerateSuccessResponse("update vehicle location success", vehicleLocation))
}

func (c *controller) Delete(ctx *fiber.Ctx) error {
	vehicleId := ctx.Params("vehicleId")

	if !validatorUtil.IsValidVehicleId(vehicleId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("delete vehicle location failed", errors.New("invalid vehicleId format")))
	}

	err := c.service.Delete(vehicleId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.Status(http.StatusNotFound).JSON(responseUtil.GenerateFailedResponse("delete vehicle location failed", err))
		} else {
			return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("delete vehicle location failed", err))
		}
	}

	return ctx.JSON(responseUtil.GenerateSuccessResponse("delete location success", nil))
}
