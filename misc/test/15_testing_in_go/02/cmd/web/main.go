package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	app, cancel := NewApp()
	defer cancel()
	// mux := app.routes()

	fmt.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(port, app)
	if err != nil {
		log.Fatal(err)
	}
}
