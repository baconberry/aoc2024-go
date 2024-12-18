package aoc

import (
	"aoc2024/util"
	"math"
	"strconv"
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

func Seventeen(lines []string) string {
	registers := util.ParseIntGrid(lines)
	var emu Emu = Emu{}
	emu.A = registers[0][0]
	emu.B = registers[1][0]
	emu.C = registers[2][0]
	program := util.ParseIntArray(lines[len(registers)+1])
	emu.program = program
	return emu.execute()
}

func (emu *Emu) execute() string {
	output := ""
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
			output += strconv.Itoa(o) + ","
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
	return output
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
