package game

type Object struct {
	Row         int
	Col         int
	Width       int
	Height      int
	Symbol      rune
	RowVelocity int
	ColVelocity int
}

func (p *Object) MoveVertically(x int) {
	p.Row += x
}

func (p *Object) MoveHorizontally(y int) {
	p.Col += y
}

func (p *Object) GetRow() int {
	return p.Row
}

func (p *Object) GetColumn() int {
	return p.Col
}

func (p *Object) GetWidth() int {
	return p.Width
}

func (p *Object) GetHeight() int {
	return p.Height
}

func (p *Object) GetSymbol() rune {
	return p.Symbol
}

func (p *Object) GetRowVelocity() int {
	return p.RowVelocity
}

func (p *Object) GetColVelocity() int {
	return p.ColVelocity
}

func (p *Object) SetRow(value int) {
	p.Row = value
}

func (p *Object) SetCol(value int) {
	p.Col = value
}

func (p *Object) SetRowVelocity(value int) {
	p.RowVelocity = value
}

func (p *Object) SetColVelocity(value int) {
	p.ColVelocity = value
}

func (p *Object) Move() {
	p.Row += p.RowVelocity
	p.Col += p.ColVelocity
}

func (p *Object) NextRow() int {
	return p.Row + p.RowVelocity
}

func (p *Object) NextColumn() int {
	return p.Col + p.ColVelocity
}
