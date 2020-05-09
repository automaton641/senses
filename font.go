package senses

import (
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type Font struct {
	Name            string
	Symbols         map[rune]*Symbol
	Source          string
	DescriptionPath string
	SymbolSize      int
	SizeFactor      int
	RealSize        int
}

func (font *Font) LoadSymbol(symbolPath string, symbolName string) {
	symbolRune, err := strconv.ParseInt(symbolName, 16, 32)
	check(err)
	file, err := os.Open(symbolPath)
	check(err)
	img, err := png.Decode(file)
	check(err)
	width := img.Bounds().Dx()
	if width < font.SymbolSize {
		panic("width < font.SymbolSize")
	}
	height := img.Bounds().Dy()
	if height < font.SymbolSize {
		panic("height < font.SymbolSize")
	}
	symbol := new(Symbol)
	symbol.Units = make([][]bool, font.SymbolSize)
	for i := range symbol.Units {
		symbol.Units[i] = make([]bool, font.SymbolSize)
	}
	for y := 0; y < font.SymbolSize; y++ {
		for x := 0; x < font.SymbolSize; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if (r+g+b)/3 < 128 {
				symbol.Units[y][x] = false
			} else {
				symbol.Units[y][x] = true
			}
		}
	}
	font.Symbols[rune(symbolRune)] = symbol
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
	font.SizeFactor = 2
	font.DescriptionPath = descriptionPath
	font.Name = jsonData["name"].(string)
	font.Source = path.Join(fontsDirectory, jsonData["source"].(string))
	font.SymbolSize = int(jsonData["symbolSize"].(float64))
	font.RealSize = font.SizeFactor * font.SymbolSize
	font.Symbols = make(map[rune]*Symbol)
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
