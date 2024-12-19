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
	optMap := make(map[rune][]string)
	memo := make(map[string]bool)
	for _, option := range options {
		r := rune(option[0])
		_, ok := optMap[r]
		if !ok {
			optMap[r] = make([]string, 0)
		}
		optMap[r] = append(optMap[r], option)
	}
	sum := 0
	for i, s := range lines[2:] {
		if len(s) == 0 {
			continue
		}
		t1 := time.Now()
		if combinationPossible(s, &optMap, &memo) {
			sum += 1
		}
		log.Println(strconv.Itoa(i), " solved after ", time.Now().Nanosecond()-t1.Nanosecond(), "ns")
	}
	return sum
}

func combinationPossible(s string, optMap *map[rune][]string, memo *map[string]bool) bool {
	if len(s) == 0 {
		return true
	}
	options, ok := (*optMap)[rune(s[0])]
	if !ok {
		return false
	}
	val, inMemo := (*memo)[s]
	if inMemo {
		return val
	}

	for _, opt := range options {
		newS, found := strings.CutPrefix(s, opt)
		if found && combinationPossible(newS, optMap, memo) {
			(*memo)[s] = true
			return true
		}
	}
	(*memo)[s] = false
	return false
}
