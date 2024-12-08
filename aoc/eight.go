package aoc

import "aoc2024/util"

type AntennaMap map[rune][]util.Point

func Eight(lines []string, part int) int {
	antenaLocations, width, height := util.GetRuneLocations(lines, isAntenna)
	aMap := make(AntennaMap)
	antinodeMap := make(map[util.Point]bool)
	for _, aloc := range antenaLocations {
		_, ok := aMap[aloc.C]
		if !ok {
			aMap[aloc.C] = make([]util.Point, 0)
		}
		aMap[aloc.C] = append(aMap[aloc.C], util.Point{X: aloc.X, Y: aloc.Y})
	}

	for _, points := range aMap {
		var antinodePoints []util.Point
		if part == 1 {
			antinodePoints = calculateAntinodePoints(points)
		} else {
			antinodePoints = antinodePointsWithinBounds(points, width, height)
		}
		for _, aP := range antinodePoints {
			if !isWithinBounds(aP, width, height) {
				continue
			}
			antinodeMap[aP] = true
		}
	}
	if part == 2 {
		for _, p := range antenaLocations {
			pLoc := util.Point{X: p.X, Y: p.Y}
			antinodeMap[pLoc] = true
		}
	}

	return len(antinodeMap)
}

func calculateAntinodePoints(points []util.Point) []util.Point {
	result := make([]util.Point, 0)
	for i, point := range points {
		for _, o := range points[i+1:] {
			result = append(result, calculateAntitondePoint(point, o, 2))
			result = append(result, calculateAntitondePoint(o, point, 2))
		}
	}
	return result
}
func antinodePointsWithinBounds(points []util.Point, width int, height int) []util.Point {
	result := make([]util.Point, 0)
	for i, point := range points {
		for _, o := range points[i+1:] {
			result = getResonantAntinodes(width, height, point, o, result)
			result = getResonantAntinodes(width, height, o, point, result)
		}
	}
	return result
}

func getResonantAntinodes(width int, height int, point util.Point, o util.Point, result []util.Point) []util.Point {
	for m := 2; m <= width*height; m++ {
		nP := calculateAntitondePoint(point, o, m)
		if !isWithinBounds(nP, width, height) {
			break
		}
		result = append(result, nP)
	}
	return result
}

func isWithinBounds(aP util.Point, width int, height int) bool {
	if aP.X < 0 || aP.X >= width || aP.Y < 0 || aP.Y >= height {
		return false
	}
	return true
}

func calculateAntitondePoint(a util.Point, b util.Point, magnitude int) util.Point {
	dX, dY := b.X-a.X, b.Y-a.Y
	nX := a.X + (magnitude * dX)
	nY := a.Y + (magnitude * dY)
	return util.Point{X: nX, Y: nY}
}

func isAntenna(c rune) bool {
	return c != '.'
}
