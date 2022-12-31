package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	ch1, err := Read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("could not read the file %v", err))
	}

	br1 := Breakup("1", ch1)
	br2 := Breakup("2", ch1)
	br3 := Breakup("3", ch1)

	for {
		if br1 == nil && br2 == nil && br3 == nil {
			break
		}
		select {
		case _, ok := <-br1:
			if !ok {
				br1 = nil
			}
		case _, ok := <-br2:
			if !ok {
				br2 = nil
			}
		case _, ok := <-br3:
			if !ok {
				br3 = nil
			}
		}
	}

	fmt.Println("Done")
}

func Breakup(worker string, ch <- chan []string) <-chan struct{} {
	out := make(chan struct{})
	consume := func() {
		for v := range ch {
			fmt.Println(worker, v)
		}
		close(out)
	}
	go consume()
	return out
}

func Read(filename string) (<-chan []string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open a file %s: %v", filename, err)
	}
	ch := make(chan []string)
	read := func() {
		defer f.Close()
		defer close(ch)
		reader := csv.NewReader(f)
		for {
			record, err := reader.Read()
			if errors.Is(err, io.EOF) {
				return
			}
			ch <- record
		}
	}
	go read()
	return ch, nil
}
