package main

import (
	"fmt"
	"time"
)

func Print(s string) {
	for x := 0; x < 5; x += 1 {
		time.Sleep(time.Second)
		fmt.Printf("%v: %v\n", s, x)
	}
}

func main() {
	// go Print("Hello")
	// go Print("World")
	// go func(){Print("idfumg")}()

	fmt.Scanln()

	for i := 0; i < 5; i += 1 {
		go func(){
			fmt.Println(i)  // i will be captured and the same address in memory will be used by
							// each goroutine
		}()
	}

	fmt.Scanln()
}