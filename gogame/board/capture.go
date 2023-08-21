package board

type adjacent int

const (
	top adjacent = iota
	bottom
	left
	right
)

/*func checkPosition(position Position, queue chan Position, toBeCapturedPositions []Position) CrossPoint {

}*/

func isAlreadyChecked(position Position, toBeCapturedPositions []Position) bool {
	for _, checkedPosition := range toBeCapturedPositions {
		if checkedPosition.CrossPoint == position.CrossPoint && checkedPosition.Row == position.Row {
			return true
		}
	}
	return false
}

func getAdjacentPositions(position Position) [4]Position {
	return [4]Position{
		Position{Row: position.Row, CrossPoint: position.CrossPoint + 1},
		Position{Row: position.Row, CrossPoint: position.CrossPoint - 1},
		Position{Row: position.Row + 1, CrossPoint: position.CrossPoint},
		Position{Row: position.Row - 1, CrossPoint: position.CrossPoint},
	}
}

// returns true if capture is still possible, false otherwise
func checkCapturePosition(boardState *State, position Position, toBeCapturedPositions []Position, queue chan Position, oppositeColor CrossPoint, lastPieceColor CrossPoint) bool {
	if boardState.IsPlaceEmpty(position) {
		return false
	}
	elementAtPosition := boardState.GetPlace(position)
	if elementAtPosition == oppositeColor {
		toBeCapturedPositions = append(toBeCapturedPositions, position)
		adjacentPositions := getAdjacentPositions(position)
		for _, adjacentPosition := range adjacentPositions {
			if !isAlreadyChecked(adjacentPosition, toBeCapturedPositions) {
				queue <- adjacentPosition
			}
		}
	}
	return true
}

func removePieces(boardState *State, toBeCapturedPositions []Position) {
	for _, position := range toBeCapturedPositions {
		boardState.removeStone(position)
	}
}

func getPointsAmount(toBeCapturedPositions []Position) int {
	return len(toBeCapturedPositions)
}

func getOppositeColor(lastPieceColor CrossPoint) CrossPoint {
	if lastPieceColor == StoneP1 {
		return StoneP2
	} else {
		return StoneP1
	}
}

func capture(toBeCapturedPositions []Position) {

}

func checkCapture(boardState *State, lastPieceColor CrossPoint, lastPiecePosition Position) {
	oppositeColor := getOppositeColor(lastPieceColor)

	adjacentPositions := getAdjacentPositions(lastPiecePosition)
	for _, adjacentPosition := range adjacentPositions {

		if boardState.GetPlace(adjacentPosition) == oppositeColor {
			var toBeCapturedPositions []Position
			queue := make(chan Position, 500)
			toBeCapturedPositions = append(toBeCapturedPositions, adjacentPosition)
			queue <- adjacentPosition

			isACapture := true
			for len(queue) != 0 && isACapture {
				nextPosition := <-queue
				isACapture = checkCapturePosition(boardState, nextPosition, toBeCapturedPositions, queue, oppositeColor, lastPieceColor)
			}

			if isACapture {
				//points := getPointsAmount(toBeCapturedPositions)
				// TODO: add points
				removePieces(boardState, toBeCapturedPositions)
			}
		}

	}
}
