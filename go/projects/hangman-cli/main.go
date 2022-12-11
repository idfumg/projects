package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"hangman/game"
	"hangman/utils"
)

var dictionary = [...]string{
	"Zombie",
	"Gopher",
	"United States of America",
	"Indonesia",
	"Apple",
	"Programming",
}

func getRandomWord() string {
	return strings.ToLower(dictionary[rand.Intn(len(dictionary))])
}

func startGame(g *game.Game) {
	for !g.IsDone() {
		g.PrintState()
		g.ChangeState(utils.GetUserInput())
	}
	g.PrintState()
	fmt.Println("You're done! :)")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	startGame(game.NewGame(getRandomWord()))
}
