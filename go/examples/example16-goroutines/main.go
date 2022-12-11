package main

import (
	"fmt"
	"time"
)

func waitAndSay(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func main() {
	go waitAndSay("World")
	go func(s string) {
		waitAndSay(s)
	}("World")
	word := "World"
	go func(){
		waitAndSay(word) // closure inside a goroutine
	}()
	fmt.Println("Hello")
	time.Sleep(4 * time.Second)
}
