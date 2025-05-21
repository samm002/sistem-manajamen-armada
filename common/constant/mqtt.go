package constant

import "fmt"

func GenerateTopic(vehicleId string) string {
	return fmt.Sprintf("/fleet/vehicle/%s/location", vehicleId)
}

var (
	MqttTopicVehicle1 = GenerateTopic(VehicleId1)
	MqttTopicVehicle2 = GenerateTopic(VehicleId2)
	MqttTopicVehicle3 = GenerateTopic(VehicleId3)
)

var MqttTopics = []string{
	MqttTopicVehicle1,
	MqttTopicVehicle2,
	MqttTopicVehicle3,
}
