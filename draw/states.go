package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/data"
	"github.com/Lytecyde/ox/player"
	"github.com/Lytecyde/ox/playfield"
	"github.com/hajimehoshi/ebiten"
)

func States(screen *ebiten.Image, matrix *playfield.Matrix, colorX, colorO color.Color) {
	for x := 0; x < matrix.Dimensions.X; x = x + 1 {
		for y := 0; y < matrix.Dimensions.Y; y = y + 1 {
			switch matrix.Fields[x][y] {
			case player.Cross:
				Cross(screen, coordinates.NewScreen(x*data.BoxSize, y*data.BoxSize), colorX)
			case player.Naught:
				Naught(screen, coordinates.NewScreen(x*data.BoxSize, y*data.BoxSize), colorO)
			}
		}
	}
}
