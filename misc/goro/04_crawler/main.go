package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	ch := make(chan string, 1024)
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go fetchUrl(ch, wg)
	}

	urls := []string{
		"http://example.com",
		"http://packtpub.com",
		"http://ya.ru",
		"http://twitter.com",
		"http://facebook.com",
	}

	wg.Add(len(urls))
	for _, url := range urls {
		ch <- url
	}
	wg.Wait()
}

func fetchUrl(ch <- chan string, wg *sync.WaitGroup) {
	for {
		url := <- ch
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Fetching error: %s (%s)\n", err, url)
		} else {
			fmt.Printf("Status: %s (%s)\n", resp.Status, url)
		}
		wg.Done()
	}
}