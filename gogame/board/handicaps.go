package board

import (
	"fmt"
)

type HandicapPositions map[string]BoardPosition
type HandicapLevels map[int][]string
type HandicapSet struct {
	positions HandicapPositions
	levels    HandicapLevels
}

var SmallBoardHandicaps = HandicapSet{
	positions: HandicapPositions{
		"a": BoardPosition{
			Row:    2,
			CrossPoint: 6,
		},
		"b": BoardPosition{
			Row:    6,
			CrossPoint: 2,
		},
		"c": BoardPosition{
			Row:    6,
			CrossPoint: 6,
		},
		"d": BoardPosition{
			Row:    2,
			CrossPoint: 2,
		},
		"e": BoardPosition{
			Row:    4,
			CrossPoint: 4,
		},
	},
	levels: HandicapLevels{
		2: []string{"a", "b"},
		3: []string{"a", "b", "c"},
		4: []string{"a", "b", "c", "d"},
		5: []string{"a", "b", "c", "d", "e"},
	},
}

var MediumBoardHandicaps = HandicapSet{
	positions: HandicapPositions{
		"a": BoardPosition{
			Row:    2,
			CrossPoint: 10,
		},
		"b": BoardPosition{
			Row:    10,
			CrossPoint: 2,
		},
		"c": BoardPosition{
			Row:    10,
			CrossPoint: 10,
		},
		"d": BoardPosition{
			Row:    2,
			CrossPoint: 2,
		},
		"e": BoardPosition{
			Row:    6,
			CrossPoint: 6,
		},
		"f": BoardPosition{
			Row:    6,
			CrossPoint: 2,
		},
		"g": BoardPosition{
			Row:    6,
			CrossPoint: 10,
		},
		"h": BoardPosition{
			Row:    2,
			CrossPoint: 6,
		},
		"i": BoardPosition{
			Row:    10,
			CrossPoint: 6,
		},
	},
	levels: HandicapLevels{
		2: []string{"a", "b"},
		3: []string{"a", "b", "c"},
		4: []string{"a", "b", "c", "d"},
		5: []string{"a", "b", "c", "d", "e"},
		6: []string{"a", "b", "c", "d", "f", "g"},
		7: []string{"a", "b", "c", "d", "e", "f", "g"},
		8: []string{"a", "b", "c", "d", "f", "g", "h", "i"},
		9: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
	},
}

var LargeBoardHandicaps = HandicapSet{
	positions: HandicapPositions{
		"a": BoardPosition{
			Row:    3,
			CrossPoint: 15,
		},
		"b": BoardPosition{
			Row:    15,
			CrossPoint: 3,
		},
		"c": BoardPosition{
			Row:    15,
			CrossPoint: 15,
		},
		"d": BoardPosition{
			Row:    3,
			CrossPoint: 3,
		},
		"e": BoardPosition{
			Row:    9,
			CrossPoint: 9,
		},
		"f": BoardPosition{
			Row:    9,
			CrossPoint: 3,
		},
		"g": BoardPosition{
			Row:    9,
			CrossPoint: 15,
		},
		"h": BoardPosition{
			Row:    3,
			CrossPoint: 9,
		},
		"i": BoardPosition{
			Row:    15,
			CrossPoint: 9,
		},
	},
	levels: HandicapLevels{
		2: []string{"a", "b"},
		3: []string{"a", "b", "c"},
		4: []string{"a", "b", "c", "d"},
		5: []string{"a", "b", "c", "d", "e"},
		6: []string{"a", "b", "c", "d", "f", "g"},
		7: []string{"a", "b", "c", "d", "e", "f", "g"},
		8: []string{"a", "b", "c", "d", "f", "g", "h", "i"},
		9: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
	},
}

func (boardState *BoardState) SetHandicap(level int) {
	if !boardState.IsEmpty() {
		fmt.Println("Game is in progress.")
		return
	}

	var set HandicapSet
	switch boardState.Size() {
	case 9:
		set = SmallBoardHandicaps
	case 13:
		set = MediumBoardHandicaps
	case 19:
		set = LargeBoardHandicaps
	default:
		fmt.Println("No handicap set available for this board size.")
		return
	}

	var positions = set.levels[level]
	if len(positions) == 0 {
		fmt.Println("Invalid handicap level.")
	} else {
		for _, position := range positions {
			fmt.Println(set.positions[position])
			boardState.Place(StoneP1, set.positions[position])
		}
	}
}
