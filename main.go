package main

import (
	"aoc2024/aoc"
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := readUntilEOF()
	args := os.Args
	day, err := strconv.Atoi(args[1])
	if err != nil {
		log.Panic("Error parsing arguments", err)
	}

	part, err := strconv.Atoi(args[2])
	if err != nil {
		log.Panic("Error parsing arguments", err)
	}
	result := 0
	switch {
	case day == 1 && part == 1:
		result = aoc.One(lines)
	case day == 1 && part == 2:
		result = aoc.OneSecond(lines)
	case day == 2 && part == 1:
		result = aoc.Two(lines)
	case day == 2 && part == 2:
		result = aoc.TwoSecond(lines)
	case day == 3 && part == 1:
		result = aoc.Three(lines)
	case day == 3 && part == 2:
		result = aoc.ThreeSecond(lines)
	case day == 4:
		result = aoc.Four(lines, part)
	case day == 5:
		result = aoc.Five(lines, part)
	case day == 6:
		result = aoc.Six(lines, part)
	case day == 7:
		result = aoc.Seven(lines, part)
	case day == 8:
		result = aoc.Eight(lines, part)
	case day == 9 && part == 1:
		result = aoc.Nine(lines)
	case day == 9 && part == 2:
		result = aoc.NineSecond(lines)
	case day == 10:
		result = aoc.Ten(lines, part)
	case day == 11:
		result = aoc.Eleven(lines, part)
	case day == 12:
		result = aoc.Twelve(lines, part)
	case day == 13:
		result = aoc.Thirteen(lines, part)
	case day == 14:
		result = aoc.Fourteen(lines, part)
	case day == 15:
		result = aoc.Fifteen(lines, part)
	case day == 16:
		result = aoc.Sixteen(lines)
	}

	log.Println("Result", result)
}

func readUntilEOF() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadFromFile() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("could not read file" + err.Error())
	}
	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
