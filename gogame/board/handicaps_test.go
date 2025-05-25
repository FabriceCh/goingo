package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetSmallHandicap(t *testing.T) {
	boardState, _ := Initialize(9)
	err := boardState.SetHandicap(1)

	assert.Error(t, err, "Setting an invalid handicap level should return an error.")
	assert.True(t, boardState.IsEmpty(), "Board should be empty after setting an invalid handicap level.")

	err = boardState.SetHandicap(2)
	assert.NoError(t, err)
	assert.False(t, boardState.IsEmpty(), "Board should not be empty after setting a valid handicap level.")
	assert.False(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 2, CrossPoint: 6}), "Handicap position should not be empty.")
	assert.False(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 6, CrossPoint: 2}), "Handicap position should not be empty.")
	assert.Equal(t, StoneP1, boardState.GetCrossPoint(BoardPosition{Row: 2, CrossPoint: 6}), "Handicap stones should belong to P1.")

	boardState, _ = Initialize(13)
	err = boardState.SetHandicap(2)
	assert.NoError(t, err)
	assert.False(t, boardState.IsEmpty(), "Board should not be empty after setting a valid handicap level.")
	assert.False(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 2, CrossPoint: 10}), "Handicap position should not be empty.")
	assert.False(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 10, CrossPoint: 2}), "Handicap position should not be empty.")
	assert.Equal(t, StoneP1, boardState.GetCrossPoint(BoardPosition{Row: 2, CrossPoint: 10}), "Handicap stones should belong to P1.")

	boardState, _ = Initialize(19)
	err = boardState.SetHandicap(3)
	assert.NoError(t, err)
	assert.False(t, boardState.IsEmpty(), "Board should not be empty after setting a valid handicap level.")
	assert.False(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 3, CrossPoint: 15}), "Handicap position should not be empty.")
	assert.False(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 15, CrossPoint: 3}), "Handicap position should not be empty.")
	assert.False(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 15, CrossPoint: 15}), "Handicap position should not be empty.")
	assert.Equal(t, StoneP1, boardState.GetCrossPoint(BoardPosition{Row: 3, CrossPoint: 15}), "Handicap stones should belong to P1.")

	boardState, _ = Initialize(9)
	_, _ = boardState.Place(StoneP1, BoardPosition{Row: 0, CrossPoint: 0})
	err = boardState.SetHandicap(2)
	assert.Error(t, err, "Setting a handicap on a non-empty board should return an error.")
	assert.True(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 2, CrossPoint: 6}),
		"The handicap should not be applied if the board is not empty.")

	boardState, _ = Initialize(10)
	err = boardState.SetHandicap(2)
	assert.Error(t, err, "If no handicap set is available an error should be returned.")
}
