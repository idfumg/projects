package main

import "fmt"

type Substractable interface {
	~int | int32 | int64 | float32 | float64
}

type MyGenericArray[T Substractable] []T
type MyAnyArray[T any] []T

func main() {
	ints := MyGenericArray[int]{}
	floats := MyGenericArray[float32]{}
	ints = append(ints, []int{1, 2, 3}...)
	floats = append(floats, []float32{1.1, 2.2, 3.3}...)
	fmt.Println(ints)   // [1 2 3]
	fmt.Println(floats) // [1.1 2.2 3.3]

	anys := MyAnyArray[any]{}
	anys = append(anys, ints, floats)
	fmt.Println(anys)

	p := &Person[int]{Name: "John"}
	c := &Car[int]{Name: "Ferrari"}
	Move(p, 10)
	Move(c, 20)
	Move[*Person[int32], int32](&Person[int32]{Name: "Mary"}, int32(11))
}

type Movable[S Substractable] interface {
	Move(S)
}

func Move[M Movable[S], S Substractable](param M, meters S) {
	param.Move(meters)
}

type Person[S Substractable] struct {
	Name string
}

func (p *Person[S]) Move(meters S) {
	fmt.Printf("%s moved %v meters\n", p.Name, meters)
}

type Car[S Substractable] struct {
	Name string
}

func (c *Car[S]) Move(meters S) {
	fmt.Printf("%s moved %+v meters\n", c.Name, meters)
}
