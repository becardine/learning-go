package main

import (
	"fmt"
	"time"
)

type Message struct {
	id   int
	Body string
}

func main() {
	channel1 := make(chan Message)
	channel2 := make(chan Message)

	//rabbitmq
	go func() {
		time.Sleep(time.Second)
		msg := Message{id: 1, Body: "Hello from RabbitMQ!"}
		channel1 <- msg
	}()

	//kafka
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- Message{id: 2, Body: "Hello from Kafka!"}
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Printf("Received from RabbitMQ: %s from channel1\n", msg1.Body)
		case msg2 := <-channel2:
			fmt.Printf("Received from Kafka: %s from channel2\n", msg2.Body)
		case <-time.After(3 * time.Second):
			println("Timeout")
			// default:
			// 	println("No channel ready")
		}
	}

}
