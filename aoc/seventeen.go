package aoc

import (
	"aoc2024/util"
	"math"
	"strconv"
	"strings"
)

type Emu struct {
	A, B, C int
	pc      int
	program []int
}
type OpCode int

const (
	adv OpCode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func Seventeen(lines []string, part int) string {
	registers := util.ParseIntGrid(lines)
	var emu = Emu{}
	emu.A = registers[0][0]
	emu.B = registers[1][0]
	emu.C = registers[2][0]
	program := util.ParseIntArray(lines[len(registers)+1])
	emu.program = program
	if part == 2 {
		for i := 1; i < 8; i++ {
			loc := heuristic(program, i)
			if loc >= 0 {
				return strconv.Itoa(loc)
			}
		}
	}
	result, _ := emu.execute(nil)
	return SplitToString(result, ",")
}

func SplitToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}

func heuristic(program []int, prevHint int) int {
	emu := Emu{program: program}
	programLen := len(program)
mainFor:
	for i := 0; i < 8; i++ {
		locI := (prevHint << 3) | i
		emu.A = locI
		emu.B = 0
		emu.C = 0
		emu.pc = 0
		localResult, _ := emu.execute(nil)
		//log.Println(locI, strconv.FormatInt(int64(i), 2), localResult)
		resLen := len(localResult)
		for i := 0; i < len(localResult); i++ {
			if localResult[resLen-1-i] != program[programLen-1-i] {
				continue mainFor
			}
		}
		if len(program) == len(localResult) {
			return locI
		}
		futureRes := heuristic(program, locI)
		if futureRes >= 0 {
			return futureRes
		}
	}
	return -1
}

type ProgramPredicate func(arr *[]int) bool

func (emu *Emu) execute(cmpFn *ProgramPredicate) ([]int, bool) {
	output := make([]int, 0)
	for emu.endNotReached() {
		opcode := emu.readOpcode()
		switch opcode {
		case adv:
			denominator := emu.readComboOperand()
			result := float64(emu.A) / math.Pow(2, float64(denominator))
			emu.A = int(result)
		case bxl:
			o := emu.readLiteralOperand()
			result := emu.B ^ o
			emu.B = result
		case bst:
			o := emu.readComboOperand() % 8
			emu.B = o
		case jnz:
			if emu.A != 0 {
				emu.pc = emu.readLiteralOperand()
				emu.pc -= 2
			}
		case bxc:
			result := emu.B ^ emu.C
			emu.B = result
		case out:
			o := emu.readComboOperand() % 8
			output = append(output, o)
			if cmpFn != nil && !(*cmpFn)(&output) {
				return output, false
			}
		case bdv:
			denominator := emu.readComboOperand()
			result := float64(emu.A) / math.Pow(2, float64(denominator))
			emu.B = int(result)
		case cdv:
			denominator := emu.readComboOperand()
			result := float64(emu.A) / math.Pow(2, float64(denominator))
			emu.C = int(result)
		}
		emu.pc += 2
	}
	return output, true
}

func (emu *Emu) readOpcode() OpCode {
	code := emu.program[emu.pc]
	code = code % 8
	return OpCode(code)
}

func (emu *Emu) readComboOperand() int {
	operand := emu.program[emu.pc+1]
	operand = operand % 8
	if operand <= 3 {
		return operand
	}
	switch operand {
	case 4:
		return emu.A
	case 5:
		return emu.B
	case 6:
		return emu.C
	case 7:
		panic("Combo operand 7 found, invalid program")
	}
	panic("Invalid operand")
}

func (emu *Emu) readLiteralOperand() int {
	return emu.program[emu.pc+1]

}

func (emu *Emu) endNotReached() bool {
	return emu.pc < len(emu.program)
}
