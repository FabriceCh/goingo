package board

import "fmt"

//Representing the game according to official rules and terminology from wikipedia
//https://en.wikipedia.org/wiki/Go_(game)

type BoardState [9][9]CrossPoint
type CrossPoint int
type BoardPosition struct {
	Row, Column int
}

const (
	Vacant CrossPoint = iota
	StoneP1
	StoneP2
)

func Initialize() BoardState {
	return BoardState{}
}

func (boardState *BoardState) Place(stone CrossPoint, position BoardPosition) {
	if position.Row >= boardState.Size() || position.Column >= boardState.Size() {
		fmt.Println("Position not on the board.")
	} else if !boardState.IsPlaceEmpty(position) {
		fmt.Println("Board position not empty.")
	} else {
		boardState[position.Row][position.Column] = stone
	}
}

func (boardState BoardState) Size() int {
	return len(boardState)
}

func (boardState BoardState) IsPlaceEmpty(position BoardPosition) bool {
	return boardState[position.Row][position.Column] == Vacant
}

func (boardState BoardState) IsEmpty() bool {
	for i, _ := range boardState {
		for j, _ := range boardState[i] {
			if !boardState.IsPlaceEmpty(BoardPosition{Row: i, Column: j}) {
				return false
			}
		}
	}
	return true
}
