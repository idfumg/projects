package main

import "fmt"

func main() {
	c := rune('ф') // int32
	fmt.Println(c) // 1092
	fmt.Printf("%T\n", c) // int32
	fmt.Println()

	c = 'ф'
	fmt.Println(c) // 1092
	fmt.Printf("%T\n", c) // int32
	fmt.Println()

	s := "Hello привет"
	fmt.Println(s) // Hello привет
	fmt.Println(len(s)) // 18 (12 symbols, but 18 runes, because of the not a single byte chars)
	fmt.Println(len("ф")) // 2
	fmt.Println()

	b := byte(65)
	fmt.Printf("%c\n", b) // A
	fmt.Printf("%T\n", b) // uint8
	fmt.Println()
}