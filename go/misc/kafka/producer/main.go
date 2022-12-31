package main

import (
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	servers := viper.GetStringSlice("kafka.servers")
	producer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New()
	app.Post("/openAccount", accountController.OpenAccount)
	app.Post("/depositFund", accountController.DepositFund)
	app.Post("/withdrawFund", accountController.WithdrawFund)
	app.Post("/closeAccount", accountController.CloseAccount)
	app.Listen(":8080")
}

// func main() {
// 	servers := []string{"localhost:9092"}
// 	producer, err := sarama.NewSyncProducer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer producer.Close()
// 	msg := sarama.ProducerMessage{
// 		Topic: "bondhello",
// 		Value: sarama.StringEncoder("Hello world!"),
// 	}
// 	partition, offset, err := producer.SendMessage(&msg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("partition=%v, offset=%v\n", partition, offset)
// }
