package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sistem-manajemen-armada/common/constant"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	producerChannel *amqp.Channel
	exchangeName    string
)

func registerProducer(channel *amqp.Channel) {
	producerChannel = channel
	exchangeName = constant.Exchange

	registerExchange(producerChannel, constant.Exchange)

	log.Printf("exchange used : %s", constant.Exchange)
}

func Publish(message []byte) error {
	if producerChannel == nil {
		return fmt.Errorf("producer channel not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := producerChannel.PublishWithContext(ctx,
		constant.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         message,
		})

	failOnError(err, "failed to publish a message")

	printMessage(message)

	return err
}

func printMessage(payload []byte) {
	var decodedPayload map[string]interface{}
	if err := json.Unmarshal(payload, &decodedPayload); err == nil {
		prettyJSON, _ := json.MarshalIndent(decodedPayload, "", "  ")

		log.Printf("[RabbitMQ Producer] - message sent:\n%s\n", string(prettyJSON))
	} else {
		log.Printf("[RabbitMQ Producer] - message sent: %s\n", string(payload))
	}
}
