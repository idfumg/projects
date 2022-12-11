package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	i int
	sync.Mutex
}

func (counter *SafeCounter) Increment() {
	counter.Lock()
	defer counter.Unlock()
	counter.i += 1
}

func (counter *SafeCounter) Decrement() {
	counter.Lock()
	defer counter.Unlock()
	counter.i -= 1
}

func (counter *SafeCounter) GetValue() int {
	counter.Lock()
	defer counter.Unlock()
	return counter.i
}

func main() {
	counter := new(SafeCounter)
	for i := 0; i < 200; i += 1 {
		go counter.Increment()
		go counter.Decrement()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(counter.GetValue())
}