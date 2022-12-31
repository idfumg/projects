package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	records, err := Read("file1.csv")
	if err != nil {
		log.Fatalf("Could not read a file: %v\n", err)
	}
	for v := range sanitize(titleize(records)) {
		fmt.Printf("%v\n", v)
	}
}

func Read(filename string) (<-chan []string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file %w", err)
	}
	ch := make(chan []string)
	read := func() {
		reader := csv.NewReader(f)
		defer f.Close()
		defer close(ch)
		for {
			record, err := reader.Read()
			if err != nil {
				return
			}
			ch <- record
		}
	}
	go read()
	return ch, nil
}

func titleize(c <-chan []string) <-chan []string {
	ch := make(chan []string)
	go func() {
		for v := range c {
			fmt.Println(v)
			v[0] = strings.Title(v[0])
			// v[1], v[2] = v[2], v[1]
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func sanitize(c <-chan []string) <-chan []string {
	ch := make(chan []string)
	go func() {
		for v := range c {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
