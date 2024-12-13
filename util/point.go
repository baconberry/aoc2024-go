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

func (p Point) Multiply(scale int) Point {
	return Point{p.X * scale, p.Y * scale}
}

func (p Point) Diff(other Point) Point {
	return Point{other.X - p.X, other.Y - p.Y}
}

func (p Point) PlusAll(d int) Point {
	p.X += d
	p.Y += d
	return p
}
