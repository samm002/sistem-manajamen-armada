package mqtt_client

import (
	"fmt"
	"log"
	"publish-data-script/app/common/constant"
	"publish-data-script/config"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
	client   mqtt.Client
	mqttDone = make(chan struct{})
)

func InitializeMqtt(validator *validator.Validate) mqtt.Client {
	brokerConnectionString := fmt.Sprintf("%s://%s:%d", config.Env.MQTT_PROTOCOL, config.Env.MQTT_BROKER_URL, config.Env.MQTT_BROKER_PORT)

	options := mqtt.NewClientOptions().AddBroker(brokerConnectionString).SetClientID(config.Env.MQTT_CLIENT_ID)

	options.SetDefaultPublishHandler(ReceivedMessageHandler)
	options.OnConnect = ConnectedHandler
	options.OnConnectionLost = ConnectionLostHandler
	options.SetUsername(config.Env.MQTT_BROKER_USERNAME)
	options.SetPassword(config.Env.MQTT_BROKER_PASSWORD)
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

	return client
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
