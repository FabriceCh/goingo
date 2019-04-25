package board

import "fmt"

//Representing the game according to official rules and terminology from wikipedia
//https://en.wikipedia.org/wiki/Go_(game)

type BoardState [19][19]CrossPoint
type CrossPoint int
type BoardPosition struct {
	row, column int
}

func (boardState BoardState) IsPlaceEmpty(position BoardPosition) bool {
	return boardState[position.row][position.column] == Vacant
}

func (boardState BoardState) IsEmpty() bool {
	for i, _ := range boardState {
		for j, _ := range boardState[i] {
			if boardState[i][j] != Vacant {
				return false
			}
		}
	}
	return true
}

const (
	Vacant CrossPoint = iota
	StoneP1
	StoneP2
)

var handicaps = map[string]BoardPosition{
	"a": BoardPosition{
		row:    3,
		column: 15,
	},
	"b": BoardPosition{
		row:    15,
		column: 3,
	},
	"c": BoardPosition{
		row:    15,
		column: 15,
	},
	"d": BoardPosition{
		row:    3,
		column: 3,
	},
	"e": BoardPosition{
		row:    9,
		column: 9,
	},
	"f": BoardPosition{
		row:    9,
		column: 3,
	},
	"g": BoardPosition{
		row:    9,
		column: 15,
	},
	"h": BoardPosition{
		row:    3,
		column: 9,
	},
	"i": BoardPosition{
		row:    15,
		column: 9,
	},
}

var boardState BoardState

func Initialize() {
	boardState = BoardState{}
}

func Place(stone CrossPoint, position BoardPosition) {
	if boardState.IsPlaceEmpty(position) {
		boardState[position.row][position.column] = stone
	} else {
		fmt.Println("Board position not empty.")
	}
}

func SetHandicap(stone CrossPoint, level int) {
	if !boardState.IsEmpty() {
		fmt.Println("Game is in progress.")
		return
	}
	var positions []string
	switch level {
	case 2:
		positions = []string{"a", "b"}
	case 3:
		positions = []string{"a", "b", "c"}
	case 4:
		positions = []string{"a", "b", "c", "d"}
	case 5:
		positions = []string{"a", "b", "c", "d", "e"}
	case 6:
		positions = []string{"a", "b", "c", "d", "f", "g"}
	case 7:
		positions = []string{"a", "b", "c", "d", "e", "f", "g"}
	case 8:
		positions = []string{"a", "b", "c", "d", "f", "g", "h", "i"}
	case 9:
		positions = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	}
	for _, position := range positions {
		Place(stone, handicaps[position])
	}
}
