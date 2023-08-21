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
	// Vacant : empty space on the point
	Vacant CrossPoint = iota
	// StoneP1 : black stone of player 1
	StoneP1
	// StoneP2 : black stone of player 2
	StoneP2
	// Wall : out of bounds represented as wall
	Wall
)

// Initialize :
func Initialize(size int) (State, error) {
	if size != 9 && size != 13 && size != 19 {
		return State{}, errors.New("invalid board size, available board sizes are: 9, 13 or 19")
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

func (boardState *State) isPositionOnBoard(position Position) bool {
	if position.Row >= boardState.Size() || position.CrossPoint >= boardState.Size() {
		return false
	}
	if position.Row < 0 || position.CrossPoint < 0 {
		return false
	}
	return true
}

// Place : place a stone on the board
func (boardState *State) Place(stone CrossPoint, position Position) error {
	if !boardState.isPositionOnBoard(position) {
		return errors.New("Position not on the board")
	} else if !boardState.IsPlaceEmpty(position) {
		return errors.New("Board position not empty")
	} else {
		boardState.Rows[position.Row].CrossPoints[position.CrossPoint] = stone
		checkCapture(boardState, stone, position)
		return nil
	}
}

// Size : returns the size of the board
func (boardState State) Size() int {
	return len(boardState.Rows)
}

// GetPlace : returns the type of the item on a given position
func (boardState State) GetPlace(position Position) CrossPoint {
	if boardState.isPositionOnBoard(position) {
		return boardState.Rows[position.Row].CrossPoints[position.CrossPoint]
	}
	return Wall
}

// IsPlaceEmpty :
func (boardState State) IsPlaceEmpty(position Position) bool {
	return boardState.GetPlace(position) == Vacant
}

// IsEmpty : function that determines if the board is empty
func (boardState State) IsEmpty() bool {
	for i := range boardState.Rows {
		for j := range boardState.Rows[i].CrossPoints {
			if !boardState.IsPlaceEmpty(Position{Row: i, CrossPoint: j}) {
				return false
			}
		}
	}
	return true
}

func (boardState State) removeStone(position Position) error {
	if !boardState.isPositionOnBoard(position) {
		return errors.New("Position not on the board")
	} else if boardState.IsPlaceEmpty(position) {
		return errors.New("There is no stone to remove at this position")
	} else {
		boardState.Rows[position.Row].CrossPoints[position.CrossPoint] = Vacant
		return nil
	}
}
