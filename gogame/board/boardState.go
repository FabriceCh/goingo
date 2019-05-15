package board

import (
	"errors"
)

//Representing the game according to official rules and terminology from wikipedia
//https://en.wikipedia.org/wiki/Go_(game)

// State : structure containing the stones on the board
type State struct {
	Rows []boardRow
}
type boardRow struct {
	CrossPoints []CrossPoint
}

// CrossPoint :
type CrossPoint int

// Position :
type Position struct {
	Row, CrossPoint int
}

const (
	Vacant CrossPoint = iota
	StoneP1
	StoneP2
	Wall
)

func Initialize(size int) (State, error) {
	if size != 9 && size != 13 && size != 19 {
		return State{}, errors.New("Invalid board size. Available board sizes are: 9, 13 or 19.")
	}
	boardState := State{
		Rows: make([]boardRow, size),
	}
	for i := 0; i < size; i++ {
		row := boardRow{
			CrossPoints: make([]CrossPoint, size),
		}
		for j := 0; j < size; j++ {
			row.CrossPoints[j] = Vacant
		}
		boardState.Rows[i] = row
	}
	return boardState, nil
}

func (boardState *State) Place(stone CrossPoint, position Position) error {
	if position.Row >= boardState.Size() || position.CrossPoint >= boardState.Size() {
		return errors.New("Position not on the board")
	} else if !boardState.IsPlaceEmpty(position) {
		return errors.New("Board position not empty")
	} else {
		boardState.Rows[position.Row].CrossPoints[position.CrossPoint] = stone
		return nil
	}
}

func (boardState State) Size() int {
	return len(boardState.Rows)
}

func (boardState State) GetPlace(position Position) CrossPoint {
	return boardState.Rows[position.Row].CrossPoints[position.CrossPoint]
}

func (boardState State) IsPlaceEmpty(position Position) bool {
	return boardState.GetPlace(position) == Vacant
}

/*
*	IsEmpty
*   function that determines if the board is empty
 */
func (boardState State) IsEmpty() bool {
	for i, _ := range boardState.Rows {
		for j, _ := range boardState.Rows[i].CrossPoints {
			if !boardState.IsPlaceEmpty(Position{Row: i, CrossPoint: j}) {
				return false
			}
		}
	}
	return true
}
