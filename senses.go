package senses

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

func DrawRectangle(image *ebiten.Image, position *Position, size *Size, color *color.NRGBA) {
	y0 := position.Y
	y1 := y0 + size.Height
	x0 := position.X
	x1 := x0 + size.Width
	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			image.Set(x, y, *color)
		}
	}
}
