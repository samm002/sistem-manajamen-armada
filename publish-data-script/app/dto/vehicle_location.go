package dto

const VehicleIdPattern = "^[A-Z]{1,2}[0-9]{1,4}[A-Z]{1,3}$"

type CreateVehicleLocationDto struct {
	VehicleId string  `json:"vehicle_id" validate:"omitempty,min=1,max=10"`
	Latitude  float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Timestamp int64   `json:"timestamp" validate:"omitempty,gt=0"`
}
