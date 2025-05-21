package mqtt_client

import (
	"fmt"
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
)

func InitializeMqtt() {
	vehicleLocationRepository = repository.NewRepository(database.DB)
	vehicleLocationService = service.NewService(vehicleLocationRepository)
	
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
}
