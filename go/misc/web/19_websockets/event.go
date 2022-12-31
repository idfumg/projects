package main

import "encoding/json"

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

const (
	EventNewMessage = "new_message"
)

type NewMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}
