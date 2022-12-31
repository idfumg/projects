package main

import (
	"context"

	"github.com/go-redis/redis/v9"
	log "github.com/sirupsen/logrus"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error! Redis is not connected: %v\n", err)
		return
	}

	ctx := context.Background()
	sub := client.Subscribe(ctx, "coords")

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Infof("%+v\n", msg)
	}
}
