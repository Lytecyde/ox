package matrix

import (
	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/player"
)

// Matrix represents a matrix on screen
type Matrix struct {
	Dimensions coordinates.Matrix
	Fields     [][]player.Type
}

func NewMatrix(dimensionx int, dimensiony int) *Matrix {
	var m Matrix
	m.Dimensions.X = dimensionx
	m.Dimensions.Y = dimensiony

	//init fields
	m.Fields = make([][]player.Type, dimensionx)
	for i := 0; i < dimensionx; i = i + 1 {
		m.Fields[i] = make([]player.Type, dimensiony)
	}

	return &m
}

func (matrix *Matrix) SetState(c coordinates.Matrix, newState player.Type) {
	matrix.Fields[c.X][c.Y] = newState
}

func (matrix *Matrix) State(c coordinates.Matrix) player.Type {
	return matrix.Fields[c.X][c.Y]
}
