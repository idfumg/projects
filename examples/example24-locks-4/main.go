package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://golang.org",
		"http://google.com",
		"http://ya.ru",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string){
			defer wg.Done()
			r, err := http.Get(url)
			if err != nil {
				fmt.Println("Error occured:", err)
			} else {
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Println("Error occured:", err)
				} else {
					fmt.Println(string(body))
				}
			}
		}(url)
	}
	wg.Wait()
}