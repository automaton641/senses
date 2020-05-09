package main

import (
	"fmt"

	"github.com/automaton641/senses"
)

func main() {
	fmt.Println("Loading font")
	senses.LoadFont("./resources/fonts", "fnt.json")
}
