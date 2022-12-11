package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"

	"snake/game"
)

func InitScreen() tcell.Screen {
	screen, err := tcell.NewScreen()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	screen.SetStyle(defStyle)

	return screen
}

func main() {
	encoding.Register()
	screen := InitScreen()
	game := game.NewGame(screen)

	for !game.IsOver() {
		game.HandleInput()
		game.UpdateState()
		game.DrawState()

		time.Sleep(50 * time.Millisecond)
	}

	game.PrintGameOverAndFinish(2)
}
