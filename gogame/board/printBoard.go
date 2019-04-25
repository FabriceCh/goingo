package board

import (
	"fmt"
)

func (boardState BoardState) ShowBoard() {
	fmt.Println("Board: ")
	fmt.Println(" _ _ _ _ _ _ _ _ _")
	for i := 0; i < 9; i++ {
		fmt.Println("|_|_|_|_|_|_|_|_|_|")
	}
}
