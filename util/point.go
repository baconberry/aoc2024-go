package util

import "math"

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

func (p Point) PlusPoint(o Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}

func (p Point) Scale(s int) Point {
	d := p
	d.X *= s
	d.Y *= s
	return d
}

func (p Point) DistanceTo(b Point) int {
	return int(math.Abs(float64(b.X-p.X)) + math.Abs(float64(b.Y-p.Y)))
}

func (p Point) Distance(other Point) int {
	dist := p.Diff(other)
	distance := 0
	if dist.X < 0 {
		distance += dist.X * -1
	} else {
		distance += dist.X
	}
	if dist.Y < 0 {
		distance += dist.Y * -1
	} else {
		distance += dist.Y
	}
	return distance
}
