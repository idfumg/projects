package game

import "math/rand"

const (
	AppleSymbol = 0x25CF
)

type Apple struct {
	Object
}

func NewApple(screenHeight int, screenWidth int) *Apple {
	return &Apple{
		Object{
			Row:         rand.Intn(screenHeight-20) + 10,
			Col:         rand.Intn(screenWidth-20) + 10,
			Width:       PartWidth,
			Height:      PartHeight,
			Symbol:      AppleSymbol,
			RowVelocity: 0,
			ColVelocity: 0,
		},
	}
}

func theyEqual(a *Apple, b *Apple) bool {
	return a.Col == b.Col && a.Row == b.Row
}

func isExists(apples []*Apple, target *Apple) bool {
	for _, apple := range apples {
		if theyEqual(apple, target) {
			return true
		}
	}
	return false
}

func NewApples(screenHeight int, screenWidth int, n int) []*Apple {
	apples := []*Apple{}
	for i := 0; i < n; i += 1 {
		apple := NewApple(screenHeight, screenWidth)
		if !isExists(apples, apple) {
			apples = append(apples, apple)
		}
	}
	return apples
}
