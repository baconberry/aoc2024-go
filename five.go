package main

import "aoc2024/util"

type IntSet map[int]bool

type Rule map[int]IntSet

func five(lines []string, part int) int {
	rules := make([]string, 0)
	pagesIdx := 0
	for i, line := range lines {
		if len(line) == 0 {
			pagesIdx = i + 1
			break
		}
		rules = append(rules, line)
	}
	pages := lines[pagesIdx:]
	rulesGrid := util.ParseIntGrid(rules)
	ruleMap := rulesFromGrid(&rulesGrid)
	pagesGrid := util.ParseIntGrid(pages)

	sum := 0
	for _, page := range pagesGrid {
		if part == 1 {
			if isSafePage(page, &ruleMap) {
				middle := len(page) / 2
				sum += page[middle]
			}
		}
		if part == 2 {
			if !isSafePage(page, &ruleMap) {
				fixedPage := fixPage(page, &ruleMap)
				middle := len(fixedPage) / 2
				sum += fixedPage[middle]
			}
		}
	}
	return sum
}

func isSafePage(page []int, r *Rule) bool {
	for i := range page {
		// if we have a match n any of the previous
		n := page[i]
		prevNums := page[:i]
		set := (*r)[n]
		for _, p := range prevNums {
			_, inSet := set[p]
			if inSet {
				return false
			}
		}
	}
	return true
}
func fixPage(page []int, r *Rule) []int {
	if isSafePage(page, r) {
		return page
	}
	if len(page) < 2 {
		return page
	}
	for i := range page {
		n := page[i]
		prevNums := page[:i]
		set := (*r)[n]
		for s, p := range prevNums {
			_, inSet := set[p]
			if inSet {
				//switch s and i
				page[i] = p
				page[s] = n
				return fixPage(page, r)
			}
		}
	}
	return page
}

func rulesFromGrid(g *[][]int) Rule {
	rules := make(Rule)

	for _, row := range *g {
		a := row[0]
		b := row[1]
		_, ok := rules[a]
		if !ok {
			rules[a] = make(IntSet)
		}
		rules[a][b] = true
	}

	return rules
}
