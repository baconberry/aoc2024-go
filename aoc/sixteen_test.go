package aoc

import (
	"strings"
	"testing"
)

func TestSixteen(t *testing.T) {
	text := "###############\n#.......#....E#\n#.#.###.#.###.#\n#.....#.#...#.#\n#.###.#####.#.#\n#.#.#.......#.#\n#.#.#####.###.#\n#...........#.#\n###.#.#####.#.#\n#...#.....#.#.#\n#.#.#.###.#.#.#\n#.....#...#.#.#\n#.###.#.#.#.#.#\n#S..#.....#...#\n###############\n"
	result := Sixteen(strings.Split(text, "\n"), 1)

	if result != 7036 {
		t.Fail()
	}
	result = Sixteen(strings.Split(text, "\n"), 2)

	if result != 45 {
		t.Fail()
	}
	textSecond := "#################\n#...#...#...#..E#\n#.#.#.#.#.#.#.#.#\n#.#.#.#...#...#.#\n#.#.#.#.###.#.#.#\n#...#.#.#.....#.#\n#.#.#.#.#.#####.#\n#.#...#.#.#.....#\n#.#.#####.#.###.#\n#.#.#.......#...#\n#.#.###.#####.###\n#.#.#...#.....#.#\n#.#.#.#####.###.#\n#.#.#.........#.#\n#.#.#.#########.#\n#S#.............#\n#################\n"
	result = Sixteen(strings.Split(textSecond, "\n"), 1)

	if result != 11048 {
		t.Fail()
	}
	result = Sixteen(strings.Split(textSecond, "\n"), 2)

	if result != 64 {
		t.Fail()
	}
}
