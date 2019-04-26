package main

import (
	"fmt"
	"git-gogame/gogame/game"
)

var currentGame game.GameState

func main() {
	fmt.Println("Welcome to the go game!")
	commands := []string{"help", "start", "handicap", "handicap", "place", "place"}
	for _, command := range commands {
		msg, shouldRender, err := execute(command)
		if err != nil {
			fmt.Println(err)
		} else {
			if msg != "" {
				fmt.Println(msg)
			}
			if shouldRender {
				currentGame.Board.ShowBoard()
			}
		}

	}
}

func execute(command string, args ...string) (msg string, shouldRender bool, err error) {
	switch command {
	case "help":
		msg = "Help command text."
		shouldRender = false
		err = nil
	case "start":
		var newGame game.GameState
		newGame, err = game.Start(9)
		if err == nil {
			currentGame = newGame
			msg = "Started a new game with a 9x9 board."
			shouldRender = true
		}
	default:
		msg, err = currentGame.ExecuteCommand(command)
		shouldRender = err == nil
	}
	return
}
