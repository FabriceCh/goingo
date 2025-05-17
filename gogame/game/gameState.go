package game

import (
	"errors"
	"fmt"
	"goingo/gogame/board"
	"strconv"
)

type GameState struct {
	isActive     bool
	board        board.BoardState
	player1      *Player
	player2      *Player
	activePlayer *Player
}

type Player struct {
	name   string
	stone  board.CrossPoint
	points int
}

func Start(size int) (GameState, error) {
	boardState, err := board.Initialize(size)
	player1 := Player{name: "Player 1 (○ )", stone: board.StoneP1, points: 0}
	player2 := Player{name: "Player 2 (● )", stone: board.StoneP2, points: 0}
	return GameState{
		isActive:     true,
		board:        boardState,
		player1:      &player1,
		player2:      &player2,
		activePlayer: &player1,
	}, err
}

func (g *GameState) ExecuteCommand(command string, args []string) (msg string, err error) {
	if !g.isActive {
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
		err = g.board.SetHandicap(level)
		msg = fmt.Sprintf("Set handicap of level %d for Player 1", level)
	case "place":
		if len(args) < 2 {
			err = errors.New("Too few arguments")
			return
		}
		row, _ := strconv.Atoi(args[0])
		crossPoint, _ := strconv.Atoi(args[1])
		err := g.Place(row, crossPoint)
		if err != nil {
			return "", err
		}
		msg = fmt.Sprintf("Placed a stone at (%d,%d) for %s", row, crossPoint, g.activePlayer.name)
	default:
		err = errors.New("Invalid command")
		msg = ""
	}
	return msg, err
}

func (g *GameState) EndTurn() {
	if g.activePlayer == g.player1 {
		g.activePlayer = g.player2
	} else {
		g.activePlayer = g.player1
	}
}

func (g GameState) Show() {
	g.board.ShowBoard()
	fmt.Printf("\n%s's turn:\n", g.activePlayer.name)
	fmt.Printf("Points: %s: %v %s: %v\n", g.player1.name, g.player1.points, g.player2.name, g.player2.points)
}

func (g GameState) GetBoard() board.BoardState {
	return g.board
}

func (g GameState) GetBoardSize() int {
	return g.board.Size()
}

func (g GameState) Place(row int, crossPoint int) error {
	points, err := g.board.Place(g.activePlayer.stone, board.BoardPosition{Row: row, CrossPoint: crossPoint})
	if err != nil {
		return err
	}
	g.activePlayer.points += points
	return nil
}
