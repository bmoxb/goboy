package cpu

import (
	"fmt"
)

// Indicies

type Index interface {
	String() string
	IndexDescription() string
}

type IndexReg16 struct{ reg Register16 }

func (i IndexReg16) String() string { return i.reg.String() }
func (i IndexReg16) IndexDescription() string {
	return "todo"
}

// Operands

type Operand interface {
	String() string
	OperandDescription() string
}

type OpReg8 struct{ reg Register8 }

func (o OpReg8) String() string { return o.reg.String() }
func (o OpReg8) OperandDescription() string {
	return "8-bit CPU register"
}

type OpReg16 struct{ reg Register16 }

func (o OpReg16) String() string { return o.reg.String() }
func (o OpReg16) OperandDescription() string {
	return "16-bit CPU register"
}

type OpImm8 struct{ value uint8 }

func (o OpImm8) String() string { return fmt.Sprintf("%d", o.value) }
func (o OpImm8) OperandDescription() string {
	return "immediate 8-bit value"
}

type OpIndex struct{ index Index }

func (o OpIndex) String() string { return "(" + o.index.String() + ")" }
func (o OpIndex) OperandDescription() string {
	return "todo"
}

// Conditions

type Condition interface {
	String() string
	ConditionDescription() string
}

type ConditionNZ struct{}

func (o ConditionNZ) String() string { return "NZ" }
func (o ConditionNZ) ConditionDescription() string {
	return "Z flag not set"
}

type ConditionN struct{}

func (o ConditionN) String() string { return "Z" }
func (o ConditionN) ConditionDescription() string {
	return "Z flag set"
}

type ConditionNC struct{}

func (o ConditionNC) String() string { return "NC" }
func (o ConditionNC) ConditionDescription() string {
	return "C flag not set"
}

type ConditionC struct{}

func (o ConditionC) String() string { return "C" }
func (o ConditionC) ConditionDescription() string {
	return "C flag set"
}

// Instructions

type Instruction interface {
	String() string
	Cycles() int
	Description() string
}

//
// CPU Control
//

type InstrNOP struct{}

func (i InstrNOP) String() string { return "NOP" }
func (i InstrNOP) Cycles() int    { return 1 }
func (i InstrNOP) Description() string {
	return "disables interrupts after instruction following this is executed"
}

type InstrSTOP struct{}

func (i InstrSTOP) String() string { return "STOP 0" }
func (i InstrSTOP) Cycles() int    { return 1 }
func (i InstrSTOP) Description() string {
	return "halt CPU & LCD display until button pressed"
}

type InstrHALT struct{}

func (i InstrHALT) String() string { return "HALT" }
func (i InstrHALT) Cycles() int    { return 1 }
func (i InstrHALT) Description() string {
	return "halt execution until an interrupt occurs"
}

//
// Jumps/Calls
//

func jumpString(name string, operand Operand, condition *Condition) string {
	if condition != nil {
		return fmt.Sprintf("%s %s, %s", name, *condition, operand)
	} else {
		return fmt.Sprintf("%s %s", name, operand)
	}
}

type InstrJP struct {
	operand   Operand
	condition *Condition
}

func (i InstrJP) String() string { return jumpString("JP", i.operand, i.condition) }
func (i InstrJP) Cycles() int {
	hl := OpIndex{index: IndexReg16{reg: RegHL{}}}
	if i.operand == hl {
		return 4
	} else {
		return 12
	}
}
func (i InstrJP) Description() string {
	s := "jump to " + i.operand.String()
	if i.condition != nil {
		s += " if " + (*i.condition).String()
	}
	return s
}

type InstrJR struct {
	operand   Operand
	condition *Condition
}

func (i InstrJR) String() string { return jumpString("JR", i.operand, i.condition) }
func (i InstrJR) Cycles() int    { return 8 }
func (i InstrJR) Description() string {
	s := fmt.Sprintf("add %s to program counter", i.operand)
	if i.condition != nil {
		s += " if " + (*i.condition).String()
	}
	return s
}

type InstrCALL struct {
	operand   Operand
	condition *Condition
}

func (i InstrCALL) String() string { return jumpString("CALL", i.operand, i.condition) }
func (i InstrCALL) Cycles() int    { return 12 }
func (i InstrCALL) Description() string {
	s := "push next instruction address to stack and jump to " + i.operand.String()
	if i.condition != nil {
		s += " if " + (*i.condition).String()
	}
	return s
}

type InstrRET struct{ condition *Condition }

func (i InstrRET) String() string {
	if i.condition != nil {
		return fmt.Sprintf("RET %s", *i.condition)
	} else {
		return "RET"
	}
}
func (i InstrRET) Cycles() int { return 8 }
func (i InstrRET) Description() string {
	s := "pop address from stack and jump to it"
	if i.condition != nil {
		s += " if " + (*i.condition).String()
	}
	return s
}

type InstrRETI struct{}

func (i InstrRETI) String() string { return "RETI" }
func (i InstrRETI) Cycles() int    { return 8 }
func (i InstrRETI) Description() string {
	return "pop address from stack stack, jump to it, and enable interrupts"
}

//
// Arithmetic/Logic
//

func binaryArithmeticCycles(operand1, operand2 Operand) int {
	switch operand1 {
	case OpReg8{reg: RegA}:
		switch operand2.(type) {
		case OpImm8:
			return 2
		}

		var indexHL Operand = OpIndex{index: IndexReg16{reg: RegHL{}}}
		if operand2 == indexHL {
			return 2
		}

		return 1

	case OpReg16{reg: RegHL{}}:
		return 2

	case OpReg16{reg: RegStackPointer{}}:
		return 4
	}

	return 0
}

type InstrADD struct{ operand1, operand2 Operand }

func (i InstrADD) String() string { return fmt.Sprintf("ADD %s, %s", i.operand1, i.operand2) }
func (i InstrADD) Cycles() int    { return binaryArithmeticCycles(i.operand1, i.operand2) }
func (i InstrADD) Description() string {
	return fmt.Sprintf("add the two operands %s and %s, result stored in the first operand %s", i.operand1, i.operand2, i.operand1)
}

type InstrADC struct{ operand Operand }

func (i InstrADC) String() string { return fmt.Sprintf("ADC A, %s", i.operand) }
func (i InstrADC) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }
func (i InstrADC) Description() string {
	return fmt.Sprintf("add operand %s and carry flag to register A, result stored in A", i.operand)
}

type InstrSUB struct{ operand Operand }

func (i InstrSUB) String() string { return fmt.Sprintf("SUB %s", i.operand) }
func (i InstrSUB) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }
func (i InstrSUB) Description() string {
	return fmt.Sprintf("subtract operand %s from register A, result stored in A", i.operand)
}

type InstrSBC struct{ operand Operand }

func (i InstrSBC) String() string { return fmt.Sprintf("SBC A, %s", i.operand) }
func (i InstrSBC) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }
func (i InstrSBC) Description() string {
	return fmt.Sprintf("subtract operand %s and carry flag from register A, result stored in A", i.operand)
}

type InstrAND struct{ operand Operand }

func (i InstrAND) String() string { return fmt.Sprintf("ADD %s", i.operand) }
func (i InstrAND) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }
func (i InstrAND) Description() string {
	return fmt.Sprintf("logical AND performed on operand %s with register A, result stored in A", i.operand)
}

type InstrOR struct{ operand Operand }

func (i InstrOR) String() string { return fmt.Sprintf("OR %s", i.operand) }
func (i InstrOR) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }
func (i InstrOR) Description() string {
	return fmt.Sprintf("logical inclusive OR performed on operand %s with register A, result stored in A", i.operand)
}

type InstrXOR struct{ operand Operand }

func (i InstrXOR) String() string { return fmt.Sprintf("XOR %s", i.operand) }
func (i InstrXOR) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }
func (i InstrXOR) Description() string {
	return fmt.Sprintf("logical exclusive OR performed on operand %s with register A, result stored in A", i.operand)
}

type InstrCP struct{ operand Operand }

func (i InstrCP) String() string { return fmt.Sprintf("CP %s", i.operand) }
func (i InstrCP) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }
func (i InstrCP) Description() string {
	return fmt.Sprintf("compared operand %s with A", i.operand)
}

func unaryArithmeticCycles(operand Operand) int {
	switch operand.(type) {
	case OpReg8:
		return 1
	}

	indexHL := OpIndex{index: IndexReg16{reg: RegHL{}}}
	if operand == indexHL {
		return 3
	}

	return 2
}

type InstrINC struct{ operand Operand }

func (i InstrINC) String() string { return fmt.Sprintf("INC %s", i.operand) }
func (i InstrINC) Cycles() int    { return unaryArithmeticCycles(i.operand) }
func (i InstrINC) Description() string {
	return fmt.Sprintf("increment register %s", i.operand)
}

type InstrDEC struct{ operand Operand }

func (i InstrDEC) String() string { return fmt.Sprintf("DEC %s", i.operand) }
func (i InstrDEC) Cycles() int    { return unaryArithmeticCycles(i.operand) }
func (i InstrDEC) Description() string {
	return fmt.Sprintf("decrement specified register %s", i.operand)
}

type InstrDAA struct{}

func (i InstrDAA) String() string { return "DAA" }
func (i InstrDAA) Cycles() int    { return 1 }
func (i InstrDAA) Description() string {
	return "decimal adjust register A (binary coded decimal representation)"
}

type InstrSCF struct{}

func (i InstrSCF) String() string { return "SCF" }
func (i InstrSCF) Cycles() int    { return 1 }
func (i InstrSCF) Description() string {
	return "set carry flag"
}

type InstrCPL struct{}

func (i InstrCPL) String() string { return "CPL" }
func (i InstrCPL) Cycles() int    { return 1 }
func (i InstrCPL) Description() string {
	return "complement register A (flip all bits)"
}

type InstrCCF struct{}

func (i InstrCCF) String() string { return "CCF" }
func (i InstrCCF) Cycles() int    { return 1 }
func (i InstrCCF) Description() string {
	return "complement carry flag (flip bit value)"
}
