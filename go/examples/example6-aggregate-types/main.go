package main

import "fmt"

type Car struct {
	NumberOfTires int
	Luxury bool
	BucketSeats bool
	Make string
	Model string
	Year int
}

func main() {
	var stringList [3]string;
	stringList[0] = "cat"
	stringList[1] = "dog"
	stringList[2] = "fish"
	fmt.Println(stringList)

	var car1 Car
	car1.NumberOfTires = 4
	car1.Luxury = false
	car1.BucketSeats = false
	car1.Make = "XC()"
	car1.Model = "Volkswagen"
	car1.Year = 2022

	car2 := Car{
		NumberOfTires: 4,
		Luxury: false,
		BucketSeats: false,
		Make: "XC90",
		Model: "Volkswagen",
		// Year: 2022,
	}
	fmt.Println(car2)
}