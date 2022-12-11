package main

import (
	"fmt"
	"time"
)

func receive(ch <-chan int) {
	for {
		value := <-ch
		fmt.Println("value received:", value)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan int, 2)
	go receive(ch)
	for i := 0; i < 10; i += 1 {
		select {
		case ch <- i:
			fmt.Println("value sent")
		// default:
		// 	fmt.Println("value ommited")
		}
	}
	time.Sleep(time.Second * 10)
}
