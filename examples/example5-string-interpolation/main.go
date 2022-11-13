package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

type User struct {
	Username string
	Age int
	FavouriteNumber float64
	OwnsADog bool
}

func main() {
	user := readUser()
	// fmt.Println(fmt.Sprintf("Your name is %s and you are %d years old", username, age))
	fmt.Printf(
		"Your name is %s and you are %d years old. Favourite number is %.2f. Dog? %t.\n",
		user.Username, 
		user.Age, 
		user.FavouriteNumber,
		user.OwnsADog)
}

func readUser() User {
	var user User
	user.Username = readValidString("What's your name?")
	user.Age = readValidInt("How old are you?")
	user.FavouriteNumber = readValidFloat("What's your favourite number?")
	user.OwnsADog = readValidBool("Do you own a dog?")
	return user
}

func prompt() {
	fmt.Print("-> ")
}

func readLine() string {
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n\r", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	return input
}

func readString(s string) string {
	fmt.Println(s)
	prompt()
	return readLine()
}

func readValidString(s string) string {
	for {
		input := readString(s)
		if input != "" {
			return input
		}
	}
}

func readInt(s string) (int, error) {
	fmt.Println(s)
	prompt()
	input := readLine()
	num, err := strconv.Atoi(input)
	return num, err
}

func readValidInt(s string) int {
	for {
		num, err := readInt(s)
		if err == nil {
			return num
		}
	}
}

func readFloat(s string) (float64, error) {
	fmt.Println(s)
	prompt()
	input := readLine()
	num, err := strconv.ParseFloat(input, 64)
	return num, err
}

func readValidFloat(s string) float64 {
	for {
		num, err := readFloat(s)
		if err == nil {
			return num
		}
	}
}

func readBool(s string) (bool, error) {
	fmt.Println(s)
	prompt()
	input := readLine()
	num, err := strconv.ParseBool(input)
	return num, err
}

func readValidBool(s string) bool {
	for {
		num, err := readBool(s)
		if err == nil {
			return num
		}
	}
}