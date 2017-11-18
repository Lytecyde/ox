package main

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
