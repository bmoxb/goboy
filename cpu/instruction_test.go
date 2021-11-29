package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmeticLogicInstrCycles(t *testing.T) {
	dataRegA := map[Operand]int{
		OpReg8{reg: RegA}:                        1,
		OpReg8{reg: RegB}:                        1,
		OpReg8{reg: RegC}:                        1,
		OpReg8{reg: RegD}:                        1,
		OpReg8{reg: RegE}:                        1,
		OpReg8{reg: RegH}:                        1,
		OpReg8{reg: RegL}:                        1,
		OpIndex{index: IndexReg16{reg: RegHL{}}}: 2,
		OpImm8{value: 0}:                         2,
	}

	dataExtended := map[[2]Operand]int{
		{OpReg16{reg: RegHL{}}, OpReg16{reg: RegBC{}}}:           2,
		{OpReg16{reg: RegHL{}}, OpReg16{reg: RegDE{}}}:           2,
		{OpReg16{reg: RegHL{}}, OpReg16{reg: RegHL{}}}:           2,
		{OpReg16{reg: RegHL{}}, OpReg16{reg: RegStackPointer{}}}: 2,
		{OpReg16{reg: RegStackPointer{}}, OpImm8{value: 0}}:      4,
	}

	fullData := map[[2]Operand]int{}
	for operand, cycles := range dataRegA {
		fullData[[2]Operand{OpReg8{reg: RegA}, operand}] = cycles
	}
	for operands, cycles := range dataExtended {
		fullData[operands] = cycles
	}

	// Test ADD instruction (with additional data):

	for operands, cycles := range fullData {
		instr := InstrAdd{operand1: operands[0], operand2: operands[1]}
		assert.Equal(t, instr.Cycles(), cycles, "Addition instruction '%s' should take %d cycles", instr, cycles)
	}

	// Test non-ADD instructions (with just register A data):

	otherBuilders := [](func(Operand) Instruction){
		func(operand Operand) Instruction { return InstrSub{operand} },
		func(operand Operand) Instruction { return InstrSbc{operand} },
		func(operand Operand) Instruction { return InstrAnd{operand} },
		func(operand Operand) Instruction { return InstrAdc{operand} },
		func(operand Operand) Instruction { return InstrXor{operand} },
		func(operand Operand) Instruction { return InstrOr{operand} },
		func(operand Operand) Instruction { return InstrCp{operand} },
	}
	for _, builder := range otherBuilders {
		for operand, cycles := range dataRegA {
			instr := builder(operand)
			assert.Equal(t, instr.Cycles(), cycles, "Arithmetic instruction '%s' should take %d cycles", instr, cycles)
		}
	}

	incDecData := map[Operand]int{
		OpReg8{reg: RegA}:                        1,
		OpReg8{reg: RegB}:                        1,
		OpReg8{reg: RegC}:                        1,
		OpReg8{reg: RegD}:                        1,
		OpReg8{reg: RegE}:                        1,
		OpReg8{reg: RegH}:                        1,
		OpReg8{reg: RegL}:                        1,
		OpIndex{index: IndexReg16{reg: RegHL{}}}: 3,
		OpReg16{reg: RegBC{}}:                    2,
		OpReg16{reg: RegDE{}}:                    2,
		OpReg16{reg: RegHL{}}:                    2,
		OpReg16{reg: RegStackPointer{}}:          2,
	}

	incDecBuilder := [](func(Operand) Instruction){
		func(operand Operand) Instruction { return InstrInc{operand} },
		func(operand Operand) Instruction { return InstrDec{operand} },
	}
	for _, builder := range incDecBuilder {
		for operand, cycles := range incDecData {
			instr := builder(operand)
			assert.Equal(t, instr.Cycles(), cycles, "Increase/decrease instruction '%s' should take %d cycles", instr, cycles)
		}
	}
}
