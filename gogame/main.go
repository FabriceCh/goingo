package main

import (
	"bufio"
	"errors"
	"fmt"
	"git-gogame/gogame/game"
	"os"
	"strconv"
	"strings"
)

var currentGame game.GameState
var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Welcome to the go game!")
	for {
		fmt.Print("→ ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		args := strings.Split(input, " ")

		msg, renderBoard, err := execute(args[0], args[1:])
		if err != nil {
			fmt.Println(err)
		} else {
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
		msg = "Help command text"
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
			msg = fmt.Sprintf("Started a new game with a %dx%d board", size, size)
			renderBoard = true
		}
	default:
		msg, err = currentGame.ExecuteCommand(command, args)
		renderBoard = err == nil
	}
	return
}
