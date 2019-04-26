package board

import "testing"

func TestInitialize(t *testing.T) {
	boardState := Initialize()
	if boardState.Size() != 9 {
		t.Errorf("Initialized board states should be 9x9")
	} else if !boardState.IsEmpty() {
		t.Errorf("Initialized board states should be empty")
	}
}

func TestIsEmpty(t *testing.T) {
	boardState := Initialize()

	if !boardState.IsEmpty() {
		t.Errorf("IsEmpty should be true when the board is empty")
	}

	boardState.Place(StoneP1, BoardPosition{Row: 0, Column: 0})

	if boardState.IsEmpty() {
		t.Errorf("IsEmpty should be false when the board is not empty")
	}
}

func TestSize(t *testing.T) {
	boardState := Initialize()
	if boardState.Size() != 9 {
		t.Errorf("Size for a 9x9 board should be 9.")
	}
}

func TestGetPlace(t *testing.T) {
	boardState := Initialize()
	boardState.Place(StoneP1, BoardPosition{Row: 1, Column: 0})
	boardState.Place(StoneP2, BoardPosition{Row: 2, Column: 0})

	if (boardState.GetPlace(BoardPosition{Row: 0, Column: 0}) != 0) {
		t.Errorf("Position (0,0) should be 0.")
	} else if (boardState.GetPlace(BoardPosition{Row: 1, Column: 0}) != 1) {
		t.Errorf("Position (1,0) should be 0.")
	} else if (boardState.GetPlace(BoardPosition{Row: 2, Column: 0}) != 2) {
		t.Errorf("Position (2,0) should be 2.")
	}
}

func TestIsPlaceEmpty(t *testing.T) {
	boardState := Initialize()
	boardState.Place(StoneP1, BoardPosition{Row: 1, Column: 0})

	if (!boardState.IsPlaceEmpty(BoardPosition{Row: 0, Column: 0})) {
		t.Errorf("Position (0,0) should be empty.")
	} else if (boardState.IsPlaceEmpty(BoardPosition{Row: 1, Column: 0})) {
		t.Errorf("Position (1,0) should not be empty.")
	}
}

func TestPlace(t *testing.T) {
	boardState := Initialize()
	boardState.Place(StoneP1, BoardPosition{Row: 0, Column: 0})

	if (boardState.GetPlace(BoardPosition{Row: 0, Column: 0}) != 1) {
		t.Errorf("Position (0,0) should be occupied by a P1 stone.")
	}

	boardState.Place(StoneP2, BoardPosition{Row: 1, Column: 0})

	if (boardState.GetPlace(BoardPosition{Row: 0, Column: 0}) != 1) {
		t.Errorf("The stone at (0,0) should not have been modified.")
	} else if (boardState.GetPlace(BoardPosition{Row: 1, Column: 0}) != 2) {
		t.Errorf("Position (1,0) should be occupied by a P2 stone.")
	}

	boardState.Place(StoneP2, BoardPosition{Row: 0, Column: 0})

	if (boardState.GetPlace(BoardPosition{Row: 0, Column: 0}) != 1) {
		t.Errorf("Position (0,0) should still be occupied by a P1 stone.")
	}

	boardState = Initialize()
	boardState.Place(StoneP1, BoardPosition{Row: 10, Column: 0})

	if !boardState.IsEmpty() {
		t.Errorf("The board should be empty after trying to place a stone on an invalid position.")
	}
}
