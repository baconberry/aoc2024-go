package aoc

import (
	"strings"
	"testing"
)

func TestNineteen(t *testing.T) {
	text := "r, wr, b, g, bwu, rb, gb, br\n\nbrwrr\nbggr\ngbbr\nrrbgbr\nubwu\nbwurrg\nbrgr\nbbrgwb\n"
	result := Nineteen(strings.Split(text, "\n"))

	if result != 6 {
		t.Fail()
	}
}
