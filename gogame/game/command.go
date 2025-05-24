package game

import (
	"fmt"
	"goingo/gogame/board"
)

type CommandName string

const (
	CommandPlace    CommandName = "place"
	CommandPass     CommandName = "pass"
	CommandHandicap CommandName = "handicap"
	CommandUndo     CommandName = "undo"
)

func StringToCommand(str string) (command CommandName, err error) {
	switch str {
	case string(CommandPlace):
		command = CommandPlace
	case string(CommandPass):
		command = CommandPass
	case string(CommandHandicap):
		command = CommandHandicap
	case string(CommandUndo):
		command = CommandUndo
	default:
		err = fmt.Errorf("could not map '%v' to any known command", str)
	}
	return
}

type Command interface {
	Execute(gameState *GameState) error
}

type BaseCommand struct {
	Name CommandName
}

// Handicap
type HandicapCommand struct {
	BaseCommand
	Level int
}

func (p HandicapCommand) Execute(g *GameState) error {
	err := g.board.SetHandicap(p.Level)
	return err
}

func NewHandicapCommand(level int) HandicapCommand {
	return HandicapCommand{
		BaseCommand: BaseCommand{
			Name: CommandHandicap,
		},
		Level: level,
	}
}

var _ Command = HandicapCommand{}       // Verify that T implements I.
var _ Command = (*HandicapCommand)(nil) // Verify that *T implements I.

// PASS
type PassCommand struct {
	BaseCommand
}

func (p PassCommand) Execute(g *GameState) error {
	g.SaveTurn()
	g.EndTurn()
	return nil
}

func NewPassCommand() PassCommand {
	return PassCommand{
		BaseCommand: BaseCommand{
			Name: CommandPass,
		},
	}
}

var _ Command = PassCommand{}       // Verify that T implements I.
var _ Command = (*PassCommand)(nil) // Verify that *T implements I.

// PLACE
type PlaceCommand struct {
	BaseCommand
	Row        int
	CrossPoint int
}

func (p PlaceCommand) Execute(g *GameState) error {
	g.SaveTurn()
	pendingBoardState := g.board.DeepCopy()
	points, err := pendingBoardState.Place(g.activePlayer.stone, board.BoardPosition{Row: p.Row, CrossPoint: p.CrossPoint})
	if err != nil {
		return err
	}

	if g.history.Contains(pendingBoardState) {
		return fmt.Errorf("this move violates the rule of prohibition of repetition: the resulting position has occured previously in the game")
	}

	g.activePlayer.points += points
	g.board = pendingBoardState
	g.EndTurn()
	return nil
}

var _ Command = PlaceCommand{}       // Verify that T implements I.
var _ Command = (*PlaceCommand)(nil) // Verify that *T implements I.

func NewPlaceCommand(row int, crossPoint int) PlaceCommand {
	return PlaceCommand{
		BaseCommand: BaseCommand{
			Name: CommandPlace,
		},
		Row:        row,
		CrossPoint: crossPoint,
	}
}

// UNDO
type UndoCommand struct {
	BaseCommand
}

func (p UndoCommand) Execute(g *GameState) error {
	previousTurn, err := g.history.Pop()
	if err != nil {
		return err
	}
	g.LoadTurn(*previousTurn)
	return nil
}

var _ Command = UndoCommand{}       // Verify that T implements I.
var _ Command = (*UndoCommand)(nil) // Verify that *T implements I.

func NewUndoCommand() UndoCommand {
	return UndoCommand{
		BaseCommand: BaseCommand{
			Name: CommandUndo,
		},
	}
}
