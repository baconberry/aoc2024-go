package aoc

import (
	"strings"
	"testing"
)

func TestEleven(t *testing.T) {
	text := "125 17"
	result := Eleven(strings.Split(text, "\n"))

	if result != 55312 {
		t.Fail()
	}
}
