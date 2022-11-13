package main

import (
	"fmt"
	"maypp/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go RabbitMQ Crash Course")
	app := App{
		Rmq: rabbitmq.NewRabbitMQService(),
	}
	if err := app.Rmq.Connect(); err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		return err
	}
	defer app.Rmq.Conn.Close()
	err := app.Rmq.Publish("Hi")
	if err != nil {
		return err
	}
	app.Rmq.Consume()
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error running an application:", err)
		return
	}
}
