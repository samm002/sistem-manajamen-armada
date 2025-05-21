package response

type LocationDto struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type GeofenceEventMessageDto struct {
	VehicleID string      `json:"vehicle_id" validate:"required"`
	Event     string      `json:"event" validate:"required,oneof=geofence_entry geofence_exit"`
	Location  LocationDto `json:"location" validate:"required,dive"`
	Timestamp int64       `json:"timestamp" validate:"required"`
}
