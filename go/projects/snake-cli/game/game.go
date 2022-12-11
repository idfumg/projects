package game

import (
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	ScreenTop   = 0
	ApplesCount = 15
)

var debugstr string

type Game struct {
	screen    tcell.Screen
	inputChan chan string
	snake     *Snake
	apples    []*Apple
	isOver    bool
}

func NewGame(screen tcell.Screen) *Game {
	rand.Seed(time.Now().UnixNano())

	width, height := screen.Size()
	inputChan := initUserInput(screen)

	return &Game{
		screen:    screen,
		inputChan: inputChan,
		snake:     NewSnake(height/2, width/2),
		apples:    NewApples(height, width, ApplesCount),
		isOver:    false,
	}
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
	// _, screenBottom := g.screen.Size()

	if key == "" {
		return
	}

	switch key {
	case "Rune[q]":
		g.screen.Fini()
		os.Exit(0)
	case "Rune[w]":
		if !g.snake.MovingDown() {
			g.snake.MoveUp()
		}
	case "Rune[s]":
		if !g.snake.MovingUp() {
			g.snake.MoveDown()
		}
	case "Rune[a]":
		if !g.snake.MovingRight() {
			g.snake.MoveLeft()
		}
	case "Rune[d]":
		if !g.snake.MovingLeft() {
			g.snake.MoveRight()
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

func (g *Game) Finish() {
	g.screen.Fini()
	os.Exit(0)
}

func (g *Game) HandleInput() {
	g.applyInput(readInputIfExists(g.inputChan))
}

func (g *Game) DrawSnake() {
	for _, o := range g.snake.parts {
		g.print(o.GetColumn(), o.GetRow(), o.GetWidth(), o.GetHeight(), o.GetSymbol())
	}
}

func (g *Game) DrawApples() {
	for _, apple := range g.apples {
		if apple == nil {
			continue
		}
		g.print(apple.GetColumn(), apple.GetRow(), apple.GetWidth(), apple.GetHeight(), apple.GetSymbol())
	}
}

func (g *Game) DrawState() {
	g.screen.Clear()

	g.PrintString(0, 0, debugstr)

	g.DrawSnake()
	g.DrawApples()

	g.screen.Show()
}

func (g *Game) show() {
	g.screen.Show()
}

func (g *Game) UpdateState() {
	screenWidth, screenHeight := g.screen.Size()

	g.snake.Move()

	if appleIdx := isSnakeMeetApples(g.snake, g.apples); appleIdx != -1 {
		if !g.snake.Grow() {
			g.isOver = true
		} else {
			g.apples[appleIdx] = nil

			newApple := NewApple(screenHeight, screenWidth)
			for isSnakeMeetApple(g.snake, newApple) {
				newApple = NewApple(screenHeight, screenWidth)
			}
			g.apples[appleIdx] = newApple
		}
	}

	if g.collidesWithWall(g.snake.GetHead()) || isHeadMeetTail(g.snake) {
		g.isOver = true
	}
}

func isSnakeMeetApples(snake *Snake, apples []*Apple) int {
	for i, apple := range apples {
		if apple == nil {
			continue
		}
		for _, part := range snake.parts {
			if part.Col == apple.Col && part.Row == apple.Row {
				return i
			}
		}
	}
	return -1
}

func isSnakeMeetApple(snake *Snake, apple *Apple) bool {
	for _, part := range snake.parts {
		if part.Col == apple.Col && part.Row == apple.Row {
			return true
		}
	}
	return false
}

func isHeadMeetTail(snake *Snake) bool {
	head := snake.GetHead()
	for _, part := range snake.parts[1:] {
		if head.Col == part.Col && head.Row == part.Row {
			return true
		}
	}
	return false
}

func (g *Game) collidesWithWall(obj IObject) bool {
	_, screenBottom := g.screen.Size()
	return obj.NextRow() < ScreenTop || obj.NextRow() >= screenBottom
}

func (g *Game) IsOver() bool {
	return g.isOver
}

func (g *Game) PrintGameOverAndFinish(seconds int) {
	screenWidth, screenHeight := g.screen.Size()
	g.PrintStringCentered(screenWidth/2, screenHeight/2, "Game Over!")
	g.show()
	time.Sleep(time.Duration(seconds) * time.Second)
	g.Finish()
}
