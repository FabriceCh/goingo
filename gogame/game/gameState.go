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
	history      History
}

type Player struct {
	name   string
	stone  board.CrossPoint
	points int
}

func NewGameState(size int) (GameState, error) {
	boardState, err := board.Initialize(size)
	player1 := &Player{name: "Player 1 (○ )", stone: board.StoneP1, points: 0}
	player2 := &Player{name: "Player 2 (● )", stone: board.StoneP2, points: 0}
	gameState := GameState{
		isActive:     true,
		board:        boardState,
		player1:      player1,
		player2:      player2,
		activePlayer: player1,
	}
	gameState.SaveTurn()
	return gameState, err
}

func (g *GameState) ExecuteCommandFromCli(commandName CommandName, args []string) (msg string, err error) {
	if !g.isActive {
		err = errors.New("no active game")
		return
	}

	var command Command
	switch commandName {
	case CommandHandicap:
		if len(args) < 1 {
			err = errors.New("too few arguments for handicap command")
			return
		}
		level, _ := strconv.Atoi(args[0])
		command = NewHandicapCommand(g.activePlayer.stone, level)
		msg = fmt.Sprintf("Set handicap of level %d for Player 1", level)
	case CommandPlace:
		if len(args) < 2 {
			err = errors.New("too few arguments for place command")
			return
		}
		row, _ := strconv.Atoi(args[0])
		crossPoint, _ := strconv.Atoi(args[1])
		command = NewPlaceCommand(g.activePlayer.stone, row, crossPoint)
		msg = fmt.Sprintf("Placed a stone at (%d,%d) for %s", row, crossPoint, g.activePlayer.name)
	case CommandPass:
		command = NewPassCommand(g.activePlayer.stone)
	case CommandUndo:
		command = NewUndoCommand()
	}

	if command == nil {
		err = errors.New("invalid command name -- if running from cli, this is a big problem")
	} else {
		err = command.Execute(g)
	}

	if err != nil {
		return "", err
	}

	return msg, err
}

func (g *GameState) switchActivePlayer() {
	if g.activePlayer == g.player1 {
		g.activePlayer = g.player2
	} else {
		g.activePlayer = g.player1
	}
}

func (g *GameState) SaveTurn() {
	newTurn := NewTurn(*g)
	g.history.Push(newTurn)
}

func (g *GameState) EndTurn() {
	g.switchActivePlayer()
}

func (g *GameState) LoadTurn(turn Turn) {
	g.board = turn.BoardState
	g.player1.points = turn.P1Points
	g.player2.points = turn.P2Points
	if turn.ActivePlayerColor == board.StoneP1 {
		g.activePlayer = g.player1
	} else {
		g.activePlayer = g.player2
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
