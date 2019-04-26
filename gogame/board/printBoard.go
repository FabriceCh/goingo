package board

import (
	"fmt"
)

func (boardState BoardState) ShowBoard() {
	fmt.Println("Board: ")
	var header = ""
	var body = "|"
	for i := 0; i < boardState.Size(); i++ {
		header += " _"
		body += "_|"
	}
	fmt.Println(header)
	for i := 0; i < boardState.Size(); i++ {
		fmt.Println(body)
	}
}
