package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/data"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func Box(screen *ebiten.Image, coordinates *coordinates.Screen, clr color.Color) {
	// draw top horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y),
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y),
		clr)

	// draw right vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y),
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y+data.BoxSize),
		clr)

	// draw left vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.X+1),
		float64(coordinates.Y),
		float64(coordinates.X+1),
		float64(coordinates.Y+data.BoxSize),
		clr)

	// draw bottom horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.X),
		float64(coordinates.Y+data.BoxSize),
		float64(coordinates.X+data.BoxSize),
		float64(coordinates.Y+data.BoxSize),
		clr)
}
