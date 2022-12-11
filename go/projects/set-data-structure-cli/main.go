package main

import "fmt"

type Set map[string]struct{}

func getSetValues(s Set) []string {
	var ans []string
	for key := range s {
		ans = append(ans, key)
	}
	return ans
}

func add(s Set, value string) Set {
	s[value] = struct{}{}
	return s
}

func main() {
	s := make(Set)
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	s["item2"] = struct{}{}
	s = add(s, "item3")
	fmt.Println(getSetValues(s))
}
