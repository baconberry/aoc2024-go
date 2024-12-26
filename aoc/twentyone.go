package aoc

import (
	"aoc2024/util"
	"strconv"
)

type PairPoint struct {
	a, b util.Point
}
type Keypad struct {
	keyPos       []util.Point
	allMovements map[PairPoint][]string
	allChars     []rune
	toAvoid      util.Point
}

func (k *Keypad) CalculateAllMovements() {
	for _, a := range k.allChars {
		for _, b := range k.allChars {
			pair := PairPoint{a: k.keyPos[a], b: k.keyPos[b]}
			allMoves := k.calculateMoves(pair.a, pair.b, make([]util.Direction, 0))
			arr := make([]string, 0)
			for _, move := range allMoves {
				moveString := dirToString(&move) + "A"
				arr = append(arr, moveString)
			}
			k.allMovements[pair] = arr
		}
	}
}

func dirToString(move *[]util.Direction) string {
	s := ""
	for _, direction := range *move {
		s += string(dirToCommand(direction))
	}
	return s
}

func (k *Keypad) calculateMoves(a util.Point, b util.Point, path []util.Direction) [][]util.Direction {
	if a == b {
		return [][]util.Direction{path}
	}
	aDist := a.DistanceTo(b)
	allPaths := make([][]util.Direction, 0)
	for _, direction := range util.CARDINAL_DIRECTIONS {
		c := a.PlusDirection(direction)
		if c.DistanceTo(b) > aDist || c == k.toAvoid {
			continue
		}
		paths := k.calculateMoves(c, b, append(path, direction))
		for _, p := range paths {
			allPaths = append(allPaths, p)
		}
	}
	return allPaths
}

func Twentyone(lines []string, part int) int {
	numpad := createNumKeyboard()
	numpad.CalculateAllMovements()
	arrows := createArrowMap()
	arrows.CalculateAllMovements()
	humanArrows := createArrowMap()
	humanArrows.toAvoid = util.Point{X: -1, Y: -1}
	humanArrows.CalculateAllMovements()
	depth := 3
	if part == 2 {
		depth = 26
	}
	currentBotPos := make([]util.Point, depth+1)
	sum := 0
	keypads := make([]Keypad, depth+1)
	memo := make(map[string]int)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		currentBotPos[depth] = numpad.keyPos['A']
		keypads[depth] = numpad

		for i := 0; i < depth; i++ {
			currentBotPos[i] = arrows.keyPos['A']
			keypads[i] = arrows
		}
		keypads[0] = humanArrows
		keyCommands := moveBotToKey(&keypads, line, depth, currentBotPos, &memo)
		codes := util.ParseIntArray(line)
		//log.Println(line, " ", keyCommands, " ", keyCommands, " ", codes[0])
		sum += keyCommands * codes[0]
	}
	return sum
}

func moveBotToKey(keypads *[]Keypad, codes string, depth int, botPos []util.Point, memo *map[string]int) int {
	if depth == 0 {
		return len(codes)
	}
	allCodes := 0
	keypad := (*keypads)[depth]
	memoKey := getMemoKeyAlt(depth, codes)
	memoRes, ok := (*memo)[memoKey]
	if ok {
		return memoRes
	}
	for _, c := range codes {
		current := (botPos)[depth]
		dest := keypad.keyPos[c]
		moves := keypad.allMovements[PairPoint{current, dest}]
		minCommand := -1
		for _, move := range moves {
			localCommand := moveBotToKey(keypads, move, depth-1, botPos, memo)
			if minCommand == -1 || localCommand < minCommand {
				minCommand = localCommand
			}
		}
		(botPos)[depth] = dest
		allCodes += minCommand
	}
	(*memo)[memoKey] = allCodes
	return allCodes
}

func getMemoKeyAlt(depth int, codes string) string {
	return strconv.Itoa(depth) + "_" + codes
}

func dirToCommand(d util.Direction) rune {
	switch d {
	case util.N:
		return 'v'
	case util.S:
		return '^'
	case util.E:
		return '>'
	case util.W:
		return '<'
	case util.NONE:
		return 'A'
	}
	return '?'
}

func createArrowMap() Keypad {
	res := Keypad{
		keyPos:       nil,
		allMovements: make(map[PairPoint][]string),
		allChars:     make([]rune, 0),
		toAvoid:      util.Point{Y: 1},
	}
	keypad := make([]util.Point, 'v'+1)

	positions := [][]rune{
		{'<', 'v', '>'},
		{' ', '^', 'A'},
	}
	for y, position := range positions {
		for x, c := range position {
			keypad[c] = util.Point{X: x, Y: y}
			if c == ' ' {
				continue
			}
			res.allChars = append(res.allChars, c)
		}
	}

	res.keyPos = keypad
	return res
}
func createNumKeyboard() Keypad {
	res := Keypad{
		keyPos:       nil,
		allMovements: make(map[PairPoint][]string),
		allChars:     make([]rune, 0),
		toAvoid:      util.Point{},
	}
	keypad := make([]util.Point, 'Z')

	positions := [][]rune{
		{' ', '0', 'A'},
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}
	for y, position := range positions {
		for x, c := range position {
			keypad[c] = util.Point{X: x, Y: y}
			if c == ' ' {
				continue
			}
			res.allChars = append(res.allChars, c)
		}
	}
	res.keyPos = keypad

	return res
}
