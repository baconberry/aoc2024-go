package aoc

import (
	"strings"
	"testing"
)

func TestEight(t *testing.T) {
	text := "............\n........0...\n.....0......\n.......0....\n....0.......\n......A.....\n............\n............\n........A...\n.........A..\n............\n............\n"
	result := Eight(strings.Split(text, "\n"))

	if result != 14 {
		t.Fail()
	}
}
