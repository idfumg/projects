package main

import (
	"fmt"
	"myapp/hlogger"
	"net/http"
)

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	logger.Println("Received http request on root url")
	
	fmt.Fprintf(w, "Welcome to the Hydra Software System")
}

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}
