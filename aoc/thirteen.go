package aoc

import (
	"aoc2024/util"
	"log"
)

type PointCost struct {
	point util.Point
	cost  int
}

func Thirteen(lines []string, part int) int {
	offset := 0
	sum := 0
	for offset < len(lines) {
		grid := util.NewGrid(util.ParseIntGrid(lines[offset:]))
		ap, bp, prize := gridToThreePoints(&grid)
		if part == 2 {
			prize.X += 10000000000000
			prize.Y += 10000000000000
		}

		a, b := calculateMinCost(ap.X, ap.Y, bp.X, bp.Y, prize.X, prize.Y)
		diff := ap.Multiply(a).Diff(prize)
		diff = diff.Diff(bp.Multiply(b))
		if diff.X == 0 && diff.Y == 0 {
			sum += (a * 3) + b
		}
		offset += 4
	}
	log.Println(offset)
	return sum
}

func calculateMinCost(ax int, ay int, bx int, by int, x int, y int) (int, int) {
	a := ((y * bx) - (x * by)) / ((bx * ay) - (ax * by))
	b := ((y * ax) - (ay * x)) / ((by * ax) - (bx * ay))
	return a, b
}

func gridToThreePoints(grid *util.Grid[int]) (util.Point, util.Point, util.Point) {
	return toPoint(grid.Row(0)), toPoint(grid.Row(1)), toPoint(grid.Row(2))
}

func toPoint(row []int) util.Point {
	return util.Point{X: row[0], Y: row[1]}
}
