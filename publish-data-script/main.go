package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"publish-data-script/app/common/constant"
	"publish-data-script/app/common/util"
	mqtt_client "publish-data-script/pkg/mqtt"
	"time"

	"github.com/go-playground/validator/v10"
)

var validatorInstance = validator.New()

func main() {
	mqttClient := mqtt_client.InitializeMqtt(validatorInstance)

	for {
		randomIndex := rand.Intn(len(constant.MqttTopics))
		topic := constant.MqttTopics[randomIndex]

		payload := util.GenerateRandomVehicleLocationData()

		encodedPayload, _ := json.Marshal(payload)

		err := mqtt_client.Publish(mqttClient, topic, encodedPayload)
		
		if err != nil {
			fmt.Printf("error encoding payload: %v\n", err)
			continue
		}

		time.Sleep(2 * time.Second)
	}
}
