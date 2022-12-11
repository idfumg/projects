package main

import (
	"log"
	"myapp/pkg/api"
	"myapp/pkg/config"
	"net/http"
)

const (
	Port = ":8080"
)

func main() {
	appConfig, err := initAll()
	if err != nil {
		log.Fatalln(err)
	}
	srv := http.Server{
		Addr:    Port,
		Handler: api.Routes(appConfig),
	}
	log.Printf("Starting application on port %s\n", Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func initAll() (*config.AppConfig, error) {
	return config.New(config.Production)
}
