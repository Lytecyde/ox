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
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
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

	keyAt = time.Time{}
)

func moveCursor(coordinates *Coordinates) {
	if time.Now().Sub(keyAt).Seconds() < 0.2 {
		return
	}

	if coordinates.x < 0 {
		return
	}

	if coordinates.x >= regularGameDimensionX {
		return
	}

	if coordinates.y < 0 {
		return
	}

	if coordinates.y >= regularGameDimensionY {
		return
	}

	cursor = coordinates

	keyAt = time.Now()
}

func chooseBox() {
	// check if box is not taken yet in matrix

	// mark box as taken in matrix
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	drawMatrix(screen, matrix, gray)

	drawBox(screen, NewCoordinates(cursor.x*boxSize, cursor.y*boxSize), red)

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		moveCursor(NewCoordinates(cursor.x, cursor.y-1))

	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		moveCursor(NewCoordinates(cursor.x, cursor.y+1))

	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		moveCursor(NewCoordinates(cursor.x-1, cursor.y))

	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		moveCursor(NewCoordinates(cursor.x+1, cursor.y))

	} else if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		chooseBox()

	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		chooseBox()

	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "TripsTrapsTrull Shapes (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
