package aoc

import (
	"aoc2024/util"
)

func Ten(lines []string) int {
	grid := util.NewGrid(util.ParseDigitGrid(lines))

	sum := 0
	grid.Iterate(func(x int, y int, n int) {
		if n == 0 {
			visited := make(map[util.Point]bool)
			peakLocations := make(map[util.Point]bool)
			explore(&grid, util.Point{X: x, Y: y}, &visited, &peakLocations)
			score := len(peakLocations)
			sum += score
		}
	})
	return sum
}
func explore(g *util.Grid[int], point util.Point, visitedSet *map[util.Point]bool, peaks *map[util.Point]bool) {
	_, visited := (*visitedSet)[point]
	if visited {
		return
	}
	height := g.GetValue(point)
	(*visitedSet)[point] = true
	for _, direction := range util.CARDINAL_DIRECTIONS {
		neighborPos := g.GetNeighbor(point, direction)
		if neighborPos != nil {
			neighborValue := g.GetValue(*neighborPos)
			if neighborValue-height == 1 {
				if neighborValue == 9 {
					(*peaks)[*neighborPos] = true
				} else {
					explore(g, *neighborPos, visitedSet, peaks)
				}
			}
		}
	}
}
