package main

import (
	"fmt"
	"git-gogame/gogame/board"
)

func main() {
	fmt.Println("Welcome to the go game!")
	var boardState = board.Initialize(9)
	boardState.ShowBoard()
}
