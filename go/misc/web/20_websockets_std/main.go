package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

const (
	port = ":12345"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWSFeed(ws *websocket.Conn) {
	fmt.Println("New ws connection from client ot orderbook feed:", ws.Request().RemoteAddr)
	for {
		payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())
		n, err := ws.Write([]byte(payload))
		if err != nil || n != len(payload) {
			fmt.Println("Write error:", err)
			delete(s.conns, ws)
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New ws connection from client:", ws.Request().RemoteAddr)
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected. EOF occurred")
				delete(s.conns, ws)
				break
			}
			fmt.Println("Write error:", err)
			continue
		}
		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("ws write error", err)
			}
		}(ws)
	}
}

func main() {
	fmt.Println("Starting ws server on port", port)
	server := NewServer()
	http.HandleFunc("/ws",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(server.handleWS)}
			s.ServeHTTP(w, req)
		})
	http.HandleFunc("/feed",
		func(w http.ResponseWriter, req *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(server.handleWSFeed)}
			s.ServeHTTP(w, req)
		})
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
