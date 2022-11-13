package main

import "fmt"

const (
	Max = 9
)

func main() {
	fmt.Printf("%5s", "X")
	for i := 1; i <= Max; i += 1 {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 1; i <= Max; i += 1 {
		fmt.Printf("%5d", i)
		for j := 1; j <= Max; j += 1 {
			fmt.Printf("%5d", j * i)
		}
		fmt.Println()
	}
}
