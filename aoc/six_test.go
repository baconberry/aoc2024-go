package aoc

import (
	"strings"
	"testing"
)

func Test_six(t *testing.T) {

	text := "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...\n"
	result := Six(strings.Split(text, "\n"))

	if result != 41 {
		t.Fail()
	}
}
