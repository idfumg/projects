package main

import (
	"fmt"
	"time"
)

func SlowCounter(n int, newDelayChannel chan int, stopChannel chan bool) {
	i := 0
	delay := time.Duration(n) * time.Second
Loop:
	for {
		select {
		case <-time.After(delay):
			fmt.Println(i)
			i += 1
		case n = <-newDelayChannel:
			fmt.Println("Delay changed to:", n)
			delay = time.Duration(n) * time.Second
		case <-stopChannel:
			fmt.Println("Timer stopped")
			break Loop
		}
	}
}

func main() {
	newDelayChannel := make(chan int)
	stopChannel := make(chan bool)

	go SlowCounter(1, newDelayChannel, stopChannel)
	time.Sleep(time.Second * 5)
	newDelayChannel <- 2
	time.Sleep(time.Second * 5)
	stopChannel <- true
	time.Sleep(time.Second * 1)
}