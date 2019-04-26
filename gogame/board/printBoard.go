package board

import (
	"fmt"
)

func (boardState BoardState) ShowBoard() {
	fmt.Println("Board: ")
<<<<<<< HEAD
	var dimension = boardState.Size()

	for i := 0; i < dimension; i++ {
		var line = ""
		var vertBarsLine = " "
		for j := 0; j < dimension; j++ {
			var elementAtPlace = BoardState.GetPlace(boardState, BoardPosition{Row: i, Column: j})
			var elementString string
			switch elementAtPlace {
			case Vacant:
				elementString = "   "
			case StoneP1:
				elementString = " ○ "
			case StoneP2:
				elementString = " ● "
			}

			line += elementString

			if j != (dimension - 1) {
				line += "———"
				vertBarsLine += "|     "
			} else {
				vertBarsLine += "|"
			}
		}
		fmt.Println(line)
		if i != (dimension - 1) {

			fmt.Println(vertBarsLine)
		}

=======
	var header = ""
	var body = "|"
	for i := 0; i < boardState.Size(); i++ {
		header += " _"
		body += "_|"
	}
	fmt.Println(header)
	for i := 0; i < boardState.Size(); i++ {
		fmt.Println(body)
>>>>>>> ec7e249719331ec4692a19d3547bd922aea71e56
	}
}
