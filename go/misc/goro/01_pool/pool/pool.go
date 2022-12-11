package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("pool has been closed")

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	closed    bool
	factory   func() (io.Closer, error)
}

func New(fn func() (io.Closer, error), size int32) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool{
		m:         sync.Mutex{},
		resources: make(chan io.Closer, size),
		closed:    false,
		factory:   fn,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		fmt.Println("Release:", "In Queue")
	default:
		fmt.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
