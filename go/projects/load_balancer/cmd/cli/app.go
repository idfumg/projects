package main

import (
	"fmt"
	"myapp/pkg/config"
	"myapp/pkg/server"
	"myapp/pkg/strategy"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Config    *config.Config
	ServerMap map[string]*server.Service
}

func NewApp(config *config.Config) *App {
	serverMap := map[string]*server.Service{}

	for _, service := range config.Services {
		replicas := []*server.Replica{}
		for _, replica := range service.Replicas {
			u, err := url.Parse(replica)
			if err != nil {
				log.Fatal(err)
			}
			proxy := httputil.NewSingleHostReverseProxy(u)
			replicas = append(replicas, &server.Replica{
				Url:   u,
				Proxy: proxy,
			})
		}
		serverMap[service.Matcher] = &server.Service{
			Name:     service.Name,
			Replicas: replicas,
			Strategy: strategy.NewStrategy(service.Strategy),
		}
	}

	return &App{
		Config:    config,
		ServerMap: serverMap,
	}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Infof("Received new request: %s\n", r.Host)
	serverList, err := app.findServerList(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	nxt, err := serverList.Next()
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	nxt.Forward(w, r)
}

func (app *App) findServerList(url string) (*server.Service, error) {
	for k, v := range app.ServerMap {
		if strings.HasPrefix(url, k) {
			log.Infof("Found services with a prefix: '%s'\n", k)
			return v, nil
		}
	}
	return nil, fmt.Errorf("could not find server list for that url")
}

func (app *App) Run(port int) error {
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: app,
	}

	return s.ListenAndServe()
}
