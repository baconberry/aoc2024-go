package util

import (
	"regexp"
	"strconv"
)

func ParseIntGrid(lines []string) [][]int {
	grid := make([][]int, 0)
	//row := make([]int, 0)

	re := regexp.MustCompile("(\\d+)*")

	for _, line := range lines {
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
