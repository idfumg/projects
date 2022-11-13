package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("RabbitMQ is connecting...")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("RabbitMQ is connected")
	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}
	_, err = r.Channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("Message was published")
	return nil
}

func (r *RabbitMQ) Consume() {
	msgs, err := r.Channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Consuming error:", err)
		return
	}
	for msg := range msgs {
		fmt.Printf("Message received: %s\n", msg.Body)
	}

}
