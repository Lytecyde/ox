package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/settings"
	"github.com/hajimehoshi/ebiten"
)

func Naught(screen *ebiten.Image, c *coordinates.Screen, clr color.Color) {
	adjustment := settings.BoxSize / 4
	width := settings.BoxSize / 2
	newC := coordinates.NewScreen(c.X+adjustment, c.Y+adjustment)

	Box(screen, newC, width, clr)
}
