package mqtt_client

import (
	"encoding/json"
	"fmt"
	"log"
	"sistem-manajemen-armada/api/dto"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ReceivedMessageHandler(client mqtt.Client, message mqtt.Message) {
	log.Printf("topic :%s\n", message.Topic())

	payload, err := validateVehicleLocationPayload(message.Payload())

	if err != nil {
		log.Printf("invalid message (payload) :%s\n", err)

		return
	} else {
		_, createErr := vehicleLocationService.Create(payload)

		if createErr != nil {
			log.Printf("failed to create vehicle location: %s\n", createErr)

			return
		}

		log.Printf("message (payload) :%s\n", message.Payload())
	}
}

func ConnectedHandler(client mqtt.Client) {
	log.Println("connected to broker")
}

func ConnectionLostHandler(client mqtt.Client, err error) {
	log.Printf("connection to broker lost, caused by :%v", err)
}

func Subscribe(client mqtt.Client, topics []string) {
	for _, topic := range topics {
		token := client.Subscribe(topic, 1, nil)
		token.Wait()
		log.Printf("subscribed to topic: %s", topic)
	}
}

func Publish(client mqtt.Client, topic string, payload string) {
	text := fmt.Sprintf("message : %s", payload)

	token := client.Publish(topic, 0, false, text)
	token.Wait()

	log.Printf("message published to topic: %s\nmessage (payload) : %s", topic, payload)
}

// function related to api
func validateVehicleLocationPayload(msg []byte) (*dto.CreateVehicleLocationDto, error) {
	var payload dto.CreateVehicleLocationDto

	if err := json.Unmarshal(msg, &payload); err != nil {
		return nil, err
	}

	if err := validate.Struct(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
