package board

import (
	"errors"
)

//Representing the game according to official rules and terminology from wikipedia
//https://en.wikipedia.org/wiki/Go_(game)

type BoardState struct {
	Rows []BoardRow
}
type BoardRow struct {
	CrossPoints []CrossPoint
}
type CrossPoint int
type BoardPosition struct {
	Row, CrossPoint int
}

const (
	Vacant CrossPoint = iota
	StoneP1
	StoneP2
)

func Initialize(size int) (BoardState, error) {
	if size != 9 && size != 13 && size != 19 {
		return BoardState{}, errors.New("Invalid board size")
	}
	boardState := BoardState{
		Rows: make([]BoardRow, size),
	}
	for i := 0; i < size; i++ {
		row := BoardRow{
			CrossPoints: make([]CrossPoint, size),
		}
		for j := 0; j < size; j++ {
			row.CrossPoints[j] = Vacant
		}
		boardState.Rows[i] = row
	}
	return boardState, nil
}

func (boardState *BoardState) Place(stone CrossPoint, position BoardPosition) error {
	if position.Row >= boardState.Size() || position.CrossPoint >= boardState.Size() {
		return errors.New("Position not on the board")
	} else if !boardState.IsPlaceEmpty(position) {
		return errors.New("Board position not empty")
	} else {
		boardState.Rows[position.Row].CrossPoints[position.CrossPoint] = stone
		return nil
	}
}

func (boardState BoardState) Size() int {
	return len(boardState.Rows)
}

func (boardState BoardState) GetPlace(position BoardPosition) CrossPoint {
	return boardState.Rows[position.Row].CrossPoints[position.CrossPoint]
}

func (boardState BoardState) IsPlaceEmpty(position BoardPosition) bool {
	return boardState.GetPlace(position) == Vacant
}

/*
*	IsEmpty
*   function that determines if the board is empty
 */
func (boardState BoardState) IsEmpty() bool {
	for i, _ := range boardState.Rows {
		for j, _ := range boardState.Rows[i].CrossPoints {
			if !boardState.IsPlaceEmpty(BoardPosition{Row: i, CrossPoint: j}) {
				return false
			}
		}
	}
	return true
}
