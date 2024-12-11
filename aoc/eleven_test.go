package aoc

import (
	"strings"
	"testing"
)

func TestEleven(t *testing.T) {
	text := "125 17"
	result := Eleven(strings.Split(text, "\n"), 1)

	if result != 55312 {
		t.Fail()
	}
	result = Eleven(strings.Split(text, "\n"), 2)

	if result != 65601038650482 {
		t.Fail()
	}
}
