package game

import (
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	ScreenTop = 0
)

type Game struct {
	screen    tcell.Screen
	paddle1   *Paddle
	paddle2   *Paddle
	ball      *Ball
	objects   []IObject
	inputChan chan string
	tempObjects []*Blank
}

func initUserInput(screen tcell.Screen) chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventKey:
				inputChan <- ev.Name()
			}
		}
	}()
	return inputChan
}

func readInputIfExists(inputChan chan string) string {
	var key string
	select {
	case key = <-inputChan:
	default:
		key = ""
	}
	return key
}

func (g *Game) applyInput(key string) {
	_, screenBottom := g.screen.Size()

	if key == "" {
		return
	}

	switch key {
	case "Rune[q]":
		g.screen.Fini()
		os.Exit(0)
	case "Rune[w]":
		if g.paddle1.Row > ScreenTop {
			g.tempObjects = append(g.tempObjects, NewBlank(g.paddle1))
			g.paddle1.MoveVertically(-1)
		}
	case "Rune[s]":
		if g.paddle1.Row+g.paddle1.Height < screenBottom {
			g.tempObjects = append(g.tempObjects, NewBlank(g.paddle1))
			g.paddle1.MoveVertically(1)
		}
	case "Up":
		if g.paddle2.Row > ScreenTop {
			g.tempObjects = append(g.tempObjects, NewBlank(g.paddle2))
			g.paddle2.MoveVertically(-1)
		}
	case "Down":
		if g.paddle2.Row+g.paddle2.Height < screenBottom {
			g.tempObjects = append(g.tempObjects, NewBlank(g.paddle2))
			g.paddle2.MoveVertically(1)
		}
	}
}

func (g *Game) PrintString(col int, row int, str string) {
	for _, c := range str {
		g.screen.SetContent(col, row, c, nil, tcell.StyleDefault)
		col += 1
	}
}

func (g *Game) PrintStringCentered(col int, row int, str string) {
	g.PrintString(col-len(str)/2, row, str)
}

func (g *Game) print(col, row, width, height int, ch rune) {
	for r := 0; r < height; r += 1 {
		for c := 0; c < width; c += 1 {
			g.screen.SetContent(col+c, row+r, ch, nil, tcell.StyleDefault)
		}
	}
}

func NewGame(screen tcell.Screen) *Game {
	rand.Seed(time.Now().UnixNano())

	width, height := screen.Size()
	paddleStart := height/2 - PaddleHeight/2

	paddle1 := NewPaddle(paddleStart, 0)
	paddle2 := NewPaddle(paddleStart, width-1)
	ball := NewBall(height/2, width/2)
	objects := append([]IObject{}, paddle1, paddle2, ball)
	inputChan := initUserInput(screen)

	return &Game{
		screen:    screen,
		paddle1:   paddle1,
		paddle2:   paddle2,
		ball:      ball,
		objects:   objects,
		inputChan: inputChan,
	}
}

func (g *Game) Finish() {
	g.screen.Fini()
	os.Exit(0)
}

func (g *Game) HandleInput() {
	g.applyInput(readInputIfExists(g.inputChan))
}

func (g *Game) DrawState() {
	// g.screen.Clear()
	for _, o := range g.tempObjects {
		g.print(o.GetColumn(), o.GetRow(), o.GetWidth(), o.GetHeight(), o.GetSymbol())
	}
	g.tempObjects = nil

	for _, o := range g.objects {
		g.print(o.GetColumn(), o.GetRow(), o.GetWidth(), o.GetHeight(), o.GetSymbol())
	}

	g.screen.Show()
}

func (g *Game) show() {
	g.screen.Show()
}

func (g *Game) UpdateState() {
	for _, o := range g.objects {
		g.tempObjects = append(g.tempObjects, NewBlank(o))
		o.Move()
	}

	if g.collidesWithWall(g.ball) {
		g.ball.ChangeRowDirection()
	} else if g.coolidesWithPaddle(g.ball) {
		g.ball.ChangeColumnDirection()
	}
}

func (g *Game) collidesWithWall(obj IObject) bool {
	_, screenBottom := g.screen.Size()
	return obj.NextRow() < ScreenTop || obj.NextRow() >= screenBottom
}

func (g *Game) coolidesWithPaddle(obj IObject) bool {
	return collidesWith(obj, g.paddle1) || collidesWith(obj, g.paddle2)
}

func (g *Game) getWinner() string {
	if g.ball.GetColumn() <= 0 {
		return "Player #2 wins"
	}
	screenRight, _ := g.screen.Size()
	if g.ball.GetColumn() >= screenRight {
		return "Player #1 wins"
	}
	return ""
}

func (g *Game) IsOver() bool {
	return g.getWinner() != ""
}

func (g *Game) PrintGameOverAndFinish(seconds int) {
	screenWidth, screenHeight := g.screen.Size()
	g.PrintStringCentered(screenWidth/2, screenHeight/2, "Game Over!")
	g.PrintStringCentered(screenWidth/2, screenHeight/2+2, g.getWinner())
	g.show()
	time.Sleep(time.Duration(seconds) * time.Second)
	g.Finish()
}
