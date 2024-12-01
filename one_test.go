package main

import (
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	lines := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
	result := one(strings.Split(lines, "\n"))

	if result != 11 {
		t.Fatal("11 !=", result)
	}
}

func TestOneSecond(t *testing.T) {
	lines := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
	result := oneSecond(strings.Split(lines, "\n"))

	if result != 31 {
		t.Fatal("31 !=", result)
	}

}
