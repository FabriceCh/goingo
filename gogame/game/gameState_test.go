package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGameState(t *testing.T) {
	gameState, err := NewGameState(9)
	assert.NoError(t, err)
	assert.Equal(t, 9, gameState.board.Size(), "A started game of size 9 should have a 9x9 game board")
	assert.Equal(t, gameState.player1, gameState.activePlayer, "The first active player should always be Player 1")

	gameState, err = NewGameState(13)
	assert.NoError(t, err)
	assert.Equal(t, 13, gameState.board.Size(), "A started game of size 13 should have a 13x13 game board")

	gameState, err = NewGameState(10)
	assert.Error(t, err, "Starting a game of invalid size should return an error")
}

func TestExecuteCommand(t *testing.T) {
	gameState, err := NewGameState(9)
	assert.NoError(t, err)

	_, err = gameState.ExecuteCommandFromCli("handicap", []string{"2"})
	assert.NoError(t, err)
	assert.False(t, gameState.board.IsEmpty(), "The game board should not be empty after putting a handicap")

	gameState, err = NewGameState(9)
	assert.NoError(t, err)

	_, err = gameState.ExecuteCommandFromCli("place", []string{"1", "1"})
	assert.NoError(t, err)
	assert.False(t, gameState.board.IsEmpty(), "The game board should not be empty after placing a stone")

	_, err = gameState.ExecuteCommandFromCli("place", []string{"1", "1"})
	assert.Error(t, err, "Placing a stone over another one should return an error")

	_, err = gameState.ExecuteCommandFromCli("invalid", []string{})
	assert.Error(t, err, "Executing an invalid command should return an error")

	_, err1 := gameState.ExecuteCommandFromCli("handicap", []string{})
	_, err2 := gameState.ExecuteCommandFromCli("handicap", []string{"asd"})
	_, err3 := gameState.ExecuteCommandFromCli("place", []string{})
	_, err4 := gameState.ExecuteCommandFromCli("place", []string{"asd"})

	assert.Error(t, err1)
	assert.Error(t, err2)
	assert.Error(t, err3)
	assert.Error(t, err4)
}

func TestExecutePassCommand(t *testing.T) {
	gameState, err := NewGameState(9)
	assert.NoError(t, err)

	_, err = gameState.ExecuteCommandFromCli("pass", []string{})
	assert.NoError(t, err)
	assert.True(t, gameState.board.IsEmpty(), "The game board should still be empty after player1 passed")
	assert.Equal(t, gameState.activePlayer, gameState.player2, "Turn should be to the next player")
}

func TestSwitchPlayerTurn(t *testing.T) {
	gameState, err := NewGameState(9)
	assert.NoError(t, err)

	gameState.switchActivePlayer()
	assert.Equal(t, gameState.player2, gameState.activePlayer, "Switching should make player 2 the active player")

	gameState.switchActivePlayer()
	assert.Equal(t, gameState.player1, gameState.activePlayer, "Switching should make player 1 the active player")
}

