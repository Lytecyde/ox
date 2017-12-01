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
	"github.com/Lytecyde/ox/data"
	"github.com/Lytecyde/ox/draw"
	"github.com/Lytecyde/ox/state"
	"github.com/hajimehoshi/ebiten"
)

var gameState = state.NewGame(data.RegularGameDimensionX, data.RegularGameDimensionY)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	draw.Matrix(screen, gameState.Matrix, gray)

	// draw cursor
	draw.Box(screen, coordinates.NewScreen(gameState.Cursor.X*data.BoxSize, gameState.Cursor.Y*data.BoxSize), data.BoxSize, red)

	draw.States(screen, gameState.Matrix, blue, green)

	//gameloop
	if !gameState.Finished {
		gameState.HandleKeyPress()
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, data.ScreenWidth, data.ScreenHeight, 1, "TripsTrapsTrull Shapes (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
