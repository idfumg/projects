package game

type IObject interface {
	MoveVertically(int)
	MoveHorizontally(int)
	GetRow() int
	GetColumn() int
	GetWidth() int
	GetHeight() int
	GetSymbol() rune
	GetRowVelocity() int
	GetColVelocity() int
	SetRow(int)
	SetCol(int)
	SetRowVelocity(int)
	SetColVelocity(int)
	Move()
	NextRow() int
	NextColumn() int
}

func collidesWhenMovingLeft(obj1 IObject, obj2 IObject) bool {
	return obj1.GetColumn() > obj2.GetColumn() && obj1.NextColumn() <= obj2.GetColumn()
}

func collidesWhenMovingRight(obj1 IObject, obj2 IObject) bool {
	return obj1.GetColumn() < obj2.GetColumn() && obj1.NextColumn() >= obj2.GetColumn()
}

func collidesMovingUpOrDown(obj1 IObject, obj2 IObject) bool {
	return obj1.NextRow() >= obj2.GetRow() && obj1.NextRow() <= obj2.GetRow()+obj2.GetHeight()
}

func collidesMovingLeftOrRight(obj1 IObject, obj2 IObject) bool {
	return collidesWhenMovingLeft(obj1, obj2) || collidesWhenMovingRight(obj1, obj2)
}

func collidesWith(obj1 IObject, obj2 IObject) bool {
	return collidesMovingUpOrDown(obj1, obj2) && collidesMovingLeftOrRight(obj1, obj2)
}
