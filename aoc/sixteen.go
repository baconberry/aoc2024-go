package aoc

import (
	"aoc2024/util"
	"math"
)

func Sixteen(lines []string) int {
	grid := util.NewGrid(util.ParseRuneGrid(lines))
	reindeerPos := grid.FindFirst('S')
	endTilePos := grid.FindFirst('E')
	pathScores := make(map[ReindeerLoc]int)
	loc := ReindeerLoc{*reindeerPos, util.E}
	endTilePathScore, _ := walkMaze(&grid, loc, *endTilePos, 0, &pathScores)
	return endTilePathScore
}

type ReindeerLoc struct {
	location util.Point
	facing   util.Direction
}

func walkMaze(g *util.Grid[rune], loc ReindeerLoc, endLoc util.Point, score int, pathScore *map[ReindeerLoc]int) (int, bool) {
	tile := g.GetOOBValue(loc.location)
	if tile == nil || *tile == '#' {
		return math.MaxInt, false
	}
	if loc.location == endLoc {
		return score, true
	}
	val, ok := (*pathScore)[loc]
	if !ok || score < val {
		(*pathScore)[loc] = score
	} else if ok && score > val {
		// if we have come here cheaper there's no point in building this path further
		return math.MaxInt, false
	}
	minPath := math.MaxInt
	for _, direction := range loc.facing.ForwardOr90Turn() {
		newScore := score
		newLoc := loc
		newLoc.facing = direction
		if direction == loc.facing {
			newScore += 1
			newLoc.location = newLoc.location.PlusDirection(direction)
		} else {
			newScore += 1000
		}
		localScore, endReached := walkMaze(g, newLoc, endLoc, newScore, pathScore)
		if endReached {
			minPath = min(localScore, minPath)
		}
	}
	return minPath, minPath < math.MaxInt
}
