package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))
	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}
	wg.Wait()
}
