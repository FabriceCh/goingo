package board

import "testing"

func TestSetSmallHandicap(t *testing.T) {
	boardState := Initialize(9)
	boardState.SetHandicap(1)

	if !boardState.IsEmpty() {
		t.Errorf("Board should be empty after setting an invalid handicap level.")
	}

	boardState.SetHandicap(2)

	if boardState.IsEmpty() {
		t.Errorf("Board should not be empty after setting a valid handicap level.")
	} else if (boardState.IsPlaceEmpty(BoardPosition{Row: 2, CrossPoint: 6}) ||
		boardState.IsPlaceEmpty(BoardPosition{Row: 6, CrossPoint: 2})) {
		t.Errorf("The selected handicap positions should not be empty.")
	} else if (boardState.GetPlace(BoardPosition{Row: 2, CrossPoint: 6}) != 1) {
		t.Errorf("The handicap should always be applied for P1.")
	}

	boardState = Initialize(19)
	boardState.SetHandicap(3)

	if boardState.IsEmpty() {
		t.Errorf("Board should not be empty after setting a valid handicap level.")
	} else if (boardState.IsPlaceEmpty(BoardPosition{Row: 3, CrossPoint: 15}) ||
		boardState.IsPlaceEmpty(BoardPosition{Row: 15, CrossPoint: 3}) ||
		boardState.IsPlaceEmpty(BoardPosition{Row: 15, CrossPoint: 15})) {
		t.Errorf("The selected handicap positions should not be empty.")
	} else if (boardState.GetPlace(BoardPosition{Row: 3, CrossPoint: 15}) != 1) {
		t.Errorf("The handicap should always be applied for P1.")
	}

	boardState = Initialize(9)
	boardState.Place(StoneP1, BoardPosition{Row: 0, CrossPoint: 0})
	boardState.SetHandicap(2)

	if (!boardState.IsPlaceEmpty(BoardPosition{Row: 2, CrossPoint: 6})) {
		t.Errorf("The handicap should not be applied if the board is not empty.")
	}
}
