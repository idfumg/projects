package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface{
	Sleep()
	Write(w io.Writer, v interface{}) (int, error)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpySleeper) Write(w io.Writer, v interface{}) (int, error) {
	return fmt.Fprint(w, v)
}

type DefaultSleeper struct {

}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *DefaultSleeper) Write(w io.Writer, v interface{}) (int, error) {
	return fmt.Fprint(w, v)
}

type SpyCountdownOperations struct {
	Calls []string
}

var (
	sleep = "sleep"
	write = "write"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(w io.Writer, v interface{}) (int, error) {
	s.Calls = append(s.Calls, write)
	return 0, nil
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Write(w, fmt.Sprintf("%d\n", i))
		s.Sleep()
	}
	s.Write(w, "Go!")
}
