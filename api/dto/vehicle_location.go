package dto

import (
	"sistem-manajemen-armada/common/constant"
	"sistem-manajemen-armada/common/dto"
	"sistem-manajemen-armada/database/model"
)

type VehicleLocationMapper interface {
	ConstructUpdatePayload() map[string]interface{}
	ToGeofenceEventMessage() model.VehicleLocation
	ToModel() model.VehicleLocation
}

const VehicleIdPattern = "^[A-Z]{1,2}[0-9]{1,4}[A-Z]{1,3}$"

type CreateVehicleLocationDto struct {
	VehicleId string  `json:"vehicle_id" validate:"required,min=5,max=10"`
	Latitude  float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Timestamp int64   `json:"timestamp" validate:"required,gt=0"`
}

type UpdateVehicleLocationDto struct {
	VehicleId string   `json:"vehicle_id" validate:"required"`
	Latitude  *float64 `json:"latitude,omitempty" validate:"omitempty,gte=-90,lte=90"`
	Longitude *float64 `json:"longitude,omitempty" validate:"omitempty,gte=-180,lte=180"`
	Timestamp *int64   `json:"timestamp,omitempty" validate:"omitempty,gt=0"`
}

type VehicleLocationDto struct {
	VehicleId string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

func ToResponse(v *model.VehicleLocation) VehicleLocationDto {
	return VehicleLocationDto{
		VehicleId: v.VehicleId,
		Latitude:  v.Latitude,
		Longitude: v.Longitude,
		Timestamp: v.Timestamp,
	}
}

func (v *CreateVehicleLocationDto) ToModel() model.VehicleLocation {
	return model.VehicleLocation{
		VehicleId: v.VehicleId,
		Latitude:  v.Latitude,
		Longitude: v.Longitude,
		Timestamp: v.Timestamp,
	}
}

func (v *UpdateVehicleLocationDto) ToModel() model.VehicleLocation {
	model := model.VehicleLocation{
		VehicleId: v.VehicleId,
	}

	if v.Latitude != nil {
		model.Latitude = *v.Latitude
	}

	if v.Longitude != nil {
		model.Longitude = *v.Longitude
	}

	if v.Timestamp != nil {
		model.Timestamp = *v.Timestamp
	}

	return model
}

func (v *VehicleLocationDto) ToModel() model.VehicleLocation {
	return model.VehicleLocation{
		VehicleId: v.VehicleId,
		Latitude:  v.Latitude,
		Longitude: v.Longitude,
		Timestamp: v.Timestamp,
	}
}

func (v *CreateVehicleLocationDto) ToGeofenceEventMessage() dto.GeofenceEventMessageDto {
	return dto.GeofenceEventMessageDto{
		VehicleId: v.VehicleId,
		Event:     constant.GeofenceEventName,
		Location: dto.LocationDto{
			Latitude:  v.Latitude,
			Longitude: v.Longitude,
		},
		Timestamp: v.Timestamp,
	}
}

func (v *UpdateVehicleLocationDto) ConstructUpdatePayload() map[string]interface{} {
	payload := make(map[string]interface{})

	if v.Latitude != nil {
		payload["latitude"] = v.Latitude
	}
	if v.Longitude != nil {
		payload["longitude"] = v.Longitude
	}
	if v.Timestamp != nil {
		payload["timestamp"] = v.Timestamp
	}

	return payload
}
