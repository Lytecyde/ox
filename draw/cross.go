package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func Cross(screen *ebiten.Image, coordinates *coordinates.Screen, boxSize int, clr color.Color) {
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y),
		float64(coordinates.X+boxSize),
		float64(coordinates.Y+boxSize),
		clr)

	ebitenutil.DrawLine(screen,
		float64(coordinates.X+boxSize),
		float64(coordinates.Y),
		float64(coordinates.X),
		float64(coordinates.Y+boxSize),
		clr)
}
