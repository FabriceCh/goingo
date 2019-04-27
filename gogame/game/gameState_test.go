package game

import "testing"

func TestStart(t *testing.T) {
	gameState, err := Start(9)
	if gameState.Board.Size() != 9 {
		t.Errorf("A started game of size 9 should have a 9x9 game board")
	}

	gameState, err = Start(13)
	if gameState.Board.Size() != 13 {
		t.Errorf("A started game of size 13 should have a 13x13 game board")
	}

	gameState, err = Start(10)
	if err == nil {
		t.Errorf("Starting a game of invalid size should return an error")
	}
}

func TestExecuteCommand(t *testing.T) {
	gameState, err := Start(9)

	gameState.ExecuteCommand("handicap", []string{"2"})
	if gameState.Board.IsEmpty() {
		t.Errorf("The game board should not be empty after putting a handicap")
	}

	gameState, _ = Start(9)
	gameState.ExecuteCommand("place", []string{"1", "1"})
	if gameState.Board.IsEmpty() {
		t.Errorf("The game board should not be empty after placing a stone")
	}

	_, err = gameState.ExecuteCommand("place", []string{"1", "1"})
	if err == nil {
		t.Errorf("Placing a stone over another one should return an error")
	}

	_, err = gameState.ExecuteCommand("invalid", []string{})
	if err == nil {
		t.Errorf("Executing an invalid command should return an error")
	}

	_, err = gameState.ExecuteCommand("handicap", []string{})
	_, err2 := gameState.ExecuteCommand("handicap", []string{"asd"})
	_, err3 := gameState.ExecuteCommand("place", []string{})
	_, err4 := gameState.ExecuteCommand("place", []string{"asd"})
	if err == nil || err2 == nil || err3 == nil || err4 == nil {
		t.Errorf("Passing invalid arguments to the command should throw an error")
	}
}
