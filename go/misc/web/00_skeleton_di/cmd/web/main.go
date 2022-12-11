package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"myapp/config"
	"myapp/server"
	"myapp/store"
)

func main() {
	logger := log.New(os.Stdout, "Service: ", log.Ldate|log.Ltime|log.Lshortfile)

	config, err := config.New()
	if err != nil {
		logger.Fatalf("Error! Could not init the config: %v\n", err)
	}
	logger.Printf("Config is read")

	// store, err := store.NewInMemory()
	store, err := store.NewPg(logger, config)
	if err != nil {
		logger.Fatalf("Error! Could not init a store: %v\n", err)
	}
	logger.Printf("Database is connected")

	s, err := server.NewServerMux(store, logger)
	if err != nil {
		logger.Fatalf("Error! Could not init an API: %v\n", err)
	}

	port := fmt.Sprintf(":%s", config.GetWebPort())
	logger.Printf("Listening on port: %s\n", port)
	logger.Fatal(http.ListenAndServe(port, server.AddCORS(s)))
}
