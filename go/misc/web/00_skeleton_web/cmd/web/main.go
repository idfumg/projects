package main

import (
	"fmt"
	"log"
	"os"

	"myapp/cache"
	"myapp/config"
	"myapp/server/rest"
	"myapp/store"
)

func main() {
	logger := log.New(os.Stdout, "Service: ", log.Ldate|log.Ltime|log.Lshortfile)

	config, err := config.New()
	if err != nil {
		logger.Fatalf("Error! Could not init the config: %v\n", err)
	}
	logger.Printf("Config is created")

	cache, err := cache.New(config.GetUseCache())
	if err != nil {
		logger.Fatalf("Error! Could not init the cache: %v\n", err)
	}
	logger.Printf("Cache is created")

	// store, err := store.NewInMemory()
	store, err := store.New(logger, config)
	if err != nil {
		logger.Fatalf("Error! Could not init a store: %v\n", err)
	}
	logger.Printf("Database is connected")

	s, err := rest.New(store, logger, cache)
	if err != nil {
		logger.Fatalf("Error! Could not init an API: %v\n", err)
	}
	logger.Printf("Rest server is created")

	port := fmt.Sprintf(":%s", config.GetWebPort())
	logger.Printf("Listening on port: %s\n", port)
	logger.Fatal(s.Serve(port))
}
