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

	err = ch.ExchangeDeclare(
		"logs_topic", // name: "" is the default exchange
		"topic",      // type: direct, topic, headers and fanout
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := BodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"logs_topic",          // exchange
		SeverityFrom(os.Args), // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
