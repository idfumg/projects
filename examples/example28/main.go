package main

import (
	"fmt"
)

func main() {
	city := "Lyon"
	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokio":
		fmt.Println("Japan")
	}
}