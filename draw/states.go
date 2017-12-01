package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/matrix"
	"github.com/Lytecyde/ox/player"
	"github.com/Lytecyde/ox/settings"
	"github.com/hajimehoshi/ebiten"
)

func States(screen *ebiten.Image, m *matrix.Matrix, colorX, colorO color.Color) {
	for x := 0; x < m.Dimensions.X; x = x + 1 {
		for y := 0; y < m.Dimensions.Y; y = y + 1 {
			c := coordinates.NewScreen(x*settings.BoxSize, y*settings.BoxSize)

			switch m.Fields[x][y] {
			case player.Cross:
				Cross(screen, c, settings.BoxSize, colorX)
			case player.Naught:
				Naught(screen, c, colorO)
			}
		}
	}
}
