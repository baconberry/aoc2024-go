package aoc

import (
	"strings"
	"testing"
)

func TestTwentyone(t *testing.T) {
	//text := "029A"
	text := "029A\n980A\n179A\n456A\n379A\n"
	result := Twentyone(strings.Split(text, "\n"))

	if result != 126384 {
		t.Fail()
	}
}
