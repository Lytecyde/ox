package main

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/data"
	"github.com/Lytecyde/ox/player"
	"github.com/Lytecyde/ox/playfield"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func drawMatrix(screen *ebiten.Image, matrix *playfield.Matrix, clr color.Color) {
	for i := 0; i < matrix.Dimensions.X; i = i + 1 {
		for j := 0; j < matrix.Dimensions.Y; j = j + 1 {
			drawBox(screen, coordinates.NewScreen(i*data.BoxSize, j*data.BoxSize), clr)
		}
	}
}

func drawStates(screen *ebiten.Image, matrix *playfield.Matrix, colorX, colorO color.Color) {
	for x := 0; x < matrix.Dimensions.X; x = x + 1 {
		for y := 0; y < matrix.Dimensions.Y; y = y + 1 {
			switch matrix.Fields[x][y] {
			case player.Cross:
				drawCross(screen, coordinates.NewScreen(x*data.BoxSize, y*data.BoxSize), colorX)
			case player.Naught:
				drawNaught(screen, coordinates.NewScreen(x*data.BoxSize, y*data.BoxSize), colorO)
			}
		}
	}
}

func drawBox(screen *ebiten.Image, coordinates *coordinates.Screen, clr color.Color) {
	// draw top horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y),
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y),
		clr)

	// draw right vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y),
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y+data.BoxSize),
		clr)

	// draw left vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.X+1),
		float64(coordinates.Y),
		float64(coordinates.X+1),
		float64(coordinates.Y+data.BoxSize),
		clr)

	// draw bottom horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y+data.BoxSize),
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y+data.BoxSize),
		clr)
}

func drawCross(screen *ebiten.Image, coordinates *coordinates.Screen, clr color.Color) {
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y),
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y+data.BoxSize),
		clr)

	ebitenutil.DrawLine(screen,
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y),
		float64(coordinates.X),
		float64(coordinates.Y+data.BoxSize),
		clr)
}

func drawNaught(screen *ebiten.Image, coordinates *coordinates.Screen, clr color.Color) {
	var adjustment int = data.BoxSize / 4

	var startX = adjustment + coordinates.X
	var startY = adjustment + coordinates.Y

	ebitenutil.DrawLine(screen,
		float64(startX),
		float64(startY),
		float64(startX+(data.BoxSize/2)),
		float64(startY),
		clr)

	ebitenutil.DrawLine(screen,
		float64(startX+(data.BoxSize/2)),
		float64(startY),
		float64(startX+(data.BoxSize/2)),
		float64(startY+(data.BoxSize/2)),
		clr)
	ebitenutil.DrawLine(screen,
		float64(startX+1),
		float64(startY),
		float64(startX+1),
		float64(startY+(data.BoxSize/2)),
		clr)

	ebitenutil.DrawLine(screen,
		float64(startX),
		float64(startY+(data.BoxSize/2)),
		float64(startX+(data.BoxSize/2)),
		float64(startY+(data.BoxSize/2)),
		clr)
}
