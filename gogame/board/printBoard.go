package board

import (
	"fmt"
)

func (boardState State) ShowBoard() {
	var dimension = boardState.Size()

	for i := 0; i < dimension; i++ {
		var line = ""
		var vertBarsLine = " "
		for j := 0; j < dimension; j++ {
			var elementAtPlace = State.GetPlace(boardState, Position{Row: i, CrossPoint: j})
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
				line += "——"
				vertBarsLine += "|    "
			} else {
				vertBarsLine += "|"
			}
		}
		fmt.Println(line)
		if i != (dimension - 1) {
			fmt.Println(vertBarsLine)
		}
	}
}
