package main

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func drawMatrix(screen *ebiten.Image, matrix *Matrix, clr color.Color) {
	for i := 0; i < matrix.dimensions.X; i = i + 1 {
		for j := 0; j < matrix.dimensions.Y; j = j + 1 {
			drawBox(screen, coordinates.NewScreen(i*boxSize, j*boxSize), clr)
		}
	}
}

func drawStates(screen *ebiten.Image, matrix *Matrix, colorX, colorO color.Color) {
	for x := 0; x < matrix.dimensions.X; x = x + 1 {
		for y := 0; y < matrix.dimensions.Y; y = y + 1 {
			switch matrix.fields[x][y] {
			case fieldStatePlayerX:
				drawCross(screen, coordinates.NewScreen(x*boxSize, y*boxSize), colorX)
			case fieldStatePlayerO:
			}
		}
	}
}

func drawBox(screen *ebiten.Image, coordinates *coordinates.Screen, clr color.Color) {
	// draw top horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y),
		float64(coordinates.X+boxSize),
		float64(coordinates.Y),
		clr)

	// draw right vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.X+boxSize),
		float64(coordinates.Y),
		float64(coordinates.X+boxSize),
		float64(coordinates.Y+boxSize),
		clr)

	// draw left vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.X+1),
		float64(coordinates.Y),
		float64(coordinates.X+1),
		float64(coordinates.Y+boxSize),
		clr)

	// draw bottom horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y+boxSize),
		float64(coordinates.X+boxSize),
		float64(coordinates.Y+boxSize),
		clr)
}

func drawCross(screen *ebiten.Image, coordinates *coordinates.Screen, clr color.Color) {
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y),
		float64(coordinates.X+boxSize),
		float64(coordinates.Y+boxSize),
		clr)

	ebitenutil.DrawLine(screen,
		float64(coordinates.X+boxSize),
		float64(coordinates.Y),
		float64(coordinates.X),
		float64(coordinates.Y+boxSize),
		clr)
}
