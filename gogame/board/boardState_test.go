package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	boardState, err := Initialize(9)
	assert.Equal(t, 9, boardState.Size())
	assert.True(t, boardState.IsEmpty())
	assert.NoError(t, err)

	boardState, err = Initialize(13)
	assert.Equal(t, 13, boardState.Size())
	assert.True(t, boardState.IsEmpty())
	assert.NoError(t, err)

	// 10 is not a valid option
	_, err = Initialize(10)
	assert.Error(t, err)
}

func TestIsEmpty(t *testing.T) {
	boardState, _ := Initialize(9)
	assert.True(t, boardState.IsEmpty())

	points, err := boardState.Place(StoneP1, BoardPosition{Row: 0, CrossPoint: 0})
	assert.False(t, boardState.IsEmpty())
	assert.Zero(t, points)
	assert.NoError(t, err)
}

func TestSize(t *testing.T) {
	boardState, _ := Initialize(9)
	assert.Equal(t, 9, boardState.Size())
}

func TestGetPlace(t *testing.T) {
	boardState, _ := Initialize(9)
	boardState.Place(StoneP1, BoardPosition{Row: 1, CrossPoint: 0})
	boardState.Place(StoneP2, BoardPosition{Row: 2, CrossPoint: 0})

	assert.Equal(t, Vacant, boardState.GetCrossPoint(BoardPosition{Row: 0, CrossPoint: 0}))
	assert.Equal(t, StoneP1, boardState.GetCrossPoint(BoardPosition{Row: 1, CrossPoint: 0}))
	assert.Equal(t, StoneP2, boardState.GetCrossPoint(BoardPosition{Row: 2, CrossPoint: 0}))
}

func TestIsPlaceEmpty(t *testing.T) {
	boardState, _ := Initialize(9)
	boardState.Place(StoneP1, BoardPosition{Row: 1, CrossPoint: 0})

	assert.True(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 0, CrossPoint: 0}))
	assert.False(t, boardState.IsCrosspointEmpty(BoardPosition{Row: 1, CrossPoint: 0}))
}

func TestPlace(t *testing.T) {
	boardState, _ := Initialize(9)
	boardState.Place(StoneP1, BoardPosition{Row: 0, CrossPoint: 0})
	assert.Equal(t, StoneP1, boardState.GetCrossPoint(BoardPosition{Row: 0, CrossPoint: 0}),
		"Position (0,0) should be occupied by a P1 stone.")

	boardState.Place(StoneP2, BoardPosition{Row: 1, CrossPoint: 0})

	assert.Equal(t, StoneP1, boardState.GetCrossPoint(BoardPosition{Row: 0, CrossPoint: 0}),
		"The stone at (0,0) should not have been modified.")
	assert.Equal(t, StoneP2, boardState.GetCrossPoint(BoardPosition{Row: 1, CrossPoint: 0}),
		"Position (1,0) should be occupied by a P2 stone.")

	_, err := boardState.Place(StoneP2, BoardPosition{Row: 0, CrossPoint: 0})
	assert.Error(t, err, "Placing a stone on another stone should return an error.")
	assert.Equal(t, StoneP1, boardState.GetCrossPoint(BoardPosition{Row: 0, CrossPoint: 0}),
		"Position (0,0) should still be occupied by a P1 stone.")

	boardState, _ = Initialize(9)
	_, err = boardState.Place(StoneP1, BoardPosition{Row: 10, CrossPoint: 0})
	assert.Error(t, err, "Placing a stone outside of the board should return an error.")
	assert.True(t, boardState.IsEmpty(),
		"The board should be empty after trying to place a stone on an invalid position.")
}

func TestDeepCopy(t *testing.T) {
	boardState, _ := Initialize(9)
	boardState.Place(StoneP1, BoardPosition{Row: 0, CrossPoint: 0})
	boardState.Place(StoneP2, BoardPosition{Row: 1, CrossPoint: 0})

	copiedBoard := boardState.DeepCopy()
	assert.Equal(t, StoneP1, copiedBoard.GetCrossPoint(BoardPosition{Row: 0, CrossPoint: 0}),
		"Position (0,0) in copied board should be occupied by a P1 stone.")
	assert.Equal(t, StoneP2, copiedBoard.GetCrossPoint(BoardPosition{Row: 1, CrossPoint: 0}),
		"Position (1,0) in copied board should be occupied by a P2 stone.")
	assert.Equal(t, Vacant, copiedBoard.GetCrossPoint(BoardPosition{Row: 1, CrossPoint: 1}),
		"Position (1,1) in copied board should be vacant.")

	boardState.Place(StoneP1, BoardPosition{Row: 2, CrossPoint: 2})
	assert.Equal(t, Vacant, copiedBoard.GetCrossPoint(BoardPosition{Row: 2, CrossPoint: 2}),
		"Position (2,2) in copied board should remain unchanged.")

}
