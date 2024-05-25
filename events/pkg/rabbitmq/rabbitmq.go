package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		panic(err)
	}
	// create a channel of rabbitmq
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	messages, err := ch.Consume("queue", "go-consumer", false, false, false, false, nil)
	if err != nil {
		return err
	}
	for m := range messages {
		out <- m
	}
	return nil
}
