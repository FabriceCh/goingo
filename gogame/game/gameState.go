package game

import (
	"errors"
	"fmt"
	"goingo/gogame/board"
	"strconv"
)

// This will eventually handle turns and points
type GameState struct {
	isActive     bool
	board        board.BoardState
	player1      Player
	player2      Player
	activePlayer Player
}

type Player struct {
	name   string
	stone  board.CrossPoint
	points int
}

func Start(size int) (GameState, error) {
	boardState, err := board.Initialize(size)
	player1 := Player{name: "Player 1", stone: board.StoneP1, points: 0}
	player2 := Player{name: "Player 2", stone: board.StoneP2, points: 0}
	return GameState{
		isActive:     true,
		board:        boardState,
		player1:      player1,
		player2:      player2,
		activePlayer: player1,
	}, err
}

func (gameState *GameState) ExecuteCommand(command string, args []string) (msg string, err error) {
	if !gameState.isActive {
		err = errors.New("No active game")
		msg = ""
		return
	}
	switch command {
	case "handicap":
		if len(args) < 1 {
			err = errors.New("Too few arguments")
			return
		}
		level, _ := strconv.Atoi(args[0])
		err = gameState.board.SetHandicap(level)
		msg = fmt.Sprintf("Set handicap of level %d for Player 1", level)
	case "place":
		if len(args) < 2 {
			err = errors.New("Too few arguments")
			return
		}
		row, _ := strconv.Atoi(args[0])
		crossPoint, _ := strconv.Atoi(args[1])
		err = gameState.board.Place(gameState.activePlayer.stone, board.BoardPosition{Row: row, CrossPoint: crossPoint})
		msg = fmt.Sprintf("Placed a stone at (%d,%d) for %s", row, crossPoint, gameState.activePlayer.name)
	default:
		err = errors.New("Invalid command")
		msg = ""
	}
	return
}

func (gameState *GameState) EndTurn() {
	if gameState.activePlayer == gameState.player1 {
		gameState.activePlayer = gameState.player2
	} else {
		gameState.activePlayer = gameState.player1
	}
}

func (gameState GameState) Show() {
	gameState.board.ShowBoard()
	fmt.Printf("\n%s's turn:\n", gameState.activePlayer.name)
}

func (gameState GameState) GetBoard() board.BoardState {
	return gameState.board
}

func (g GameState) GetBoardSize() int {
	return g.board.Size()
}
