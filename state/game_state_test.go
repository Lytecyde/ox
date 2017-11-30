package state

import (
	"testing"
	"time"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/data"
	"github.com/stretchr/testify/assert"
)

func Test_NewGameState_SetsMatrixSizeFromParameters(t *testing.T) {
	// Arrange
	x, y := 100, 100

	// Act
	gameState := NewGame(x, y)

	// Assert
	assert.Equal(t, x, gameState.Matrix.Dimensions.X)
	assert.Equal(t, y, gameState.Matrix.Dimensions.Y)
}

func Test_moveCursor_MovesCursorToGivenCoordinates(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	gameState.KeyAt = time.Now().Add(-(data.HumanReactionSec * 1000) * time.Millisecond)
	c := coordinates.NewMatrix(2, 2)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, c.X, gameState.Cursor.X)
	assert.Equal(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfNegativeX(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	c := coordinates.NewMatrix(-1, 2)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfOffScreenX(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	c := coordinates.NewMatrix(101, 2)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfNegativeY(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	c := coordinates.NewMatrix(1, -1)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfOffScreenY(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	c := coordinates.NewMatrix(1, 101)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_InCaseOfTimeIsWithinHumanPerceptionLimit(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	duration := time.Duration(1 * time.Second)
	gameState.KeyAt = time.Now().Add(-duration)
	c := coordinates.NewMatrix(5, 5)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.Equal(t, c.X, gameState.Cursor.X)
	assert.Equal(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_InCaseOfTimeIsOutOfHumanPerceptionLimit(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	duration := time.Duration(100 * time.Millisecond)
	gameState.KeyAt = time.Now().Add(-duration)
	c := coordinates.NewMatrix(5, 5)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}
