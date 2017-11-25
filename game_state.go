package main

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// GameState represents game state
type GameState struct {
	matrix *Matrix
	cursor *Coordinates
	keyAt  time.Time
}

// NewGameState returns new instance
func NewGameState(gameDimensionX, gameDimensionY int) *GameState {
	return &GameState{
		matrix: NewMatrix(gameDimensionX, gameDimensionY),
		cursor: NewCoordinates(0, 0),
	}
}

func (gameState *GameState) moveCursor(coordinates *Coordinates) {
	if time.Now().Sub(gameState.keyAt).Seconds() < 0.2 {
		return
	}

	if coordinates.x < 0 {
		return
	}

	if coordinates.x >= gameState.matrix.dimensions.x {
		return
	}

	if coordinates.y < 0 {
		return
	}

	if coordinates.y >= gameState.matrix.dimensions.y {
		return
	}

	gameState.cursor = coordinates

	gameState.keyAt = time.Now()
}

func (gameState *GameState) moveCursorUp() {
	gameState.moveCursor(NewCoordinates(gameState.cursor.x, gameState.cursor.y-1))
}

func (gameState *GameState) moveCursorDown() {
	gameState.moveCursor(NewCoordinates(gameState.cursor.x, gameState.cursor.y+1))
}

func (gameState *GameState) moveCursorLeft() {
	gameState.moveCursor(NewCoordinates(gameState.cursor.x-1, gameState.cursor.y))
}

func (gameState *GameState) moveCursorRight() {
	gameState.moveCursor(NewCoordinates(gameState.cursor.x+1, gameState.cursor.y))
}

func (gameState GameState) drawMatrix(screen *ebiten.Image) {
	drawMatrix(screen, gameState.matrix, gray)
}

func (gameState GameState) drawCursor(screen *ebiten.Image) {
	drawBox(screen, NewCoordinates(gameState.cursor.x*boxSize, gameState.cursor.y*boxSize), red)
}

func (gameState GameState) drawStates(screen *ebiten.Image) {
	drawStates(screen, gameState.matrix, blue, green)
}

func (gameState *GameState) handleKeyPress() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		gameState.moveCursorUp()

	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		gameState.moveCursorDown()

	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		gameState.moveCursorLeft()

	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		gameState.moveCursorRight()

	} else if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		gameState.setMark()

	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		gameState.setMark()

	}
}

func (gameState *GameState) setMark() {
	// check if box is not taken yet in matrix

	// mark box as taken in matrix
	gameState.matrix.setState(*gameState.cursor, fieldStatePlayerX)
}
