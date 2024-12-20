package aoc

import (
	"aoc2024/util"
	"math"
)

func Twenty(lines []string, part int) int {
	return twenty(lines, 100)
}

func twenty(lines []string, saveGoal int) int {
	grid := util.NewGrid(util.ParseRuneGrid(lines))
	start := grid.FindFirst('S')
	end := grid.FindFirst('E')
	lowerBound := dijkstra(&grid, *start, *end)
	distFromEnd := minDistanceFromSource(&grid, *end)
	distFromStart := minDistanceFromSource(&grid, *start)
	lowerBound = lowerBound - saveGoal
	sum := 0
	for point, startDist := range distFromStart {
		if startDist == math.MaxInt {
			continue
		}
		for _, direction := range util.CARDINAL_DIRECTIONS {
			dx, dy := direction.CoordinatesDiff()
			dP := util.Point{X: dx, Y: dy}
			dP = dP.Scale(2)
			dP = point.PlusPoint(dP)
			dPEndDist, ok := distFromEnd[dP]
			if ok && dPEndDist < math.MaxInt {
				totalDist := 2 + startDist + dPEndDist
				if totalDist <= lowerBound {
					sum += 1
				}
			}
		}
	}
	return sum
}

func minDistanceFromSource(g *util.Grid[rune], startPoint util.Point) map[util.Point]int {
	distanceMap := make(map[util.Point]int)
	bounds := g.BoundPoint()
	unvisitedSet := make(map[util.Point]bool)
	for y := 0; y < bounds.Y; y++ {
		for x := 0; x < bounds.X; x++ {
			p := util.Point{X: x, Y: y}
			distanceMap[p] = math.MaxInt
			val := g.GetValue(p)
			if val != '#' {
				unvisitedSet[p] = true
			}
		}
	}
	distanceMap[startPoint] = 0
	current := startPoint
	for len(unvisitedSet) > 0 {
		delete(unvisitedSet, current)
		nextDist := distanceMap[current] + 1
		for _, direction := range cardinalSEPriority {
			p := current.PlusDirection(direction)
			pval := g.GetOOBValue(p)
			if pval == nil || *pval == '#' {
				delete(unvisitedSet, current)
				continue
			}
			if distanceMap[p] > nextDist {
				distanceMap[p] = nextDist
			}
		}
		minDistNeighbor := current
		minDist := math.MaxInt
		for point, _ := range unvisitedSet {
			if distanceMap[point] < minDist {
				minDistNeighbor = point
				minDist = distanceMap[point]
			}
		}
		current = minDistNeighbor
	}
	return distanceMap
}
