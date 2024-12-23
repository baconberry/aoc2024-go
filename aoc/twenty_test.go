package aoc

import (
	"strings"
	"testing"
)

func Test_twenty(t *testing.T) {
	text := "###############\n#...#...#.....#\n#.#.#.#.#.###.#\n#S#...#.#.#...#\n#######.#.#.###\n#######.#.#...#\n#######.#.###.#\n###..E#...#...#\n###.#######.###\n#...###...#...#\n#.#####.#.###.#\n#.#...#.#.#...#\n#.#.#.#.#.#.###\n#...#...#...###\n###############\n"
	result := twenty(strings.Split(text, "\n"), 40, 1, 2)

	if result != 2 {
		t.Fail()
	}
	result = twenty(strings.Split(text, "\n"), 70, 2, 20)

	if result != 41 {
		t.Fail()
	}
}
