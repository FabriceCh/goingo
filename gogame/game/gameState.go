package game

import (
	"errors"
	"git-gogame/gogame/board"
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

func (gameState GameState) ExecuteCommand(command string, args ...string) (msg string, err error) {
	switch command {
	case "handicap":
		err = gameState.Board.SetHandicap(2)
		msg = "Set handicap of level 2 for Player 1"
	case "place":
		err = gameState.Board.Place(board.StoneP1, board.BoardPosition{Row: 0, CrossPoint: 0})
		msg = "Placed a stone at (0,0) for Player 1"
	default:
		err = errors.New("Invalid command")
		msg = ""
	}
	return
}
