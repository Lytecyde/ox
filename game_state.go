package main

import (
	"fmt"
	"time"

	"github.com/Lytecyde/ox/coordinates"
	"github.com/Lytecyde/ox/player"
	"github.com/hajimehoshi/ebiten"
)

// GameState represents game state
type GameState struct {
	matrix    *Matrix
	cursor    *coordinates.Matrix
	keyAt     time.Time
	turnOf    player.Type
	winner    player.Type
	EndOfGame bool
}

// NewGameState returns new instance
func NewGameState(gameDimensionX, gameDimensionY int) *GameState {
	return &GameState{
		matrix:    NewMatrix(gameDimensionX, gameDimensionY),
		cursor:    coordinates.NewMatrix(0, 0),
		turnOf:    player.Cross,
		EndOfGame: false,
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

func (gameState *GameState) handleKeyPress() {
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

func (gameState *GameState) setMark() {

	if gameState.isBoxTaken() {
		return
	}

	gameState.markBox()

	gameState.isWinMessage()

	gameState.switchPlayers()
}

func (gameState *GameState) isBoxTaken() bool {
	return gameState.matrix.state(*gameState.cursor) != player.None
}

func (gameState *GameState) markBox() {
	gameState.matrix.setState(*gameState.cursor, gameState.turnOf)
}

func (gameState *GameState) switchPlayers() {
	gameState.turnOf = alter(gameState.turnOf)
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

////////////////WIN package stuff is here
func (gamesState *GameState) isWinMessage() {
	var name [3]string
	name[0] = "continue"
	name[1] = "Crosses"
	name[2] = "Naughts"
	if gameState.isWin() {
		fmt.Println(name[gameState.turnOf] + " wins!")
		return
	}
	fmt.Println("game continues")
}

func (gamesState *GameState) isWin() bool {
	if int(getWinner(gameState.turnOf)) > int(player.None) {
		gameState.winner = getWinner(gameState.turnOf)
		gameState.EndOfGame = true
		return true
	}
	return false
}

func getWinner(p player.Type) player.Type {
	//TODO: remove debug comments
	var winConditions int = 4
	allWinConditions := make([]bool, winConditions)
	allWinConditions[0] = isDiagonalDownWin(p)
	//fmt.Println("diagdn:" + strconv.FormatBool(isDiagonalDownWin(p)))
	allWinConditions[1] = isDiagonalUpWin(p)
	//fmt.Println("diagup:" + strconv.FormatBool(isDiagonalUpWin(p)))
	allWinConditions[2] = isColumnWin(p)
	//fmt.Println("col:" + strconv.FormatBool(isColumnWin(p)))
	allWinConditions[3] = isRowWin(p)
	//fmt.Println("row:" + strconv.FormatBool(isRowWin(p)))

	if isOneTrue(allWinConditions) == false {
		return player.None
	}
	return p
}

func isOneTrue(all []bool) bool {
	var oneTrue bool = false
	for i := 0; i < len(all); i = i + 1 {
		oneTrue = oneTrue || all[i]
	}
	return oneTrue
}

func isDiagonalDownWin(p player.Type) bool {
	win := make([]bool, regularGameDimensionX)
	var i int = 0
	var y int = 0
	for x := 0; x < regularGameDimensionX; x = x + 1 {
		y = x
		if gameState.matrix.fields[x][y] == p {
			win[i] = true
			i++
		}
	}
	i = 0
	return isAllTrue(win)
}

func isDiagonalUpWin(p player.Type) bool {
	var i int = 0
	win := make([]bool, regularGameDimensionX)
	var y int = 0
	for x := 0; x < regularGameDimensionX; x = x + 1 {
		y = regularGameDimensionY - 1 - x
		if (x == y) && (gameState.matrix.fields[x][y] == p) {
			win[i] = true
			i++
		}

	}
	return isAllTrue(win)
}

func isColumnWin(p player.Type) bool {
	win := make([]bool, regularGameDimensionX)
	//winNum := make([]int, regularGameDimensionY)
	for x := 0; x < regularGameDimensionX; x = x + 1 {
		for y := 0; y < regularGameDimensionY; y = y + 1 {
			fmt.Print("nr of player on the field")
			fmt.Println(int(gameState.matrix.fields[x][y]))
			if (int(gameState.matrix.fields[x][y])) == (int(p)) {
				win[y] = true

			}
			//winNum[x] = winNum[x] + int(gameState.matrix.fields[x][y])
		}
		//wincheck
		if isAllTrue(win) {
			return true
		}
		/*win = make([]bool, regularGameDimensionX)
		//winNum check
		if winNum[x] == 3 {
			return true
		}
		if winNum[x] == 6 {
			return true
		}*/

	}
	return false
}

func isRowWin(p player.Type) bool {
	win := make([]bool, regularGameDimensionX)
	for y := 0; y < regularGameDimensionY; y = y + 1 {
		for x := 0; x < regularGameDimensionX; x = x + 1 {
			if (int(gameState.matrix.fields[x][y])) == (int(p)) {
				win[x] = true
			}
		}
		if isAllTrue(win) {
			return true
		}
		win = make([]bool, regularGameDimensionX)
	}
	return false
}

func isAllTrue(all []bool) bool {
	var allTrue bool = true
	for i := 0; i < len(all); i = i + 1 {
		allTrue = allTrue && all[i]
	}
	return allTrue
}
