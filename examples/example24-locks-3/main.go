package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i += 1 {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Println("Work done for:", i)
		}(i)
	}

	wg.Wait()
}