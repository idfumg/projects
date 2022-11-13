package main

import (
	"log"
	"net/http"
	"os"

	"myapp/server"
	"myapp/config"
	"myapp/store"
)

func main() {
	logger := log.New(os.Stdout, "Service: ", log.Ldate|log.Ltime|log.Lshortfile)

	config, err := config.New()
	if err != nil {
		log.Fatalf("Error! Could not init the config: %v\n", err)
	}

	// store, err := store.NewStoreMemory()
	store, err := store.NewStorePg(logger, config)
	if err != nil {
		log.Fatalf("Error! Could not init a store: %v\n", err)
	}

	s, err := server.NewServerMux(store, logger)
	if err != nil {
		log.Fatalf("Error! Could not init an API: %v\n", err)
	}

	log.Fatal(http.ListenAndServe(":50051", server.AddCORS(s)))
}
