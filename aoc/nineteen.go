package aoc

import (
	"aoc2024/util"
	"log"
	"strconv"
	"strings"
	"time"
)

func Nineteen(lines []string) int {
	options := util.ParseStringArray(lines[0])
	sum := 0
	for i, s := range lines[2:] {
		if len(s) == 0 {
			continue
		}
		t1 := time.Now()
		if combinationPossible(s, &options) {
			sum += 1
		}
		log.Println(strconv.Itoa(i), " solved after ", time.Now().Nanosecond()-t1.Nanosecond(), "ns")
	}
	return sum
}

func combinationPossible(s string, options *[]string) bool {
	if len(s) == 0 {
		return true
	}

	for _, opt := range *options {
		newS, found := strings.CutPrefix(s, opt)
		if found && combinationPossible(newS, options) {
			return true
		}
	}
	return false
}
