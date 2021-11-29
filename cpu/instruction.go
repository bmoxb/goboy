package cpu

import (
	"fmt"
)

type Operand interface {
	String() string
	operand()
}

type OpReg8 struct{ reg Register8 }

func (o OpReg8) String() string { return o.reg.String() }
func (o OpReg8) operand()       {}

type OpReg16 struct{ reg Register16 }

func (o OpReg16) String() string { return o.reg.String() }
func (o OpReg16) operand()       {}

type OpImmediate8 struct{ value uint8 }

func (o OpImmediate8) String() string { return fmt.Sprintf("%d", o.value) }
func (o OpImmediate8) operand()       {}

type Instruction interface {
	String() string
	Cycles() int
}

//
// CPU Control
//

type InstrNop struct{}

func (i InstrNop) String() string { return "nop" }
func (i InstrNop) Cycles() int    { return 1 }

//
// Arithmetic/Logic
//

func binaryArithmeticCycles(operand1, operand2 Operand) int {
	switch operand1 {
	case OpReg8{reg: RegA}:
		switch operand2.(type) {
		case OpImmediate8:
			return 2
		}

		var hl Operand = OpReg16{reg: RegHL{}}
		if operand2 == hl {
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

type InstrAdd struct{ operand1, operand2 Operand }

func (i InstrAdd) String() string { return fmt.Sprintf("ADD %s, %s", i.operand1, i.operand2) }
func (i InstrAdd) Cycles() int    { return binaryArithmeticCycles(i.operand1, i.operand2) }

type InstrAdc struct{ operand Operand }

func (i InstrAdc) String() string { return fmt.Sprintf("ADC A, %s", i.operand) }
func (i InstrAdc) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }

type InstrSub struct{ operand Operand }

func (i InstrSub) String() string { return fmt.Sprintf("SUB %s", i.operand) }
func (i InstrSub) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }

type InstrSbc struct{ operand Operand }

func (i InstrSbc) String() string { return fmt.Sprintf("SBC A, %s", i.operand) }
func (i InstrSbc) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }

type InstrAnd struct{ operand Operand }

func (i InstrAnd) String() string { return fmt.Sprintf("ADD %s", i.operand) }
func (i InstrAnd) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }

type InstrOr struct{ operand Operand }

func (i InstrOr) String() string { return fmt.Sprintf("OR %s", i.operand) }
func (i InstrOr) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }

type InstrXor struct{ operand Operand }

func (i InstrXor) String() string { return fmt.Sprintf("XOR %s", i.operand) }
func (i InstrXor) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }

type InstrCp struct{ operand Operand }

func (i InstrCp) String() string { return fmt.Sprintf("CP %s", i.operand) }
func (i InstrCp) Cycles() int    { return binaryArithmeticCycles(OpReg8{reg: RegA}, i.operand) }

func unaryArithmeticCycles(operand Operand) int {
	switch operand.(type) {
	case OpReg8:
		return 1
	}
	hl := OpReg16{reg: RegHL{}}
	if operand == hl {
		return 3
	}
	return 2
}

type InstrInc struct{ operand Operand }

func (i InstrInc) String() string { return fmt.Sprintf("INC %s", i.operand) }
func (i InstrInc) Cycles() int    { return unaryArithmeticCycles(i.operand) }

type InstrDec struct{ operand Operand }

func (i InstrDec) String() string { return fmt.Sprintf("DEC %s", i.operand) }
func (i InstrDec) Cycles() int    { return unaryArithmeticCycles(i.operand) }

type InstrDaa struct{}

func (i InstrDaa) String() string { return "DAA" }
func (i InstrDaa) Cycles() int    { return 1 }

type InstrScf struct{}

func (i InstrScf) String() string { return "SCF" }
func (i InstrScf) Cycles() int    { return 1 }

type InstrCpl struct{}

func (i InstrCpl) String() string { return "CPL" }
func (i InstrCpl) Cycles() int    { return 1 }

type InstrCcf struct{}

func (i InstrCcf) String() string { return "CCF" }
func (i InstrCcf) Cycles() int    { return 1 }
