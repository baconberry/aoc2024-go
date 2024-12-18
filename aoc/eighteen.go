package aoc

import (
	"aoc2024/util"
	"math"
)

type PointPath []util.Point

func Eighteen(lines []string) int {
	return eighteen(lines, 70, 70, 1024)
}
func eighteen(lines []string, xb, yb, byteLimit int) int {
	positions := util.ParseIntGrid(lines)
	grid := util.InitGrid(util.Point{X: xb + 1, Y: yb + 1}, '.')
	byteLimit = min(len(positions), byteLimit)
	for i := 0; i < byteLimit; i++ {
		grid.SetValue(util.Point{X: positions[i][0], Y: positions[i][1]}, '#')
	}
	grid.Print(func(r rune) string {
		return string(r)
	})
	//visitedScore := make(map[util.Point]int)
	//return findPath(&grid, util.Point{}, util.Point{X: xb, Y: yb}, 0, &visitedScore)
	return dijkstra(&grid, util.Point{}, util.Point{X: xb, Y: yb})
}

var cardinalSEPriority = []util.Direction{util.E, util.S, util.N, util.W}

func findPath(g *util.Grid[rune], point util.Point, endPoint util.Point, steps int, visited *map[util.Point]int) int {
	val := g.GetOOBValue(point)
	if val == nil || *val == '#' {
		return math.MaxInt
	}

	if point == endPoint {
		return steps
	}
	prevSteps, ok := (*visited)[point]
	if ok && steps > prevSteps {
		return math.MaxInt
	}
	(*visited)[point] = steps
	minSteps := math.MaxInt
	for _, direction := range cardinalSEPriority {
		newPoint := point.PlusDirection(direction)
		localSteps := findPath(g, newPoint, endPoint, steps+1, visited)
		if localSteps < math.MaxInt {
			minSteps = min(localSteps, minSteps)
		}
	}

	return minSteps
}

func pointInPath(p *PointPath, u *util.Point) bool {
	for i := len(*p) - 1; i >= 0; i-- {
		if *u == (*p)[i] {
			return true
		}
	}
	return false
}

type PointDistance struct {
	point    util.Point
	distance int
}

func dijkstra(g *util.Grid[rune], startPoint util.Point, endPoint util.Point) int {
	distanceMap := make(map[util.Point]int)
	bounds := g.BoundPoint()
	unvisitedSet := make(map[util.Point]bool)
	for y := 0; y < bounds.Y; y++ {
		for x := 0; x < bounds.X; x++ {
			p := util.Point{X: x, Y: y}
			distanceMap[p] = math.MaxInt
			unvisitedSet[p] = true
		}
	}
	distanceMap[startPoint] = 0
	current := startPoint
	for {
		delete(unvisitedSet, current)
		if current == endPoint {
			return distanceMap[current]
		}
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
		if current == minDistNeighbor {
			panic("Invalid state")
		}
		current = minDistNeighbor
	}
}
