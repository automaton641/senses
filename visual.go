package senses

import "github.com/hajimehoshi/ebiten"

type Visual struct {
	position   *Position
	size       *Size
	Theme      *Theme
	Proportion float64
	growRatio  float64
}

func NewVisual(theme *Theme) *Visual {
	visual := new(Visual)
	visual.size = NewSize(1, 1)
	visual.position = NewPosition(0, 0)
	visual.Theme = theme
	visual.growRatio = 1
	visual.Proportion = 1
	return visual
}

func (visual *Visual) Arrange() {
}

func (visual *Visual) Draw(screen *ebiten.Image) {
	DrawRectangle(screen, visual.position, visual.size, visual.Theme.BackgroundColor)
}

func (visual *Visual) GrowRatio() float64 {
	return visual.growRatio
}

func (visual *Visual) Position() *Position {
	return visual.position
}

func (visual *Visual) Size() *Size {
	return visual.size
}
