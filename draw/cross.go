package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/data"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func Cross(screen *ebiten.Image, coordinates *coordinates.Screen, clr color.Color) {
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y),
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y+data.BoxSize),
		clr)

	ebitenutil.DrawLine(screen,
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y),
		float64(coordinates.X),
		float64(coordinates.Y+data.BoxSize),
		clr)
}
