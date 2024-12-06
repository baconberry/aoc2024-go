package aoc

import (
	"strings"
	"testing"
)

func Test_six(t *testing.T) {

	text := "....#.....\n.........#\n..........\n..#.......\n.......#..\n..........\n.#..^.....\n........#.\n#.........\n......#...\n"
	result := Six(strings.Split(text, "\n"), 1)

	if result != 41 {
		t.Fail()
	}

	result = Six(strings.Split(text, "\n"), 2)

	if result != 6 {
		t.Fail()
	}
}
