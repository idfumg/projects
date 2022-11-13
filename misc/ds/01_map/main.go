package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m == nil) // true
	m = map[string]string{}
	fmt.Println(m == nil) // false
	fmt.Println(m) // map[]
	fmt.Println(len(m)) // 0
	m = make(map[string]string, 5) // preallocate memory for 5 elements
	fmt.Println(len(m)) // 0
	m = map[string]string{"a": "Foo", "b": "Bar"}
	fmt.Println(m) // map[a:Foo b:Bar]
	fmt.Println(len(m)) // 2
	m["a"] = "SuperFoo"
	fmt.Println(m) // map[a:SuperFoo b:Bar]
	fmt.Println(m["a"]) // SuperFoo
	delete(m, "a")
	fmt.Println(m) // map[b:Bar]
	m["a"] = "NewFoo"

	for name, value := range m {
		fmt.Printf("name: %s, value: %s\n", name, value)
	}

	name, ok := m["a"]
	if ok {
		fmt.Println("Name:", name) // Name: NewFoo
	}
}