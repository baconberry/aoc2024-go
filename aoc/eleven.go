package aoc

import (
	"aoc2024/util"
	"strconv"
)

func Eleven(lines []string) int {
	arr := util.ParseIntArray(lines[0])
	stringArr := make([]string, len(arr))
	for i, n := range arr {
		stringArr[i] = strconv.Itoa(n)
	}
	result := stringArr
	for i := 0; i < 25; i++ {
		result = blink(result)
	}
	return len(result)
}

func blink(arr []string) []string {
	result := make([]string, 0)
	for _, s := range arr {
		result = append(result, blinkNumber(s)...)
	}
	return result
}

func blinkNumber(s string) []string {
	result := make([]string, 0)
	if s == "0" {
		return append(result, "1")
	}
	if len(s)%2 == 0 {
		middle := len(s) / 2
		l := removeLeadingZeroes(s[:middle])
		r := removeLeadingZeroes(s[middle:])
		result = append(result, l)
		result = append(result, r)
		return result
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		panic("Could not convert string to number")
	}
	return append(result, strconv.Itoa(n*2024))
}

func removeLeadingZeroes(s string) string {
	if len(s) == 1 {
		return s
	}
	if s[0] == '0' {
		return removeLeadingZeroes(s[1:])
	}
	return s
}
