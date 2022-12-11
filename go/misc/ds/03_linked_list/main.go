package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	cmp    func(i, j T) bool
	head   *Node[T]
	length int
}

func NewLinkedList[T any](cmp func(i, j T) bool) *LinkedList[T] {
	return &LinkedList[T]{
		cmp: cmp,
	}
}

func (l *LinkedList[T]) Size() int {
	return l.length
}

func (l *LinkedList[T]) Insert(value T) {
	l.head = &Node[T]{value: value, next: l.head}
	l.length += 1
}

func (l *LinkedList[T]) Print() {
	for n := l.head; n != nil; n = n.next {
		fmt.Printf("%v ", n.value)
	}
	fmt.Println()
}

func (l *LinkedList[T]) Remove(key T) {
	if l.head == nil {
		return
	}
	if l.cmp(l.head.value, key) {
		l.head = l.head.next
		l.length -= 1
		return
	}
	for n := l.head; n.next != nil; n = n.next {
		if l.cmp(n.next.value, key) {
			n.next = n.next.next
			l.length -= 1
			return
		}
	}
}

func main() {
	L := NewLinkedList(func(i, j int) bool {
		return i == j
	})

	L.Insert(48)
	L.Insert(18)
	L.Insert(16)
	L.Insert(11)
	L.Insert(7)
	L.Insert(2)
	L.Print()

	L.Remove(11)
	L.Print()

	L.Remove(48)
	L.Print()

	L.Remove(99)
	L.Print()
	fmt.Println(L.Size())

	L = NewLinkedList(func(i, j int) bool {
		return i == j
	})
	L.Remove(99)
}
