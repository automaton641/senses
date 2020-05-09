package senses

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type SymbolGame struct {
}

const symbolSize int = 12

var pixels []byte
var random *rand.Rand
var clr color.RGBA = color.RGBA{R: 0, G: 0, B: 0, A: 255}
var units [][]bool

func SaveSymbol() {
	var img *image.RGBA = image.NewRGBA(image.Rect(0, 0, symbolSize, symbolSize))
	for y := 0; y < symbolSize; y++ {
		for x := 0; x < symbolSize; x++ {
			if units[y][x] {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}
	file, err := os.Create("symbol.png")
	check(err)
	png.Encode(file, img)
	file.Close()
	fmt.Println("symbol saved")
}

func (g *SymbolGame) Update(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		units[y][x] = !units[y][x]
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		SaveSymbol()
	}
	return nil
}

func (g *SymbolGame) Draw(screen *ebiten.Image) {
	for y := 0; y < screen.Bounds().Dy(); y++ {
		for x := 0; x < screen.Bounds().Dx(); x++ {
			if units[y][x] {
				screen.Set(x, y, color.White)
			} else {
				screen.Set(x, y, color.Black)
			}
		}
	}
}

func (g *SymbolGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return symbolSize, symbolSize
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SymbolMain() {
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)
	units = make([][]bool, symbolSize)
	for i := range units {
		units[i] = make([]bool, symbolSize)
	}
	ebiten.SetWindowSize(512, 512)
	ebiten.SetWindowTitle("Symbol")
	if err := ebiten.RunGame(&SymbolGame{}); err != nil {
		log.Fatal(err)
	}
}
