package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
)

// go get -u github.com/inancgumus/prettyslice


func main() {
	books := [...]string{
		"Dracula",
		"1984",
		"Island",
	}
	fmt.Printf("books: %#v\n", books)

	var games []string
	fmt.Printf("games: %#v\n", games)
	fmt.Printf("games: %T\n", games)
	fmt.Printf("games len: %d\n", len(games))
	fmt.Printf("games is nil: %t\n", games == nil)

	var arr []int
	arr = append(arr, 1)
	arr = append(arr, 2, 3)
	arr2 := append(arr, 4)
	fmt.Println(arr, arr2)
	arr[0] = 9
	fmt.Println(arr, arr2)

	var todo []string
	todo = append(todo, "sing")
	todo = append(todo, "run", "code", "play")
	s.Show("todo", todo)
	todo = append(todo, []string{"walk", "read"}...)
	s.Show("todo", todo)
	todo = append(todo, "fly")
	s.Show("todo", todo)

	// create a brand new slice by creating new slice object and appending to it
	newTodo := append([]string(nil), todo...)
	s.Show("NewTodo", newTodo)
}
