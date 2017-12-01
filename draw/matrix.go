package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/matrix"
	"github.com/Lytecyde/ox/settings"
	"github.com/hajimehoshi/ebiten"
)

func Matrix(screen *ebiten.Image, m *matrix.Matrix, clr color.Color) {
	for i := 0; i < m.Dimensions.X; i = i + 1 {
		for j := 0; j < m.Dimensions.Y; j = j + 1 {
			c := coordinates.NewScreen(i*settings.BoxSize, j*settings.BoxSize)

			Box(screen, c, settings.BoxSize, clr)
		}
	}
}
