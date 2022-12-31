package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
)

type Manager struct {
	clients  ClientList
	mu       sync.RWMutex
	handlers map[string]EventHandler
	otps     RetensionMap
}

func NewManager(ctx context.Context) *Manager {
	return &Manager{
		clients:  ClientList{},
		mu:       sync.RWMutex{},
		handlers: configureEventHandlers(),
		otps:     NewRetentionMap(ctx, 5*time.Second),
	}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	otp := r.URL.Query().Get("otp")
	if otp == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !m.otps.VerifyOTP(otp) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}
	client := NewClient(conn, m)
	m.addClient(client)
	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) loginHandler(w http.ResponseWriter, r *http.Request) {
	type userLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req userLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Errorf("Error on login: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Username == "test" && req.Password == "test" {
		type response struct {
			OTP string `json:"otp"`
		}
		otp := m.otps.NewOTP()
		resp := response{
			OTP: otp.Key,
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Error(err)
			return
		}
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func configureEventHandlers() map[string]EventHandler {
	m := map[string]EventHandler{}
	m[EventNewMessage] = newMessageHandler
	return m
}

func newMessageHandler(event Event, c *Client) error {
	log.Info("NewMessageHandler: ", event)
	for client := range c.manager.clients {
		client.dispatchCh <- event
	}
	return nil
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("there is no such event type: %v", event.Type)
	}
}

func (m *Manager) addClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.clients[client]; ok {
		client.conn.Close()
		delete(m.clients, client)
	}
}

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	switch origin {
	case "http://localhost:8080":
		return true
	default:
		return false
	}
}
