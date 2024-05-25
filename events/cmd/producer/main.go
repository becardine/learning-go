package main

import "github.com/becardine/learning-go/events/utils/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, World from Go!", "amq.direct")
}
