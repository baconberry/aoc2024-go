package aoc

import (
	"strings"
	"testing"
)

func TestTwelve(t *testing.T) {
	text := "AAAA\nBBCD\nBBCC\nEEEC\n"
	result := Twelve(strings.Split(text, "\n"), 1)

	if result != 140 {
		t.Fail()
	}
	result = Twelve(strings.Split(text, "\n"), 2)
	if result != 80 {
		t.Fail()
	}
	text = "RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE\n"
	result = Twelve(strings.Split(text, "\n"), 1)

	if result != 1930 {
		t.Fail()
	}
	result = Twelve(strings.Split(text, "\n"), 2)

	if result != 1206 {
		t.Fail()
	}
}
