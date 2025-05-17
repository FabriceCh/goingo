package board

import (
	"fmt"
)

func (b *BoardState) ShowBoard() {
	var dimension = b.Size()

	for i := range dimension {
		var line = ""
		var vertBarsLine = " "
		for j := 0; j < dimension; j++ {
			var elementAtPlace = b.GetPlace(BoardPosition{Row: i, CrossPoint: j})
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
