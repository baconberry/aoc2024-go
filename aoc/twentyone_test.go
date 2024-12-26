package aoc

import (
	"strings"
	"testing"
)

func TestTwentyone(t *testing.T) {
	//text := "029A"
	text := "029A\n980A\n179A\n456A\n379A\n"
	result := Twentyone(strings.Split(text, "\n"), 1)

	if result != 126384 {
		t.Fail()
	}
	result = Twentyone(strings.Split(text, "\n"), 2)

	if result != 154115708116294 {
		t.Fail()
	}
}
