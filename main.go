// Copyright 2017 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

const regularGameDimensionX = 3
const regularGameDimensionY = 3

var matrix = NewMatrix(regularGameDimensionX, regularGameDimensionY)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	drawMatrix(screen, matrix)

	return nil
}

const boxSize = 150

func drawMatrix(screen *ebiten.Image, matrix *Matrix) {
	for i := 0; i < regularGameDimensionX; i = i + 1 {
		for j := 0; j < regularGameDimensionY; j = j + 1 {
			drawBox(screen, NewCoordinates(i*boxSize, j*boxSize))
		}

	}
	drawCross(screen, NewCoordinates(0, 0))

}

func drawBox(screen *ebiten.Image, coordinates *Coordinates) {
	// draw up horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y),
		float64(coordinates.x+boxSize),
		float64(coordinates.y),
		gray)

	// draw right vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.x+boxSize),
		float64(coordinates.y),
		float64(coordinates.x+boxSize),
		float64(coordinates.y+boxSize),
		gray)

	// draw left vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y),
		float64(coordinates.x),
		float64(coordinates.y+boxSize),
		gray)

	// draw down horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y+boxSize),
		float64(coordinates.x+boxSize),
		float64(coordinates.y+boxSize),
		gray)
}

var gray = color.RGBA{0x80, 0x80, 0x80, 0x80}
var red = color.RGBA{0x80, 0x0, 0x0, 0x80}

func drawCross(screen *ebiten.Image, coordinates *Coordinates) {
	ebitenutil.DrawLine(screen, float64(coordinates.x), float64(coordinates.y),
		float64(boxSize), float64(boxSize), red)
	ebitenutil.DrawLine(screen, float64(boxSize), float64(coordinates.y),
		float64(coordinates.x), float64(boxSize), red)
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "TripsTrapsTrull Shapes (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
