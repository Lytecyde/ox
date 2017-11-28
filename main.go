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

	"github.com/Lytecyde/ox/coordinates"
	"github.com/hajimehoshi/ebiten"
)

const (
	boxSize = 150

	regularGameDimensionX = 3
	regularGameDimensionY = 3

	borderMargin = 2

	screenWidth  = boxSize*regularGameDimensionX + borderMargin
	screenHeight = boxSize*regularGameDimensionY + borderMargin
)

var gameState = NewGameState(regularGameDimensionX, regularGameDimensionY)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	drawMatrix(screen, gameState.matrix, gray)

	// draw cursor
	drawBox(screen, coordinates.NewScreen(gameState.cursor.X*boxSize, gameState.cursor.Y*boxSize), red)

	drawStates(screen, gameState.matrix, blue, green)

	//gameloop
	if !gameState.EndOfGame {
		gameState.handleKeyPress()
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "TripsTrapsTrull Shapes (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
