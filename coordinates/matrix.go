package coordinates

// Matrix represents coordinates in matrix
type Matrix struct {
	coordinates
}

// NewMatrix returns instance
func NewMatrix(x int, y int) *Matrix {
	c := coordinates{
		X: x,
		Y: y,
	}

	return &Matrix{c}
}
