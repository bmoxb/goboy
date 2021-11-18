package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterAccess(t *testing.T) {
	c := Cpu{
		programCounter: 0,
		stackPointer:   0,
		reg:            map[Register]uint8{REG_A: 0xAA, REG_F: 0xBB, REG_B: 0x12, REG_C: 0x34, REG_D: 1, REG_E: 2, REG_H: 0, REG_L: 0},
	}

	accessData := map[[2]Register]uint16{
		{REG_A, REG_F}: 0xAABB,
		{REG_B, REG_C}: 0x1234,
		{REG_D, REG_E}: 0x102,
		{REG_H, REG_L}: 0,
	}

	for key, val := range accessData {
		assert.Equal(t, c.Reg16(key[0], key[1]), val, "Access pair of 8-bit registers as single 16-bit value")
	}

	panicData := [4][2]Register{{REG_A, REG_B}, {REG_B, REG_E}, {REG_D, REG_H}, {REG_H, REG_F}}

	for _, pair := range panicData {
		assert.Panics(t, func() { c.Reg16(pair[0], pair[1]) }, "Should panic when attempting to access invalid pair of 8-bit registers")
	}
}

func TestFlagAccess(t *testing.T) {
	c := Cpu{
		programCounter: 0,
		stackPointer:   0,
		reg:            map[Register]uint8{REG_F: 0b10100101},
	}

	data := map[Flag]bool{
		FLAG_ZERO:       true,
		FLAG_SUBTRACT:   false,
		FLAG_HALF_CARRY: true,
		FLAG_CARRY:      false,
	}

	for flag, expected := range data {
		assert.Equal(t, c.Flag(flag), expected, "Access individual CPU flag")
	}
}
