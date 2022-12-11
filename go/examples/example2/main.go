package main

import (
	"fmt"
	"myapp/packageone"
)

func main() {
	newString := packageone.PublicVar
	fmt.Println(newString)
	packageone.Exported()
	
}