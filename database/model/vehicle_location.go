package model

import _ "ariga.io/atlas-provider-gorm/gormschema"

type VehicleLocation struct {
	VehicleId string  `gorm:"size:10;not null;uniqueIndex:idx_vehicleid_timestamp"`
	Latitude  float64 `gorm:"not null" json:"latitude"`
	Longitude float64 `gorm:"not null" json:"longitude"`
	Timestamp int64   `gorm:"not null;uniqueIndex:idx_vehicleid_timestamp"`
}
