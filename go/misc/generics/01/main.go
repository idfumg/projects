package main

import "fmt"

func print1(values ...interface{}) { // doesn't support static type checking
	fmt.Println(values...)
}

func print2[T interface{}](values ...T) { // infer a type only once and use it
	fmt.Println(values)
}

func print3[T any](values ...T) { // replacement for the interface{}
	fmt.Println(values)
}

type MyType interface {
	int | string
}

func print4[T MyType](values ...T) {
	fmt.Println(values)
}

func main() {
	print1(1, 2, 3)       // 1 2 3
	print1("1", "2", "3") // 1 2 3
	print1(1, 2.0, "3")   // 1 2 3

	print2(1, 2, 3)       // [1 2 3]
	print2("1", "2", "3") // [1 2 3]
	// print2(1, 2.0, "3") // default type float64 of 2.0 does not match inferred type int for T

	print3(1, 2, 3)       // [1 2 3]
	print3("1", "2", "3") // [1 2 3]
	// print3(1, 2.0, "3")   // default type float64 of 2.0 does not match inferred type int for T

	print4(1, 2, 3)       // [1 2 3]
	print4("1", "2", "3") // [1 2 3]
	// print4(1, 2.0, "3")   // default type float64 of 2.0 does not match inferred type int for T
}
