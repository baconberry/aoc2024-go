package aoc

import (
	"strings"
	"testing"
)

func TestTen(t *testing.T) {
	text := "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732\n"
	result := Ten(strings.Split(text, "\n"), 1)

	if result != 36 {
		t.Fail()
	}
	result = Ten(strings.Split(text, "\n"), 2)

	if result != 81 {
		t.Fail()
	}

}
