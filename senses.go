package senses

import (
	"encoding/json"
	"fmt"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Size struct {
	Width  int
	Height int
}

type Position struct {
	X int
	Y int
}

type Symbol struct {
	Size  Size
	Units [][]bool
}

type Font struct {
	Name            string
	Symbols         map[string]*Symbol
	Source          string
	DescriptionPath string
	SymbolSize      int
}

func (font *Font) LoadSymbol(symbolPath string, symbolName string) {
	file, err := os.Open(symbolPath)
	check(err)
	img, err := png.Decode(file)
	check(err)
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	symbol := new(Symbol)
	symbol.Units = make([][]bool, height)
	for i := range symbol.Units {
		symbol.Units[i] = make([]bool, width)
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if (r+g+b)/3 < 128 {
				symbol.Units[y][x] = true
			} else {
				symbol.Units[y][x] = false
			}
		}
	}
	font.Symbols[symbolName] = symbol
}

func LoadFont(fontsDirectory, descriptionFileName string) *Font {
	descriptionPath := path.Join(fontsDirectory, descriptionFileName)
	fontFileData, err := ioutil.ReadFile(descriptionPath)
	check(err)
	var jsonData map[string]interface{}
	err = json.Unmarshal(fontFileData, &jsonData)
	check(err)
	fmt.Println(jsonData)
	font := new(Font)
	font.DescriptionPath = descriptionPath
	font.Name = jsonData["name"].(string)
	font.Source = path.Join(fontsDirectory, jsonData["source"].(string))
	font.SymbolSize = int(jsonData["symbolSize"].(float64))
	files, err := ioutil.ReadDir(font.Source)
	check(err)
	for _, file := range files {
		var fileName string = file.Name()
		tokens := strings.Split(fileName, ".")
		if len(tokens) >= 2 {
			if tokens[1] == "png" {
				font.LoadSymbol(path.Join(font.Source, fileName), tokens[0])
			}
		}
	}
	return font
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
