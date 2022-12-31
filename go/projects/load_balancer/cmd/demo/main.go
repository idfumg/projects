package main

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	port = flag.Int("port", 8080, "port to listen to")
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello from: %d", *port)))
}

func main() {
	flag.Parse()
	log.Infof("Starting server on port: %d", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), http.HandlerFunc(Home))
}
