package aoc

import (
	"strings"
	"testing"
)

func TestFourteen(t *testing.T) {
	text := "p=0,4 v=3,-3\np=6,3 v=-1,-3\np=10,3 v=-1,2\np=2,0 v=2,-1\np=0,0 v=1,3\np=3,0 v=-2,-2\np=7,6 v=-1,-3\np=3,0 v=-1,-2\np=9,3 v=2,3\np=7,3 v=-1,2\np=2,4 v=2,-3\np=9,5 v=-3,-3\n"
	result := QuadrantRobots(strings.Split(text, "\n"), 11, 7, 100)

	if result != 12 {
		t.Fail()
	}
}
