package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult string `json:"round_result"`
}

var winMessages = []string {
	"Good job!",
	"Nice work!",
	"You should buy a lottery ticket!",
}

var loseMessages = []string {
	"To bad!",
	"Try again!",
	"This is just not your day!",
}

var drawMessages = []string {
	"Great mind think alike",
	"Oh no. Try again",
	"Nobody wins, but you can try again",
}

func getComputerChoice() (int, string) {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	switch (computerValue) {
	case ROCK: return computerValue, "Computer chose ROCK";
	case PAPER: return computerValue, "Computer chose PAPER";
	case SCISSORS: return computerValue, "Computer chose SCISSORS";
	}
	return 0, ""
}

func getRandomMesssage(messages []string) string {
	rand.Seed(time.Now().UnixNano())
	return messages[rand.Intn(3)]
}

func getRandomWinMessage() string {
	return getRandomMesssage(winMessages)
}

func getRandomLoseMessage() string {
	return getRandomMesssage(loseMessages)
}

func getRandomDrawMessage() string {
	return getRandomMesssage(drawMessages)
}

func getRoundResult(player int, computer int) (string, string) {
	if (player == computer) { return getRandomDrawMessage(), "It's a draw" }
	if (player == (computer + 1) % 3) { return getRandomWinMessage(), "Player wins!" }
	return getRandomLoseMessage(), "Computer wins!"
}

func PlayRound(playerValue int) Round {
	computerValue, computerChoice := getComputerChoice()
	message, roundResult := getRoundResult(playerValue, computerValue)
	return Round{
		Message: message,
		ComputerChoice: computerChoice,
		RoundResult: roundResult,
	}
}