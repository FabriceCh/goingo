package main

import (
	"fmt"
	"goingo/gogame/board"
)

func main() {
	fmt.Println("Welcome to the go game!")
	var boardState = board.Initialize()
	boardState.SetHandicap(5)
	boardState.Place(board.StoneP2, board.BoardPosition{Row: 0, Column: 0})
	boardState.ShowBoard()
}
