package main

import (
	"flag"
	"myapp/pkg/config"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	port       = flag.Int("port", 8080, "port to listen to")
	configPath = flag.String("config", "config.yml", "yaml config file path")
)

func main() {
	flag.Parse()

	f, err := os.Open(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	config, err := config.New(f)
	if err != nil {
		log.Fatal(err)
	}

	app := NewApp(config)

	log.Infof("Start listening on port: %d", *port)
	err = app.Run(*port)
	if err != nil {
		log.Fatal(err)
	}
}
