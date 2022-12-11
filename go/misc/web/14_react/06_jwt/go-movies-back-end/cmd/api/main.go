package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	cfg := NewCfg()
	log.Println("Config parsed")

	conn, cancel, err := connectToDB(cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	log.Println("Postgres connected")

	app := NewApp(conn, cfg)
	log.Println("App created")

	log.Println("Starting api on port", port)
	err = http.ListenAndServe(port, app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
