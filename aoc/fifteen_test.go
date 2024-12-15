package aoc

import (
	"strings"
	"testing"
)

func TestFifteen(t *testing.T) {
	text := "########\n#..O.O.#\n##@.O..#\n#...O..#\n#.#.O..#\n#...O..#\n#......#\n########\n\n<^^>>>vv<v>>v<<\n"
	result := Fifteen(strings.Split(text, "\n"))

	if result != 2028 {
		t.Fail()
	}
}
