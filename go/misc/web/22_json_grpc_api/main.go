package main

import (
	"myapp/api"
	"myapp/config"
	"myapp/service"
)

func main() {
	cfg := config.NewConfig()

	svc :=
		service.NewWithLogsService(
			service.NewWithMetricsService(
				service.NewPriceFetcher(cfg), cfg), cfg)

	server := api.NewServer(":8080", svc)
	server.Run()
}
