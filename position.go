package senses

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) *Position {
	position := new(Position)
	position.X = x
	position.Y = y
	return position
}

func (position *Position) Update(x, y int) {
	position.X = x
	position.Y = y
}

func (position *Position) UpdateFrom(other *Position) {
	position.X = other.X
	position.Y = other.Y
}
