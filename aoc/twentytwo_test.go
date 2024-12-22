package aoc

import (
	"strings"
	"testing"
)

func TestTwentytwo(t *testing.T) {

	text := "1\n10\n100\n2024\n"
	result := Twentytwo(strings.Split(text, "\n"), 1)

	if result != 37327623 {
		t.Fail()
	}
	text = "1\n2\n3\n2024\n"
	result = Twentytwo(strings.Split(text, "\n"), 2)

	if result != 23 {
		t.Fail()
	}
}

//func TestTwentytwoBase(t *testing.T) {
//	secretNumber := 123
//	for i := 0; i < 10; i++ {
//		log.Println(secretNumber)
//		secretNumber = iterateSecretNumber(secretNumber)
//	}
//}
