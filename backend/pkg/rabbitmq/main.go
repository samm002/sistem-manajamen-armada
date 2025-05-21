package rabbitmq

func InitializeRabbitMq() {
	conn := initiateConnection()

	consumerChannel := createChannel(conn)

	producerChannel := createChannel(conn)

	// Worker service yang membaca data dari geofence_alerts queue
	// Service subscribe pada topic yang telah dibind dengan exchange
	go registerConsumer(consumerChannel)

	registerProducer(producerChannel)
}
