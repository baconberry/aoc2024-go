package aoc

import "aoc2024/util"

func Twentytwo(lines []string) int {
	secretNumbers := util.ParseIntGrid(lines)
	sum := 0
	for _, number := range secretNumbers {
		sum += iterateSecretNumber(number[0], 2000)
	}
	return sum
}

func iterateSecretNumber(secretNumber int, iteration int) int {
	if iteration == 0 {
		return secretNumber
	}
	secretNumber = (secretNumber * 64) ^ secretNumber
	secretNumber = secretNumber % 16777216
	secretNumber = (secretNumber / 32) ^ secretNumber
	secretNumber = secretNumber % 16777216
	secretNumber = (secretNumber * 2048) ^ secretNumber
	secretNumber = secretNumber % 16777216
	return iterateSecretNumber(secretNumber, iteration-1)
}
