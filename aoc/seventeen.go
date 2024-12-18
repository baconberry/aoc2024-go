package aoc

import (
	"aoc2024/util"
	"log"
	"math"
	"runtime"
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
	var emu Emu = Emu{}
	emu.A = registers[0][0]
	emu.B = registers[1][0]
	emu.C = registers[2][0]
	program := util.ParseIntArray(lines[len(registers)+1])
	emu.program = program
	if part == 2 {
		return strconv.Itoa(solveProgram(program, emu.B, emu.C))
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

func solveProgram(program []int, b, c int) int {
	ichan := make(chan int, 1000)
	programLen := len(program)
	cmpFn := func(arr *[]int) bool {
		if len(*arr) > programLen {
			return false
		}
		for i := 0; i < len(*arr); i++ {
			if (*arr)[i] != program[i] {
				return false
			}
		}

		return true
	}
	gResult := math.MinInt
	log.Println("Starting parallel computation with cpus: ", runtime.NumCPU())
	multiplier := 1000000
	endRange := 1000000
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			var emu Emu = Emu{}
			emu.program = program
			for {
				start := <-ichan
				start = start * multiplier
				for i := start; i <= start+endRange; i++ {
					emu.pc = 0
					emu.A = i
					emu.B = b
					emu.C = c
					result, found := emu.execute((*ProgramPredicate)(&cmpFn))
					if found {
						resLen := len(result)
						if resLen == len(program) {
							if gResult == math.MinInt {
								gResult = i
							}
						}
					}
				}
			}

		}()
	}
	for i := 0; i < math.MaxInt; i++ {
		if gResult > math.MinInt {
			return gResult
		}
		ichan <- i
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
