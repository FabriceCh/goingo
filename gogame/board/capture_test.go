package board

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testWhiteCapturesBlack(t *testing.T, blackPiecesPositions []BoardPosition, whitePiecesPositions []BoardPosition) {
	// set this to false to remove board output from tests
	visuals := false

	boardState, _ := Initialize(9)

	// place black stones
	for _, pos := range blackPiecesPositions {
		boardState.Place(StoneP1, pos)
	}

	lastWhiteStoneIndex := len(whitePiecesPositions) - 1
	// place white stones until black's are almost captured
	for i := range lastWhiteStoneIndex {
		boardState.Place(StoneP2, whitePiecesPositions[i])
	}

	for _, pos := range blackPiecesPositions {
		if boardState.IsPlaceEmpty(pos) {
			t.Errorf("Position (%d,%d) should not be empty yet.", pos.Row, pos.CrossPoint)
		}
	}

	if visuals {
		boardState.ShowBoard()
	}

	// make the capturing move
	boardState.Place(StoneP2, whitePiecesPositions[lastWhiteStoneIndex])

	if visuals {
		boardState.ShowBoard()
	}
	for _, pos := range blackPiecesPositions {
		if !boardState.IsPlaceEmpty(pos) {
			t.Errorf("Position (%d,%d) should be empty because of a capture.", pos.Row, pos.CrossPoint)
		}
	}

	for _, pos := range whitePiecesPositions {
		if boardState.GetCrossPoint(pos) != StoneP2 {
			t.Errorf("Position (%d,%d) should still have white stones.", pos.Row, pos.CrossPoint)
		}
	}

}

func TestSimpleSingleCapture(t *testing.T) {
	/*
	    |    |    |
	 ——   —— ● ——   ——
	    |    |    |
	 —— ● —— ○ —— ● ——
	    |    |    |
	 ——   —— x ——   ——
	    |    |    |
	*/

	blackPiecesPositions := []BoardPosition{
		BoardPosition{Row: 2, CrossPoint: 2},
	}

	whitePiecesPositions := []BoardPosition{
		BoardPosition{Row: 1, CrossPoint: 2},
		BoardPosition{Row: 2, CrossPoint: 1},
		BoardPosition{Row: 2, CrossPoint: 3},
		// capturing move:
		BoardPosition{Row: 3, CrossPoint: 2},
	}

	testWhiteCapturesBlack(t, blackPiecesPositions, whitePiecesPositions)
}

func TestSingleWallCapture(t *testing.T) {
	/*
	   —— ● —— ○ —— x ——
	 |    |    |    |
	   ——   —— ● ——   ——
	 |    |    |    |
	*/

	blackPiecesPositions := []BoardPosition{
		BoardPosition{Row: 0, CrossPoint: 2},
	}

	whitePiecesPositions := []BoardPosition{
		BoardPosition{Row: 0, CrossPoint: 1},
		BoardPosition{Row: 1, CrossPoint: 2},
		// capturing move:
		BoardPosition{Row: 0, CrossPoint: 3},
	}

	testWhiteCapturesBlack(t, blackPiecesPositions, whitePiecesPositions)
}

func TestSingleCornerCapture(t *testing.T) {
	/*
		○ —— ● ——
		|    |
		x ——
		|
	*/

	blackPiecesPositions := []BoardPosition{
		BoardPosition{Row: 0, CrossPoint: 0},
	}

	whitePiecesPositions := []BoardPosition{
		BoardPosition{Row: 0, CrossPoint: 1},
		// capturing move:
		BoardPosition{Row: 1, CrossPoint: 0},
	}

	testWhiteCapturesBlack(t, blackPiecesPositions, whitePiecesPositions)
}

func TestGroupCapture(t *testing.T) {
	/*
		  —— ● —— ● ——   ——
		|    |    |    |
		● —— ○ —— ○ —— ● ——
		|    |    |    |
		● —— ○ —— ○ —— x ——
		|    |    |    |
		● —— ○ —— ● ——   ——
		|    |    |    |
		  —— ● ——   ——   ——
		|    |    |    |
	*/

	blackPiecesPositions := []BoardPosition{
		BoardPosition{Row: 1, CrossPoint: 1},
		BoardPosition{Row: 2, CrossPoint: 1},
		BoardPosition{Row: 3, CrossPoint: 1},
		BoardPosition{Row: 1, CrossPoint: 2},
		BoardPosition{Row: 2, CrossPoint: 2},
	}

	whitePiecesPositions := []BoardPosition{
		BoardPosition{Row: 1, CrossPoint: 0},
		BoardPosition{Row: 2, CrossPoint: 0},
		BoardPosition{Row: 3, CrossPoint: 0},
		BoardPosition{Row: 0, CrossPoint: 1},
		BoardPosition{Row: 4, CrossPoint: 1},
		BoardPosition{Row: 0, CrossPoint: 2},
		BoardPosition{Row: 3, CrossPoint: 2},
		BoardPosition{Row: 1, CrossPoint: 3},
		// capturing move:
		BoardPosition{Row: 2, CrossPoint: 3},
	}

	testWhiteCapturesBlack(t, blackPiecesPositions, whitePiecesPositions)

}

func TestEyeCapture(t *testing.T) {
	/*
	   ——   ——   ——   ——   ——   ——   ——   ——
	 |    |    |    |    |    |    |    |    |
	   ——   —— ● —— ● —— ● ——   ——   ——   ——
	 |    |    |    |    |    |    |    |    |
	   —— ● —— ○ —— ○ —— ○ —— ● ——   ——   ——
	 |    |    |    |    |    |    |    |    |
	   —— ● —— ○ —— x —— ○ —— ● ——   ——   ——
	 |    |    |    |    |    |    |    |    |
	   —— ● —— ○ —— ○ —— ○ —— ● ——   ——   ——
	 |    |    |    |    |    |    |    |    |
	   ——   —— ● —— ● —— ● ——   ——   ——   ——
	 |    |    |    |    |    |    |    |    |
	   ——   ——   ——   ——   ——   ——   ——   ——
	 |    |    |    |    |    |    |    |    |
	   ——   ——   ——   ——   ——   ——   ——   ——
	 |    |    |    |    |    |    |    |    |
	   ——   ——   ——   ——   ——   ——   ——   ——

	*/

	blackPiecesPositions := []BoardPosition{
		BoardPosition{Row: 2, CrossPoint: 2},
		BoardPosition{Row: 3, CrossPoint: 2},
		BoardPosition{Row: 4, CrossPoint: 2},
		BoardPosition{Row: 2, CrossPoint: 3},
		BoardPosition{Row: 4, CrossPoint: 3},
		BoardPosition{Row: 2, CrossPoint: 4},
		BoardPosition{Row: 3, CrossPoint: 4},
		BoardPosition{Row: 4, CrossPoint: 4},
	}

	whitePiecesPositions := []BoardPosition{
		BoardPosition{Row: 2, CrossPoint: 1},
		BoardPosition{Row: 3, CrossPoint: 1},
		BoardPosition{Row: 4, CrossPoint: 1},
		BoardPosition{Row: 1, CrossPoint: 2},
		BoardPosition{Row: 1, CrossPoint: 3},
		BoardPosition{Row: 1, CrossPoint: 4},
		BoardPosition{Row: 2, CrossPoint: 5},
		BoardPosition{Row: 3, CrossPoint: 5},
		BoardPosition{Row: 4, CrossPoint: 5},
		BoardPosition{Row: 5, CrossPoint: 2},
		BoardPosition{Row: 5, CrossPoint: 3},
		BoardPosition{Row: 5, CrossPoint: 4},
		// capturing move:
		BoardPosition{Row: 3, CrossPoint: 3},
	}

	testWhiteCapturesBlack(t, blackPiecesPositions, whitePiecesPositions)

}

func TestMultipleCapturedGroups(t *testing.T) {
	/*
			● —— ○ —— ○ —— x —— ○ —— ○ —— ● ——
		 	|    |    |    |    |    |    |
			  —— ● —— ● —— ○ —— ● —— ● ——   ——
		 	|    |    |    |    |    |    |
		   	  ——   ——   —— ● ——   ——   ——   ——
		 	|    |    |    |    |    |    |
	*/
	blackPiecesPositions := []BoardPosition{
		BoardPosition{Row: 0, CrossPoint: 1},
		BoardPosition{Row: 0, CrossPoint: 2},
		BoardPosition{Row: 1, CrossPoint: 3},
		BoardPosition{Row: 0, CrossPoint: 4},
		BoardPosition{Row: 0, CrossPoint: 5},
	}

	whitePiecesPositions := []BoardPosition{
		BoardPosition{Row: 0, CrossPoint: 0},
		BoardPosition{Row: 1, CrossPoint: 1},
		BoardPosition{Row: 1, CrossPoint: 2},
		BoardPosition{Row: 2, CrossPoint: 3},
		BoardPosition{Row: 1, CrossPoint: 4},
		BoardPosition{Row: 1, CrossPoint: 5},
		BoardPosition{Row: 0, CrossPoint: 6},
		// capturing move:
		BoardPosition{Row: 0, CrossPoint: 3},
	}

	testWhiteCapturesBlack(t, blackPiecesPositions, whitePiecesPositions)
}

func testFindGroup(t *testing.T, blackPiecesPositions []BoardPosition, whitePiecesPositions []BoardPosition, empty bool) {
	boardState, _ := Initialize(9)
	for _, pos := range blackPiecesPositions {
		boardState.Place(StoneP1, pos)
	}
	for _, pos := range whitePiecesPositions {
		boardState.Place(StoneP2, pos)
	}
	group := boardState.findGroup(StoneP1, blackPiecesPositions[0])
	if empty {
		assert.Empty(t, group, fmt.Sprintf("group should be empty but was of length %v", len(group)))
		if len(group) != 0 {
			t.Errorf("group should be empty but was of length %v", len(group))
		}
	} else if len(group) != len(blackPiecesPositions) {
		t.Errorf("found group is not the right length, expected %v, got %v", len(blackPiecesPositions), len(group))
	}
}

func TestGroupShouldBeEmpty(t *testing.T) {

	blackPiecesPositions := []BoardPosition{
		{Row: 2, CrossPoint: 2},
		{Row: 1, CrossPoint: 2},
	}

	whitePiecesPositions := []BoardPosition{
		{Row: 2, CrossPoint: 1},
		{Row: 2, CrossPoint: 3},
		// capturing move:
		{Row: 3, CrossPoint: 2},
	}
	testFindGroup(t, blackPiecesPositions, whitePiecesPositions, true)
}
