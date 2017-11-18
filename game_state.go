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
		matrix: NewMatrix(regularGameDimensionX, regularGameDimensionY),
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

	if coordinates.x >= regularGameDimensionX {
		return
	}

	if coordinates.y < 0 {
		return
	}

	if coordinates.y >= regularGameDimensionY {
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
