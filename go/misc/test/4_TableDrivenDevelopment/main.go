package main

import "math"

func IsArmstrong(n int) bool {
	a := n % 10
	b := n / 10 % 10;
	c := n / 100;
	return n == int(math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3))
}

func main() {

}