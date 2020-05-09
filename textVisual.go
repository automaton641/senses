package senses

import (
	"github.com/hajimehoshi/ebiten"
)

type TextVisual struct {
	*Visual
	Text string
}

func NewTextVisual(text string, theme *Theme) *TextVisual {
	textVisual := new(TextVisual)
	textVisual.Visual = NewVisual(theme)
	textVisual.Text = text
	return textVisual
}

func (textVisual *TextVisual) DrawRune(screen *ebiten.Image, r rune, x0, y0 int) {
	symbol := textVisual.Theme.Font.Symbols[r]
	if symbol == nil {
		panic("symbol == nil")
	}
	position := new(Position)
	size := NewSize(textVisual.Theme.Font.SizeFactor, textVisual.Theme.Font.SizeFactor)
	for y := y0; y < y0+textVisual.Theme.Font.RealSize; y += textVisual.Theme.Font.SizeFactor {
		for x := x0; x < x0+textVisual.Theme.Font.RealSize; x += textVisual.Theme.Font.SizeFactor {
			//fmt.Println(y, x)
			//fmt.Println((y-y0)/textVisual.Theme.Font.SizeFactor, (x-x0)/textVisual.Theme.Font.SizeFactor)
			if symbol.Units[(y-y0)/textVisual.Theme.Font.SizeFactor][(x-x0)/textVisual.Theme.Font.SizeFactor] {
				position.Update(x, y)
				DrawRectangle(screen, position, size, textVisual.Theme.ForegroundColor)
			}
		}
	}
}

func (textVisual *TextVisual) Draw(screen *ebiten.Image) {
	textVisual.Visual.Draw(screen)
	//fmt.Println("TextVisual drawing...")
	font := textVisual.Theme.Font
	text := textVisual.Text
	textWidth := len(text) * font.RealSize
	textHeight := font.SymbolSize * font.SizeFactor
	y0 := textVisual.position.Y + textVisual.size.Height/2 - textHeight/2
	x0 := textVisual.position.X + textVisual.size.Width/2 - textWidth/2
	for _, r := range text {
		textVisual.DrawRune(screen, r, x0, y0)
		x0 += font.RealSize
	}

}
