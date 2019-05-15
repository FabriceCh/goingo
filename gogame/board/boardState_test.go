package board

import "testing"

func TestInitialize(t *testing.T) {
	boardState, err := Initialize(9)
	if boardState.Size() != 9 {
		t.Errorf("Initialized board states should be 9x9")
	} else if !boardState.IsEmpty() {
		t.Errorf("Initialized board states should be empty")
	}

	boardState, err = Initialize(13)
	if boardState.Size() != 13 {
		t.Errorf("Initialized board states should be 13x13")
	} else if !boardState.IsEmpty() {
		t.Errorf("Initialized board states should be empty")
	}

	boardState, err = Initialize(10)
	if err == nil {
		t.Errorf("Initializing with invalid board size should return an error")
	}
}

func TestIsEmpty(t *testing.T) {
	boardState, _ := Initialize(9)

	if !boardState.IsEmpty() {
		t.Errorf("IsEmpty should be true when the board is empty")
	}

	boardState.Place(StoneP1, BoardPosition{Row: 0, CrossPoint: 0})

	if boardState.IsEmpty() {
		t.Errorf("IsEmpty should be false when the board is not empty")
	}
}

func TestSize(t *testing.T) {
	boardState, _ := Initialize(9)
	if boardState.Size() != 9 {
		t.Errorf("Size for a 9x9 board should be 9.")
	}
}

func TestGetPlace(t *testing.T) {
	boardState, _ := Initialize(9)
	boardState.Place(StoneP1, BoardPosition{Row: 1, CrossPoint: 0})
	boardState.Place(StoneP2, BoardPosition{Row: 2, CrossPoint: 0})

	if (boardState.GetPlace(BoardPosition{Row: 0, CrossPoint: 0}) != Vacant) {
		t.Errorf("Position (0,0) should be 0.")
	} else if (boardState.GetPlace(BoardPosition{Row: 1, CrossPoint: 0}) != StoneP1) {
		t.Errorf("Position (1,0) should be 1.")
	} else if (boardState.GetPlace(BoardPosition{Row: 2, CrossPoint: 0}) != StoneP2) {
		t.Errorf("Position (2,0) should be 2.")
	}
}

func TestIsPlaceEmpty(t *testing.T) {
	boardState, _ := Initialize(9)
	boardState.Place(StoneP1, BoardPosition{Row: 1, CrossPoint: 0})

	if (!boardState.IsPlaceEmpty(BoardPosition{Row: 0, CrossPoint: 0})) {
		t.Errorf("Position (0,0) should be empty.")
	} else if (boardState.IsPlaceEmpty(BoardPosition{Row: 1, CrossPoint: 0})) {
		t.Errorf("Position (1,0) should not be empty.")
	}
}

func TestPlace(t *testing.T) {
	boardState, _ := Initialize(9)
	boardState.Place(StoneP1, BoardPosition{Row: 0, CrossPoint: 0})

	if (boardState.GetPlace(BoardPosition{Row: 0, CrossPoint: 0}) != 1) {
		t.Errorf("Position (0,0) should be occupied by a P1 stone.")
	}

	boardState.Place(StoneP2, BoardPosition{Row: 1, CrossPoint: 0})

	if (boardState.GetPlace(BoardPosition{Row: 0, CrossPoint: 0}) != 1) {
		t.Errorf("The stone at (0,0) should not have been modified.")
	} else if (boardState.GetPlace(BoardPosition{Row: 1, CrossPoint: 0}) != 2) {
		t.Errorf("Position (1,0) should be occupied by a P2 stone.")
	}

	err := boardState.Place(StoneP2, BoardPosition{Row: 0, CrossPoint: 0})

	if err == nil {
		t.Errorf("Placing a stone on another stone should return an error.")
	} else if (boardState.GetPlace(BoardPosition{Row: 0, CrossPoint: 0}) != 1) {
		t.Errorf("Position (0,0) should still be occupied by a P1 stone.")
	}

	boardState, _ = Initialize(9)
	err = boardState.Place(StoneP1, BoardPosition{Row: 10, CrossPoint: 0})

	if err == nil {
		t.Errorf("Placing a stone outside of the board should return an error.")
	} else if !boardState.IsEmpty() {
		t.Errorf("The board should be empty after trying to place a stone on an invalid position.")
	}
}
