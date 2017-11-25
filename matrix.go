package main

import (
	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/player"
)

// Matrix represents a matrix on screen
type Matrix struct {
	dimensions coordinates.Matrix
	fields     [][]player.Type
}

// NewMatrix returns instance
func NewMatrix(dimensionx int, dimensiony int) *Matrix {
	var m Matrix
	m.dimensions.X = dimensionx
	m.dimensions.Y = dimensiony

	//init fields
	m.fields = make([][]player.Type, dimensionx)
	for i := 0; i < dimensionx; i = i + 1 {
		m.fields[i] = make([]player.Type, dimensiony)
	}

	return &m
}

func (matrix *Matrix) setState(c coordinates.Matrix, newState player.Type) {
	matrix.fields[c.X][c.Y] = newState
}

func (matrix *Matrix) state(c coordinates.Matrix) player.Type {
	return matrix.fields[c.X][c.Y]
}
