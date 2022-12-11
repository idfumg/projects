package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	total int
	count int
	mutex *sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(size int) *Barrier {
	lockToUse := &sync.Mutex{}
	condToUse := sync.NewCond(lockToUse)
	return &Barrier{total: size, count: size, mutex: lockToUse, cond: condToUse}
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count--
	if b.count == 0 {
		b.count = b.total
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.mutex.Unlock()
}

func start(name string, timeout int, barrier *Barrier) {
	for {
		fmt.Printf("%s is working\n", name)
		time.Sleep(time.Duration(timeout) * time.Second)
		fmt.Printf("%s is waiting on the barrier\n", name)
		barrier.Wait()
	}
}

func main() {
	barrier := NewBarrier(2)
	go start("red", 2, barrier)
	go start("blue", 4, barrier)
	time.Sleep(60 * time.Second)
}
