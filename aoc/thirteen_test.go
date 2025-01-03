package aoc

import (
	"strings"
	"testing"
)

func TestThirteen(t *testing.T) {
	text := "Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400\n\nButton A: X+26, Y+66\nButton B: X+67, Y+21\nPrize: X=12748, Y=12176\n\nButton A: X+17, Y+86\nButton B: X+84, Y+37\nPrize: X=7870, Y=6450\n\nButton A: X+69, Y+23\nButton B: X+27, Y+71\nPrize: X=18641, Y=10279\n"
	result := Thirteen(strings.Split(text, "\n"), 1)

	if result != 480 {
		t.Fail()
	}
}
