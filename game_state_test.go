package main

import (
	"testing"
	"time"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/stretchr/testify/assert"
)

func Test_NewGameState_SetsMatrixSizeFromParameters(t *testing.T) {
	// Arrange
	x, y := 100, 100

	// Act
	gameState := NewGameState(x, y)

	// Assert
	assert.Equal(t, x, gameState.matrix.dimensions.X)
	assert.Equal(t, y, gameState.matrix.dimensions.Y)
}

func Test_moveCursor_MovesCursorToGivenCoordinates(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	c := coordinates.NewMatrix(2, 2)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.Equal(t, c.X, gameState.cursor.X)
	assert.Equal(t, c.Y, gameState.cursor.Y)
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfNegativeX(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	c := coordinates.NewMatrix(-1, 2)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.cursor.X)
	assert.NotEqual(t, c.Y, gameState.cursor.Y)
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfOffScreenX(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	c := coordinates.NewMatrix(101, 2)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.cursor.X)
	assert.NotEqual(t, c.Y, gameState.cursor.Y)
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfNegativeY(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	c := coordinates.NewMatrix(1, -1)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.cursor.X)
	assert.NotEqual(t, c.Y, gameState.cursor.Y)
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfOffScreenY(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	c := coordinates.NewMatrix(1, 101)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.cursor.X)
	assert.NotEqual(t, c.Y, gameState.cursor.Y)
}

func Test_moveCursor_InCaseOfTimeIsWithinHumanPerceptionLimit(t *testing.T) {
	// Arrange
	gameState := NewGameState(10, 10)
	duration := time.Duration(1 * time.Second)
	gameState.keyAt = time.Now().Add(-duration)
	c := coordinates.NewMatrix(5, 5)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.Equal(t, c.X, gameState.cursor.X)
	assert.Equal(t, c.Y, gameState.cursor.Y)
}

func Test_moveCursor_InCaseOfTimeIsOutOfHumanPerceptionLimit(t *testing.T) {
	// Arrange
	gameState := NewGameState(10, 10)
	duration := time.Duration(100 * time.Millisecond)
	gameState.keyAt = time.Now().Add(-duration)
	c := coordinates.NewMatrix(5, 5)

	// Act
	gameState.moveCursor(c)

	// Assert
	assert.NotEqual(t, c.X, gameState.cursor.X)
	assert.NotEqual(t, c.Y, gameState.cursor.Y)
}
