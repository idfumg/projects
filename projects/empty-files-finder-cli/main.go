package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please, enter a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	var names []string

	for _, file := range files {
		if file.Size() == 0 {
			names = append(names, file.Name())
		}
	}

	var namesBytes []byte

	for _, name := range names {
		namesBytes = append(namesBytes, name...)
		namesBytes = append(namesBytes, '\n')
	}

	err = ioutil.WriteFile("out.txt", namesBytes, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", namesBytes)
}