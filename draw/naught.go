package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/data"
	"github.com/hajimehoshi/ebiten"
)

func Naught(screen *ebiten.Image, c *coordinates.Screen, clr color.Color) {
	adjustment := data.BoxSize / 4
	width := data.BoxSize / 2

	Box(screen, coordinates.NewScreen(c.X+adjustment, c.Y+adjustment), width, clr)
}
