package board

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateNumberString(n int) string {
	if n <= 0 {
		return ""
	}

	var sb strings.Builder
	for i := range n {
		sb.WriteString(strconv.Itoa(i))
		if i != n {
			sb.WriteString("    ")
		}
	}
	return "   " + sb.String()
}

func (b *BoardState) ShowBoard() {
	var dimension = b.Size()

	fmt.Println(GenerateNumberString(dimension))
	for i := range dimension {
		var line = fmt.Sprintf("%v ", i)
		var vertBarsLine = "   "
		for j := range dimension {
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
		fmt.Printf("%v %v \n", line, i)
		if i != (dimension - 1) {
			fmt.Println(vertBarsLine)
		}
	}
	fmt.Println(GenerateNumberString(dimension))
}
