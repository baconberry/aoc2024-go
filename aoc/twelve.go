package aoc

import "aoc2024/util"

func Twelve(lines []string, part int) int {
	grid := util.NewGrid(util.ParseRuneGrid(lines))

	visited := make(map[util.Point]bool)
	total := 0
	grid.Iterate(func(x int, y int, c rune) {
		region := util.InitGrid(grid.BoundPoint(), '.')
		area, perimeter := getAreaAndRegion(&grid, c, util.Point{X: x, Y: y}, &visited, &region)
		if area == 0 {
			return
		}
		if part == 1 {
			total += area * perimeter
		} else {
			//region.Print(func(t rune) string {
			//	return string(t)
			//})
			sides := calculateSides(&region, c)
			total += sides * area

		}
	})
	return total
}

func calculateSides(grid *util.Grid[rune], toFind rune) int {
	sides := 0
	for i := 0; i < grid.Height(); i++ {
		initPoint := util.Point{Y: i}
		sides += getGroups(grid, initPoint, toFind, util.E, util.N, 0)
		sides += getGroups(grid, initPoint, toFind, util.E, util.S, 0)
	}
	for i := 0; i < grid.Width(); i++ {
		initPoint := util.Point{X: i}
		sides += getGroups(grid, initPoint, toFind, util.S, util.E, 0)
		sides += getGroups(grid, initPoint, toFind, util.S, util.W, 0)
	}
	return sides
}

func getGroups(grid *util.Grid[rune], point util.Point, find rune, forward util.Direction, compare util.Direction, groupCounter int) int {
	val := grid.GetOOBValue(point)
	if val == nil {
		//if groupCounter > 0 {
		//	return 1
		//}
		return 0
	}
	sum := 0
	if *val == find {
		groupSize := 0
		compareVal := grid.GetOOBValue(point.PlusDirection(compare))
		if compareVal == nil || *compareVal != find {
			if groupCounter == 0 {
				sum += 1
			}
			groupSize = groupCounter + 1
		}

		sum += getGroups(grid, point.PlusDirection(forward), find, forward, compare, groupSize)
	} else {
		sum += getGroups(grid, point.PlusDirection(forward), find, forward, compare, 0)
	}
	return sum
}

func getAreaAndRegion(grid *util.Grid[rune], c rune, point util.Point, visited *map[util.Point]bool, region *util.Grid[rune]) (int, int) {
	newRune := grid.GetValue(point)
	if newRune != c {
		return 0, 1
	}
	if (*visited)[point] {
		return 0, 0
	}

	(*visited)[point] = true

	region.SetValue(point, c)
	area := 1 // itself
	perimeter := 0
	bounds := grid.BoundPoint()
	directionalPerimeter := make(map[util.Direction]bool)
	for _, direction := range util.CARDINAL_DIRECTIONS {
		newPoint := point.PlusDirection(direction)
		if newPoint.WithinBounds(bounds) {
			na, np := getAreaAndRegion(grid, newRune, newPoint, visited, region)
			if np == 1 {
				directionalPerimeter[direction] = true
			}
			area += na
			perimeter += np
		} else {
			perimeter += 1
			directionalPerimeter[direction] = true
		}
	}
	return area, perimeter
}
