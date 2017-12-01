package draw

import (
	"image/color"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/data"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func Naught(screen *ebiten.Image, coordinates *coordinates.Screen, clr color.Color) {
	var adjustment int = data.BoxSize / 4

	var startX = adjustment + coordinates.X
	var startY = adjustment + coordinates.Y

	ebitenutil.DrawLine(screen,
		float64(startX),
		float64(startY),
		float64(startX+(data.BoxSize/2)),
		float64(startY),
		clr)

	ebitenutil.DrawLine(screen,
		float64(startX+(data.BoxSize/2)),
		float64(startY),
		float64(startX+(data.BoxSize/2)),
		float64(startY+(data.BoxSize/2)),
		clr)
	ebitenutil.DrawLine(screen,
		float64(startX+1),
		float64(startY),
		float64(startX+1),
		float64(startY+(data.BoxSize/2)),
		clr)

	ebitenutil.DrawLine(screen,
		float64(startX),
		float64(startY+(data.BoxSize/2)),
		float64(startX+(data.BoxSize/2)),
		float64(startY+(data.BoxSize/2)),
		clr)
}
