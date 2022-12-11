package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

type Result struct {
	url   string
	urls  []string
	err   error
	depth int
}

var fetched map[string]bool

func fetch(url string, depth int, ch chan<- *Result) {
	urls, err := findLinks(url)
	ch<- &Result{
		url,
		urls,
		err, 
		depth,
	}
}

func Crawl(url string, depth int) {
	ch := make(chan *Result)
	go fetch(url, depth, ch)
	fetched[url] = true

	for fetching := 1; fetching > 0; fetching -= 1 {
		result := <-ch
		if result.err != nil {
			continue
		}
		fmt.Printf("Found %s\n", result.url)
		if result.depth > 0 {
			for _, u := range result.urls {
				if fetched[u] {
					continue
				}
				go fetch(u, result.depth - 1, ch)
				fetched[u] = true
				fetching += 1
			}
		}
	}

	close(ch)
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	fetched = make(map[string]bool)
	now := time.Now()
	Crawl("http://andcloud.io", 2)
	fmt.Println("Time taken:", time.Since(now))
}
