package aoc

import (
	"aoc2024/util"
	"log"
	"math"
)

func Eighteen(lines []string, part int) int {
	return eighteen(lines, 70, 70, 1024, part)
}
func eighteen(lines []string, xb, yb, byteLimit, part int) int {
	positions := util.ParseIntGrid(lines)
	grid := util.InitGrid(util.Point{X: xb + 1, Y: yb + 1}, '.')
	byteLimit = min(len(positions), byteLimit)
	for i := 0; i < byteLimit; i++ {
		grid.SetValue(util.Point{X: positions[i][0], Y: positions[i][1]}, '#')
	}
	//grid.Print(func(r rune) string {
	//	return string(r)
	//})
	if part == 2 {
		for i := byteLimit; i < len(positions); i++ {
			block := util.Point{X: positions[i][0], Y: positions[i][1]}
			grid.SetValue(block, '#')
			steps := dijkstra(&grid, util.Point{}, util.Point{X: xb, Y: yb})
			if steps == math.MaxInt {
				log.Println("Blocked with:", block)
				return math.MaxInt
			}
		}
	}
	return dijkstra(&grid, util.Point{}, util.Point{X: xb, Y: yb})
}

var cardinalSEPriority = []util.Direction{util.E, util.S, util.N, util.W}

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
			log.Print("Invalid state, returning inf")
			return math.MaxInt
		}
		current = minDistNeighbor
	}
}
