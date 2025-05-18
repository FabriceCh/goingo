package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecute(t *testing.T) {
	_, _, err := execute("start", []string{"9"})
	assert.NoError(t, err)
	assert.Equal(t, 9, currentGame.GetBoardSize(), "Starting a game of size 9 should create a 9x9 board")

	_, _, err = execute("start", []string{"13"})
	assert.NoError(t, err)
	assert.Equal(t, 13, currentGame.GetBoardSize(), "Starting a game of size 13 should create a 13x13 board")

	_, _, err1 := execute("start", []string{})
	_, _, err2 := execute("start", []string{"1"})
	assert.Error(t, err1, "Missing argument should return an error")
	assert.Error(t, err2, "Invalid board size should return an error")

	_, _, err = execute("invalid", []string{})
	assert.Error(t, err, "Executing an invalid command should return an error")
}
