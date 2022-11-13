package main

import (
	"log"
	"myapp/server"
	"myapp/service"
	"net/http"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
	}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := i.store[name]
	return score, ok
}

func (i *InMemoryPlayerStore) GetLeague() 	service.League {
	var ans []service.Player
	for name, wins := range i.store {
		ans = append(ans, service.Player{Name: name, Wins: wins})
	}
	return ans
}

func main() {
	server := server.NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":8080", server))
}
