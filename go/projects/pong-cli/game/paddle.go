package game

const (
	PaddleSymbol          = '█'
	PaddleWidth           = 1
	PaddleHeight          = 8
	InitPaddleRowVelocity = 0
	InitPaddleColVelocity = 0
)

type Paddle struct {
	Object
}

func NewPaddle(row int, col int) *Paddle {
	return &Paddle{
		Object{
			Row:         row,
			Col:         col,
			Width:       PaddleWidth,
			Height:      PaddleHeight,
			Symbol:      PaddleSymbol,
			RowVelocity: InitPaddleRowVelocity,
			ColVelocity: InitPaddleColVelocity,
		},
	}
}
