package rabbitmq

import (
	"log"
	"strings"
	"strconv"
	"time"
)

func WorkConsume() {
	conn, err := getConn()
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"world2", // name
		true,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for i := 0; i < 1; i++ {
		go func(i int) {
			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,   // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)
			failOnError(err, "Failed to register a consumer")

			for d := range msgs {
				msg := string(d.Body)
				interval, _ := strconv.Atoi(strings.Split(msg, "_")[1])
				log.Printf("==%d==Received a message: %s   Received",i, d.Body)
				slp := interval
				time.Sleep(time.Second*time.Duration(slp))

				if d.Ack(true) == nil {
					log.Printf("==%d==Received a message: %s   Finished",i, d.Body)
				}
			}
		}(i)
	}

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
