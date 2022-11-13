package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, slices, maps, functions, channels)
// interface type

func main() {
	x := 10
	fmt.Println(x) // 10

	p := &x;
	fmt.Println(p) // [addr]

	*p = 20;
	fmt.Println(x) // 20

	changePointerValue(p);
	fmt.Println(x) // 25

	var animals []string;
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	animals = append(animals, "horse")
	fmt.Println(animals)
	for i, value := range animals {
		fmt.Print(i, ", ", value, "; ")
	}
	fmt.Println("\nZero element is:", animals[0]) // dog
	fmt.Println("First two positions:", animals[0:2]) // dog cat
	fmt.Printf("Animals has %d elements\n", len(animals)) // 4
	fmt.Println("String are sorted:", sort.StringsAreSorted(animals)) // false
	sort.Strings(animals)
	fmt.Println("String are sorted:", sort.StringsAreSorted(animals)) // true
	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals) // cat horse fish

	intMap := make(map[string]int) // not sorted
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5
	for key, value := range intMap {
		fmt.Print(key, " ", value, " ")
	}
	fmt.Println()
	delete(intMap, "five")
	delete(intMap, "five")
	for key, value := range intMap {
		fmt.Print(key, " ", value, " ")
	}
	fmt.Println()
	_, ok := intMap["five"]
	if (ok) {
		fmt.Printf("five exists in a map")
	} else {
		fmt.Printf("five does not exist in a map")
	}
}

func changePointerValue(num *int) {
	*num = 25;
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a) - 1]
	a[len(a) - 1] = ""
	a = a[0:len(a) - 1]
	return a
}