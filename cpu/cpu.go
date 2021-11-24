package cpu

import (
	"github.com/WiredSound/goboy/memory"
)

const INSTRUCTIONS_PER_SECOND = 1_050_000 // 1.05 MHz
const CYCLE_DURATION = 1.0 / INSTRUCTIONS_PER_SECOND

type Cpu struct {
	programCounter uint16
	stackPointer   uint16
	reg            map[Register8]uint8
}

func New() Cpu {
	return Cpu{
		programCounter: 0x100,
		stackPointer:   0xFFFE,
		reg:            map[Register8]uint8{RegA: 0, RegB: 0, RegC: 0, RegD: 0, RegE: 0, RegF: 0, RegH: 0, RegL: 0},
	}
}

func (c Cpu) Reg8(r Register8) uint8 {
	return c.reg[r]
}

func (c Cpu) Reg16(r Register16) uint16 {
	return (uint16(c.Reg8(r.MostSigComponent())) << 8) + uint16(c.Reg8(r.LeastSigComponent()))
}

func (c Cpu) Flag(f Flag) bool {
	flags := c.Reg8(RegF)

	switch f {
	case FlagZero:
		return (flags & 0b10000000) > 0
	case FlagSubtract:
		return (flags & 0b01000000) > 0
	case FlagHalfCarry:
		return (flags & 0b00100000) > 0
	case FlagCarry:
		return (flags & 0b00010000) > 0
	}

	return false
}

func (c Cpu) Tick(mem *memory.Memory) {}
