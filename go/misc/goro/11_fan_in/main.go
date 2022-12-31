package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	ch1, err := Read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("could not read file1 %v", err))
	}

	ch2, err := Read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("could not read file1 %v", err))
	}

	quit := make(chan struct{})

	ch := Merge1(ch1, ch2)

	consume := func() {
		for v := range ch {
			fmt.Println(v)
		}
		close(quit)
	}

	go consume()

	<-quit
}

func Merge1(ins ...<-chan []string) <-chan []string {
	var wg sync.WaitGroup

	out := make(chan []string)

	send := func(in <-chan []string) {
		for n := range in {
			out <- n
		}
		wg.Done()
	}

	done := func() {
		wg.Wait()
		close(out)
	}

	wg.Add(len(ins))

	for _, in := range ins {
		go send(in)
	}

	go done()

	return out
}

func Merge2(ins ...<-chan []string) <-chan []string {
	out := make(chan []string)
	done := make(chan struct{}, len(ins))
	send := func(in <-chan []string) {
		defer func() {
			done <- struct{}{}
		}()
		for v := range in {
			out <- v
		}
	}
	wait := func() {
		cnt := len(ins)
		for range done {
			cnt--
			if cnt == 0 {
				break
			}
		}
		close(out)
		close(done)
	}
	for _, in := range ins {
		go send(in)
	}

	go wait()

	return out
}

func Read(filename string) (<-chan []string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %v", filename, err)
	}
	ch := make(chan []string)
	reader := csv.NewReader(f)
	read := func() {
		for {
			record, err := reader.Read()
			if errors.Is(err, io.EOF) {
				close(ch)
				return
			}
			ch <- record
		}
	}
	go read()
	return ch, nil
}
