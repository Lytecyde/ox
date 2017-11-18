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
