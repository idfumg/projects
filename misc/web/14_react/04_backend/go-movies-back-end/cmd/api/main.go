package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	app := NewApp()
	app.Domain = "example.com"

	log.Println("Starting api on port", port)
	err := http.ListenAndServe(port, app.routes())
	if err != nil {
		log.Fatal(err)
	}
}