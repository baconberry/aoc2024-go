package aoc

import (
	"aoc2024/util"
	"log"
)

func Fourteen(lines []string, part int) int {
	if part == 2 {
		return XmasTree(lines, 101, 103)
	}
	return QuadrantRobots(lines, 101, 103, 100)
}

func QuadrantRobots(lines []string, width int, height int, limit int) int {
	grid := util.ParseIntGrid(lines)
	middleX := (width + 0) / 2
	middleY := (height + 0) / 2
	quadrantCounts := make([]int, 4)
	for _, row := range grid {
		px, py := row[0], row[1]
		vx, vy := row[2], row[3]
		px, py = fixRobotPosition(px, py, width, height)
		lx, ly := robotLastPosition(px, py, vx, vy, limit, width, height)
		if lx == middleX || ly == middleY {
			continue
		}
		//binary notation quadrant
		one, two := lx > middleX, ly > middleY
		quadrant := 0
		if two {
			quadrant += 2
		}
		if one {
			quadrant += 1
		}
		quadrantCounts[quadrant] += 1
	}
	total := quadrantCounts[1] * quadrantCounts[2] * quadrantCounts[3] * quadrantCounts[0]
	return total
}

type Robot struct {
	initPoint util.Point
	velocity  util.Point
}
type RobotTime struct {
	robot Robot
	time  int
}

func XmasTree(lines []string, width int, height int) int {
	grid := util.ParseIntGrid(lines)
	robots := make(map[Robot][]util.Point)
	for _, row := range grid {
		px, py := row[0], row[1]
		vx, vy := row[2], row[3]
		px, py = fixRobotPosition(px, py, width, height)
		robot := Robot{util.Point{X: px, Y: py}, util.Point{X: vx, Y: vy}}
		robots[robot] = make([]util.Point, 1)
		robots[robot][0] = robot.initPoint
	}
	time := 1
	for {
		if time%10 == 0 {
			log.Println("Rendering time ", time)
		}
		treeSet := make(map[util.Point]bool)
		for r, _ := range robots {
			np := robots[r][time-1]
			nx, ny := np.X, np.Y
			nx += r.velocity.X
			ny += r.velocity.Y
			lx, ly := fixRobotPosition(nx, ny, width, height)
			treeSet[util.Point{X: lx, Y: ly}] = true
			robots[r] = append(robots[r], util.Point{X: lx, Y: ly})
		}
		if isXmasTree(treeSet) {
			break
		}
		time += 1
	}
	return time
}

func isXmasTree(grid map[util.Point]bool) bool {
	for point, _ := range grid {
		sx, ex := point.X, point.X+1
		for i := 0; i < 3; i++ {
			sx = sx - i
			ex = ex + i
			if !rangeMatches(&grid, sx, point.Y+i, ex) {
				return false
			}
		}
		return true
	}
	return false
}

func rangeMatches(treeSet *map[util.Point]bool, sx int, sy int, ex int) bool {
	for i := sx; i < ex; i++ {
		point := util.Point{X: i, Y: sy}
		if !(*treeSet)[point] {
			return false
		}
	}
	return true
}

func fixRobotPosition(nx int, ny int, width int, height int) (int, int) {
	if nx < 0 {
		nx = width + nx
	}
	if ny < 0 {
		ny = height + ny
	}
	nx = nx % width
	ny = ny % height
	return nx, ny
}
func robotLastPosition(x int, y int, vx int, vy int, limit int, width int, height int) (int, int) {
	if limit == 0 {
		return x, y
	}
	x += vx
	y += vy
	x, y = fixRobotPosition(x, y, width, height)
	lx, ly := robotLastPosition(x, y, vx, vy, limit-1, width, height)
	return lx, ly
}
