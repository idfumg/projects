package main

import "fmt"

type PriorityQueue[T any] struct {
	cmp   func(i, j T) bool
	array []T
}

func NewPriorityQueue[T any](cmp func(i, j T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		cmp: cmp,
	}
}

func MakePriorityQueue[T any](cmp func(i, j T) bool, array []T) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		cmp:   cmp,
		array: array,
	}
	for i := pq.Size() / 2; i >= 0; i-- {
		pq.heapifyDown(i, len(pq.array))
	}
	return pq
}

func HeapSort[T any](cmp func(i, j T) bool, array []T) {
	pq := MakePriorityQueue(cmp, array)
	for i := pq.Size() - 1; i >= 1; i-- {
		pq.swap(0, i)
		pq.heapifyDown(0, i)
	}
}

func (pq *PriorityQueue[T]) Size() int {
	return len(pq.array)
}

func (pq *PriorityQueue[T]) Push(key T) {
	pq.array = append(pq.array, key)
	pq.heapifyUp(len(pq.array) - 1)
}

func (pq *PriorityQueue[T]) Pop() T {
	if len(pq.array) == 0 {
		panic("Error! Size of PriorityQueue is zero!")
	}
	ans := pq.array[0]
	pq.array[0] = pq.array[len(pq.array)-1]
	pq.array = pq.array[:len(pq.array)-1]
	pq.heapifyDown(0, len(pq.array))
	return ans
}

func (pq *PriorityQueue[T]) Top() T {
	return pq.array[0]
}

func (pq *PriorityQueue[T]) heapifyUp(idx int) {
	for pq.cmp(pq.array[parent(idx)], pq.array[idx]) {
		pq.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

func (pq *PriorityQueue[T]) heapifyDown(idx, n int) {
	l, r := left(idx), right(idx)
	for {
		next := idx
		if l < n && pq.cmp(pq.array[next], pq.array[l]) {
			next = l
		}
		if r < n && pq.cmp(pq.array[next], pq.array[r]) {
			next = r
		}
		if next == idx {
			break
		}
		pq.swap(next, idx)
		l, r, idx = left(next), right(next), next
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.array[i], pq.array[j] = pq.array[j], pq.array[i]
}

func main() {
	pq := NewPriorityQueue(func(i, j int) bool {
		return i > j
	})
	fmt.Println(pq)
	for _, v := range []int{10, 20, 30, 5, 1, 50} {
		pq.Push(v)
	}
	fmt.Println(pq)
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Println()

	pq = MakePriorityQueue(func(i, j int) bool {
		return i > j
	}, []int{10, 20, 30, 5, 1, 50})

	fmt.Println(pq)
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Printf("%d ", pq.Pop())
	fmt.Println()

	arr := []int{10, 20, 30, 5, 1, 50}
	HeapSort(func(i, j int) bool {
		return i < j
	}, arr)
	fmt.Println(arr)

	pq2 := MakePriorityQueue(func(i, j float64) bool {
		return i > j
	}, []float64{10.0, 20.0, 30.0, 5.0, 1.0, 50.0})

	fmt.Println(pq2)
	fmt.Printf("%f ", pq2.Pop())
	fmt.Printf("%f ", pq2.Pop())
	fmt.Printf("%f ", pq2.Pop())
	fmt.Printf("%f ", pq2.Pop())
	fmt.Printf("%f ", pq2.Pop())
	fmt.Printf("%f ", pq2.Pop())
	fmt.Println()

	pq3 := MakePriorityQueue(func(i, j string) bool {
		return i > j
	}, []string{"a", "d", "c", "b"})

	fmt.Println(pq2)
	fmt.Printf("%s ", pq3.Pop())
	fmt.Printf("%s ", pq3.Pop())
	fmt.Printf("%s ", pq3.Pop())
	fmt.Printf("%s ", pq3.Pop())
	fmt.Println()
}
