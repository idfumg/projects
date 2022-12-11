package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MapCounter struct {
	m map[int]int
	sync.RWMutex
}

func runWriters(mc *MapCounter, n int) {
	for i := 0; i < n; i += 1 {
		mc.Lock()
		mc.m[i] = i * 10
		mc.Unlock()
		time.Sleep(time.Second * 1)
	}
}

func runReaders(mc *MapCounter, n int) {
	for {
		mc.RLock()
		value := mc.m[rand.Intn(n)]
		mc.RUnlock()
		fmt.Println(value)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mc := MapCounter{m: make(map[int]int)}
	go runWriters(&mc, 10)
	go runReaders(&mc, 10)
	go runReaders(&mc, 10)
	time.Sleep(time.Second * 10)
}