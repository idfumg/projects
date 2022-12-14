package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint hit: HomePage")
}

func Setup() {
	http.HandleFunc("/", HomePage)
}

func main() {
	Setup()
	log.Fatal(http.ListenAndServe(":8080", nil))
}