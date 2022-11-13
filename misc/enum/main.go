package main

import "fmt"

type BodyPart int

const (
	Head BodyPart = iota
	Shoulder
	Knee
	Toe
)

type ErrorType int

const (
	NotFound ErrorType = iota
	NotImplemented
)

func (b BodyPart) String() string {
	return []string{"Head", "Shoulder", "Knee", "Toe"}[b]
}

func main() {
	toe := Toe
	fmt.Printf("%T %s\n", toe, toe)

	err := NotImplemented
	fmt.Printf("%T %d\n", err, err)

	// toe = err // conversion error
}