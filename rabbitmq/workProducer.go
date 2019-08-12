package rabbitmq

import (
	"github.com/streadway/amqp"
	"fmt"
	"time"
)

func WorkProducer() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"world2", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")


	body := "Hello World! : "

	for i:= 5; i<100000;i++ {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing {
				ContentType: "text/plain",
				Body:        []byte(body+fmt.Sprintf("_%d", i)),
				DeliveryMode: 2,
			})
		failOnError(err, "Failed to publish a message")
		time.Sleep(time.Second*5)
	}
}
