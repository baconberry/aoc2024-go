package main

import (
	"aoc2024/aoctypes"
	"aoc2024/util"
)

func two(lines []string) int {
	grid := util.ParseIntGrid(lines)
	grid = aoctypes.IntGrid(grid)
	result := 0
	for _, row := range grid {
		if isRowSafe(row) {
			result += 1
		}
	}

	return result
}

func twoSecond(lines []string) int {
	grid := util.ParseIntGrid(lines)
	grid = aoctypes.IntGrid(grid)
	result := 0
mainFor:
	for _, row := range grid {
		if isRowSafe(row) {
			result += 1
			continue
		}

		row := aoctypes.IntRow(row)
		for i := 0; i < len(row); i++ {
			newRow := row.RemoveElement(i)
			if isRowSafe(newRow) {
				result += 1
				continue mainFor
			}
		}
	}

	return result
}

func isRowSafe(row []int) bool {
	current := row[0]
	isAsc := true
	for i, n := range row[1:] {
		if i == 0 {
			isAsc = n-current > 0
		}
		diff := 0
		if isAsc {
			diff = n - current
		} else {
			diff = current - n
		}
		if diff < 0 {
			return false
		}
		if diff < 1 || diff > 3 {
			return false
		}
		current = n
	}

	return true
}
