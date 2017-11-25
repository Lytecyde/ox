package coordinates

// Screen represents coordinates in screen
type Screen struct {
	coordinates
}

// NewScreenCoordinates returns instance
func NewScreen(x int, y int) *Screen {
	c := coordinates{
		X: x,
		Y: y,
	}

	return &Screen{c}
}
