package util

import (
	"math"
	"sistem-manajemen-armada/common/constant"
)

// Memakai rumus haversine untuk menghitung jarak koordinat latitude dan longitude
func CalculateCoordinateDistance(latitude1, longitude1, latitude2, longitude2 float64) float64 {
	latitudeDifference := (latitude2 - latitude1) * math.Pi / 180.0
	longitudeDifference := (longitude2 - longitude1) * math.Pi / 180.0

	latitude1 = latitude1 * math.Pi / 180.0
	latitude2 = latitude2 * math.Pi / 180.0

	haversine := math.Sin(latitudeDifference/2)*math.Sin(latitudeDifference/2) +
		math.Cos(latitude1)*math.Cos(latitude2)*math.Sin(longitudeDifference/2)*math.Sin(longitudeDifference/2)

	centralAngle := 2 * math.Atan2(math.Sqrt(haversine), math.Sqrt(1-haversine))

	return constant.EarthRadius * centralAngle // meter
}
