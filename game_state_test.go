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

func Test_moveCursor_DoesNotMoveCursor_InCaseOfNegativeX(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(-1, 2)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.x == coordinates.x {
		t.Fatal("invalid x")
	}

	if gameState.cursor.y == coordinates.y {
		t.Fatal("invalid y")
	}
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfOffScreenX(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(101, 2)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.x == coordinates.x {
		t.Fatal("invalid x")
	}

	if gameState.cursor.y == coordinates.y {
		t.Fatal("invalid y")
	}
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfNegativeY(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(1, -1)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.x == coordinates.x {
		t.Fatal("invalid x")
	}

	if gameState.cursor.y == coordinates.y {
		t.Fatal("invalid y")
	}
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfOffScreenY(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(1, 101)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.x == coordinates.x {
		t.Fatal("invalid x")
	}

	if gameState.cursor.y == coordinates.y {
		t.Fatal("invalid y")
	}
}
