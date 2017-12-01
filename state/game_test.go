package state

import (
	"testing"
	"time"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/stretchr/testify/assert"
)

func Test_NewGameState_SetsMatrixSizeFromParameters_InCaseOfSuccess(t *testing.T) {
	// Arrange
	x, y := 100, 100

	// Act
	gameState := NewGame(x, y)

	// Assert
	assert.Equal(t, x, gameState.Matrix.Dimensions.X)
	assert.Equal(t, y, gameState.Matrix.Dimensions.Y)
}

func Test_moveCursor_MovesCursorToGivenCoordinates_InCaseOfSuccess(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	c := coordinates.NewMatrix(2, 2)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, c.X, gameState.Cursor.X)
	assert.Equal(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_ReturnsError_InCaseOfNegativeX(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	c := coordinates.NewMatrix(-1, 2)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "x less than 0", err.Error())
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_ReturnsError_InCaseOfOffScreenX(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	c := coordinates.NewMatrix(101, 2)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "out of the matrix", err.Error())
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_ReturnsError_InCaseOfNegativeY(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	c := coordinates.NewMatrix(1, -1)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "y less than 0", err.Error())
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_ReturnsError_InCaseOfOffScreenY(t *testing.T) {
	// Arrange
	gameState := NewGame(100, 100)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	c := coordinates.NewMatrix(1, 101)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "out of the matrix", err.Error())
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_MovesCursor_InCaseOfSuccess(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	c := coordinates.NewMatrix(5, 5)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, c.X, gameState.Cursor.X)
	assert.Equal(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursor_ReturnsError_InCaseOfTimeIsOutOfHumanPerceptionLimit(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	duration := time.Duration(100 * time.Millisecond)
	gameState.KeyAt = time.Now().Add(-duration)
	c := coordinates.NewMatrix(5, 5)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "less than human reaction time limit", err.Error())
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}
