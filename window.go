package senses

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Window struct {
	Title     string
	Size      *Size
	Container *Container
	Theme     *Theme
}

func (window *Window) SetContainer(container *Container) {
	container.position.Update(0, 0)
	container.size.Update(window.Size.Width, window.Size.Height)
	window.Container = container
}

func (window *Window) Update(screen *ebiten.Image) error {
	return nil
}

func (window *Window) Draw(screen *ebiten.Image) {
	window.Container.Draw(screen)
}

func (window *Window) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return window.Size.Width, window.Size.Height
}

func (window *Window) Show() {
	ebiten.SetWindowSize(window.Size.Width, window.Size.Height)
	ebiten.SetWindowTitle(window.Title)
	if err := ebiten.RunGame(window); err != nil {
		log.Fatal(err)
	}
}

func NewWindow(title string, width int, height int) *Window {
	window := new(Window)
	window.Title = title
	window.Size = NewSize(width, height)
	backgroundColor := &color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	foregroundColor := &color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	window.Theme = NewTheme(LoadFont("./resources/fonts", "fnt.json"), backgroundColor, foregroundColor)
	return window
}
