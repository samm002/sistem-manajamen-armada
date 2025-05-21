package util

import (
	"publish-data-script/app/common/constant"
	"publish-data-script/app/dto"
	"time"
)

func GenerateRandomVehicleLocationData() dto.CreateVehicleLocationDto {
	return dto.CreateVehicleLocationDto{
		VehicleId: GenerateRandomVehicleId(),
		Latitude:  GenerateRandomCoordinate(constant.MaxLatitude),
		Longitude: GenerateRandomCoordinate(constant.MaxLongitude),
		Timestamp: time.Now().Unix(),
	}
}
