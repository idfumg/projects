package main

import (
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting web server on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}