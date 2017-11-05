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

var matrix *Matrix = NewMatrix(regularGameDimensionX, regularGameDimensionY)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	drawMatrix(screen, matrix)

	return nil
}

func drawMatrix(screen *ebiten.Image, matrix *Matrix) {
	var x int = 1
	var y int = 1

	for i := 1; i < regularGameDimensionX; i = i + 1 {
		x = i
		for j := 1; j < regularGameDimensionY; j = j + 1 {
			y = j
			c := NewCoordinates(x, y)
			drawBox(screen, c)
		}

	}

}

func drawBox(screen *ebiten.Image, coordinates *Coordinates) {
	ebitenutil.DrawRect(screen,
		50,
		50,
		float64(50*coordinates.x),
		float64(50*coordinates.y),
		color.RGBA{0x80, 0x80, 0x80, 0x80})
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "TripsTrapsTrull Shapes (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
