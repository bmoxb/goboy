package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmeticLogicInstrCycles(t *testing.T) {
	assert.Equal(t, InstrAdd{operand1: OpReg8{reg: RegA}, operand2: RegB}.Cycles(), 1)
	assert.Equal(t, InstrAdd{operand1: OpReg8{reg: RegA}, operand2: OpReg16{reg: RegHL{}}}.Cycles(), 2)
	assert.Equal(t, InstrAdd{operand1: OpReg16{reg: RegHL{}}, operand2: OpReg16{reg: RegDE{}}}.Cycles(), 2)
}
