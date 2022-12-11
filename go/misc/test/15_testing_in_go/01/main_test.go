package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrimeWithTableTests(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
		msg   string
	}{
		{"ZeroIsNotPrime", 0, false, "0 is not a prime number, by definition!"},
		{"OneIsNotPrime", 1, false, "1 is not a prime number, by definition!"},
		{"SevenIsPrime", 7, true, "7 is a prime number!"},
		{"MinusSevenIsNotPrime", -7, false, "Negative numbers are not prime, by definition!"},
		{"EightIsNotPrime", 8, false, "8 is not prime number because it is divisible by 2"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, msg := isPrime(test.value)
			if got != test.want {
				t.Errorf("with value %d, got %v, but expected %v", test.value, got, test.want)
			}
			if msg != test.msg {
				t.Error("wrong message returned", msg)
			}
		})
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	prompt(w)
	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> got %v", out)
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Enter the whole number") {
		t.Errorf("incorrect intro: %v", out)
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"ValidPrimeNumber", "7", "7 is a prime number!"},
		{"qForExit", "q", ""},
		{"QForExit", "Q", ""},
		{"WrongInputAlpha", "abc", "Please enter a whole number!"},
		{"WrongInputFloat", "7.7", "Please enter a whole number!"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.input)
			scanner := bufio.NewScanner(reader)
			got, _ := checkNumbers(scanner)

			if !strings.EqualFold(got, test.want) {
				t.Errorf("wrong result, got %v, want %v", got, test.want)
			}
		})
	}
}

func Test_readUserInput(t *testing.T) {
	doneCh := make(chan bool)
	stdin := &bytes.Buffer{}
	stdin.Write([]byte("1\nq\n"))
	stdout := &bytes.Buffer{}
	go readUserInput(stdin, stdout, doneCh)
	<-doneCh
	close(doneCh)
}
