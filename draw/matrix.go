package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/playfield"
	"github.com/Lytecyde/ox/settings"
	"github.com/hajimehoshi/ebiten"
)

func Matrix(screen *ebiten.Image, matrix *playfield.Matrix, clr color.Color) {
	for i := 0; i < matrix.Dimensions.X; i = i + 1 {
		for j := 0; j < matrix.Dimensions.Y; j = j + 1 {
			c := coordinates.NewScreen(i*settings.BoxSize, j*settings.BoxSize)

			Box(screen, c, settings.BoxSize, clr)
		}
	}
}
