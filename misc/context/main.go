package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	sleepAndTalk(ctx, 3*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, duration time.Duration, msg string) {
	select {
	case <-time.After(duration):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
