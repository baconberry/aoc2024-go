package util

type Point struct {
	X, Y int
}

func (p Point) PlusDirection(d Direction) Point {
	dx, dy := d.CoordinatesDiff()
	return Point{p.X + dx, p.Y + dy}
}

func (p Point) WithinBounds(bound Point) bool {
	if p.X < 0 || p.X >= bound.X || p.Y < 0 || p.Y >= bound.Y {
		return false
	}
	return true
}
