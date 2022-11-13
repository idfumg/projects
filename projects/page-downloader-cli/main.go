package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func getFilename(url string) string {
	return strings.Split(url, "//")[1] + ".txt"
}

func download(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status code: %d\n", resp.StatusCode)
	if resp.StatusCode != 200 {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	filename := getFilename(url)
	fmt.Printf("Writing response body to %s\n\n", filename)
	if err := ioutil.WriteFile(filename, bodyBytes, 0666); err != nil {
		fmt.Println(err)
	}
}

func main() {
	urls := os.Args[1:]
	wg := sync.WaitGroup{}

	wg.Add(len(urls))

	for _, url := range urls {
		go download(url, &wg)
	}

	wg.Wait()
}
