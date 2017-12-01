package state

import (
	"github.com/Lytecyde/ox/data"
	"github.com/Lytecyde/ox/player"
)

func WriteMessage(gamesState *Game) {
	var name [3]string
	name[0] = "continue"
	name[1] = "Crosses"
	name[2] = "Naughts"
	if isWin(gamesState) {
		return
	}
}

func isWin(gamesState *Game) bool {
	if int(getWinner(gamesState.TurnOf, *gamesState)) > int(player.None) {
		gamesState.Winner = getWinner(gamesState.TurnOf, *gamesState)
		gamesState.EndOfGame = true
		return true
	}
	return false
}

func getWinner(p player.Type, gamesState Game) player.Type {
	var winConditions int = 4
	allWinConditions := make([]bool, winConditions)
	allWinConditions[0] = isDiagonalDownWin(p, &gamesState)
	allWinConditions[1] = isDiagonalUpWin(p, &gamesState)
	allWinConditions[2] = isColumnWin(p, &gamesState)
	allWinConditions[3] = isRowWin(p, &gamesState)

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

func isDiagonalDownWin(p player.Type, gamesState *Game) bool {
	win := make([]bool, data.RegularGameDimensionX)
	var i int = 0
	var y int = 0
	for x := 0; x < data.RegularGameDimensionX; x = x + 1 {
		y = x
		if gamesState.Matrix.Fields[x][y] == p {
			win[i] = true
			i++
		}
	}
	i = 0
	return isAllTrue(win)
}

func isDiagonalUpWin(p player.Type, gamesState *Game) bool {
	var i int = 0
	win := make([]bool, data.RegularGameDimensionX)
	var y int = 0
	for x := 0; x < data.RegularGameDimensionX; x = x + 1 {
		y = data.RegularGameDimensionY - 1 - x
		if (x == y) && (gamesState.Matrix.Fields[x][y] == p) {
			win[i] = true
			i++
		}
	}
	return isAllTrue(win)
}

func isColumnWin(p player.Type, gamesState *Game) bool {
	win := make([]bool, data.RegularGameDimensionX)
	for x := 0; x < data.RegularGameDimensionX; x = x + 1 {
		for y := 0; y < data.RegularGameDimensionY; y = y + 1 {
			if (int(gamesState.Matrix.Fields[x][y])) == (int(p)) {
				win[y] = true
			}
		}

		if isAllTrue(win) {
			return true
		}
		win = make([]bool, data.RegularGameDimensionX)
	}
	return false
}

func isRowWin(p player.Type, gamesState *Game) bool {
	win := make([]bool, data.RegularGameDimensionX)
	for y := 0; y < data.RegularGameDimensionY; y = y + 1 {
		for x := 0; x < data.RegularGameDimensionX; x = x + 1 {
			if (int(gamesState.Matrix.Fields[x][y])) == (int(p)) {
				win[x] = true
			}
		}
		if isAllTrue(win) {
			return true
		}
		win = make([]bool, data.RegularGameDimensionX)
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
