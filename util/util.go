package util

import (
	"regexp"
	"strconv"
)

type Grid[T any] struct {
	grid          [][]T
	height, width int
}

func NewGrid[T any](grid [][]T) Grid[T] {
	height := len(grid)
	width := 0
	if height > 0 {
		width = len(grid[0])
	}
	return Grid[T]{grid, height, width}
}

func (g *Grid[T]) Iterate(fn func(int, int, T)) {
	for y, row := range g.grid {
		for x, t := range row {
			fn(x, y, t)
		}
	}
}

func (g *Grid[T]) GetNeighbor(loc Point, dir Direction) *Point {
	newPoint := loc.PlusDirection(dir)
	if newPoint.WithinBounds(g.BoundPoint()) {
		return &newPoint
	}
	return nil
}

func (g Grid[T]) BoundPoint() Point {
	return Point{g.width, g.height}
}

func (g *Grid[T]) GetValue(point Point) T {
	return g.grid[point.Y][point.X]
}

func (g *Grid[T]) Row(i int) []T {
	return g.grid[i]
}

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
func ParseDigitGrid(lines []string) [][]int {
	grid := make([][]int, 0)
	re := regexp.MustCompile("(\\d)")

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

type RuneLocation struct {
	C    rune
	X, Y int
}

func GetRuneLocations(lines []string, runeFilter func(c rune) bool) ([]RuneLocation, int, int) {
	runeLocations := make([]RuneLocation, 0)
	height := len(lines)
	width := 0
	for y, line := range lines {
		if width == 0 {
			width = len(line)
		}
		if len(line) == 0 {
			height = y
			break
		}
		for x, c := range line {
			if runeFilter(c) {
				runeLocations = append(runeLocations, RuneLocation{c, x, y})
			}
		}
	}
	return runeLocations, width, height
}

func ParseSingleDigitArray(line string) []int {
	re := regexp.MustCompile("(\\d)")
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
