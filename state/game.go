package state

import (
	"fmt"
	"time"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/matrix"
	"github.com/Lytecyde/ox/player"
	"github.com/Lytecyde/ox/settings"
	"github.com/hajimehoshi/ebiten"
	"github.com/juju/errors"
)

type Game struct {
	Matrix        *matrix.Matrix
	Cursor        *coordinates.Matrix
	KeyAt         time.Time
	currentPlayer player.Type
	winner        player.Type
	Finished      bool
}

// NewGame returns new instance
func NewGame(gameDimensionX, gameDimensionY int) *Game {
	return &Game{
		Matrix:        matrix.NewMatrix(gameDimensionX, gameDimensionY),
		Cursor:        coordinates.NewMatrix(0, 0),
		KeyAt:         time.Now(),
		currentPlayer: player.Cross,
		winner:        player.None,
	}
}

func (gameState *Game) moveCursor(c *coordinates.Matrix) error {
	if time.Now().Sub(gameState.KeyAt).Seconds() < settings.HumanReactionSeconds {
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

	gameState.Matrix.SetState(*gameState.Cursor, gameState.currentPlayer)

	gameState.currentPlayer = player.Switch(gameState.currentPlayer)

}

func (gameState Game) isBoxTaken() bool {
	return gameState.Matrix.State(*gameState.Cursor) != player.None
}

func (gamesState *Game) SelectWinner() {
	name := [...]string{"none", "Crosses", "Naughts"}
	if gamesState.isWinner(player.Cross) {
		gamesState.winner = player.Cross
		fmt.Println(name[gamesState.winner])
		gamesState.Finished = true
	}
	if gamesState.isWinner(player.Naught) {
		gamesState.winner = player.Naught
		fmt.Println(name[gamesState.winner])
		gamesState.Finished = true
	}
}

func (gameState Game) isWinner(p player.Type) bool {
	return gameState.isDiagonalDownWin(p) ||
		gameState.isDiagonalUpWin(p) ||
		gameState.isColumnWin(p) ||
		gameState.isRowWin(p)
}

func (gamesState Game) isDiagonalDownWin(p player.Type) bool {
	win := make([]bool, settings.MatrixWidth)
	i := 0
	y := 0
	for x := 0; x < settings.MatrixWidth; x = x + 1 {
		y = x
		if gamesState.Matrix.Fields[x][y] == p {
			win[i] = true
			i++
		}
	}

	return isAllTrue(win)
}

func (gamesState Game) isDiagonalUpWin(p player.Type) bool {
	win := make([]bool, settings.MatrixWidth)
	y := 0
	i := 0
	for x := 0; x < settings.MatrixWidth; x = x + 1 {
		y = settings.MatrixHeight - 1 - x
		if gamesState.Matrix.Fields[x][y] == p {
			win[i] = true
			i++
		}
	}

	return isAllTrue(win)
}

func (gamesState Game) isColumnWin(p player.Type) bool {
	win := make([]bool, settings.MatrixWidth)
	for x := 0; x < settings.MatrixWidth; x = x + 1 {
		for y := 0; y < settings.MatrixHeight; y = y + 1 {
			if gamesState.Matrix.Fields[x][y] == p {
				win[y] = true
			}
		}

		if isAllTrue(win) {
			return true
		}

		win = make([]bool, settings.MatrixWidth)
	}

	return false
}

func (gamesState Game) isRowWin(p player.Type) bool {
	win := make([]bool, settings.MatrixWidth)
	for y := 0; y < settings.MatrixHeight; y = y + 1 {
		for x := 0; x < settings.MatrixWidth; x = x + 1 {
			if gamesState.Matrix.Fields[x][y] == p {
				win[x] = true
			}
		}

		if isAllTrue(win) {
			return true
		}

		win = make([]bool, settings.MatrixWidth)
	}

	return false
}

func isAllTrue(all []bool) bool {
	for i := 0; i < len(all); i = i + 1 {
		if !all[i] {
			return false
		}
	}

	return true
}
