package cpu

import (
	"fmt"
)

type Operand interface {
	String() string
}

type OpReg8 struct{ reg Register8 }

func (o OpReg8) String() string { return o.reg.String() }

type OpReg16 struct{ reg Register16 }

func (o OpReg16) String() string { return o.reg.String() }

type OpImmediate8 struct{ value uint8 }

func (o OpImmediate8) String() string { return fmt.Sprintf("%d", o.value) }

type Instruction interface {
	String() string
	Cycles() int
}

// CPU Control

type InstrNop struct{}

func (i InstrNop) String() string { return "nop" }
func (i InstrNop) Cycles() int    { return 1 }

// Arithmetic/Logic

type InstrAdd struct{ operand1, operand2 Operand }

func (i InstrAdd) String() string { return fmt.Sprintf("add %s, %s", i.operand1, i.operand2) }
func (i InstrAdd) Cycles() int {
	switch i.operand1 {
	case OpReg8{reg: RegA}:
		switch i.operand2.(type) {
		case OpImmediate8:
			return 2
		}

		var hl Operand = OpReg16{reg: RegHL{}}
		if i.operand2 == hl {
			return 2
		}

		return 1

	case OpReg16{reg: RegHL{}}:
		return 2

	case RegStackPointer{}:
		return 4
	}
	return 0
}

type InstrAdc struct{ operand Operand }

func (i InstrAdc) String() string { return fmt.Sprintf("adc %s", i.operand) }
func (i InstrAdc) Cycles() int    { return 2 }

type InstrSub struct{ operand1, operand2 Operand }

func (i InstrSub) String() string { return fmt.Sprintf("sub %s, %s", i.operand1, i.operand2) }
func (i InstrSub) Cycles() int    { return 2 }

type InstrInc struct{ operand Operand }

func (i InstrInc) String() string { return fmt.Sprintf("inc %s", i.operand) }
func (i InstrInc) Cycles() int    { return 2 }

type InstrDec struct{ operand Operand }

func (i InstrDec) String() string { return fmt.Sprintf("dec %s", i.operand) }
func (i InstrDec) Cycles() int    { return 2 }
