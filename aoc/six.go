package aoc

import (
	"aoc2024/util"
	"strconv"
)

func Six(lines []string, part int) int {
	if part == 1 {
		return sixOne(lines)
	}
	return sixTwo(lines)
}
func sixOne(lines []string) int {
	grid, iX, iY, height, width := parseGrid(lines)
	posSet := make(map[string]bool)
	posSet[strconv.Itoa(iY)+"_"+strconv.Itoa(iX)] = true
	path := make([]string, 0)
	path = walk(&grid, util.N, iX, iY, path, width, height)

	for _, s := range path {
		posSet[s] = true
	}

	return len(posSet)
}

type Point struct {
	x, y      int
	direction util.Direction
}

func sixTwo(lines []string) int {
	grid, iX, iY, height, width := parseGrid(lines)

	sum := 0
	for y, rows := range grid {
		for x, c := range rows {
			if c == '.' {
				grid[y][x] = '#'
				posSet := make(map[Point]bool)
				if hasCycle(&posSet, &grid, util.N, iX, iY, width, height) {
					sum += 1
				}
				grid[y][x] = '.'
			}
		}
	}

	return sum

}

func parseGrid(lines []string) ([][]rune, int, int, int, int) {
	grid := util.ParseRuneGrid(lines)
	iX, iY := getInitPos(&grid)
	height := len(grid)
	width := 0
	if height > 0 {
		width = len(grid[0])
	}
	return grid, iX, iY, height, width
}

func walk(grid *[][]rune, d util.Direction, x int, y int, path []string, width int, height int) []string {
	dX, dY := d.CoordinatesDiff()
	nX, nY := x+dX, y+dY
	if nX < 0 || nX >= width || nY < 0 || nY >= height {
		return path
	}
	direction := d
	if (*grid)[nY][nX] == '#' {
		return walk(grid, direction.Right(), x, y, path, width, height)
	}
	nPath := append(path, strconv.Itoa(nY)+"_"+strconv.Itoa(nX))
	return walk(grid, d, nX, nY, nPath, width, height)
}

func hasCycle(posSet *map[Point]bool, grid *[][]rune, d util.Direction, x int, y int, width int, height int) bool {
	dX, dY := d.CoordinatesDiff()
	nX, nY := x+dX, y+dY
	if nX < 0 || nX >= width || nY < 0 || nY >= height {
		return false
	}
	direction := d
	if (*grid)[nY][nX] == '#' {
		return hasCycle(posSet, grid, direction.Right(), x, y, width, height)
	}
	point := Point{nX, nY, d}
	_, ok := (*posSet)[point]
	if ok {
		return true
	}
	(*posSet)[point] = true
	return hasCycle(posSet, grid, d, nX, nY, width, height)
}

func getInitPos(g *[][]rune) (int, int) {
	for y, rows := range *g {
		for x, c := range rows {
			if c == '^' {
				return x, y
			}
		}
	}
	return -1, -1
}
