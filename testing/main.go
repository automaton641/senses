package main

import (
	"fmt"

	"github.com/automaton641/senses"
)

func main() {
	fmt.Println("Hello senses")
	window := senses.NewWindow("Senses", 1280, 720)
	container := senses.NewContainer(senses.Vertical, window.Theme)
	text := senses.NewTextVisual("<{HELLO}, [SENSES]>;", window.Theme)
	window.SetContainer(container)
	container.Add(text)
	window.Show()
}
