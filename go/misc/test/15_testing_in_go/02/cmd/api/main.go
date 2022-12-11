package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8090"

func main() {
	app, cancel := NewApp()
	defer cancel()

	fmt.Printf("Starting api on port %s...\n", port)
	err := http.ListenAndServe(port, app)
	if err != nil {
		log.Fatal(err)
	}

}
