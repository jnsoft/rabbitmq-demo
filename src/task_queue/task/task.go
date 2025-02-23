package main

import (
	"context"
	"log"
	"os"
	"time"

	. "github.com/jnsoft/rabbitmqdemo/src/misc"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial(RABBITMQ_CON_STRING)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := BodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // persistence guarantees aren't strong, a stronger guarantee is to use publisher confirms.
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
