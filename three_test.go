package main

import (
	"strings"
	"testing"
)

func Test_three(t *testing.T) {
	text := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := three(strings.Split(text, "\n"))

	if result != 161 {
		t.Fail()
	}

}

func Test_threeSecond(t *testing.T) {
	text := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result := threeSecond(strings.Split(text, "\n"))

	if result != 48 {
		t.Fail()
	}

}
