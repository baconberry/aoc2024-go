package aoc

import (
	"aoc2024/util"
	"log"
	"strconv"
)

type Trail struct {
	path []util.Point
}

func (t Trail) String() string {
	s := ""
	for _, point := range t.path {
		s = s + strconv.Itoa(point.X) + "_" + strconv.Itoa(point.Y) + ";"
	}
	return s
}

func Ten(lines []string, part int) int {
	grid := util.NewGrid(util.ParseDigitGrid(lines))

	sum := 0
	grid.Iterate(func(x int, y int, n int) {
		if n == 0 {
			visited := make(map[util.Point]bool)
			peakLocations := make(map[util.Point]bool)
			score := 0
			startPoint := util.Point{X: x, Y: y}
			if part == 1 {
				explore(&grid, startPoint, &visited, &peakLocations)
				score = len(peakLocations)
			} else {
				trails := make(map[string]bool)
				trail := Trail{make([]util.Point, 0)}
				trail.path = append(trail.path, startPoint)
				exploreTrails(&grid, startPoint, util.NONE, trail, &trails)
				score = len(trails)
				log.Println(score)
			}
			sum += score
		}
	})
	return sum
}

func exploreTrails(g *util.Grid[int], point util.Point, prevDirection util.Direction, trail Trail, trailSet *map[string]bool) {
	height := g.GetValue(point)
	for _, direction := range util.CARDINAL_DIRECTIONS {
		if direction == prevDirection.Inverse() {
			continue
		}
		neighborPos := g.GetNeighbor(point, direction)
		if neighborPos != nil {
			neighborValue := g.GetValue(*neighborPos)
			if neighborValue-height == 1 {
				newTrail := Trail{append(trail.path, *neighborPos)}
				if neighborValue == 9 {
					(*trailSet)[newTrail.String()] = true
				} else {
					exploreTrails(g, *neighborPos, direction, newTrail, trailSet)
				}
			}
		}
	}
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
