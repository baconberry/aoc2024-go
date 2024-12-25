package aoc

import (
	"strings"
	"testing"
)

func TestTwentyfive(t *testing.T) {
	text := "#####\n.####\n.####\n.####\n.#.#.\n.#...\n.....\n\n#####\n##.##\n.#.##\n...##\n...#.\n...#.\n.....\n\n.....\n#....\n#....\n#...#\n#.#.#\n#.###\n#####\n\n.....\n.....\n#.#..\n###..\n###.#\n###.#\n#####\n\n.....\n.....\n.....\n#....\n#.#..\n#.#.#\n#####\n"
	result := Twentyfive(strings.Split(text, "\n"))

	if result != 3 {
		t.Fail()
	}
}
