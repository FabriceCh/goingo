package main

import "testing"

func TestExecute(t *testing.T) {
	execute("start", []string{"9"})
	if currentGame.GetBoard().Size() != 9 {
		t.Errorf("Starting a game of size 9 should create a 9x9 board")
	}

	execute("start", []string{"13"})
	if currentGame.GetBoard().Size() != 13 {
		t.Errorf("Starting a game of size 13 should create a 13x13 board")
	}

	_, _, err := execute("start", []string{})
	_, _, err2 := execute("start", []string{"1"})
	if err == nil || err2 == nil {
		t.Errorf("Passing invalid arguments should return an error")
	}

	_, _, err = execute("invalid", []string{})
	if err == nil {
		t.Errorf("Executing an invalid command should return an error")
	}
}
