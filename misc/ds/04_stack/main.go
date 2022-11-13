package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() T {
	ans := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return ans
}

func main() {
	s := Stack[int]{}

	s.Push(17)
	s.Push(10)
	s.Push(7)
	s.Push(5)
	fmt.Println(s)
	value1 := s.Pop()
	fmt.Println(value1)
	fmt.Println(s)
}
