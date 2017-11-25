package main

import (
	"github.com/Lytecyde/ox/coordinates"
)

// Matrix represents a matrix on screen
type Matrix struct {
	dimensions coordinates.Matrix
	fields     [][]fieldState
}

// NewMatrix returns instance
func NewMatrix(dimensionx int, dimensiony int) *Matrix {
	var m Matrix
	m.dimensions.X = dimensionx
	m.dimensions.Y = dimensiony

	//init fields
	m.fields = make([][]fieldState, dimensionx)
	for i := 0; i < dimensionx; i = i + 1 {
		m.fields[i] = make([]fieldState, dimensiony)
	}

	return &m
}

func (matrix *Matrix) setState(c coordinates.Matrix, newState fieldState) {
	matrix.fields[c.X][c.Y] = newState
}
