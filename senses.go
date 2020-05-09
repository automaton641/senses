package senses

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
)

type Size struct {
	Width  int
	Height int
}

type Position struct {
	X int
	Y int
}

type Glyph struct {
	Size  Size
	Units [][]bool
}

type Font struct {
	Name   string
	Glyphs map[rune]*Glyph
}

func LoadFont(path string) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)

	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

type Theme struct {
	BacgroundColor  *color.Color
	ForegroundColor *color.Color
	Font            Font
}

type Visual struct {
	Proportion float64
	GrowRatio  float64
}

type Window struct {
	Position *Position
	Size     *Size
}
