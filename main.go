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

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480

	regularGameDimensionX = 3
	regularGameDimensionY = 3

	boxSize = 150
)

var gameState = NewGameState(regularGameDimensionX, regularGameDimensionY)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	gameState.drawMatrix(screen)

	gameState.drawCursor(screen)

	gameState.handleKeyPress()

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "TripsTrapsTrull Shapes (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
