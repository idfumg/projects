package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// wait := WaitGroups()
	// wait := ErrGroups()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	wait := ErrGroupsCtx(ctx)
	<-wait
}

func WaitGroups() <-chan struct{} {
	ch := make(chan struct{}, 1)
	var wg sync.WaitGroup
	for _, file := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		file := file
		wg.Add(1)
		read := func() {
			defer wg.Done()
			ch, err := Read(file)
			if err != nil {
				fmt.Printf("error reading file: %v", err)
				return
			}
			for line := range ch {
				fmt.Println(line)
			}
		}
		go read()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func ErrGroups() <-chan struct{} {
	ch := make(chan struct{}, 1)
	var g errgroup.Group
	for _, file := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		file := file
		g.Go(func() error {
			ch, err := Read(file)
			if err != nil {
				return fmt.Errorf("error reading %s: %w", file, err)
			}
			for line := range ch {
				fmt.Println(line)
			}
			return nil
		})
	}
	go func() {
		if err := g.Wait(); err != nil {
			fmt.Println("Error reading files: %w", err)
		}
		close(ch)
	}()
	return ch
}

func ErrGroupsCtx(ctx context.Context) <-chan struct{} {
	ch := make(chan struct{}, 1)
	g, ctx := errgroup.WithContext(ctx)
	for _, file := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		file := file
		g.Go(func() error {
			ch, err := Read(file)
			if err != nil {
				return fmt.Errorf("error reading %s: %w", file, err)
			}
			for {
				select {
				case line, ok := <-ch:
					if !ok {
						return nil
					}
					fmt.Println(line)
				case <-ctx.Done():
					fmt.Printf("Context is completed %v\n", ctx.Err())
					return ctx.Err()
				}
			}
		})
	}
	go func() {
		if err := g.Wait(); err != nil {
			fmt.Println("Error reading files: %w", err)
		}
		close(ch)
	}()
	return ch
}

func Read(filename string) (<-chan []string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	ch := make(chan []string)
	read := func() {
		reader := csv.NewReader(f)
		time.Sleep(time.Millisecond)
		for {
			record, err := reader.Read()
			if errors.Is(err, io.EOF) {
				break
			}
			ch <- record
		}
		close(ch)
	}
	go read()
	return ch, nil
}
