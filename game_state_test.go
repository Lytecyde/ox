package main

import (
	"testing"
)

func Test_NewGameState_SetsMatrixSizeFromParameters(t *testing.T) {
	// Arrange
	x, y := 100, 100

	// Act
	gameState := NewGameState(x, y)

	// Assert
	if gameState.matrix.dimensionx != x {
		t.Fatal("invalid x")
	}

	if gameState.matrix.dimensiony != y {
		t.Fatal("invalid y")
	}
}

func Test_moveCursor_MovesCursorToGivenCoordinates(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(2, 2)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.x != coordinates.x {
		t.Fatal("invalid x")
	}

	if gameState.cursor.y != coordinates.y {
		t.Fatal("invalid y")
	}
}
