package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:4]
	c := a[:3]
	d := a[3:]

	fmt.Printf("a: %v, b: %v, c: %v, d: %v\n", a, b, c, d) // a: [1 2 3 4 5 6], b: [3 4], c: [1 2 3], d: [4 5 6]
	fmt.Printf("Capacity of b: %v\n", cap(b))              // 4 till the end of the slice
	fmt.Printf("What b actually sees: %v\n", b[:cap(b)])   // [3 4 5 6]

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[2:4]
	s2[0] = 10
	fmt.Printf("s1: %v, s2: %v\n", s1, s2)

	t1 := []int{1, 2, 3, 4, 5, 6}
	t2 := make([]int, 2)
	ncopied := copy(t2, t1)
	t2[0] = 10
	fmt.Printf("t1: %v, t2: %v, ncopied: %d\n", t1, t2, ncopied)

	p1 := []int{1, 2, 3}
	p1 = append(p1, 4, 5, 6)
	p2 := []int{7, 8, 9}
	p1 = append(p1, p2...)
	fmt.Println(p1) // [1 2 3 4 5 6 7 8 9]

	// remove subslice
	r1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r1 = append(r1[:2], r1[4:]...)
	fmt.Println(r1) // [1 2 5 6]
}
