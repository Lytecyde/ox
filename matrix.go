package main

// Matrix represents a matrix on screen
type Matrix struct {
	dimensions Coordinates
	fields     [][]fieldState
}

// NewMatrix returns instance
func NewMatrix(dimensionx int, dimensiony int) *Matrix {
	var m Matrix
	m.dimensions.x = dimensionx
	m.dimensions.y = dimensiony

	//init fields
	m.fields = make([][]fieldState, dimensionx)
	for i := 0; i < dimensionx; i = i + 1 {
		m.fields[i] = make([]fieldState, dimensiony)
	}

	return &m
}

func (matrix *Matrix) setState(coordinates Coordinates, newState fieldState) {
	println(coordinates.String(), newState)
	matrix.fields[coordinates.x][coordinates.y] = newState
}
