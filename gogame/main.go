package main

import (
	"bufio"
	"errors"
	"fmt"
	"goingo/gogame/game"
	"os"
	"strconv"
	"strings"
)

var currentGame game.GameState
var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Welcome to the go game!\nIf you are lost, you can type in \"help\" at any time to get the list of available commands.\nTo start a game, enter \"start <X>\" with X being the size of the board.")
	for {
		fmt.Print("→ ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		args := strings.Split(input, " ")

		command := args[0]
		msg, renderBoard, err := execute(command, args[1:])
		if err != nil {
			fmt.Println(err)
		} else {
			if command == "place" {
				currentGame.EndTurn()
			}
			if msg != "" {
				fmt.Println(msg)
			}
			if renderBoard {
				currentGame.Show()
			}
		}
	}
}

func execute(command string, args []string) (msg string, renderBoard bool, err error) {
	switch command {
	case "exit":
		os.Exit(0)
	case "help":
		msg = "Available commands:\n\n help: print this message.\n start <X>: start a game. X is the size of the board.\n place <X> <Y>: place a stone at given position if possible.\n exit: leave the game."
		renderBoard = false
		err = nil
	case "start":
		if len(args) < 1 {
			err = errors.New("Too few arguments")
			return
		}
		size, _ := strconv.Atoi(args[0])
		var newGame game.GameState
		newGame, err = game.Start(size)
		if err == nil {
			currentGame = newGame
			msg = fmt.Sprintf("Started a new game with a %dx%d board.\nYou can place a stone using the command \"place <X> <Y>\".", size, size)
			renderBoard = true
		}
	default:
		msg, err = currentGame.ExecuteCommand(command, args)
		renderBoard = err == nil
	}
	return
}
