package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func drawMatrix(screen *ebiten.Image, matrix *Matrix, clr color.Color) {
	for i := 0; i < regularGameDimensionX; i = i + 1 {
		for j := 0; j < regularGameDimensionY; j = j + 1 {
			drawBox(screen, NewCoordinates(i*boxSize, j*boxSize), clr)
		}
	}
}

func drawStates(screen *ebiten.Image, matrix *Matrix, colorX, colorO color.Color) {
	for x := 0; x < matrix.dimensions.x; x = x + 1 {
		for y := 0; y < matrix.dimensions.y; y = y + 1 {
			switch matrix.fields[x][y] {
			case fieldStatePlayerX:
				drawCross(screen, NewCoordinates(x*boxSize, y*boxSize), colorX)
			case fieldStatePlayerO:
			}
		}
	}
}

func drawBox(screen *ebiten.Image, coordinates *Coordinates, clr color.Color) {
	// draw top horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y),
		float64(coordinates.x+boxSize),
		float64(coordinates.y),
		clr)

	// draw right vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.x+boxSize),
		float64(coordinates.y),
		float64(coordinates.x+boxSize),
		float64(coordinates.y+boxSize),
		clr)

	// draw left vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.x+1),
		float64(coordinates.y),
		float64(coordinates.x+1),
		float64(coordinates.y+boxSize),
		clr)

	// draw bottom horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y+boxSize),
		float64(coordinates.x+boxSize),
		float64(coordinates.y+boxSize),
		clr)
}

func drawCross(screen *ebiten.Image, coordinates *Coordinates, clr color.Color) {
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y),
		float64(coordinates.x+boxSize),
		float64(coordinates.y+boxSize),
		clr)

	ebitenutil.DrawLine(screen,
		float64(coordinates.x+boxSize),
		float64(coordinates.y),
		float64(coordinates.x),
		float64(coordinates.y+boxSize),
		clr)
}
