package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()
	doneCh := make(chan bool)
	go readUserInput(os.Stdin, os.Stdout, doneCh)
	<-doneCh
	close(doneCh)
	fmt.Println("Bye bye")
}

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("------------")
	fmt.Println("Enter the whole number, and we'll tell you if it's	a prime")
	prompt(os.Stdout)
}

func prompt(w io.Writer) {
	w.Write([]byte("-> "))
}

func readUserInput(r io.Reader, w io.Writer, doneCh chan<- bool) {
	scanner := bufio.NewScanner(r)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneCh <- true
			return
		}

		w.Write([]byte(res))
		prompt(w)
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number, by definition!", n)
	}

	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
