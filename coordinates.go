package main

type Coordinates struct {
	x int
	y int
}

func NewCoordinates(x int, y int) *Coordinates {
	return &Coordinates{
		x: x,
		y: y,
	}
}
