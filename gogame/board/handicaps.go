package board

import (
	"errors"
)

// HandicapPositions :
type HandicapPositions map[string]Position

// HandicapLevels :
type HandicapLevels map[int][]string

// HandicapSet :
type HandicapSet struct {
	positions HandicapPositions
	levels    HandicapLevels
}

// SmallBoardHandicaps :
var SmallBoardHandicaps = HandicapSet{
	positions: HandicapPositions{
		"a": Position{
			Row:        2,
			CrossPoint: 6,
		},
		"b": Position{
			Row:        6,
			CrossPoint: 2,
		},
		"c": Position{
			Row:        6,
			CrossPoint: 6,
		},
		"d": Position{
			Row:        2,
			CrossPoint: 2,
		},
		"e": Position{
			Row:        4,
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

// MediumBoardHandicaps :
var MediumBoardHandicaps = HandicapSet{
	positions: HandicapPositions{
		"a": Position{
			Row:        2,
			CrossPoint: 10,
		},
		"b": Position{
			Row:        10,
			CrossPoint: 2,
		},
		"c": Position{
			Row:        10,
			CrossPoint: 10,
		},
		"d": Position{
			Row:        2,
			CrossPoint: 2,
		},
		"e": Position{
			Row:        6,
			CrossPoint: 6,
		},
		"f": Position{
			Row:        6,
			CrossPoint: 2,
		},
		"g": Position{
			Row:        6,
			CrossPoint: 10,
		},
		"h": Position{
			Row:        2,
			CrossPoint: 6,
		},
		"i": Position{
			Row:        10,
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

// LargeBoardHandicaps :
var LargeBoardHandicaps = HandicapSet{
	positions: HandicapPositions{
		"a": Position{
			Row:        3,
			CrossPoint: 15,
		},
		"b": Position{
			Row:        15,
			CrossPoint: 3,
		},
		"c": Position{
			Row:        15,
			CrossPoint: 15,
		},
		"d": Position{
			Row:        3,
			CrossPoint: 3,
		},
		"e": Position{
			Row:        9,
			CrossPoint: 9,
		},
		"f": Position{
			Row:        9,
			CrossPoint: 3,
		},
		"g": Position{
			Row:        9,
			CrossPoint: 15,
		},
		"h": Position{
			Row:        3,
			CrossPoint: 9,
		},
		"i": Position{
			Row:        15,
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

// SetHandicap :
func (boardState *State) SetHandicap(level int) error {
	if !boardState.IsEmpty() {
		return errors.New("Board is not empty")
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
		return errors.New("No handicap set available for this board size")
	}

	var positions = set.levels[level]
	if len(positions) == 0 {
		return errors.New("Invalid handicap level")
	}
	for _, position := range positions {
		boardState.Place(StoneP1, set.positions[position])
	}
	return nil

}
