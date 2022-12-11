package main

import "fmt"

const ARRAY_SIZE = 100

type HashTable[T any] struct {
	hash  func(value T) int
	cmp   func(i, j T) bool
	array [ARRAY_SIZE]*bucket[T]
}

func NewHashTable[T any](hash func(value T) int, cmp func(i, j T) bool) *HashTable[T] {
	result := &HashTable[T]{
		hash: hash,
		cmp:  cmp,
	}
	for i := range result.array {
		result.array[i] = &bucket[T]{}
	}
	return result
}

func (h *HashTable[T]) Insert(key T) {
	index := h.hash(key)
	h.array[index].insert(h.cmp, key)
}

func (h *HashTable[T]) Search(key T) bool {
	index := h.hash(key)
	return h.array[index].search(h.cmp, key)
}

func (h *HashTable[T]) Delete(key T) {
	index := h.hash(key)
	h.array[index].delete(h.cmp, key)
}

type bucket[T any] struct {
	head *bucketNode[T]
}

type bucketNode[T any] struct {
	key  T
	next *bucketNode[T]
}

func (b *bucket[T]) insert(cmp func(i, j T) bool, key T) {
	if b.search(cmp, key) {
		return
	}

	n := &bucketNode[T]{key: key}
	n.next = b.head
	b.head = n
}

func (b *bucket[T]) search(cmp func(i, j T) bool, key T) bool {
	for n := b.head; n != nil; n = n.next {
		if cmp(key, n.key) {
			return true
		}
	}
	return false
}

func (b *bucket[T]) delete(cmp func(i, j T) bool, key T) {
	if b.head == nil {
		return
	}

	n1 := b.head

	if cmp(n1.key, key) {
		b.head = nil
		return
	}

	n2 := b.head.next

	for n1 != nil && n2 != nil {
		if cmp(n2.key, key) {
			n1.next = n2.next
			return
		}
		n1 = n1.next
		n2 = n2.next
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum = (sum*256 + int(v)) % ARRAY_SIZE
	}
	return sum % ARRAY_SIZE
}

func main() {
	t := NewHashTable[string](
		func(value string) int {
			return hash(value)
		},
		func(i, j string) bool {
			return i == j
		})

	t.Insert("RANDY")
	t.Insert("RANDY2")
	t.Insert("RANDY3")
	t.Insert("RANDY3")
	t.Insert("RANDY4")

	fmt.Println(t)

	fmt.Println("RANDY in a hash table is true:", t.Search("RANDY"))
	t.Delete("RANDY")
	fmt.Println("RANDY in a hash table is false:", t.Search("RANDY"))
}
