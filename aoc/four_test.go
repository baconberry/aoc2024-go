package aoc

import (
	"strings"
	"testing"
)

func Test_four(t *testing.T) {
	text := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX\n"
	result := Four(strings.Split(text, "\n"), 1)

	if result != 18 {
		t.Fail()
	}

	result = Four(strings.Split(text, "\n"), 2)

	if result != 9 {
		t.Fail()
	}
}
