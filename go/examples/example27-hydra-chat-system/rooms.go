package hydrachat

import (
	"fmt"
	"io"
	"sync"
)

type Room struct {
	Name        string
	RoomChannel chan string
	Clients     map[chan<- string]struct{}
	Quit        chan struct{}
	*sync.RWMutex
}

func CreateRoom(name string) *Room {
	r := &Room{
		Name:        name,
		RoomChannel: make(chan string),
		Clients:     make(map[chan<- string]struct{}),
		Quit:        make(chan struct{}),
		RWMutex:     new(sync.RWMutex),
	}
	r.Run()
	return r
}

func (r *Room) Run() {
	logger.Println("Starting chat room", r.Name)
	go func() {
		for msg := range r.RoomChannel {
			r.broadcastMsg(msg)
		}
	}()
}

func (r *Room) RemoveClient(wc chan<- string) {
	logger.Println("Removing client")
	r.Lock()
	close(wc)
	delete(r.Clients, wc)
	r.Unlock()

	select {
	case <-r.Quit:
		if len(r.Clients) == 0 {
			close(r.RoomChannel)
		}
	default:
	}
}

func (r *Room) AddClient(c io.ReadWriteCloser) {
	r.Lock()
	wc, done := StartClient(r.RoomChannel, c, r.Quit)
	r.Clients[wc] = struct{}{}
	r.Unlock()

	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

func (r *Room) broadcastMsg(msg string) {
	r.RLock()
	defer r.RUnlock()
	fmt.Println("Received message:", msg)
	for wc, _ := range r.Clients {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}
