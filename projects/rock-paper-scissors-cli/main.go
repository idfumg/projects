package main

import "myapp/game"

func main() {
	game := game.Create()
	go game.Rounds()
	game.ClearScreen()
	game.PrintIntro()

	for {
		game.RoundChan <- 1
		<-game.RoundChan
		if game.Round.RoundNumber > 3 {
			break
		}
		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}

	game.PrintSummary()
}
