package main

import (
	"fmt"
	"time"
)

func TickCounter(ticker *time.Ticker) chan bool {
	stopChannel := make(chan bool)
	go func(){
		i := 0
Loop:
		for {
			select {
			case t := <-ticker.C:
				i += 1
				fmt.Println("Count", i, "at", t)
			case <- stopChannel:
				fmt.Println("Stop TickCounter goroutine")
				break Loop
			}
		}
	}()
	return stopChannel
}

func main() {
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	stopChannel := TickCounter(ticker)
	time.Sleep(time.Second * time.Duration(5))
	ticker.Stop()
	stopChannel <- true
	time.Sleep(time.Second * time.Duration(2))
}