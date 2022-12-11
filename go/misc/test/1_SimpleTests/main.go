package main

import "fmt"

type Number interface {
	float64
}

func divide[T Number](a, b T) (T, error) {
	if (b == 0) {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

func main() {

}
