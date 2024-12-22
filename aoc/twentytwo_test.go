package aoc

import (
	"log"
	"strings"
	"testing"
)

func TestTwentytwo(t *testing.T) {

	text := "1\n10\n100\n2024\n"
	result := Twentytwo(strings.Split(text, "\n"))

	if result != 37327623 {
		t.Fail()
	}
}
func TestTwentytwoBase(t *testing.T) {
	secretNumber := 123
	for i := 0; i < 10; i++ {
		log.Println(secretNumber)
		secretNumber = iterateSecretNumber(secretNumber, 1)
	}
}
