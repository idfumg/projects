package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	n := 0

	run := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exiting")
				close(ch)
				return
			case <-time.After(time.Millisecond * 300):
				fmt.Println(n)
				n++
			}
		}
	}

	quit := func(cancel func()) {
		time.Sleep(time.Second * 2)
		fmt.Println("goodbye")
		cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())

	go run(ctx)
	go quit(cancel)

	<-ch
}
