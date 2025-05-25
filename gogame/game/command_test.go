package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"goingo/gogame/board"
)

func TestStringToCommand(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantCmd CommandName
		wantErr bool
	}{
		{"Place command", "place", CommandPlace, false},
		{"Pass command", "pass", CommandPass, false},
		{"Handicap command", "handicap", CommandHandicap, false},
		{"Undo command", "undo", CommandUndo, false},
		{"Unknown command", "unknown", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, err := StringToCommand(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantCmd, gotCmd)
		})
	}
}

func TestNewHandicapCommand_Execute(t *testing.T) {
	gameState, err := NewGameState(19)
	assert.NoError(t, err)

	command := NewHandicapCommand(3)
	err = command.Execute(&gameState)

	assert.NoError(t, err)
	assert.False(t, gameState.board.IsEmpty(), "Board should not be empty after setting a valid handicap level.")
	assert.False(t, gameState.board.IsCrosspointEmpty(board.BoardPosition{Row: 3, CrossPoint: 15}), "Handicap position should not be empty.")
	assert.False(t, gameState.board.IsCrosspointEmpty(board.BoardPosition{Row: 15, CrossPoint: 3}), "Handicap position should not be empty.")
	assert.False(t, gameState.board.IsCrosspointEmpty(board.BoardPosition{Row: 15, CrossPoint: 15}), "Handicap position should not be empty.")
	fmt.Println(gameState.board.GetCrossPoint(board.BoardPosition{3, 15}))
	assert.Equal(t, board.StoneP1, gameState.board.GetCrossPoint(board.BoardPosition{Row: 3, CrossPoint: 15}), "Handicap stones should belong to P1.")
}

func TestNewPassCommand_Execute(t *testing.T) {
	gameState, err := NewGameState(9)
	assert.NoError(t, err)

	activePlayerBeforePass := gameState.activePlayer
	boardBeforePass := gameState.board.DeepCopy()
	command := NewPassCommand()
	err = command.Execute(&gameState)
	activePlayerAfterPass := gameState.activePlayer
	boardAfterPass := gameState.board.DeepCopy()

	assert.NoError(t, err)
	assert.NotEqual(t, activePlayerBeforePass, activePlayerAfterPass)
	assert.True(t, boardBeforePass.IsEqual(&boardAfterPass), "Board should remain the same afteer a player passes their turn.")
}

func TestNewPlaceCommand_Execute(t *testing.T) {
	gameState, err := NewGameState(9)
	assert.NoError(t, err)

	command := NewPlaceCommand(0, 0)
	err = command.Execute(&gameState)

	assert.NoError(t, err)
	assert.Equal(t, board.StoneP1, gameState.board.GetCrossPoint(board.BoardPosition{Row: 0, CrossPoint: 0}), "Crosspoint 0,0 should be filled with player 1 stone.")

	command = NewPlaceCommand(1, 0)
	err = command.Execute(&gameState)

	assert.NoError(t, err)
	assert.Equal(t, board.StoneP2, gameState.board.GetCrossPoint(board.BoardPosition{Row: 1, CrossPoint: 0}), "Crosspoint 0,0 should be filled with player 1 stone.")
	// TODO: Add more assertions based on expected state after Place execution
}

func TestNewUndoCommand_Execute(t *testing.T) {
	gameState, err := NewGameState(9)
	assert.NoError(t, err)

	// Save the current turn to have something to undo
	gameState.SaveTurn()

	command := NewUndoCommand()
	err = command.Execute(&gameState)

	assert.NoError(t, err)
	// TODO: Add more assertions based on expected state after Undo execution
}
