package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	var text1 = "Hello, world!";

	text2 := "Hello, world!";

	var text3 string;
	text3 = "Hello, world!";

	say(text1)
	say(text2)
	say(text3)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("-->")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if (userInput == "quit") {
			break
		}
		fmt.Println(doctor.Response(userInput))
	}
}

func say(text string) {
	fmt.Println(text)
}