package aoc

import "aoc2024/util"

func Fourteen(lines []string) int {
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
	return robotLastPosition(x, y, vx, vy, limit-1, width, height)
}
