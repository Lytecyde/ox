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

// +build example

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

type Coordinates struct {
	x int
	y int
}

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	c := newCoordinates(1,1)
	drawBox(screen, c)
	
	return nil
}

func newCoordinates(x int, y int) Coordinates {
	return Coordinates{
		x:x,
		y:y,	
	}
}

func drawBox(screen *ebiten.Image, coordinates Coordinates){
	ebitenutil.DrawRect(screen, float64(50*coordinates.x), float64(50*coordinates.y), 100, 100, color.RGBA{0x80, 0x80, 0x80, 0x80})
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Shapes (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
