package aoc

import (
	"strings"
	"testing"
)

func TestSeventeen(t *testing.T) {
	text := "Register A: 729\nRegister B: 0\nRegister C: 0\n\nProgram: 0,1,5,4,3,0\n"
	result := Seventeen(strings.Split(text, "\n"))

	if result != "4,6,3,5,6,3,5,2,1,0," {
		t.Fail()
	}
}
