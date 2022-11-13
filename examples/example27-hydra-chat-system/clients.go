package hydrachat

import (
	"bufio"
	"io"
)

type Client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

func StartClient(roomChannel chan string, cn io.ReadWriteCloser, quit chan struct{}) (chan string, chan struct{}) {
	client := new(Client)
	client.Reader = bufio.NewReader(cn)
	client.Writer = bufio.NewWriter(cn)
	client.wc = make(chan string)
	done := make(chan struct{})

	go func() {
		scanner := bufio.NewScanner(client.Reader)
		for scanner.Scan() {
			logger.Println(scanner.Text())
			roomChannel <- scanner.Text()
		}
		done <- struct{}{}
	}()

	client.writeMonitor()

	go func() {
		select {
		case <-quit:
			cn.Close()
		case <-done:
		}
	}()

	return client.wc, done
}
