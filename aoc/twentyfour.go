package aoc

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type Gate struct {
	a, b, operation, output string
}

var printedDot = false

func Twentyfour(lines []string, part int) int {
	wireState := make(map[string]bool)
	operationsLineOffset := 0
	//newInitState := getRandomState()
	operationsLineOffset = parseWireState(lines, operationsLineOffset, wireState)
	//operationsLineOffset = 91
	untriggeredGates := make(map[Gate]bool)
	parseGates(lines, operationsLineOffset, untriggeredGates)
	triggeredGates := make(map[Gate]bool)
	triggerGates(untriggeredGates, wireState, triggeredGates)
	num := parseBinaryString(wireState, 'z')
	if part == 2 {
		if !printedDot {
			printDotFile(wireState, triggeredGates)
			printedDot = true
		}
		xnum := parseBinaryString(wireState, 'x')
		ynum := parseBinaryString(wireState, 'y')
		xy := xnum + ynum
		if xy != num {
			log.Println("sum", strconv.FormatInt(int64(xy), 2))
			log.Println("res", strconv.FormatInt(int64(num), 2))
			log.Fatal("Numbers(sum) do not match", xnum, ynum, xy, num)
		}
	}
	return int(num)
}

func getRandomState() []string {
	lines := make([]string, 0)
	for i := 0; i < 45; i++ {
		num := ""
		if i < 10 {
			num += "0"
		}
		num += strconv.Itoa(i)
		if rand.Int()%2 == 1 {
			lines = append(lines, "x"+num+": 1")
		} else {
			lines = append(lines, "x"+num+": 0")
		}
		if rand.Int()%2 == 1 {
			lines = append(lines, "y"+num+": 1")
		} else {
			lines = append(lines, "y"+num+": 0")
		}
	}
	return lines
}

func printDotFile(wireState map[string]bool, gates map[Gate]bool) {
	gateCounter := 0
	for gate, _ := range gates {
		op := gate.operation + strconv.Itoa(gateCounter)
		println(gate.a, "--", op)
		println(gate.b, "--", op)
		println(op, "--", gate.output)
		gateCounter += 1
	}
}

func parseBinaryString(wireState map[string]bool, startsWith rune) int64 {
	binaryString := createBinaryString(wireState, startsWith)
	num, err := strconv.ParseInt(binaryString, 2, 64)
	//log.Println(startsWith, binaryString)
	if err != nil {
		log.Fatal("Could not parse binary number", binaryString)
	}
	return num
}

func createBinaryString(wireState map[string]bool, startsWith rune) string {
	zstate := make([]bool, 64)
	for wire, b := range wireState {
		if rune(wire[0]) != startsWith {
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
