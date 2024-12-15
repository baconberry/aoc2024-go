package aoc

import (
	"aoc2024/util"
)

func Fifteen(lines []string, part int) int {
	preGrid := util.ParseRuneGrid(lines)
	if part == 2 {
		newLines := expandGrid(preGrid)
		preGrid = util.ParseRuneGrid(newLines)
	}
	grid := util.NewGrid(preGrid)
	initRobotPos := grid.FindFirst('@')
	if initRobotPos == nil {
		panic("Robot not found")
	}

	moveScript := util.NewGrid(util.ParseRuneGrid(lines[grid.Height()+1:]))
	robotPos := *initRobotPos
	moveScript.Iterate(func(x, y int, c rune) {
		direction := moveScriptToDirection(c)
		//log.Println("Executing move:", string(c), x, y)
		if part == 1 || direction == util.E || direction == util.W {
			robotPos = moveRobot(&grid, robotPos, direction)
		} else {
			if moveRobotRange(&grid, robotPos, 1, direction) {
				grid.SetValue(robotPos, '.')
				newPos := robotPos.PlusDirection(direction)
				grid.SetValue(newPos, '@')
				robotPos = newPos
			}
		}
		//printRuneGrid(&grid)
		//log.Println("--------------------")
	})
	sum := 0
	grid.Iterate(func(x, y int, c rune) {
		if part == 1 && c == 'O' {
			sum += (y * 100) + x
		}
		if part == 2 && c == '[' {
			sum += (y * 100) + x
		}
	})

	return sum
}

func expandGrid(grid [][]rune) []string {
	lines := make([]string, 0)
	for _, row := range grid {
		line := ""
		for _, r := range row {
			switch r {
			case '#':
				line += "##"
			case 'O':
				line += "[]"
			case '.':
				line += ".."
			case '@':
				line += "@."
			}
		}
		lines = append(lines, line)
	}
	return lines
}

func printRuneGrid(grid *util.Grid[rune]) {
	grid.Print(func(c rune) string {
		return string(c)
	})
}

/*
*
call this only when doing N or S direction
*/
func moveRobotRange(g *util.Grid[rune], start util.Point, dist int, direction util.Direction) bool {
	if direction != util.N && direction != util.S {
		panic("moveRobotRange called with non N/S direction")
	}
	if dist == 0 {
		return true
	}
	if g.GetValue(start) == '.' {
		start.X += 1
		return moveRobotRange(g, start, dist-1, direction)
	}
	endPos := util.Point{X: start.X + dist - 1, Y: start.Y}
	endVal := g.GetValue(endPos)
	if endVal == '.' {
		return moveRobotRange(g, start, dist-1, direction)
	}

	if gridAnyRangeMatches(g, start.X, start.Y, start.X+dist, '#') {
		return false
	}
	if gridAllRangeMatches(g, start.X, start.Y, start.X+dist, '.') {
		return true
	}
	val := g.GetValue(start)
	newDist := dist
	newStart := util.Point{X: start.X, Y: start.Y}
	if val == ']' {
		newStart.X -= 1
		newDist += 1
		start.X -= 1
		dist += 1
	}
	if endVal == '[' {
		newDist += 1
		dist += 1
	}
	newStart = newStart.PlusDirection(direction)
	canMove := moveRobotRange(g, newStart, newDist, direction)
	if canMove {
		for i := start.X; i < start.X+dist; i++ {
			nPos := util.Point{X: i, Y: start.Y}
			valB := g.GetValue(nPos)
			g.SetValue(nPos, '.')
			nPos = nPos.PlusDirection(direction)
			g.SetValue(nPos, valB)
		}
		return true
	}
	return false
}

func gridAnyRangeMatches(g *util.Grid[rune], sx, y, ex int, c rune) bool {
	for i := sx; i < ex; i++ {
		val := g.GetOOBValue(util.Point{X: i, Y: y})
		if val != nil && *val == c {
			return true
		}
	}
	return false
}
func gridAllRangeMatches(g *util.Grid[rune], sx, y, ex int, c rune) bool {
	for i := sx; i < ex; i++ {
		val := g.GetOOBValue(util.Point{X: i, Y: y})
		if val == nil || *val != c {
			return false
		}
	}
	return true
}
func moveRobot(g *util.Grid[rune], point util.Point, direction util.Direction) util.Point {
	val := g.GetOOBValue(point)
	nextPos := point.PlusDirection(direction)
	nextVal := g.GetOOBValue(nextPos)
	if nextVal == nil || *nextVal == '#' {
		return point
	}
	if *nextVal == '.' {
		// do the move
		g.SetValue(nextPos, *val)
		g.SetValue(point, '.')
		return nextPos
	}
	movedPos := moveRobot(g, nextPos, direction)
	if movedPos != nextPos {
		return moveRobot(g, point, direction)
	}
	return point
}

func moveScriptToDirection(c rune) util.Direction {
	switch c {
	case '<':
		return util.W
	case '>':
		return util.E
	case '^':
		return util.N
	case 'v':
		return util.S
	}
	return util.NONE
}
