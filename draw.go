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

func drawBox(screen *ebiten.Image, coordinates *Coordinates, clr color.Color) {
	// draw up horizontal
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

	// draw down horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y+boxSize),
		float64(coordinates.x+boxSize),
		float64(coordinates.y+boxSize),
		clr)
}

func drawCross(screen *ebiten.Image, coordinates *Coordinates, clr color.Color) {
	ebitenutil.DrawLine(screen, float64(coordinates.x), float64(coordinates.y),
		float64(boxSize), float64(boxSize), clr)
	ebitenutil.DrawLine(screen, float64(boxSize), float64(coordinates.y),
		float64(coordinates.x), float64(boxSize), clr)
}
