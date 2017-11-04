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

type Coordinates struct {
	x int
	y int
}

const regularGameDimensionX = 3
const regularGameDimensionY = 3

type Matrix struct {
	dimensionx int
	dimensiony int
	fields     [][]int
}

var matrix Matrix = newMatrix(regularGameDimensionX, regularGameDimensionY)

func update(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		return nil
	}

	drawMatrix(screen, matrix)

	return nil
}

func drawMatrix(screen *ebiten.Image, matrix Matrix) {
	var x int = 1
	var y int = 1

	for i:=1;i<regularGameDimensionX;i = i + 1 {
		x = i
		for j:=1;j<regularGameDimensionY;j = j +1 {
			y = j
			c := newCoordinates(x, y) 	
			drawBox(screen, c)
		}
	
	}
	
}

func newCoordinates(x int, y int) Coordinates {
	return Coordinates{
		x: x,
		y: y,
	}
}

func newMatrix(dimensionx int, dimensiony int) Matrix {
	var m Matrix
	m.dimensionx = dimensionx
	m.dimensiony = dimensiony
	//init fields
	m.fields = make([][]int, dimensionx)
	for i := 0; i < dimensionx; i = i + 1 {
		m.fields[i] = make([]int, dimensiony)
	}

	return m
}

func drawBox(screen *ebiten.Image, coordinates Coordinates) {
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
