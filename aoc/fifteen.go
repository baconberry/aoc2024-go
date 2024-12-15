package aoc

import (
	"aoc2024/util"
)

func Fifteen(lines []string) int {
	grid := util.NewGrid(util.ParseRuneGrid(lines))
	initRobotPos := grid.FindFirst('@')
	if initRobotPos == nil {
		panic("Robot not found")
	}

	moveScript := util.NewGrid(util.ParseRuneGrid(lines[grid.Height()+1:]))
	robotPos := *initRobotPos
	moveScript.Iterate(func(_, _ int, c rune) {
		direction := moveScriptToDirection(c)
		robotPos = moveRobot(&grid, robotPos, direction)
		//printRuneGrid(&grid)
		//log.Println("--------------------")
	})
	sum := 0
	grid.Iterate(func(x, y int, c rune) {
		if c == 'O' {
			sum += (y * 100) + x
		}
	})

	return sum
}

func printRuneGrid(grid *util.Grid[rune]) {
	grid.Print(func(c rune) string {
		return string(c)
	})
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
	//movePos := moveRobot(g, nextPos, &point, direction)
	//if *movePos == nextPos {
	//	g.SetValue(point, '.')
	//	g.SetValue(nextPos, *val)
	//	return &nextPos
	//	//nextVal := g.GetOOBValue(nextPos)
	//	//if nextVal != nil && *nextVal == '.' {
	//	//}
	//}
	//return &point
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
