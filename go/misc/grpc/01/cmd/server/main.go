package main

import (
	"log"
	"myapp/internal/db"
	"myapp/internal/rocket"
	"myapp/internal/transport/grpc"
)

func Run() error {
	rocketStore, err := db.NewStore()
	if err != nil {
		return err
	}

	rocketService, err := rocket.NewService(rocketStore)
	if err != nil {
		return err
	}

	rocketHandler := grpc.NewHandler(rocketService)
	if err != nil {
		return err
	}

	if err := rocketHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
