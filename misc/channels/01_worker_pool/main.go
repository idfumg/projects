package main

import (
	"fmt"
	"net/http"
)

type Site struct {
	Url string
}

type Result struct {
	Url    string
	Status int
}

func crawl(w int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		fmt.Printf("Worker #%d is fetching %s\n", w, site.Url)
		resp, err := http.Get(site.Url)
		if err != nil {
			fmt.Println(err)
		} else {
			results <- Result{
				Url:    site.Url,
				Status: resp.StatusCode,
			}
		}
	}
}

func main() {
	jobs := make(chan Site, 2)
	results := make(chan Result, 2)

	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, results)
	}

	urls := []string{
		"https://yandex.ru",
		"https://google.com",
		"https://example.com",
		"https://youtube.com",
	}

	for _, url := range urls {
		jobs <- Site{
			Url: url,
		}
	}
	close(jobs)

	for i := 1; i <= 4; i++ {
		result := <- results
		fmt.Println(result.Url, "->", result.Status)
	}
}
