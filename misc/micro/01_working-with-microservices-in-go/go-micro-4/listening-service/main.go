package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to RabbitMQ")

	defer rabbitConn.Close()

	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		log.Panic(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	counts := 0

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ is not ready yet...")
			counts++
		} else {
			return c, nil
		}

		if counts > 10 {
			fmt.Println(err)
			return nil, err
		}

		backOff := time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
	}
}
