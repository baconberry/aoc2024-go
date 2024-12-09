package aoc

import (
	"strings"
	"testing"
)

func TestNine(t *testing.T) {
	text := "2333133121414131402"
	result := Nine(strings.Split(text, "\n"))

	if result != 1928 {
		t.Fail()
	}

}
