package board

import "testing"

func testWhiteCapturesBlack(t *testing.T, blackPiecesPositions []Position, whitePiecesPositions []Position) {
	// set this to false to remove board output from tests
	visuals := true

	boardState, _ := Initialize(9)

	// place black stones
	for _, pos := range blackPiecesPositions {
		boardState.Place(StoneP1, pos)
	}

	lastWhiteStoneIndex := len(whitePiecesPositions) - 1
	// place white stones until black's are almost captured
	for i := 0; i < lastWhiteStoneIndex; i++ {
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
		if boardState.GetPlace(pos) != StoneP2 {
			t.Errorf("Position (%d,%d) should still have white stones.", pos.Row, pos.CrossPoint)
		}
	}

	//TODO: also check if the points for the right player went up by the correct amount

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

	blackPiecesPositions := []Position{
		Position{Row: 2, CrossPoint: 2},
	}

	whitePiecesPositions := []Position{
		Position{Row: 1, CrossPoint: 2},
		Position{Row: 2, CrossPoint: 1},
		Position{Row: 2, CrossPoint: 3},
		// capturing move:
		Position{Row: 3, CrossPoint: 2},
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

	blackPiecesPositions := []Position{
		Position{Row: 0, CrossPoint: 2},
	}

	whitePiecesPositions := []Position{
		Position{Row: 0, CrossPoint: 1},
		Position{Row: 1, CrossPoint: 2},
		// capturing move:
		Position{Row: 0, CrossPoint: 3},
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

	blackPiecesPositions := []Position{
		Position{Row: 0, CrossPoint: 0},
	}

	whitePiecesPositions := []Position{
		Position{Row: 0, CrossPoint: 1},
		// capturing move:
		Position{Row: 1, CrossPoint: 0},
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

	blackPiecesPositions := []Position{
		Position{Row: 1, CrossPoint: 1},
		Position{Row: 2, CrossPoint: 1},
		Position{Row: 3, CrossPoint: 1},
		Position{Row: 1, CrossPoint: 2},
		Position{Row: 2, CrossPoint: 2},
	}

	whitePiecesPositions := []Position{
		Position{Row: 1, CrossPoint: 0},
		Position{Row: 2, CrossPoint: 0},
		Position{Row: 3, CrossPoint: 0},
		Position{Row: 0, CrossPoint: 1},
		Position{Row: 4, CrossPoint: 1},
		Position{Row: 0, CrossPoint: 2},
		Position{Row: 3, CrossPoint: 2},
		Position{Row: 1, CrossPoint: 3},
		// capturing move:
		Position{Row: 2, CrossPoint: 3},
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

	blackPiecesPositions := []Position{
		Position{Row: 2, CrossPoint: 2},
		Position{Row: 3, CrossPoint: 2},
		Position{Row: 4, CrossPoint: 2},
		Position{Row: 2, CrossPoint: 3},
		Position{Row: 4, CrossPoint: 3},
		Position{Row: 2, CrossPoint: 4},
		Position{Row: 3, CrossPoint: 4},
		Position{Row: 4, CrossPoint: 4},
	}

	whitePiecesPositions := []Position{
		Position{Row: 2, CrossPoint: 1},
		Position{Row: 3, CrossPoint: 1},
		Position{Row: 4, CrossPoint: 1},
		Position{Row: 1, CrossPoint: 2},
		Position{Row: 1, CrossPoint: 3},
		Position{Row: 1, CrossPoint: 4},
		Position{Row: 2, CrossPoint: 5},
		Position{Row: 3, CrossPoint: 5},
		Position{Row: 4, CrossPoint: 5},
		Position{Row: 5, CrossPoint: 2},
		Position{Row: 5, CrossPoint: 3},
		Position{Row: 5, CrossPoint: 4},
		// capturing move:
		Position{Row: 3, CrossPoint: 3},
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
	blackPiecesPositions := []Position{
		Position{Row: 0, CrossPoint: 1},
		Position{Row: 0, CrossPoint: 2},
		Position{Row: 1, CrossPoint: 3},
		Position{Row: 0, CrossPoint: 4},
		Position{Row: 0, CrossPoint: 5},
	}

	whitePiecesPositions := []Position{
		Position{Row: 0, CrossPoint: 0},
		Position{Row: 1, CrossPoint: 1},
		Position{Row: 1, CrossPoint: 2},
		Position{Row: 2, CrossPoint: 3},
		Position{Row: 1, CrossPoint: 4},
		Position{Row: 1, CrossPoint: 5},
		Position{Row: 0, CrossPoint: 6},
		// capturing move:
		Position{Row: 0, CrossPoint: 3},
	}

	testWhiteCapturesBlack(t, blackPiecesPositions, whitePiecesPositions)
}
