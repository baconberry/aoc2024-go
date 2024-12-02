package main

import (
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
		result = one(lines)
	case day == 1 && part == 2:
		result = oneSecond(lines)
	case day == 2 && part == 1:
		result = two(lines)
	case day == 2 && part == 2:
		result = twoSecond(lines)
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

func readFromFile() []string {
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
