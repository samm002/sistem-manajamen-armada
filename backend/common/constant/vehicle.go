package constant

const (
	VehicleId1        = "AB1234CDE"
	VehicleId2        = "L5432AB"
	VehicleId3        = "H7890SM"
	GeofenceEventName = "geofence_entry"
	GeofenceLatitude  = -6.201000
	GeofenceLongitude = 106.817000
)

// Untuk trigger geofence diperkirakan butuh kisaran :
// Latitude  antara -6.201450  hingga -6.200550
// Longitude antara 106.816547 hingga 106.817453
// Pengujian berhasil (<50 m terhadap geofence) adalah : "latitude":-6.201200 & "longitude": 106.816800
