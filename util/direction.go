package util

type Direction int

const (
	N Direction = iota
	S
	E
	W
	NE
	NW
	SE
	SW
	NONE
)

var ALL_DIRECTIONS = [...]Direction{N, S, E, W, NE, NW, SE, SW}
var CARDINAL_DIRECTIONS = [...]Direction{N, S, E, W}

func (d Direction) CoordinatesDiff() (int, int) {
	switch d {
	case N:
		return 0, -1
	case S:
		return 0, 1
	case E:
		return 1, 0
	case W:
		return -1, 0
	case NE:
		return 1, -1
	case NW:
		return -1, -1
	case SE:
		return 1, 1
	case SW:
		return -1, 1
	default:
		return 0, 0 // return 0, 0 if an invalid direction is given
	}
}

func (d Direction) Right() Direction {
	switch d {
	case N:
		return E
	case E:
		return S
	case S:
		return W
	case W:
		return N
	}
	return NONE
}
