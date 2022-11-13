package main

import (
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://yandex.ru",
}

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(resp.Status)
}

func crawl() {
	var wg = &sync.WaitGroup{}

	wg.Add(len(urls))

	for _, url := range urls {
		go fetch(url, wg)
	}

	wg.Wait()
}

func main() {
	crawl()
}
