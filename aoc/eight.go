package aoc

import "aoc2024/util"

type AntennaMap map[rune][]util.Point

func Eight(lines []string) int {
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
		antinodePoints := calculateAntinodePoints(points)
		for _, aP := range antinodePoints {
			if aP.X < 0 || aP.X >= width || aP.Y < 0 || aP.Y >= height {
				continue
			}
			antinodeMap[aP] = true
		}
	}

	return len(antinodeMap)
}

func calculateAntinodePoints(points []util.Point) []util.Point {
	result := make([]util.Point, 0)
	for i, point := range points {
		for _, o := range points[i+1:] {
			result = append(result, calculateAntitondePoint(point, o))
			result = append(result, calculateAntitondePoint(o, point))
		}
	}
	return result
}

func calculateAntitondePoint(a util.Point, b util.Point) util.Point {
	dX, dY := b.X-a.X, b.Y-a.Y
	nX := a.X + (2 * dX)
	nY := a.Y + (2 * dY)
	return util.Point{X: nX, Y: nY}
}

func isAntenna(c rune) bool {
	return c != '.'
}
