package main

// Matrix represents a matrix on screen
type Matrix struct {
	Coordinates
	fields [][]int
}

// NewMatrix returns instance
func NewMatrix(dimensionx int, dimensiony int) *Matrix {
	var m Matrix
	m.x = dimensionx
	m.y = dimensiony

	//init fields
	m.fields = make([][]int, dimensionx)
	for i := 0; i < dimensionx; i = i + 1 {
		m.fields[i] = make([]int, dimensiony)
	}

	return &m
}
