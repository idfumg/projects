package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	s := &SpySleeper{}
	Countdown(buffer, s)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if s.Calls != 3 {
		t.Errorf("not enough calls to sleeper, got %q, want 3", s.Calls)
	}
}

func TestCountdownOperations(t *testing.T) {
	buffer := &bytes.Buffer{}

	s := &SpyCountdownOperations{}
	Countdown(buffer, s)

	want := []string{
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}

	if !reflect.DeepEqual(s.Calls, want) {
		t.Errorf("not enough calls to sleeper, got %q, want %q", s.Calls, want)
	}
}
