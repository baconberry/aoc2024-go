package aoc

import (
	"aoc2024/util"
	"log"
	"strconv"
)

type KeyLock struct {
	colCount []int
	isLock   bool
}

func (k *KeyLock) String() string {
	s := ""
	if k.isLock {
		s += "L"
	} else {
		s += "K"
	}
	for _, h := range k.colCount {
		s += "_" + strconv.Itoa(h)
	}
	return s
}

func Twentyfive(lines []string) int {
	locks, keys := parseKeyLocks(lines)
	sum := 0
	lockHeight := 7
	validCombos := make(map[string]map[string]bool)
	for _, lock := range locks {
	nextKey:
		for _, key := range keys {
			for k := 0; k < len(lock.colCount); k++ {
				if key.colCount[k]+lock.colCount[k] > lockHeight {
					continue nextKey
				}
			}
			_, ok := validCombos[lock.String()]
			if !ok {
				validCombos[lock.String()] = make(map[string]bool)
			}
			validCombos[lock.String()][key.String()] = true
			sum += 1
		}
	}
	asum := 0
	for _, keys := range validCombos {
		asum += len(keys)
	}
	log.Println("sum", sum, "asum", asum)
	return asum
}

func parseKeyLocks(lines []string) ([]KeyLock, []KeyLock) {
	keys := make([]KeyLock, 0)
	locks := make([]KeyLock, 0)
	linesOffset := 0
	for linesOffset < len(lines) {
		grid := util.ParseRuneGrid(lines[linesOffset:])
		linesOffset += len(grid) + 1
		firstHashCount := 0
		lastHashCount := 0
		colCount := make([]int, len(grid[0]))
		for y, row := range grid {
			hashCount := 0
			for x, c := range row {
				if c == '#' {
					colCount[x] += 1
					hashCount += 1
				}
			}
			if y == 0 {
				firstHashCount = hashCount
			}
			lastHashCount = hashCount
		}
		k := KeyLock{colCount, firstHashCount > lastHashCount}
		if k.isLock {
			locks = append(locks, k)
		} else {
			keys = append(keys, k)
		}
	}
	return locks, keys
}
