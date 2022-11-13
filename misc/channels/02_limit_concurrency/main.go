package main

import (
	"fmt"
	"time"
)

const (
	concurrencyLevel = 3
)

var (
	requestIDs = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

func main() {
	q := make(chan bool, concurrencyLevel)

	for _, id := range requestIDs {
		q <- true
		go func(id int) {
			// defer func() { <-q }()
			makeRequest(id)
			<-q
		}(id)
	}

	for i := 0; i < concurrencyLevel; i++ {
		q <- true
	}
}

func makeRequest(id int) {
	time.Sleep(time.Second * 2)
	fmt.Println(id)
}
