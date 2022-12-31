package main

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	readMessageLimit = 512
	pongWait         = 10 * time.Second
	pingInterval     = (pongWait * 9) / 10
)

type ClientList map[*Client]bool

type Client struct {
	conn       *websocket.Conn
	manager    *Manager
	dispatchCh chan Event
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		conn:       conn,
		manager:    manager,
		dispatchCh: make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()
	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Fatal(err)
		return
	}
	c.conn.SetPongHandler(c.pongHandler)
	c.conn.SetReadLimit(readMessageLimit)
	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("Error reading message: %v", err)
			}
			break
		}
		var requestEvent Event
		if err := json.Unmarshal(payload, &requestEvent); err != nil {
			log.Errorf("Error marshaling the event: %v", err)
			break
		}
		if err := c.manager.routeEvent(requestEvent, c); err != nil {
			log.Errorf("Error routing the event: %v", err)
			break
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()
	ticker := time.NewTicker(pingInterval)
	for {
		select {
		case event, ok := <-c.dispatchCh:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Errorf("Failed to close a connection: %v", err)
				}
				return
			}
			data, err := json.Marshal(event)
			if err != nil {
				log.Errorf("Error marshaling the event: %v", err)
				continue
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Errorf("Failed to send a message: %v", err)
				continue
			}

		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Errorf("Ping error occured: %v", err)
				return
			}
		}
	}
}

func (c *Client) pongHandler(appData string) error {
	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}
