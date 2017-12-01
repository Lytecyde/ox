package state

import (
	"fmt"
	"time"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/player"
	"github.com/Lytecyde/ox/playfield"
	"github.com/hajimehoshi/ebiten"
	"github.com/juju/errors"
)

type Game struct {
	Matrix    *playfield.Matrix
	Cursor    *coordinates.Matrix
	KeyAt     time.Time
	TurnOf    player.Type
	Winner    player.Type
	EndOfGame bool
}

// NewGame returns new instance
func NewGame(gameDimensionX, gameDimensionY int) *Game {
	return &Game{
		Matrix:    playfield.NewMatrix(gameDimensionX, gameDimensionY),
		Cursor:    coordinates.NewMatrix(0, 0),
		KeyAt:     time.Now(),
		TurnOf:    player.Cross,
		Winner:    player.None,
		EndOfGame: false,
	}
}

func (gameState *Game) moveCursor(c *coordinates.Matrix) error {
	if time.Now().Sub(gameState.KeyAt).Seconds() < 0.2 {
		return errors.Errorf("less than human reaction time limit")
	}

	if c.X < 0 {
		return errors.Errorf("x less than 0")
	}

	if c.X >= gameState.Matrix.Dimensions.X {
		return errors.Errorf("out of the matrix")
	}

	if c.Y < 0 {
		return errors.Errorf("y less than 0")
	}

	if c.Y >= gameState.Matrix.Dimensions.Y {
		return errors.Errorf("out of the matrix")
	}

	gameState.Cursor = c

	gameState.KeyAt = time.Now()

	return nil
}

func (gameState *Game) moveCursorUp() {
	gameState.moveCursor(coordinates.NewMatrix(gameState.Cursor.X, gameState.Cursor.Y-1))
}

func (gameState *Game) moveCursorDown() {
	gameState.moveCursor(coordinates.NewMatrix(gameState.Cursor.X, gameState.Cursor.Y+1))
}

func (gameState *Game) moveCursorLeft() {
	gameState.moveCursor(coordinates.NewMatrix(gameState.Cursor.X-1, gameState.Cursor.Y))
}

func (gameState *Game) moveCursorRight() {
	gameState.moveCursor(coordinates.NewMatrix(gameState.Cursor.X+1, gameState.Cursor.Y))
}

func (gameState *Game) HandleKeyPress() {

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		gameState.moveCursorUp()

	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		gameState.moveCursorDown()

	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		gameState.moveCursorLeft()

	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		gameState.moveCursorRight()

	} else if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyEnter) {
		gameState.setMark()

	}
}

func (gameState *Game) setMark() {
	if gameState.isBoxTaken() {
		return
	}

	gameState.markBox()
	gameState.switchPlayers()
}

func (gameState *Game) isBoxTaken() bool {
	return gameState.Matrix.State(*gameState.Cursor) != player.None
}

func (gameState *Game) markBox() {
	gameState.Matrix.SetState(*gameState.Cursor, gameState.TurnOf)
}

func (gameState *Game) switchPlayers() {
	gameState.TurnOf = alter(gameState.TurnOf)
}

func alter(t player.Type) player.Type {
	switch t {
	case player.Cross:
		return player.Naught
	case player.Naught:
		return player.Cross
	}

	panic(fmt.Sprintf("invalid player type: %d", t))
}
