package main

import (
	"aoc2024/util"
)

func four(lines []string, part int) int {
	if part == 1 {
		return fourFirst(lines)
	}
	return fourSecond(lines)
}
func fourFirst(lines []string) int {
	grid, height, width := parseFour(lines)
	search := []rune{'X', 'M', 'A', 'S'}
	sum := 0
	for y, rows := range grid {
		for x, _ := range rows {
			for _, direction := range util.ALL_DIRECTIONS {
				sum += walkRuneGrid(&grid, x, y, direction, search, width, height)
			}
		}
	}
	return sum
}

func fourSecond(lines []string) int {
	grid, height, width := parseFour(lines)
	mas := []rune{'M', 'A', 'S'}
	sam := []rune{'S', 'A', 'M'}
	sum := 0
	for y, rows := range grid {
		for x, _ := range rows {
			if isAnyMatch(&grid, x, y, width, height, util.SE, &mas, &sam) {
				if isAnyMatch(&grid, x+2, y, width, height, util.SW, &mas, &sam) ||
					isAnyMatch(&grid, x, y+2, width, height, util.NE, &mas, &sam) {
					sum += 1
				}
			}
		}
	}
	return sum
}

func isAnyMatch(grid *[][]rune, x int, y int, width int, height int, direction util.Direction, searches ...*[]rune) bool {
	for _, search := range searches {
		if walkRuneGrid(grid, x, y, direction, *search, width, height) > 0 {
			return true
		}
	}
	return false
}

func parseFour(lines []string) ([][]rune, int, int) {
	grid := util.ParseRuneGrid(lines)
	height := len(grid)
	width := 0
	if height > 0 {
		width = len(grid[0])
	}
	return grid, height, width
}

func walkRuneGrid(grid *[][]rune, x int, y int, d util.Direction, search []rune, width int, height int) int {
	if len(search) == 0 {
		return 1
	}
	if x < 0 || x >= width || y < 0 || y >= height {
		return 0
	}
	sum := 0
	if (*grid)[y][x] == search[0] {
		dx, dy := d.CoordinatesDiff()
		nx, ny := x+dx, y+dy
		sum += walkRuneGrid(grid, nx, ny, d, search[1:], width, height)
	}
	return sum
}
