package state

import (
	"testing"
	"time"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/player"
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
	c := coordinates.NewMatrix(5, 5)

	// Act
	err := gameState.moveCursor(c)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "less than human reaction time limit", err.Error())
	assert.NotEqual(t, c.X, gameState.Cursor.X)
	assert.NotEqual(t, c.Y, gameState.Cursor.Y)
}

func Test_moveCursorUp_MovesCursorUp_InCaseOfSuccess(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	gameState.Cursor = coordinates.NewMatrix(5, 5)

	// Act
	gameState.moveCursorUp()

	// Assert
	assert.Equal(t, 5, gameState.Cursor.X)
	assert.Equal(t, 4, gameState.Cursor.Y)
}

func Test_moveCursorDown_MovesCursorDown_InCaseOfSuccess(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	gameState.Cursor = coordinates.NewMatrix(5, 5)

	// Act
	gameState.moveCursorDown()

	// Assert
	assert.Equal(t, 5, gameState.Cursor.X)
	assert.Equal(t, 6, gameState.Cursor.Y)
}

func Test_moveCursorLeft_MovesCursorLeft_InCaseOfSuccess(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	gameState.Cursor = coordinates.NewMatrix(5, 5)

	// Act
	gameState.moveCursorLeft()

	// Assert
	assert.Equal(t, 4, gameState.Cursor.X)
	assert.Equal(t, 5, gameState.Cursor.Y)
}

func Test_moveCursorRight_MovesCursorRight_InCaseOfSuccess(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.KeyAt = time.Now().Add(-1 * time.Second)
	gameState.Cursor = coordinates.NewMatrix(5, 5)

	// Act
	gameState.moveCursorRight()

	// Assert
	assert.Equal(t, 6, gameState.Cursor.X)
	assert.Equal(t, 5, gameState.Cursor.Y)
}

func Test_setMark_DoesNotMark_InCaseOfAlreadyMarked(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.Cursor = coordinates.NewMatrix(5, 5)
	gameState.Matrix.SetState(*gameState.Cursor, player.Naught)
	gameState.currentPlayer = player.Cross

	// Act
	gameState.setMark()

	// Assert
	assert.Equal(t, player.Naught, gameState.Matrix.State(*gameState.Cursor))
}

func Test_setMark_Marks_InCaseOfSuccess(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.Cursor = coordinates.NewMatrix(5, 5)
	gameState.Matrix.SetState(*gameState.Cursor, player.None)
	gameState.currentPlayer = player.Cross

	// Act
	gameState.setMark()

	// Assert
	assert.Equal(t, player.Cross, gameState.Matrix.State(*gameState.Cursor))
}

func Test_isBoxTaken_ReturnsFalse_InCaseOfBoxNotTaken(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.Cursor = coordinates.NewMatrix(5, 5)
	gameState.Matrix.SetState(*gameState.Cursor, player.None)

	// Act
	taken := gameState.isBoxTaken()

	// Assert
	assert.False(t, taken)
}

func Test_isBoxTaken_ReturnsTrue_InCaseOfBoxTanel(t *testing.T) {
	// Arrange
	gameState := NewGame(10, 10)
	gameState.Cursor = coordinates.NewMatrix(5, 5)
	gameState.Matrix.SetState(*gameState.Cursor, player.Cross)

	// Act
	taken := gameState.isBoxTaken()

	// Assert
	assert.True(t, taken)
}
