package rabbitmq

import (
	"log"
	"sistem-manajemen-armada/common/constant"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	consumerDone = make(chan struct{})
)

func registerConsumer(channel *amqp.Channel) {
	err := channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	registerExchange(channel, constant.Exchange)
	log.Printf("exchange used : %s", constant.Exchange)

	queue := registerQueue(channel, constant.GeofenceAlertQueue)

	bindQueue(channel, constant.GeofenceAlertQueue, constant.Exchange)

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("[RabbitMQ Consumer Worker Service] - received a message: %s", d.Body)
		}

	}()

	log.Printf("waiting for messages, to exit press CTRL + C")
	<-consumerDone
}
