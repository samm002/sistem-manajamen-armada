package mqtt_client

import (
	"fmt"
	"log"
	"sistem-manajemen-armada/api/repository"
	"sistem-manajemen-armada/api/service"
	"sistem-manajemen-armada/common/constant"
	"sistem-manajemen-armada/config"
	"sistem-manajemen-armada/database"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-playground/validator/v10"
)

var (
	validate                  = validator.New()
	vehicleLocationRepository repository.Repository
	vehicleLocationService    service.Service
	client                    mqtt.Client
	mqttDone                  = make(chan struct{})
)

func InitializeMqtt(validator *validator.Validate) {
	vehicleLocationRepository = repository.NewRepository(database.DB)
	vehicleLocationService = service.NewService(vehicleLocationRepository, validator)

	brokerConnectionString := fmt.Sprintf("%s://%s:%d", config.Env.MQTT_PROTOCOL, config.Env.BROKER_URL, config.Env.BROKER_PORT)

	options := mqtt.NewClientOptions().AddBroker(brokerConnectionString).SetClientID(config.Env.MQTT_CLIENT_ID)

	options.SetDefaultPublishHandler(ReceivedMessageHandler)
	options.OnConnect = ConnectedHandler
	options.OnConnectionLost = ConnectionLostHandler
	options.SetUsername(config.Env.BROKER_USERNAME)
	options.SetPassword(config.Env.BROKER_PASSWORD)
	options.SetKeepAlive(2 * time.Second)
	options.SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(options)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	Subscribe(client, constant.MqttTopics)

	go func() {
		log.Printf("mqtt client listening to %d subscribed topics :\n", len(constant.MqttTopics))

		for _, topic := range constant.MqttTopics {
			log.Printf("- %s", topic)
		}

		<-mqttDone
		log.Println("MQTT client shutting down...")
	}()

	log.Printf("waiting for messages, to exit press CTRL + C")
}

func Shutdown() {
	if client != nil && client.IsConnected() {
		client.Disconnect(250)
		log.Println("MQTT client disconnected")
	}

	if mqttDone != nil {
		close(mqttDone)
	}
}
