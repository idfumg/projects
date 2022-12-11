package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}

	m.Store(0, 1)
	m.Store(1, 1)
	m.Store(1, 2)

	m.Range(func(key, value any) bool {
		fmt.Print(key, ":", value, ", ")
		return true
	})
	fmt.Println()

	value, ok := m.Load(0)
	fmt.Println(value, ok) // 0, true

	m.Delete(0)
	value, ok = m.Load(0)
	fmt.Println(value, ok) // <nil>, false

	value, ok = m.LoadAndDelete(1)
	fmt.Println(value, ok) // 0, true	
}