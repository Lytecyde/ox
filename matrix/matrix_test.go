package matrix

import (
	"testing"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/player"
	"github.com/stretchr/testify/assert"
)

func Test_NewMatrix_ReturnsInstance_InCaseOfSuccess(t *testing.T) {
	// Arrange
	x, y := 2, 3

	// Act
	m := NewMatrix(x, y)

	// Assert
	assert.Equal(t, x, m.Dimensions.X)
	assert.Equal(t, y, m.Dimensions.Y)
}

func Test_SetState_SetsState_InCaseOfSuccess(t *testing.T) {
	// Arrange
	m := NewMatrix(10, 10)
	c := coordinates.NewMatrix(5, 5)

	// Act
	m.SetState(*c, player.Cross)

	// Assert
	assert.Equal(t, player.Cross, m.State(*c))

}

func Test_State_ReturnsState_InCaseOfSuccess(t *testing.T) {
	// Arrange
	m := NewMatrix(10, 10)
	c := coordinates.NewMatrix(5, 5)
	m.Fields[c.X][c.Y] = player.Naught

	// Act
	state := m.State(*c)

	// Assert
	assert.Equal(t, player.Naught, state)
}

func Test_Load_LoadsState_InCaseOfSuccess(t *testing.T) {
	// Arrange
	m := NewMatrix(2, 2)

	// Act
	m.Load([]player.Type{
		player.None, player.Naught,
		player.Cross, player.None,
	})

	// Assert
	assert.Equal(t, player.None, m.Fields[0][0])
	assert.Equal(t, player.Naught, m.Fields[0][1])
	assert.Equal(t, player.Cross, m.Fields[1][0])
	assert.Equal(t, player.None, m.Fields[1][1])
}
