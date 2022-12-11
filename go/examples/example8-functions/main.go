package main

import "fmt"

type Animal struct {
	Name string
	Sound string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func main() {
	total := sumMany(1, 2, 3, 4)
	fmt.Println(total)

	dog := Animal {
		Name: "dog",
		Sound: "bark",
		NumberOfLegs: 4,
	}
	dog.Says()
}

func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	return total
}