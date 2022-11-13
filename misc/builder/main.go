package main

import (
	"fmt"
	"net/http"
	"time"
)

type httpClientBuilder struct {
	timeout time.Duration
}

func NewBuilder() *httpClientBuilder {
	return &httpClientBuilder{
		timeout: 5 * time.Second,
	}
}

func (b *httpClientBuilder) Build() *http.Client {
	return &http.Client{
		Timeout: b.timeout,
	}
}

func (b *httpClientBuilder) Timeout(t time.Duration) *httpClientBuilder {
	b.timeout = t
	return b
}

func main() {
	c := NewBuilder().Timeout(5 * time.Second).Build()
	res, err := c.Head("http://hoani.net")
	if err != nil {
		fmt.Printf("error ocuried %v\n", err)
	} else {
		fmt.Printf("status %v\n", res.Status)
	}
}
