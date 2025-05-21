package util

import (
	"encoding/json"
	"sistem-manajemen-armada/api/dto"

	"github.com/go-playground/validator/v10"
)

func GenerateGeofenceEventMessage(payload *dto.CreateVehicleLocationDto, validate *validator.Validate) (*[]byte, error) {
	geofenceEventMessage := payload.ToGeofenceEventMessage()

	err := validate.Struct(geofenceEventMessage)
	if err != nil {
		return nil, err
	}

	encodedPayload, err := json.Marshal(geofenceEventMessage)
	if err != nil {
		return nil, err
	}

	return &encodedPayload, nil
}
