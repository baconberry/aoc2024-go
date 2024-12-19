package util

import (
	"cmp"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Grid[T cmp.Ordered] struct {
	grid          [][]T
	height, width int
}

func NewGrid[T cmp.Ordered](grid [][]T) Grid[T] {
	height := len(grid)
	width := 0
	if height > 0 {
		width = len(grid[0])
	}
	return Grid[T]{grid, height, width}
}

func InitGrid[T cmp.Ordered](bounds Point, initValue T) Grid[T] {
	grid := make([][]T, bounds.Y)
	for y := 0; y < bounds.Y; y++ {
		grid[y] = make([]T, bounds.X)
		for x := 0; x < bounds.X; x++ {
			grid[y][x] = initValue
		}
	}
	return NewGrid(grid)
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
func (g *Grid[T]) GetOOBValue(point Point) *T {
	if point.WithinBounds(g.BoundPoint()) {
		return &g.grid[point.Y][point.X]
	}
	return nil
}

func (g *Grid[T]) AugmentAndInit(augmentBy int, initValue T) Grid[T] {
	bounds := g.BoundPoint()
	bounds.X += augmentBy
	bounds.Y += augmentBy
	return InitGrid(bounds, initValue)
}

func (g *Grid[T]) SetValue(point Point, c T) {
	g.grid[point.Y][point.X] = c
}

func (g *Grid[T]) Print(fn func(tVal T) string) {
	for _, row := range g.grid {
		s := ""
		for _, t := range row {
			s = s + " " + fn(t)
		}
		log.Println(s)
	}
}

func (g *Grid[T]) Grid() [][]T {
	return g.grid
}

func (g *Grid[T]) Width() int {
	return g.width
}
func (g *Grid[T]) Height() int {
	return g.height
}

func (g *Grid[T]) Row(i int) []T {
	return g.grid[i]
}

func (g *Grid[T]) FindFirst(val T) *Point {
	for y, row := range g.grid {
		for x, t := range row {
			if t == val {
				return &Point{x, y}
			}
		}
	}
	return nil
}

func ParseIntGrid(lines []string) [][]int {
	grid := make([][]int, 0)
	re := regexp.MustCompile("(-?\\d+)")

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
func ParseStringArray(line string) []string {
	arr := make([]string, 0)
	for _, s := range strings.Split(line, ",") {
		arr = append(arr, strings.Trim(s, " "))
	}
	return arr
}
