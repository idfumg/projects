package strategy

import (
	"fmt"
	"sync"

	"myapp/pkg/server"
)

type Strategy interface {
	Next(servers []*server.Replica) (*server.Replica, error)
}

func NewStrategy(name string) Strategy {
	switch name {
	case "RoundRobin":
		return &RoundRobin{}
	}
	return &RoundRobin{}
}

type RoundRobin struct {
	Current int32
	m       sync.Mutex
}

func (r *RoundRobin) Next(servers []*server.Replica) (*server.Replica, error) {
	if len(servers) == 0 {
		return nil, fmt.Errorf("you provided zero services")
	}
	r.m.Lock()
	idx := r.Current
	r.Current = (r.Current + 1) % int32(len(servers))
	r.m.Unlock()
	return servers[idx], nil
}
