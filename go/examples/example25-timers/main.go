package main

import (
	"fmt"
	"time"
)

func SlowCounter(n int) {
	i := 0
	d := time.Duration(n) * time.Second
	for {
		// t := time.NewTimer(d)
		// <-t.C
		<-time.After(d)
		fmt.Println(i)
		i += 1
	}
}

func main() {
	go SlowCounter(2)
	time.Sleep(time.Second * 10)
}