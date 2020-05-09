package senses

type Size struct {
	Width  int
	Height int
}

func NewSize(width, height int) *Size {
	size := new(Size)
	size.Width = width
	size.Height = height
	return size
}

func (size *Size) Update(width, height int) {
	size.Width = width
	size.Height = height
}

func (size *Size) UpdateFrom(other *Size) {
	size.Width = other.Width
	size.Height = other.Height
}
