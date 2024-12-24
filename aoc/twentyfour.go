package aoc

import (
	"log"
	"strconv"
	"strings"
)

type Gate struct {
	a, b, operation, output string
}

func Twentyfour(lines []string) int {
	wireState := make(map[string]bool)
	operationsLineOffset := 0
	operationsLineOffset = parseWireState(lines, operationsLineOffset, wireState)
	untriggeredGates := make(map[Gate]bool)
	parseGates(lines, operationsLineOffset, untriggeredGates)
	triggeredGates := make(map[Gate]bool)
	triggerGates(untriggeredGates, wireState, triggeredGates)
	binaryString := createBinaryString(wireState)
	num, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		log.Fatal("Could not parse binary number", binaryString)
	}
	return int(num)
}

func createBinaryString(wireState map[string]bool) string {
	zstate := make([]bool, 128)
	for wire, b := range wireState {
		if wire[0] != 'z' {
			continue
		}
		pos, err := strconv.Atoi(wire[1:])
		if err != nil {
			log.Fatal("Colud not parse z position", err)
		}
		zstate[pos] = b
	}
	binaryString := ""
	for _, b := range zstate {
		if b {
			binaryString = "1" + binaryString
		} else {
			binaryString = "0" + binaryString
		}
	}
	return binaryString
}

func triggerGates(untriggeredGates map[Gate]bool, wireState map[string]bool, triggeredGates map[Gate]bool) {
	for len(untriggeredGates) > 0 {
		for gate, _ := range untriggeredGates {
			_, aok := wireState[gate.a]
			_, bok := wireState[gate.b]
			if aok && bok {
				wireState[gate.output] = performOperation(&wireState, gate.operation, gate.a, gate.b)
				delete(untriggeredGates, gate)
				triggeredGates[gate] = true
			}
		}
	}
}

func parseGates(lines []string, operationsLineOffset int, untriggeredGates map[Gate]bool) {
	for i := operationsLineOffset; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		parts := strings.Split(lines[i], " ")
		gate := Gate{parts[0], parts[2], parts[1], parts[4]}
		untriggeredGates[gate] = false
	}
}

func parseWireState(lines []string, operationsLineOffset int, wireState map[string]bool) int {
	for i, line := range lines {
		if len(line) == 0 {
			operationsLineOffset = i + 1
			break
		}
		parts := strings.Split(line, ":")
		val := false
		if strings.Contains(parts[1], "1") {
			val = true
		}
		wireState[parts[0]] = val
	}
	return operationsLineOffset
}

func performOperation(wireState *map[string]bool, op string, a string, b string) bool {
	aval := (*wireState)[a]
	bval := (*wireState)[b]
	switch op {
	case "XOR":
		return aval != bval
	case "AND":
		return aval && bval
	case "OR":
		return aval || bval
	}
	panic("Invalid state operation not implemented " + op)
}
