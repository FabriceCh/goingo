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
			Column: 6,
		},
		"b": BoardPosition{
			Row:    6,
			Column: 2,
		},
		"c": BoardPosition{
			Row:    6,
			Column: 6,
		},
		"d": BoardPosition{
			Row:    2,
			Column: 2,
		},
		"e": BoardPosition{
			Row:    4,
			Column: 4,
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
			Column: 10,
		},
		"b": BoardPosition{
			Row:    10,
			Column: 2,
		},
		"c": BoardPosition{
			Row:    10,
			Column: 10,
		},
		"d": BoardPosition{
			Row:    2,
			Column: 2,
		},
		"e": BoardPosition{
			Row:    6,
			Column: 6,
		},
		"f": BoardPosition{
			Row:    6,
			Column: 2,
		},
		"g": BoardPosition{
			Row:    6,
			Column: 10,
		},
		"h": BoardPosition{
			Row:    2,
			Column: 6,
		},
		"i": BoardPosition{
			Row:    10,
			Column: 6,
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
			Column: 15,
		},
		"b": BoardPosition{
			Row:    15,
			Column: 3,
		},
		"c": BoardPosition{
			Row:    15,
			Column: 15,
		},
		"d": BoardPosition{
			Row:    3,
			Column: 3,
		},
		"e": BoardPosition{
			Row:    9,
			Column: 9,
		},
		"f": BoardPosition{
			Row:    9,
			Column: 3,
		},
		"g": BoardPosition{
			Row:    9,
			Column: 15,
		},
		"h": BoardPosition{
			Row:    3,
			Column: 9,
		},
		"i": BoardPosition{
			Row:    15,
			Column: 9,
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
