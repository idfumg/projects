package game

type Snake struct {
	parts []*Object
}

const (
	PartWidth               = 1
	PartHeight              = 1
	PartSymbol              = 0x2588
	SnakeDefaultRowVelocity = 0
	SnakeDefaultColVelocity = 0
	Speed                   = 1
)

func NewSnake(row int, col int) *Snake {
	parts := []*Object{}

	init := &Object{
		Row:         row,
		Col:         col,
		Width:       PartWidth,
		Height:      PartHeight,
		Symbol:      PartSymbol,
		RowVelocity: 0,
		ColVelocity: -Speed,
	}

	parts = append(parts, init)
	for i := 1; i < 10; i += 1 {
		obj := *init
		obj.Col = col + i
		obj.ColVelocity = 0
		parts = append(parts, &obj)
	}

	return &Snake{
		parts,
	}
}

func (s *Snake) Grow() bool {
	part := s.getNewSnakePart()
	s.parts = append(s.parts, part)
	return part != nil
}

func (s *Snake) getNewSnakePart() *Object {
	if len(s.parts) == 1 {
		head := s.GetHead()
		if head.RowVelocity == 1 {
			return &Object{
				Row:         head.Row - 1,
				Col:         head.Col,
				Width:       PartWidth,
				Height:      PartHeight,
				Symbol:      PartSymbol,
				RowVelocity: 0,
				ColVelocity: 0,
			}
		} else if head.RowVelocity == -1 {
			return &Object{
				Row:         head.Row + 1,
				Col:         head.Col,
				Width:       PartWidth,
				Height:      PartHeight,
				Symbol:      PartSymbol,
				RowVelocity: 0,
				ColVelocity: 0,
			}
		} else if head.ColVelocity == 1 {
			return &Object{
				Row:         head.Row,
				Col:         head.Col - 1,
				Width:       PartWidth,
				Height:      PartHeight,
				Symbol:      PartSymbol,
				RowVelocity: 0,
				ColVelocity: 0,
			}
		} else if head.ColVelocity == -1 {
			return &Object{
				Row:         head.Row,
				Col:         head.Col + 1,
				Width:       PartWidth,
				Height:      PartHeight,
				Symbol:      PartSymbol,
				RowVelocity: 0,
				ColVelocity: 0,
			}
		}
	}
	a := s.parts[len(s.parts)-1]
	b := s.parts[len(s.parts)-2]
	if a.Row == b.Row+1 {
		return &Object{
			Row:         a.Row + 1,
			Col:         a.Col,
			Width:       PartWidth,
			Height:      PartHeight,
			Symbol:      PartSymbol,
			RowVelocity: 0,
			ColVelocity: 0,
		}
	} else if a.Row == b.Row-1 {
		return &Object{
			Row:         a.Row - 1,
			Col:         a.Col,
			Width:       PartWidth,
			Height:      PartHeight,
			Symbol:      PartSymbol,
			RowVelocity: 0,
			ColVelocity: 0,
		}
	} else if a.Col == b.Col+1 {
		return &Object{
			Row:         a.Row,
			Col:         a.Col + 1,
			Width:       PartWidth,
			Height:      PartHeight,
			Symbol:      PartSymbol,
			RowVelocity: 0,
			ColVelocity: 0,
		}
	} else if a.Col == b.Col-1 {
		return &Object{
			Row:         a.Row,
			Col:         a.Col - 1,
			Width:       PartWidth,
			Height:      PartHeight,
			Symbol:      PartSymbol,
			RowVelocity: 0,
			ColVelocity: 0,
		}
	}
	i, j := s.getPlaceToPutPart(a)
	if i == -1 || j == -1 {
		return nil
	}
	return &Object{
		Row:         i,
		Col:         j,
		Width:       PartWidth,
		Height:      PartHeight,
		Symbol:      PartSymbol,
		RowVelocity: 0,
		ColVelocity: 0,
	}
}

func (s *Snake) getPlaceToPutPart(item *Object) (int, int) {
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, direction := range directions {
		i := direction[0] + item.Row
		j := direction[1] + item.Col
		if s.isEmptyPlace(i, j) {
			return i, j
		}
	}
	return -1, -1
}

func (s *Snake) isEmptyPlace(i int, j int) bool {
	for _, part := range s.parts {
		if part.Row == i && part.Col == j {
			return true
		}
	}
	return true
}

func (s *Snake) GetHead() *Object {
	return s.parts[0]
}

func (s *Snake) MoveUp() {
	p := s.GetHead()
	p.RowVelocity = -Speed
	p.ColVelocity = 0
}

func (s *Snake) MoveDown() {
	p := s.GetHead()
	p.RowVelocity = Speed
	p.ColVelocity = 0
}

func (s *Snake) MoveLeft() {
	p := s.GetHead()
	p.RowVelocity = 0
	p.ColVelocity = -Speed
}

func (s *Snake) MoveRight() {
	p := s.GetHead()
	p.RowVelocity = 0
	p.ColVelocity = Speed
}

func (s *Snake) MovingUp() bool {
	p := s.GetHead()
	return p.RowVelocity == -Speed
}

func (s *Snake) MovingDown() bool {
	p := s.GetHead()
	return p.RowVelocity == Speed
}

func (s *Snake) MovingLeft() bool {
	p := s.GetHead()
	return p.ColVelocity == -Speed
}

func (s *Snake) MovingRight() bool {
	p := s.GetHead()
	return p.ColVelocity == Speed
}

func (s *Snake) Move() {
	for i := len(s.parts) - 1; i > 0; i -= 1 {
		s.parts[i].Row = s.parts[i-1].Row
		s.parts[i].Col = s.parts[i-1].Col
	}
	s.parts[0].Move()
}
