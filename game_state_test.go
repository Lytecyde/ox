package main

import (
	"testing"
	"time"
)

func Test_NewGameState_SetsMatrixSizeFromParameters(t *testing.T) {
	// Arrange
	x, y := 100, 100

	// Act
	gameState := NewGameState(x, y)

	// Assert
	if gameState.matrix.x != x {
		t.Fatal("invalid x")
	}

	if gameState.matrix.y != y {
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
	if !gameState.cursor.Equal(*coordinates) {
		t.Fatal("invalid coordinates")
	}
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfNegativeX(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(-1, 2)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.Equal(*coordinates) {
		t.Fatal("invalid coordinates")
	}
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfOffScreenX(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(101, 2)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.Equal(*coordinates) {
		t.Fatal("invalid coordinates")
	}
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfNegativeY(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(1, -1)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.Equal(*coordinates) {
		t.Fatal("invalid coordinates")
	}
}

func Test_moveCursor_DoesNotMoveCursor_InCaseOfOffScreenY(t *testing.T) {
	// Arrange
	gameState := NewGameState(100, 100)
	coordinates := NewCoordinates(1, 101)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.Equal(*coordinates) {
		t.Fatal("invalid coordinates")
	}
}

func Test_moveCursor_InCaseOfTimeIsWithinHumanPerceptionLimit(t *testing.T) {
	// Arrange
	gameState := NewGameState(10, 10)
	duration := time.Duration(1 * time.Second)
	gameState.keyAt = time.Now().Add(-duration)
	coordinates := NewCoordinates(5, 5)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if !gameState.cursor.Equal(*coordinates) {
		t.Fatal("invalid coordinates")
	}
}

func Test_moveCursor_InCaseOfTimeIsOutOfHumanPerceptionLimit(t *testing.T) {
	// Arrange
	gameState := NewGameState(10, 10)
	duration := time.Duration(100 * time.Millisecond)
	gameState.keyAt = time.Now().Add(-duration)
	coordinates := NewCoordinates(5, 5)

	// Act
	gameState.moveCursor(coordinates)

	// Assert
	if gameState.cursor.Equal(*coordinates) {
		t.Fatal("invalid coordinates")
	}
}
