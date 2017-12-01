package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/data"
	"github.com/Lytecyde/ox/playfield"
	"github.com/hajimehoshi/ebiten"
)

func Matrix(screen *ebiten.Image, matrix *playfield.Matrix, clr color.Color) {
	for i := 0; i < matrix.Dimensions.X; i = i + 1 {
		for j := 0; j < matrix.Dimensions.Y; j = j + 1 {
			Box(screen, coordinates.NewScreen(i*data.BoxSize, j*data.BoxSize), data.BoxSize, clr)
		}
	}
}
