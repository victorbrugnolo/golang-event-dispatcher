package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/victorbrugnolo/golang-utils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "default-queue")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		_ = msg.Ack(false)
	}
}
