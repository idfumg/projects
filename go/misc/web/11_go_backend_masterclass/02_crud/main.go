package main

import (
	"log"
	"os"

	"myapp/config"
	"myapp/server"
	"myapp/store"
)

func main() {
	logger := log.New(os.Stdout, "Service: ", log.Ldate|log.Ltime|log.Lshortfile)

	config, err := config.New()
	if err != nil {
		log.Fatalf("Error! Could not init the config: %v\n", err)
	}

	store, err := store.NewPg(logger, config)
	if err != nil {
		log.Fatalf("Error! Could not init a store: %v\n", err)
	}

	s, err := server.NewServerMux(store, logger)
	if err != nil {
		log.Fatalf("Error! Could not init an API: %v\n", err)
	}

	log.Printf("Listening on port: %s\n", config.GetAppPort())
	log.Fatal(s.Serve(":"+config.GetAppPort(), server.AddCORS(s)))
}
