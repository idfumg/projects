package main

import (
	"fmt"
	"math"
)

// Non-generic
func sumOfInt(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans += num
	}
	return ans
}

type Number interface {
	int64 | float64
}

// Generic
func sumOf[T Number](nums []T) T {
	var ans T
	for _, num := range nums {
		ans += num
	}
	return ans
}

func Has[T comparable](arr []T, key T) bool {
	for _, value := range arr {
		if value == key {
			return true
		}
	}
	return false
}

func NewEmptyList[T any]() []T {
	return make([]T, 0)
}

func PrintThings[A, B any, C ~int | ~uint](a1, a2 A, b B, c C) { // ~int == the current type or the underlying type is an `int`
	fmt.Println(a1, a2, b, c)
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values ...T) Set[T] {
	ans := make(Set[T], len(values))
	for _, value := range values {
		ans[value] = struct{}{}
	}
	return ans
}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}

func Has2[T any](list []T, value T, equal func(a, b T) bool) bool {
	for _, v := range list {
		if equal(v, value) {
			return true
		}
	}
	return false
}

type Equalizer[T any] interface {
	Equal(rhs T) bool
}

type Id int

func (id Id) Equal(rhs Id) bool {
	return id == rhs
}

func Has3[T Equalizer[T]](list []T, value T) bool {
	for _, v := range list {
		if value.Equal(v) {
			return true
		}
	}
	return false
}

type CombinedType interface {
	~int | uint
	Number
	IsValid() bool
}

type Currency interface {
	~int | ~int64
	ISO4127Code() string
	Decimal() int
}

func PrintBalance[T Currency](b T) {
	balance := float64(b) / math.Pow10(b.Decimal())
	fmt.Printf("%.*f %s\n", b.Decimal(), balance, b.ISO4127Code())
}

type NZD int64

func (b NZD) ISO4127Code() string {
	return "NZD"
}

func (b NZD) Decimal() int {
	return 2
}

type Num interface {
	int8 | int16 | int32 | int64 | int
}

func BubbleSort[T Num](arr []T) []T {
	n, swapped := len(arr), true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}

func main() {
	fmt.Println("sumOfInt:", sumOfInt([]int{1, 2, 3}))
	fmt.Println("sumOf:", sumOf([]int64{1, 2, 3}))
	fmt.Println("Has 5:", Has([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println("Has 6:", Has([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println("NewEmptyList:", NewEmptyList[int]()) // providing the type explicitly dua to inability to defer it
	PrintThings(1, 2, "3", 4)
	// PrintThings(1, 2.0, 3, 4) // it's wrong because of different types of a1 & a2
	// PrintThings(1.0, 2.0, 3, 4) // now the type are the same, so it's okay

	intSet := NewSet(1, 2, 3, 4, 5)
	fmt.Println("SetHas 4:", intSet.Has(5))
	stringSet := NewSet("a", "b", "c")
	fmt.Println("SetHas \"b\":", stringSet.Has("b"))

	intEqual := func(a, b int) bool {
		return a == b
	}
	fmt.Println("Has2:", Has2([]int{1, 2, 3, 4, 5}, 4, intEqual))

	fmt.Println("Has3:", Has3([]Id{1, 2, 3, 4, 5}, 4))

	PrintBalance(NZD(250))

	fmt.Println(BubbleSort([]int{1, 4, 3, 2, 6, 5}))
}
