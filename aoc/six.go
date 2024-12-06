package aoc

import (
	"aoc2024/util"
	"strconv"
)

func Six(lines []string) int {
	grid := util.ParseRuneGrid(lines)
	iX, iY := getInitPos(&grid)
	height := len(grid)
	width := 0
	if height > 0 {
		width = len(grid[0])
	}
	posSet := make(map[string]bool)
	posSet[strconv.Itoa(iY)+"_"+strconv.Itoa(iX)] = true
	path := make([]string, 0)
	path = walk(&grid, util.N, iX, iY, path, width, height)

	for _, s := range path {
		posSet[s] = true
	}

	return len(posSet)
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
