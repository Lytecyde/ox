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

func (m *Matrix) SetState(c coordinates.Matrix, newState player.Type) {
	m.Fields[c.X][c.Y] = newState
}

func (m *Matrix) State(c coordinates.Matrix) player.Type {
	return m.Fields[c.X][c.Y]
}

func (m *Matrix) Load(states []player.Type) {
	row, col := 0, 0
	for _, state := range states {
		m.Fields[row][col] = state
		col++
		if col >= len(m.Fields[row]) {
			row++
			col = 0
		}
	}
}
