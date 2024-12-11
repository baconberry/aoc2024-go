package aoc

import (
	"aoc2024/util"
	"strconv"
)

func Eleven(lines []string, part int) int {
	arr := util.ParseIntArray(lines[0])
	stringArr := make([]string, len(arr))
	for i, n := range arr {
		stringArr[i] = strconv.Itoa(n)
	}
	limit := 25
	if part == 2 {
		limit = 75
	}
	memo := make(map[BlinkN]int)
	sum := 0
	for _, s := range stringArr {
		_, n := blinkNTimes(s, limit, &memo)
		sum += n
	}
	return sum
}

type BlinkN struct {
	s     string
	times int
}

func blinkNTimes(s string, times int, memo *map[BlinkN]int) ([]string, int) {
	if times == 0 {
		return nil, 1
	}
	blink := BlinkN{s, times}
	v, ok := (*memo)[blink]
	if ok {
		return nil, v
	}

	firstBlink := blinkNumber(s)
	result := make([]string, 0)
	sum := 0
	for _, part := range firstBlink {
		arr, n := blinkNTimes(part, times-1, memo)
		if n >= 0 {
			sum += n
		} else {
			result = append(result, arr...)
		}
	}
	total := sum + len(result)
	(*memo)[blink] = total
	return result, total
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
