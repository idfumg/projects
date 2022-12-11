package main

import (
	"log"
	"myapp/pkg/repository"
	"myapp/pkg/repository/dbrepo"
	"net/http"
)

type application struct {
	http.Handler
	DB        repository.DatabaseRepo
	JWTSecret string
	Domain    string
}

func NewApp() (*application, func()) {
	cmd := NewArgs()
	log.Println("Command line arguments parsed")

	conn, err := connectToDB(cmd.DSN)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Postgres connected")

	cancel := func() {
		conn.Close()
	}

	app := &application{
		DB:        &dbrepo.PostgresDBRepo{DB: conn},
		JWTSecret: cmd.JWTSecret,
		Domain:    cmd.Domain,
	}

	app.Handler = app.routes()
	log.Println("Routes created")

	return app, cancel
}
