package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func RouteProducer(routKey string ) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_route",   // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//body := bodyFrom(routKey)
	err = ch.Publish(
		"logs_route", // exchange
		routKey,     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(routKey),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", routKey)
}