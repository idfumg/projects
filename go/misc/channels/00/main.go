package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	SIZE = 10
)

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func CalculateValue() chan int {
	intChan := make(chan int)
	go func(){
		number := RandomNumber(SIZE)
		intChan <- number
		close(intChan)
	}()
	return intChan
}

func main() {
	intChan := CalculateValue()
	value := <- intChan
	fmt.Println(value)
}
