package main

import (
	"myapp/pkg/repository/dbrepo"
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "827ccb0eea8a706c4c34a16891f84e7b"
	os.Exit(m.Run())
}
