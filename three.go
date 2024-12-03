package main

import (
	"aoc2024/util"
	"regexp"
	"strings"
)

func three(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += mulLine(line)
	}
	return sum
}
func mulLine(line string) int {
	re := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	parts := re.FindAllString(line, -1)
	sum := 0
	for _, part := range parts {
		arr := util.ParseIntArray(part)
		sum += arr[0] * arr[1]
	}
	return sum
}

func threeSecond(lines []string) int {
	line := strings.Join(lines, " ")
	re := regexp.MustCompile("don't\\(\\)")
	reDo := regexp.MustCompile("do\\(\\)")
	sum := 0
	parts := re.Split(line, -1)
	if len(parts) == 1 {
		return mulLine(parts[0])
	}
	sum += mulLine(parts[0])
	for _, part := range parts[1:] {
		enabledParts := reDo.Split(part, -1)
		if len(enabledParts) > 1 {
			for _, enabledPart := range enabledParts[1:] {
				sum += mulLine(enabledPart)
			}
		}
	}
	return sum
}
