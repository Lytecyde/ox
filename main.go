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
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480

	regularGameDimensionX = 3
	regularGameDimensionY = 3

	boxSize = 150
)

var (
	matrix = NewMatrix(regularGameDimensionX, regularGameDimensionY)

	cursor = NewCoordinates(0, 0)

	gray  = color.RGBA{0x80, 0x80, 0x80, 0x80}
	red   = color.RGBA{0x80, 0x0, 0x0, 0x80}
	green = color.RGBA{0x0, 0x80, 0x0, 0x80}

	keyAt = time.Time{}
)

func moveCursor(coordinates *Coordinates) {
	if time.Now().Sub(keyAt).Seconds() < 0.5 {
		return
	}

	if coordinates.x < 0 {
		return
	}

	if coordinates.x >= regularGameDimensionX*boxSize {
		return
	}

	if coordinates.y < 0 {
		return
	}

	if coordinates.y >= regularGameDimensionY*boxSize {
		return
	}

	cursor = coordinates

	keyAt = time.Now()
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	drawMatrix(screen, matrix, gray)

	drawBox(screen, cursor, red)

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		moveCursor(NewCoordinates(cursor.x, cursor.y-boxSize))

	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		moveCursor(NewCoordinates(cursor.x, cursor.y+boxSize))

	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		moveCursor(NewCoordinates(cursor.x-boxSize, cursor.y))

	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		moveCursor(NewCoordinates(cursor.x+boxSize, cursor.y))
	}

	return nil
}

func drawMatrix(screen *ebiten.Image, matrix *Matrix, clr color.Color) {
	for i := 0; i < regularGameDimensionX; i = i + 1 {
		for j := 0; j < regularGameDimensionY; j = j + 1 {
			drawBox(screen, NewCoordinates(i*boxSize, j*boxSize), clr)
		}

	}
}

func drawBox(screen *ebiten.Image, coordinates *Coordinates, clr color.Color) {
	// draw up horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y),
		float64(coordinates.x+boxSize),
		float64(coordinates.y),
		clr)

	// draw right vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.x+boxSize),
		float64(coordinates.y),
		float64(coordinates.x+boxSize),
		float64(coordinates.y+boxSize),
		clr)

	// draw left vertical
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y),
		float64(coordinates.x),
		float64(coordinates.y+boxSize),
		clr)

	// draw down horizontal
	ebitenutil.DrawLine(screen,
		float64(coordinates.x),
		float64(coordinates.y+boxSize),
		float64(coordinates.x+boxSize),
		float64(coordinates.y+boxSize),
		clr)
}

func drawCross(screen *ebiten.Image, coordinates *Coordinates, clr color.Color) {
	ebitenutil.DrawLine(screen, float64(coordinates.x), float64(coordinates.y),
		float64(boxSize), float64(boxSize), clr)
	ebitenutil.DrawLine(screen, float64(boxSize), float64(coordinates.y),
		float64(coordinates.x), float64(boxSize), clr)
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "TripsTrapsTrull Shapes (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
