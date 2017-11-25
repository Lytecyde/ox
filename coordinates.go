package main

import (
	"fmt"
)

// Coordinates represents coordinates on screen
type Coordinates struct {
	x int
	y int
}

// NewCoordinates returns instance
func NewCoordinates(x int, y int) *Coordinates {
	return &Coordinates{
		x: x,
		y: y,
	}
}

func (coordinates Coordinates) Equal(other Coordinates) bool {
	return coordinates.x == other.x && coordinates.y == other.y
}

func (coordinates Coordinates) String() string {
	return fmt.Sprintf("x=%d, y=%d", coordinates.x, coordinates.y)
}
