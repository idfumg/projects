package main

import "fmt"

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) Dequeue() T {
	ans := q.items[0]
	q.items = q.items[1:]
	return ans
}

func main() {
	q := Queue[int]{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q)

	q.Dequeue()
	fmt.Println(q)

	q.Dequeue()
	fmt.Println(q)
}
