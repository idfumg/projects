package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(filename string) error {
	fd, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	fd.Close()
	return nil
}

func check1() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}

func check2() {
	err := fileChecker("not_here.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("That file doesn't exist")
		}
	}
}

func main() {
	check1()
	check2()
}
