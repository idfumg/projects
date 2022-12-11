package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, chs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)
	consume := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			select {
			case out <- v:
			case <-done:
				return
			}
		}
	}
	for _, c := range chs {
		wg.Add(1)
		go consume(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	done := make(chan struct{})
	in := generator(done, 2, 3, 4, 5)
	ch1 := square(done, in)
	ch2 := square(done, in)
	// for value := range merge(ch1, ch2) {
	// 	fmt.Print(value, " ")
	// }
	out := merge(done, ch1, ch2)
	fmt.Println(<-out)
	close(done)
	time.Sleep(2 * time.Second)
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}
