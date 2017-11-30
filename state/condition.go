package state

import (
	"fmt"

	"github.com/Lytecyde/ox/data"
	"github.com/Lytecyde/ox/player"
)

var regularGameDimensionX int = data.RegularGameDimensionX
var regularGameDimensionY int = data.RegularGameDimensionY

func WriteMessage(gamesState *Game) {
	var name [3]string
	name[0] = "continue"
	name[1] = "Crosses"
	name[2] = "Naughts"
	if isWin(gamesState) {
		fmt.Println(name[gamesState.TurnOf] + " wins!")
		return
	}
	fmt.Println("game continues")
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
	//TODO: remove debug comments
	var winConditions int = 4
	allWinConditions := make([]bool, winConditions)
	allWinConditions[0] = isDiagonalDownWin(p, &gamesState)
	//fmt.Println("diagdn:" + strconv.FormatBool(isDiagonalDownWin(p)))
	allWinConditions[1] = isDiagonalUpWin(p, &gamesState)
	//fmt.Println("diagup:" + strconv.FormatBool(isDiagonalUpWin(p)))
	allWinConditions[2] = isColumnWin(p, &gamesState)
	//fmt.Println("col:" + strconv.FormatBool(isColumnWin(p)))
	allWinConditions[3] = isRowWin(p, &gamesState)
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

func isDiagonalDownWin(p player.Type, gamesState *Game) bool {
	win := make([]bool, regularGameDimensionX)
	var i int = 0
	var y int = 0
	for x := 0; x < regularGameDimensionX; x = x + 1 {
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
	win := make([]bool, regularGameDimensionX)
	var y int = 0
	for x := 0; x < regularGameDimensionX; x = x + 1 {
		y = regularGameDimensionY - 1 - x
		if (x == y) && (gamesState.Matrix.Fields[x][y] == p) {
			win[i] = true
			i++
		}
	}
	return isAllTrue(win)
}

func isColumnWin(p player.Type, gamesState *Game) bool {
	win := make([]bool, regularGameDimensionX)
	for x := 0; x < regularGameDimensionX; x = x + 1 {
		for y := 0; y < regularGameDimensionY; y = y + 1 {
			fmt.Print("nr of player on the field")
			fmt.Println(int(gamesState.Matrix.Fields[x][y]))
			if (int(gamesState.Matrix.Fields[x][y])) == (int(p)) {
				win[y] = true
			}
		}

		if isAllTrue(win) {
			return true
		}
		win = make([]bool, regularGameDimensionX)
	}
	return false
}

func isRowWin(p player.Type, gamesState *Game) bool {
	win := make([]bool, regularGameDimensionX)
	for y := 0; y < regularGameDimensionY; y = y + 1 {
		for x := 0; x < regularGameDimensionX; x = x + 1 {
			if (int(gamesState.Matrix.Fields[x][y])) == (int(p)) {
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
