package utils

import (
	"fmt"
	"strings"
)

func GetCount(s []rune, ch rune) int {
	cnt := 0
	for i := range s {
		if s[i] == ch {
			cnt += 1
		}
	}
	return cnt
}

func isValidInput(s string) bool {
	return len(s) == 1 && (s[0] >= 'a' && s[0] <= 'z')
}

func getInput() string {
	var s string
	fmt.Scanf("%s", &s)
	return strings.ToLower(strings.TrimSpace(s))
}

func GetUserInput() rune {
	for {
		fmt.Printf("-> ")
		if s := getInput(); isValidInput(s) {
			return rune(s[0])
		}
	}
}
