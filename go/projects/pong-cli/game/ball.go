package game

import (
	"math/rand"
)

const (
	BallSymbol = '●'
	BallWidth  = 1
	BallHeight = 1
)

type Ball struct {
	Object
}

func getModifier() int {
	x := rand.Intn(2)
	if x == 0 {
		return 1
	}
	return -1
}

func getInitVelocity() int {
	return rand.Intn(3) + 1
}

func NewBall(row int, col int) *Ball {
	return &Ball{
		Object{
			Row:         row,
			Col:         col,
			Width:       BallWidth,
			Height:      BallHeight,
			Symbol:      BallSymbol,
			RowVelocity: getInitVelocity() * getModifier(),
			ColVelocity: getInitVelocity() * getModifier(),
		},
	}
}

func (b *Ball) ChangeRowDirection() {
	b.RowVelocity *= -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (b *Ball) ChangeColumnDirection() {
	b.ColVelocity *= -1

	b.RowVelocity /= abs(b.RowVelocity)
	b.RowVelocity *= getInitVelocity()
	b.ColVelocity /= abs(b.ColVelocity)
	b.ColVelocity *= getInitVelocity()
}
