package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalln(res.Status)
	}

	io.Copy(os.Stdout, res.Body)

	err = res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
