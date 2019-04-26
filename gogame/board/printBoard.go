package board

import (
	"fmt"
)

func (boardState BoardState) ShowBoard() {
	for i := range boardState.Rows {
		crossPoints := ""
		separators := ""
		for j, crossPoint := range boardState.Rows[i].CrossPoints {
			if j < boardState.Size()-1 {
				crossPoints += fmt.Sprintf("%d—", crossPoint)
			} else {
				crossPoints += fmt.Sprintf("%d", crossPoint)
			}
			separators += "| "
		}

		fmt.Println(crossPoints)
		if i < boardState.Size()-1 {
			fmt.Println(separators)
		}
	}
}
