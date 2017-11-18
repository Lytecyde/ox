package main

import (
	"time"
)

type GameState struct {
	matrix *Matrix
	cursor *Coordinates
	keyAt  time.Time
}

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
