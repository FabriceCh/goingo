package game

import (
	"fmt"
	"goingo/gogame/board"
)

type History struct {
	turns []Turn
}

type Turn struct {
	BoardState        board.BoardState
	P1Points          int
	P2Points          int
	ActivePlayerColor board.CrossPoint
}

func NewTurn(gameState GameState) Turn {
	return Turn{
		BoardState:        gameState.board.DeepCopy(),
		P1Points:          gameState.player1.points,
		P2Points:          gameState.player2.points,
		ActivePlayerColor: gameState.activePlayer.stone,
	}
}

func (h *History) Push(turn Turn) {
	h.turns = append(h.turns, turn)
}

func (h *History) Pop() (*Turn, error) {
	if len(h.turns) < 1 {
		return nil, fmt.Errorf("")
	}
	t := h.turns[len(h.turns)-1]
	h.turns = h.turns[:len(h.turns)-1]
	return &t, nil
}
