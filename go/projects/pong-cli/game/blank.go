package game

type Blank struct {
	Object
}

func NewBlank(obj IObject) *Blank {
	return &Blank{
		Object{
			Row:         obj.GetRow(),
			Col:         obj.GetColumn(),
			Width:       obj.GetWidth(),
			Height:      obj.GetHeight(),
			Symbol:      ' ',
			RowVelocity: 0,
			ColVelocity: 0,
		},
	}
}
