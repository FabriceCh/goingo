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
	Wall
)

func Initialize(size int) (BoardState, error) {
	if size != 9 && size != 13 && size != 19 {
		return BoardState{}, errors.New("invalid board size: available board sizes are 9, 13 and 19")
	}
	boardState := BoardState{
		Rows: make([]BoardRow, size),
	}
	for i := range size {
		row := BoardRow{
			CrossPoints: make([]CrossPoint, size),
		}
		for j := range size {
			row.CrossPoints[j] = Vacant
		}
		boardState.Rows[i] = row
	}
	return boardState, nil
}

func (b *BoardState) Capture(position BoardPosition) {
	b.Rows[position.Row].CrossPoints[position.CrossPoint] = Vacant
}

func (b *BoardState) Place(stone CrossPoint, position BoardPosition) (int, error) {
	if !b.IsWithinBounds(position) {
		return 0, errors.New("position not on the board")
	} else if !b.IsPlaceEmpty(position) {
		return 0, errors.New("board position not empty")
	} else {
		b.Rows[position.Row].CrossPoints[position.CrossPoint] = stone
		points := b.CheckCapture(stone, position)
		return points, nil
	}
}

func (b *BoardState) Size() int {
	return len(b.Rows)
}

func (b *BoardState) IsWithinBounds(pos BoardPosition) bool {
	return pos.Row < b.Size() && pos.Row > -1 && pos.CrossPoint < b.Size() && pos.CrossPoint > -1
}

func (b *BoardState) GetPlace(position BoardPosition) CrossPoint {
	if !b.IsWithinBounds(position) {
		return Wall
	}
	return b.Rows[position.Row].CrossPoints[position.CrossPoint]
}

func (b *BoardState) IsPlaceEmpty(position BoardPosition) bool {
	if !b.IsWithinBounds(position) {
		return false
	}
	return b.GetPlace(position) == Vacant
}

/*
*	IsEmpty
*   function that determines if the board is empty
 */
func (b *BoardState) IsEmpty() bool {
	for i := range b.Rows {
		for j := range b.Rows[i].CrossPoints {
			if !b.IsPlaceEmpty(BoardPosition{Row: i, CrossPoint: j}) {
				return false
			}
		}
	}
	return true
}

func (b *BoardState) DeepCopy() BoardState {
	copiedBoardState := BoardState{
		Rows: make([]BoardRow, b.Size()),
	}
	for i := range b.Size() {
		row := BoardRow{
			CrossPoints: make([]CrossPoint, b.Size()),
		}
		for j := range b.Size() {
			row.CrossPoints[j] = b.GetPlace(BoardPosition{Row: i, CrossPoint: j})
		}
		copiedBoardState.Rows[i] = row
	}
	return copiedBoardState
}
