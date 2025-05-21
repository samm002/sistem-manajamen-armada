package validator

import (
	"publish-data-script/app/dto"
	"regexp"
)

func IsValidVehicleId(vehicleId string) bool {
	valid, _ := regexp.MatchString(dto.VehicleIdPattern, vehicleId)

	return valid
}
