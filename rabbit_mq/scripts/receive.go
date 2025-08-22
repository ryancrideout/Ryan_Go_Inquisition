package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}

func main() {
	/*
		Notes:
		1) First we establish a connection
		2) Then we open a channel
		3) Then we declare a queue
		4) Then we consume messages from the queue.
	*/
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	messages, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for message := range messages {
			log.Println("Received a message:", string(message.Body))
		}
	}()

	log.Println("[*] Waiting for messages. To exit press CTRL+C.")

	<-forever
}
