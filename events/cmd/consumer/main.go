package main

import (
	"fmt"

	"github.com/becardine/learning-go/events/utils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out, "queue")
	for m := range out {
		fmt.Println(string(m.Body))
		m.Ack(false) // acknowledge the message to rabbitmq
	}
}
