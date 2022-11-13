package server

import (
	"encoding/json"
	"fmt"
	"myapp/service"
	"net/http"
	"strings"
)

const (
	jsonContentType = "application/json"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
	GetLeague() service.League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store: store,
	}
	p.Handler = p.setupRouter()
	return p
}

func (p *PlayerServer) setupRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	return router
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := parsePlayer(r.URL.Path)
	if player == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, ok := p.store.GetPlayerScore(player)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}

func parsePlayer(path string) string {
	const prefix = "/players/"
	return strings.TrimPrefix(path, prefix)
}
