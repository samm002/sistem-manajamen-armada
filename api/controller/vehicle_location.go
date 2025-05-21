package controller

import (
	"errors"
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
	FindHistory(ctx *fiber.Ctx) error
	FindLatestLocationById(ctx *fiber.Ctx) error
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
		return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("create vehicle location failed", errors.New("invalid request body")))
	}

	if payload.VehicleId == "" {
		payload.VehicleId = util.GenerateRandomVehicleId()
	}

	if !validatorUtil.IsValidVehicleId(payload.VehicleId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("create vehicle location failed", errors.New("invalid requestId format")))
	}

	if err := c.validator.Struct(&payload); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("create vehicle location failed", errors.New("validation failed")))
	}

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

func (c *controller) FindHistory(ctx *fiber.Ctx) error {
	vehicleId := ctx.Params("vehicleId")
	start := ctx.Query("start")
	end := ctx.Query("end")

	if !validatorUtil.IsValidVehicleId(vehicleId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("get vehicle location history failed", errors.New("invalid vehicleId format")))
	}

	convertedStart, err := util.ConvertStringToIntPointer(start)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			responseUtil.GenerateFailedResponse("get vehicle location history failed", errors.New("invalid start format")),
		)
	}

	convertedEnd, err := util.ConvertStringToIntPointer(end)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			responseUtil.GenerateFailedResponse("get vehicle location history failed", errors.New("invalid end format")),
		)
	}

	if err := validatorUtil.IsValidDateRange(convertedStart, convertedEnd); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("get vehicle location history failed", err))
	}

	vehicleLocation, err := c.service.FindHistory(vehicleId, convertedStart, convertedEnd)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.Status(http.StatusNotFound).JSON(responseUtil.GenerateFailedResponse("get vehicle location history failed", err))
		} else {
			return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("get vehicle location history failed", err))
		}
	}

	return ctx.JSON(responseUtil.GenerateSuccessResponse("get vehicle location history success", vehicleLocation))
}

func (c *controller) FindLatestLocationById(ctx *fiber.Ctx) error {
	vehicleId := ctx.Params("vehicleId")

	if !validatorUtil.IsValidVehicleId(vehicleId) {
		return ctx.Status(fiber.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("get vehicle location latest location failed", errors.New("invalid vehicleId format")))
	}

	vehicleLocation, err := c.service.FindLatestLocationById(vehicleId)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.Status(http.StatusNotFound).JSON(responseUtil.GenerateFailedResponse("get vehicle location latest location failed", err))
		} else {
			return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("get vehicle location latest location failed", err))
		}
	}

	return ctx.JSON(responseUtil.GenerateSuccessResponse("get vehicle location latest location success", vehicleLocation))
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
		return ctx.Status(http.StatusBadRequest).JSON(responseUtil.GenerateFailedResponse("validation failed", err))
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
