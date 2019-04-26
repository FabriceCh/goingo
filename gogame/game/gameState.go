package game

import (
	"errors"
	"fmt"
	"git-gogame/gogame/board"
	"strconv"
)

// This will eventually handle turns and points
type GameState struct {
	Board board.BoardState
}

func Start(size int) (GameState, error) {
	boardState, err := board.Initialize(size)
	return GameState{
		Board: boardState,
	}, err
}

func (gameState GameState) ExecuteCommand(command string, args []string) (msg string, err error) {
	switch command {
	case "handicap":
		if len(args) < 1 {
			err = errors.New("Too few arguments")
			return
		}
		level, _ := strconv.Atoi(args[0])
		err = gameState.Board.SetHandicap(level)
		msg = fmt.Sprintf("Set handicap of level %d for Player 1", level)
	case "place":
		if len(args) < 2 {
			err = errors.New("Too few arguments")
			return
		}
		row, _ := strconv.Atoi(args[0])
		crossPoint, _ := strconv.Atoi(args[1])
		err = gameState.Board.Place(board.StoneP1, board.BoardPosition{Row: row, CrossPoint: crossPoint})
		msg = fmt.Sprintf("Placed a stone at (%d,%d) for Player 1", row, crossPoint)
	default:
		err = errors.New("Invalid command")
		msg = ""
	}
	return
}

func (gameState GameState) Show() {
	gameState.Board.ShowBoard()
}
