package senses

import "image/color"

type Theme struct {
	BackgroundColor *color.NRGBA
	ForegroundColor *color.NRGBA
	Font            *Font
}

func NewTheme(font *Font, backgroundColor, foregroundColor *color.NRGBA) *Theme {
	theme := new(Theme)
	theme.BackgroundColor = backgroundColor
	theme.ForegroundColor = foregroundColor
	theme.Font = font
	return theme
}
