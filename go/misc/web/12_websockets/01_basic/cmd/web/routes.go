package main

import (
	"net/http"

	"github.com/bmizerany/pat"

	"myapp/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
