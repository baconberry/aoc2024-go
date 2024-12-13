package aoc

import (
	"aoc2024/util"
	"log"
	"math"
)

type PointCost struct {
	point util.Point
	cost  int
}

func Thirteen(lines []string) int {
	offset := 0
	sum := 0
	for offset < len(lines) {
		grid := util.NewGrid(util.ParseIntGrid(lines[offset:]))
		a, b, prize := gridToThreePoints(&grid)
		aCost := PointCost{a, 3}
		bCost := PointCost{b, 1}
		altCost := math.MaxInt
		for i := 0; i < 101; i++ {
			altCost = min(getCost(aCost, bCost, prize, i), altCost)
		}
		if altCost < math.MaxInt {
			sum += altCost
		}
		offset += 4
	}
	log.Println(offset)
	return sum
}

func getCost(a PointCost, b PointCost, prize util.Point, aScale int) int {
	endPoint := a.point.Multiply(aScale)
	diffPoint := endPoint.Diff(prize)
	if diffPoint.X == 0 && diffPoint.Y == 0 {
		//only a is needed
		return a.cost * aScale
	}
	numerator := diffPoint.X
	if numerator < 0 {
		return math.MaxInt
	}
	bScale := numerator / b.point.X
	if bScale == 0 {
		return math.MaxInt
	}
	if bScale > 100 {
		return math.MaxInt
	}
	bDiff := diffPoint.Diff(b.point.Multiply(bScale))
	if bDiff.X == 0 && bDiff.Y == 0 {
		aCost := a.cost * aScale
		bCost := bScale * b.cost
		return aCost + bCost
	}
	return math.MaxInt
}

func gridToThreePoints(grid *util.Grid[int]) (util.Point, util.Point, util.Point) {
	return toPoint(grid.Row(0)), toPoint(grid.Row(1)), toPoint(grid.Row(2))
}

func toPoint(row []int) util.Point {
	return util.Point{X: row[0], Y: row[1]}
}
