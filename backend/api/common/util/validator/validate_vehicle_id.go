package validator

import (
	"regexp"
	"sistem-manajemen-armada/api/dto"
)

func IsValidVehicleId(vehicleId string) bool {
	valid, _ := regexp.MatchString(dto.VehicleIdPattern, vehicleId)

	return valid
}
