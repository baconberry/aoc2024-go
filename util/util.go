package util

import (
	"regexp"
	"strconv"
)

func ParseIntGrid(lines []string) [][]int {
	grid := make([][]int, 0)
	re := regexp.MustCompile("(\\d+)")

	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		parts := re.FindAllString(line, -1)
		if len(parts) == 1 && parts[0] == "" {
			break
		}
		row := make([]int, len(parts))
		for i, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				panic("Could not convert string to int " + err.Error())
			}
			row[i] = n
		}
		grid = append(grid, row)
	}
	return grid
}

func ParseIntArray(line string) []int {
	re := regexp.MustCompile("(\\d+)")
	parts := re.FindAllString(line, -1)
	arr := make([]int, 0)
	if len(parts) == 1 && parts[0] == "" {
		return arr
	}
	for _, part := range parts {
		strconv.Atoi(part)
		n, err := strconv.Atoi(part)
		if err != nil {
			panic("Could not convert string to int " + err.Error())
		}
		arr = append(arr, n)
	}
	return arr
}

func ParseRuneGrid(lines []string) [][]rune {
	grid := make([][]rune, 0)
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		grid = append(grid, []rune(line))
	}
	return grid
}
