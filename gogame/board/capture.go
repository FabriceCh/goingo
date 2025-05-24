package board

import (
	"goingo/shared"
)

type posInfo struct {
	Position BoardPosition
	Status   CrossPoint
}

func (board *BoardState) getNeighbours(pos BoardPosition) []BoardPosition {
	return []BoardPosition{
		{pos.Row + 1, pos.CrossPoint},
		{pos.Row - 1, pos.CrossPoint},
		{pos.Row, pos.CrossPoint - 1},
		{pos.Row, pos.CrossPoint + 1},
	}
}

type Set map[BoardPosition]struct{}

func addInSet(s Set, bp BoardPosition) {
	s[bp] = struct{}{}
}

func (s *Set) IsEmpty() bool {
	return len(s.ToArray()) == 0
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

func (board *BoardState) CheckCapture(lastPieceColor CrossPoint, lastPiecePosition BoardPosition) (points int) {

	opponentColor := getOpponentColor(lastPieceColor)

	var groups []Set
	adjacentPositions := board.getNeighbours(lastPiecePosition)
	for _, adjacentPos := range adjacentPositions {
		if !isPositionInAGroup(adjacentPos, groups) {
			groups = append(groups, board.findGroup(opponentColor, adjacentPos))
		}
	}

	points = 0
	for _, g := range groups {
		for pos := range g {
			board.Capture(pos)
			points++
		}
	}
	return points
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

		switch board.GetCrossPoint(pos) {
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
