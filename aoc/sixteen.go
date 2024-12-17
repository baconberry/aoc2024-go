package aoc

import (
	"aoc2024/util"
	"math"
)

func Sixteen(lines []string, part int) int {
	grid := util.NewGrid(util.ParseRuneGrid(lines))
	reindeerPos := grid.FindFirst('S')
	endTilePos := grid.FindFirst('E')
	pathScores := make(map[ReindeerLoc]int)
	loc := ReindeerLoc{*reindeerPos, util.E}
	path := make(ReindeerPathType, 0)
	reachedPaths := make([]ReindeerPath, 0)
	endTilePathScore, _ := walkMaze(&grid, loc, *endTilePos, 0, &pathScores, path, &reachedPaths)
	if part == 2 {
		bestPaths := calculateBestPaths(reachedPaths, endTilePathScore)
		bestTiles := make(map[util.Point]bool)
		markBestTiles(bestPaths, &bestTiles)
		markGridTiles(&grid, &bestTiles)
		//grid.Print(func(r rune) string {
		//	return string(r)
		//})
		return len(bestTiles) - 1
	}
	return endTilePathScore
}

func markGridTiles(g *util.Grid[rune], m *map[util.Point]bool) {
	for loc, _ := range *m {
		g.SetValue(loc, 'O')
	}
}

func markBestTiles(paths []ReindeerPath, m *map[util.Point]bool) {
	for _, path := range paths {
		for _, loc := range path.path {
			if !(*m)[loc] {
				(*m)[loc] = true
			}
		}
	}
}

func calculateBestPaths(paths []ReindeerPath, score int) []ReindeerPath {
	result := make([]ReindeerPath, 0)
	for _, path := range paths {
		//pathScore := calculatePathScore(path)
		if path.score == score {
			result = append(result, path)
		}
	}
	return result
}

type ReindeerLoc struct {
	location util.Point
	facing   util.Direction
}

type ReindeerPathType []ReindeerLoc
type ReindeerPath struct {
	path  []util.Point
	score int
}

func walkMaze(g *util.Grid[rune], loc ReindeerLoc, endLoc util.Point, score int, pathScore *map[ReindeerLoc]int, path ReindeerPathType, reachedPaths *[]ReindeerPath) (int, bool) {
	tile := g.GetOOBValue(loc.location)
	if tile == nil || *tile == '#' {
		return math.MaxInt, false
	}
	if loc.location == endLoc {
		locCopy := loc
		pathCopy := path
		pathCopy = append(pathCopy, locCopy)
		//pathCopy[locCopy] = true
		*reachedPaths = append(*reachedPaths, ReindeerPath{pathToPoints(pathCopy), score})
		return score, true
	}
	//locCopy := loc
	pathCopy := path
	pathCopy = append(pathCopy, loc)
	// if we are in the same location but facing the other way it's also more expensive
	inverseLoc := loc
	inverseLoc.facing = inverseLoc.facing.Inverse()
	wasHereBefore := isLocInPath(&path, inverseLoc)
	if wasHereBefore {
		return math.MaxInt, false
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
		localScore, endReached := walkMaze(g, newLoc, endLoc, newScore, pathScore, pathCopy, reachedPaths)
		if endReached {
			minPath = min(localScore, minPath)
		}
	}
	return minPath, minPath < math.MaxInt
}

func pathToPoints(pathCopy ReindeerPathType) []util.Point {
	arr := make([]util.Point, len(pathCopy))
	for _, loc := range pathCopy {
		arr = append(arr, loc.location)
	}
	return arr
}

func isLocInPath(arr *ReindeerPathType, loc ReindeerLoc) bool {
	for i := len(*arr) - 1; i >= 0; i-- {
		if (*arr)[i] == loc {
			return true
		}
	}
	return false
}
