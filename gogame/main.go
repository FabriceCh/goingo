package main

import (
	"fmt"
	"goingo/gogame/board"
)

func main() {
	fmt.Println("Welcome to the go game!")
<<<<<<< HEAD
	var boardState = board.Initialize()
	boardState.SetHandicap(5)
	boardState.Place(board.StoneP2, board.BoardPosition{Row: 0, Column: 0})
=======
	var boardState, _ = board.Initialize(9)
>>>>>>> ec7e249719331ec4692a19d3547bd922aea71e56
	boardState.ShowBoard()
}
