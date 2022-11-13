package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func getNumBetween2And10() int {
	return rand.Intn(8) + 2
}

func getAnswer(first, second, third int) int {
	return first * second - third;
}

func playGame(first, second, third int, reader *bufio.Reader) {
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10 ", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by ", first, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply th result by ", second, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", third, prompt)
	reader.ReadString('\n')

	fmt.Println("The answer is", getAnswer(first, second, third))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	playGame(
		getNumBetween2And10(), 
		getNumBetween2And10(), 
		getNumBetween2And10(), 
		bufio.NewReader(os.Stdin))
}