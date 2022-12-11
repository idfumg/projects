package main

import (
	"fmt"
	"io"
	"log"
	"myapp/pool"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 10
	maxResources  = 2
)

var idCount int32 = 0

type dbConnection struct {
	id int32
}

func (d *dbConnection) Close() error {
	fmt.Println("Closing database connection")
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCount, 1)
	return &dbConnection{id}, nil
}

func main() {
	p, err := pool.New(createConnection, maxResources)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func(q int) {
			handle(p, q)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Shutdown the program")
	p.Close()
}

func handle(p *pool.Pool, query int) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(500) * time.Millisecond)
	fmt.Printf("QUI[%d], CID[%d]\n", query, conn.(*dbConnection).id)
}
