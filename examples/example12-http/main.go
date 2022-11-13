package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	path := os.Args[0]
	args = os.Args[1:]
	fmt.Println(path)
	for i, v := range args {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	firstName := flag.String("first_name", "", "First name as a string")
	score := flag.Int("score", 0, "Score as an int")
	flag.Parse()
	fmt.Printf("First name: %s, score: %d\n", *firstName, *score)

}