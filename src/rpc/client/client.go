package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/jnsoft/rabbitmqdemo/src/misc"
	amqp "github.com/rabbitmq/amqp091-go"
)

func fibonacciRPC(n int) (res int, err error) {
	conn, err := amqp.Dial(RABBITMQ_CON_STRING)
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	corrId := RandomString(32)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	FailOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			FailOnError(err, "Failed to convert body to integer")
			break
		}
	}

	return
}

func main() {
	n := bodyFrom(os.Args)
	log.Printf(" [x] Requesting fib(%d)", n)

	res, err := fibonacciRPC(n)
	FailOnError(err, "Failed to handle RPC request")

	log.Printf(" [.] Got %d", res)
}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	FailOnError(err, "Failed to convert arg to integer")
	return n
}
