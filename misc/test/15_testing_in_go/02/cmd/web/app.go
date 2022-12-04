package main

import (
	"encoding/gob"
	"log"
	"myapp/pkg/data"
	"myapp/pkg/repository"
	"myapp/pkg/repository/dbrepo"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	http.Handler
	Session *scs.SessionManager
	DB      repository.DatabaseRepo
}

func NewApp() (*application, func()) {
	cmd := NewArgs()

	conn, err := connectToDB(cmd.DSN)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Postgres connected")

	session := getSession()
	if session == nil {
		log.Fatal(session)
	}
	log.Println("Session attached")

	gob.Register(data.User{})

	cancel := func() {
		conn.Close()
	}

	app := &application{
		Session: session,
		DB:      &dbrepo.PostgresDBRepo{DB: conn},
	}

	app.Handler = app.routes()

	return app, cancel
}
