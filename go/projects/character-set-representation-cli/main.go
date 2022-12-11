package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n",
		"literal", "dec", "hex", "encoded", strings.Repeat("-", 45))
	for ch := 'A'; ch <= 'Z'; ch += 1{
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x\n", ch, ch, ch, string(ch))
	}
}