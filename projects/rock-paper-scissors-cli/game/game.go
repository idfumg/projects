package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Reader      *bufio.Reader
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

func Create() Game {
	displayChan := make(chan string)
	roundChan := make(chan int)

	return Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Reader:      bufio.NewReader(os.Stdin),
		Round: Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}
}

func (g *Game) Rounds() {
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber += round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Print(msg)
			g.DisplayChan <- ""
		}
	}
}

func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins the game")
	fmt.Println()
}

func (g *Game) readInput() string {
	input, _ := g.Reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	input = strings.ToLower(input)
	return input
}

func (g *Game) getUserInput() (string, int) {
	playerChoice := g.readInput()

	playerValue := -1
	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}
	return playerChoice, playerValue
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())

	g.DisplayChan <- fmt.Sprintf("\nRound %d\n", g.Round.RoundNumber)
	<-g.DisplayChan
	g.DisplayChan <- "------\n"
	<-g.DisplayChan
	g.DisplayChan <- "Please enter rock, paper or scissors -> "
	<-g.DisplayChan

	playerChoice, playerValue := g.getUserInput()
	if (playerValue == -1) {
		g.DisplayChan <- "Wrong value!"
		<-g.DisplayChan
		return false;
	}

	g.DisplayChan <- fmt.Sprintf("\nPlayer chose %s\n", strings.ToUpper(playerChoice))
	<-g.DisplayChan

	computerValue := rand.Intn(3)
	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose ROCK\n"
		break
	case PAPER:
		g.DisplayChan <- "Computer chose PAPER\n"
		break
	case SCISSORS:
		g.DisplayChan <- "Computer chose SCISSORS\n"
		break
	}
	<-g.DisplayChan

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!\n"
		<-g.DisplayChan
		return false
	}

	if computerValue == ((playerValue+1)%3) {
		g.playerWins()
	} else {
		g.computerWins()
	}

	return true
}

func (g *Game) playerWins() {
	g.Round.ComputerScore += 1
	g.DisplayChan <- "Computer wins\n"
	<-g.DisplayChan
}

func (g *Game) computerWins() {
	g.Round.PlayerScore += 1
	g.DisplayChan <- "Player wins\n"
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	fmt.Println("\n\nFinal score")
	fmt.Println("-----------")
	fmt.Printf("Player: %d/3, Computer: %d/3\n",
		g.Round.PlayerScore,
		g.Round.ComputerScore)
	if g.Round.PlayerScore > g.Round.ComputerScore {
		fmt.Println("Player wins game!")
	} else {
		fmt.Println("Computer wins game!")
	}
}