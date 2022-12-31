package main

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	run()
}

func run() {
	setupAPI()
	log.Info("Staring webserver on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func setupAPI() {
	ctx := context.Background()
	manager := NewManager(ctx)
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.Handle("/ws", http.HandlerFunc(manager.serveWS))
	http.Handle("/login", http.HandlerFunc(manager.loginHandler))
}
