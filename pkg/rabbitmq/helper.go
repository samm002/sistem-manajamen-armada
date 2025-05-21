package rabbitmq

import (
	"log"
	"sistem-manajemen-armada/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}

func initiateConnection() *amqp.Connection {
	connection, err := amqp.Dial(config.Env.RABBITMQ_URL)

	failOnError(err, "failed connecting to rabbitmq broker")

	log.Println("connected to rabbitmq broker")

	return connection
}

func createChannel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()

	failOnError(err, "failed to open a channel")

	log.Println("open channel success")

	return channel
}

func registerQueue(channel *amqp.Channel, queueName string) *amqp.Queue {
	queue, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "failed to declare a queue")

	return &queue
}

func registerExchange(channel *amqp.Channel, exchangeName string) {
	err := channel.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare an exchange")
}

func bindQueue(channel *amqp.Channel, queueName string, exchangeName string) {
	err := channel.QueueBind(
		queueName,    // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)

	failOnError(err, "Failed to bind a queue")
}
