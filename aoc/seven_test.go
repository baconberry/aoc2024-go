package aoc

import (
	"strings"
	"testing"
)

func TestSeven(t *testing.T) {
	text := "190: 10 19\n3267: 81 40 27\n83: 17 5\n156: 15 6\n7290: 6 8 6 15\n161011: 16 10 13\n192: 17 8 14\n21037: 9 7 18 13\n292: 11 6 16 20\n"
	result := Seven(strings.Split(text, "\n"), 1)

	if result != 3749 {
		t.Fail()
	}

	result = Seven(strings.Split(text, "\n"), 2)

	if result != 11387 {
		t.Fail()
	}

}
