package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	quit := make(chan struct{})

	produce := func() {
		for i := 0; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("All completed")
		close(ch)
	}

	consume := func() {
		for v := range ch {
			fmt.Println(time.Now(), "received ", v)
		}
		close(quit)
	}

	go produce()
	go consume()

	<-quit
}
