package main

import "fmt"

type Node[T any] struct {
	Key   T
	Left  *Node[T]
	Right *Node[T]
}

type OrderedSet[T any] struct {
	cmp  func(i, j T) int
	root *Node[T]
}

func NewOrderedSet[T any](cmp func(i, j T) int) *OrderedSet[T] {
	return &OrderedSet[T]{
		cmp: cmp,
	}
}

func (s *OrderedSet[T]) Insert(key T) {
	s.root = s.insert(s.root, key)
}

func (s *OrderedSet[T]) insert(root *Node[T], key T) *Node[T] {
	if root == nil {
		return &Node[T]{Key: key}
	} else if s.cmp(key, root.Key) < 0 {
		root.Left = s.insert(root.Left, key)
	} else if s.cmp(key, root.Key) > 0 {
		root.Right = s.insert(root.Right, key)
	}
	return root
}

func (s *OrderedSet[T]) Inorder() {
	s.inorder(s.root)
}

func (s *OrderedSet[T]) inorder(root *Node[T]) {
	if root == nil {
		return
	}
	s.inorder(root.Left)
	fmt.Printf("%v ", root.Key)
	s.inorder(root.Right)
}

func (s *OrderedSet[T]) Search(key T) bool {
	return s.search(s.root, key)
}

func (s *OrderedSet[T]) search(root *Node[T], key T) bool {
	if root == nil {
		return false
	}
	if s.cmp(key, root.Key) < 0 {
		return s.search(root.Left, key)
	} else if s.cmp(key, root.Key) > 0 {
		return s.search(root.Right, key)
	} else {
		return true
	}
}

func main() {
	set := NewOrderedSet(func(i, j int) int{
		return i - j
	})
	fmt.Printf("%+v\n", set)

	set.Insert(90)
	set.Insert(92)
	set.Insert(95)
	fmt.Printf("%+v\n", set)

	set.Insert(200)
	set.Insert(201)
	set.Insert(300)
	set.Insert(250)
	fmt.Printf("%+v\n", set)

	set.Inorder()
	fmt.Println()

	fmt.Println("Search 201 =", set.Search(201))
	fmt.Println("Search 203 =", set.Search(203))
}
