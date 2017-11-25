package main

import (
	"time"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/player"
	"github.com/hajimehoshi/ebiten"
)

// GameState represents game state
type GameState struct {
	matrix *Matrix
	cursor *coordinates.Matrix
	keyAt  time.Time
	turnOf player.Type
}

// NewGameState returns new instance
func NewGameState(gameDimensionX, gameDimensionY int) *GameState {
	return &GameState{
		matrix: NewMatrix(gameDimensionX, gameDimensionY),
		cursor: coordinates.NewMatrix(0, 0),
		turnOf: player.Cross,
	}
}

func (gameState *GameState) moveCursor(c *coordinates.Matrix) {
	if time.Now().Sub(gameState.keyAt).Seconds() < 0.2 {
		return
	}

	if c.X < 0 {
		return
	}

	if c.X >= gameState.matrix.dimensions.X {
		return
	}

	if c.Y < 0 {
		return
	}

	if c.Y >= gameState.matrix.dimensions.Y {
		return
	}

	gameState.cursor = c

	gameState.keyAt = time.Now()
}

func (gameState *GameState) moveCursorUp() {
	gameState.moveCursor(coordinates.NewMatrix(gameState.cursor.X, gameState.cursor.Y-1))
}

func (gameState *GameState) moveCursorDown() {
	gameState.moveCursor(coordinates.NewMatrix(gameState.cursor.X, gameState.cursor.Y+1))
}

func (gameState *GameState) moveCursorLeft() {
	gameState.moveCursor(coordinates.NewMatrix(gameState.cursor.X-1, gameState.cursor.Y))
}

func (gameState *GameState) moveCursorRight() {
	gameState.moveCursor(coordinates.NewMatrix(gameState.cursor.X+1, gameState.cursor.Y))
}

func (gameState GameState) drawMatrix(screen *ebiten.Image) {
	drawMatrix(screen, gameState.matrix, gray)
}

func (gameState GameState) drawCursor(screen *ebiten.Image) {
	drawBox(screen, coordinates.NewScreen(gameState.cursor.X*boxSize, gameState.cursor.Y*boxSize), red)
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

	gameState.matrix.setState(*gameState.cursor, gameState.turnOf)
	gameState.turnOf = alter(gameState.turnOf)
}

func alter(t player.Type) player.Type {
	var out player.Type = player.Cross
	switch t {
	case player.Cross:
		out = player.Naught
	case player.Naught:
		out = player.Cross
	}
	return out
}
