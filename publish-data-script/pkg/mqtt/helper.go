package mqtt_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"publish-data-script/app/common/util"
	validatorUtil "publish-data-script/app/common/util/validator"
	"publish-data-script/app/dto"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ReceivedMessageHandler(client mqtt.Client, message mqtt.Message) {
	_, err := validateVehicleLocationPayload(message.Payload())

	if err != nil {
		log.Printf("invalid message (payload) :%s\n", err)
	}

	log.Printf("[MQTT Client Message Handler] - received message (payload) from topic %s :\n%s", message.Topic(), message.Payload())
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

func Publish(client mqtt.Client, topic string, payload interface{}) error {
	token := client.Publish(topic, 1, false, payload)
	token.Wait()

	if token.Error() != nil {
		return fmt.Errorf("failed to publish MQTT message: %w", token.Error())
	}

	log.Printf("[MQTT Publisher] - message published to topic: %s\nmessage (payload) : %s", topic, payload)

	return nil
}

// function related to api
func validateVehicleLocationPayload(msg []byte) (*dto.CreateVehicleLocationDto, error) {
	var payload dto.CreateVehicleLocationDto

	if payload.VehicleId == "" {
		payload.VehicleId = util.GenerateRandomVehicleId()
	}

	if payload.Timestamp == 0 {
		payload.Timestamp = time.Now().Unix()
	}

	if !validatorUtil.IsValidVehicleId(payload.VehicleId) {
		return nil, errors.New("invalid requestId format")
	}

	if err := json.Unmarshal(msg, &payload); err != nil {
		return nil, err
	}

	if err := validate.Struct(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
