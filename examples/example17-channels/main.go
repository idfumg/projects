package main

import "fmt"

func waitAndSay(ch chan bool, s string) {
	if b := <-ch; b { // wait a value in the channel
		fmt.Println(s)
	}
	ch <- true
}

func SayHelloMultipleTImes(ch chan string, n int) {
	for n > 0 {
		ch <- "hello"
		n -= 1
	}
	close(ch)
}

func main() {
	ch := make(chan bool)
	go waitAndSay(ch, "World!")
	fmt.Println("Hello")
	ch <- true // sedn to the channel
	<-ch       // wait a value from the channel

	// Buffered Channels
	ch2 := make(chan string, 3)
	ch2 <- "hello"
	ch2 <- "hello2"
	ch2 <- "hello3"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// wait values from a channel with the range keyword
	ch3 := make(chan string)
	go SayHelloMultipleTImes(ch3, 5)
	for s := range ch3 {
		fmt.Println(s)
	}
	v, ok := <-ch3
	fmt.Println("Is channel closed?", !ok, "value:", v)
}