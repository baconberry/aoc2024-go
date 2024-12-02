package main

import (
	"strings"
	"testing"
)

func Test_two(t *testing.T) {
	text := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n"
	result := two(strings.Split(text, "\n"))

	if result != 2 {
		t.Fail()
	}

}

func Test_twoSecond(t *testing.T) {
	text := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9\n"
	result := twoSecond(strings.Split(text, "\n"))

	if result != 4 {
		t.Fail()
	}

}
