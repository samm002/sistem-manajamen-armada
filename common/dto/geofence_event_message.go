package dto

type LocationDto struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type GeofenceEventMessageDto struct {
	VehicleId string      `json:"vehicle_id" validate:"required"`
	Event     string      `json:"event" validate:"required,oneof=geofence_entry geofence_exit"`
	Location  LocationDto `json:"location" validate:"required"`
	Timestamp int64       `json:"timestamp" validate:"required"`
}

type CreateGeofenceEventMessageDto struct {
	VehicleId string  `json:"vehicle_id" validate:"required,min=5,max=10"`
	Event     string  `json:"event" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Timestamp int64   `json:"timestamp" validate:"required,gt=0"`
}
