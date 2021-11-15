package emu

import "fmt"

type Register uint8

const (
	REG_A Register = iota
	REG_B
	REG_C
	REG_D
	REG_E
	REG_F
	REG_H
	REG_L
)

func (r Register) String() string {
	switch r {
	case REG_A:
		return "A"
	case REG_B:
		return "B"
	case REG_C:
		return "C"
	case REG_D:
		return "D"
	case REG_E:
		return "E"
	case REG_F:
		return "F"
	case REG_H:
		return "H"
	case REG_L:
		return "L"
	}
	return ""
}

type cpu struct {
	programCounter uint16
	stackPointer   uint16
	reg            map[Register]uint8
}

func NewCpu() cpu {
	return cpu{
		stackPointer:   0,
		programCounter: 0,
		reg: map[Register]uint8{REG_A: 0, REG_B: 0, REG_C: 0, REG_D: 0, REG_E: 0, REG_F: 0, REG_H: 0, REG_L: 0},
	}
}

func (c cpu) Reg8(r Register) uint8 {
	return c.reg[r]
}

func (c cpu) Reg16(l Register, r Register) uint16 {
	pair := [2]Register{l, r}

	allowedPairs := map[[2]Register]bool{
		{REG_A, REG_F}: true,
		{REG_B, REG_C}: true,
		{REG_D, REG_E}: true,
		{REG_H, REG_L}: true,
	}

	_, validPair := allowedPairs[pair]

	if !validPair {
		panic(fmt.Sprintf("Invalid combination of 8-bit registers to access: %s%s", l, r))
	}

	return (uint16(c.Reg8(l)) << 8) + uint16(c.Reg8(r))
}
