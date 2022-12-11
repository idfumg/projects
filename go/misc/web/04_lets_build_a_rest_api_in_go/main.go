package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Coaster struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID           string `json:"id"`
	InPark       string `json:"in_park"`
	Height       int    `json:"height"`
}

type Store struct {
	mutex    sync.RWMutex
	Coasters []*Coaster
}

func NewStore() *Store {
	return &Store{
		Coasters: []*Coaster{
			{
				Name:         "Fury 325",
				Manufacturer: "B+M",
				ID:           "id1",
				InPark:       "Carowinds",
				Height:       99,
			},
		},
	}
}

func (s *Store) GetAll() []*Coaster {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.Coasters
}

func (s *Store) FindIdx(c *Coaster) int {
	for i, coaster := range s.Coasters {
		if coaster.ID == c.ID {
			return i
		}
	}
	return -1
}

func CoastersGet(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonBytes, err := json.Marshal(s.Coasters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}

func CoastersPost(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
			return
		}

		var coaster Coaster
		err = json.Unmarshal(body, &coaster)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		coaster.ID = fmt.Sprintf("%d", time.Now().UnixNano())

		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.Coasters = append(s.Coasters, &coaster)
	}
}

func CoastersPut(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		ct := r.Header.Get("Content-Type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
			return
		}

		var c Coaster
		err = json.Unmarshal(body, &c)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if idx := s.FindIdx(&c); idx != -1 {
			s.Coasters[idx] = &c
		}
	}
}

func CoastersHandler(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			CoastersGet(s)(w, r)
			return
		case "POST":
			CoastersPost(s)(w, r)
			return
		case "PUT":
			CoastersPut(s)(w, r)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("method not allowed"))
			return
		}
	}
}

type Admin struct {
	passw string
}

func NewAdmin() *Admin {
	passw := os.Getenv("ADMIN_PASSWORD")
	if passw == "" {
		panic("required environment ADMIN_PASSWORD is not set")
	}
	return &Admin{passw: passw}
}

func AdminHandler(admin *Admin) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, passw, ok := r.BasicAuth()
		if !ok || user != "admin" || passw != admin.passw {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - unauthorized"))
			return
		}
		w.Write([]byte("<html><h1>Super secret admin portal</h1></html>"))
	}
}

func main() {
	store := NewStore()
	admin := NewAdmin()
	http.HandleFunc("/coasters", CoastersHandler(store))
	http.HandleFunc("/admin", AdminHandler(admin))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
