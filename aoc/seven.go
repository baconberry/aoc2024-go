package aoc

import (
	"aoc2024/util"
	"strconv"
)

func Seven(lines []string, part int) int {
	grid := util.ParseIntGrid(lines)
	sum := 0
	operands := make([]func(a int, b int) int, 2)
	operands[0] = mulFn
	operands[1] = sumFn
	if part == 2 {
		operands = append(operands, concatFn)
	}
	for _, row := range grid {
		target := row[0]
		if hasCombination(target, row[1:], operands...) {
			sum += target
		}
	}
	return sum
}

func hasCombination(target int, row []int, operands ...func(a int, b int) int) bool {
	if target == 0 {
		return true
	}
	if target > 0 && len(row) == 1 {
		return target == row[0]
	}
	if target < 0 || len(row) < 2 {
		return false
	}
	result := false
	for _, operand := range operands {
		newArr := make([]int, 1)
		newArr[0] = operand((row)[0], (row)[1])
		result = result || hasCombination(target, append(newArr, row[2:]...), operands...)
	}

	return result
}

func mulFn(a int, b int) int {
	return a * b
}
func sumFn(a int, b int) int {
	return a + b
}

func concatFn(a int, b int) int {
	r, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		panic("Could not concatenate ")
	}
	return r
}
