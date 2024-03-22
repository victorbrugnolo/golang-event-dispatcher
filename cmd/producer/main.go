package main

import "github.com/victorbrugnolo/golang-utils/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, "Hello world", "amq.direct")
}
