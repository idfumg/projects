package main

import (
	"log"
	"myapp/work"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	log.Println(n.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(10)

	var wg sync.WaitGroup
	wg.Add(5 * len(names))
	for i := 0; i < 5; i++ {
		for _, name := range names {
			go func(n string){
				p.Add(&namePrinter{n})
				wg.Done()
			}(name)
		}
	}
	wg.Wait()
	p.Shutdown()
}