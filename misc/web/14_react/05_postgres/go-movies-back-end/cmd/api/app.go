package main

import (
	"database/sql"
	"myapp/internal/repository"
	"myapp/internal/repository/dbrepo"
)

type application struct {
	Domain string
	DB     repository.DatabaseRepo
}

func NewApp(conn *sql.DB) *application {
	return &application{
		Domain: "example.com",
		DB:     &dbrepo.PostgresDBRepo{DB: conn},
	}
}
