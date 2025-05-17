package board

import (
	"goingo/shared"
)

type posInfo struct {
	Position BoardPosition
	Status   CrossPoint
}

func (board *BoardState) getNeighbours(pos BoardPosition) []BoardPosition {
	var up, down, left, right BoardPosition
	up = BoardPosition{Row: pos.Row + 1, CrossPoint: pos.CrossPoint}
	down = BoardPosition{Row: pos.Row - 1, CrossPoint: pos.CrossPoint}
	left = BoardPosition{Row: pos.Row, CrossPoint: pos.CrossPoint - 1}
	right = BoardPosition{Row: pos.Row, CrossPoint: pos.CrossPoint + 1}
	return []BoardPosition{up, down, left, right}
}

type Set map[BoardPosition]struct{}

func addInSet(s Set, bp BoardPosition) {
	s[bp] = struct{}{}
}

func (s *Set) ToArray() []BoardPosition {
	arr := make([]BoardPosition, 0)
	for pos := range *s {
		arr = append(arr, pos)
	}
	return arr
}

func getOpponentColor(playerColor CrossPoint) CrossPoint {
	if playerColor == StoneP1 {
		return StoneP2
	} else {
		return StoneP1
	}
}

func (board *BoardState) CheckCapture(lastPieceColor CrossPoint, lastPiecePosition BoardPosition) {

	opponentColor := getOpponentColor(lastPieceColor)

	var groups []Set
	adjacentPositions := board.getNeighbours(lastPiecePosition)
	for _, adjacentPos := range adjacentPositions {
		// skip to avoid duplicated groups
		if isPositionInAGroup(adjacentPos, groups) {
			continue
		}
		groups = append(groups, board.findGroup(opponentColor, adjacentPos))
	}

	for _, g := range groups {
		for pos := range g {
			board.Capture(pos)
		}
	}
}

func isPositionInAGroup(pos BoardPosition, groups []Set) bool {
	for _, g := range groups {
		_, found := g[pos]
		if found {
			return true
		}
	}
	return false
}

func (board *BoardState) findGroup(playerColor CrossPoint, lastPos BoardPosition) Set {
	queue := shared.NewQueue[BoardPosition]()
	seen := Set{}
	group := Set{}
	queue.Insert(lastPos)

	for !queue.IsEmpty() {
		pos := queue.Pop()

		// do nothing if position was already seen
		_, alreadySeen := seen[pos]
		if alreadySeen {
			continue
		}
		addInSet(seen, pos)

		switch board.GetPlace(pos) {
		case playerColor:
			// add to group and continue searching
			addInSet(group, pos)
			queue.InsertMany(board.getNeighbours(pos))
		case Vacant:
			// the whole group is not capturable, so return empty group
			return Set{}
		default:
			// wall or other player color
		}
	}
	return group
}
