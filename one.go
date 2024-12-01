package main

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
)

func splitList(lines []string) ([]int, []int, error) {
	listLeft := make([]int, len(lines))
	listRight := make([]int, len(lines))
	re := regexp.MustCompile("(\\d+)\\s+(\\d+)")
	for i, line := range lines {
		parts := re.FindStringSubmatch(line)
		//parts := re.Split(line, -1)
		left, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		right, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, nil, err
		}
		listLeft[i] = left
		listRight[i] = right
	}
	return listLeft, listRight, nil
}

func one(lines []string) int {
	left, right, err := splitList(lines)
	if err != nil {
		log.Fatal("error parsing lists", err)
		return 0
	}
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))
	log.Println("Length of list", len(left))
	distance := 0
	for i, l := range left {
		r := right[i]
		distance += int(math.Abs(float64(l - r)))
	}
	return distance
}

func oneSecond(lines []string) int {
	left, right, err := splitList(lines)
	if err != nil {
		log.Fatal("error parsing lists", err)
		return 0
	}
	freqMap := make(map[int]int)
	sum := 0
	for _, l := range left {
		freq, ok := freqMap[l]
		if ok {
			sum += freq
			continue
		}
		f := frequency(l, right)
		f *= l
		freqMap[l] = f
		sum += f
	}
	return sum
}

func frequency(i int, arr []int) int {
	f := 0
	for _, n := range arr {
		if n == i {
			f += 1
		}
	}
	return f

}
