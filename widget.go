package senses

import "github.com/hajimehoshi/ebiten"

type Widget interface {
	Arrange()
	GrowRatio() float64
	Draw(image *ebiten.Image)
	Position() *Position
	Size() *Size
}
