package aoc

import "aoc2024/util"

func Twelve(lines []string) int {
	grid := util.NewGrid(util.ParseRuneGrid(lines))

	visited := make(map[util.Point]bool)
	total := 0
	grid.Iterate(func(x int, y int, c rune) {
		area, perimeter := getAreaAndPerimeter(&grid, c, util.Point{X: x, Y: y}, &visited)
		total += area * perimeter
	})
	return total
}

func getAreaAndPerimeter(grid *util.Grid[rune], c rune, point util.Point, visited *map[util.Point]bool) (int, int) {
	newRune := grid.GetValue(point)
	if newRune != c {
		return 0, 1
	}
	if (*visited)[point] {
		return 0, 0
	}

	(*visited)[point] = true

	area := 1 // itself
	perimeter := 0
	bounds := grid.BoundPoint()
	for _, direction := range util.CARDINAL_DIRECTIONS {
		newPoint := point.PlusDirection(direction)
		if newPoint.WithinBounds(bounds) {
			na, np := getAreaAndPerimeter(grid, newRune, newPoint, visited)
			area += na
			perimeter += np
		} else {
			perimeter += 1
		}
	}

	return area, perimeter
}
