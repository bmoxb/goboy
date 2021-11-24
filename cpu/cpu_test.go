package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagAccess(t *testing.T) {
	c := Cpu{
		programCounter: 0,
		stackPointer:   0,
		reg:            map[Register8]uint8{RegF: 0b10100101},
	}

	data := map[Flag]bool{
		FlagZero:      true,
		FlagSubtract:  false,
		FlagHalfCarry: true,
		FlagCarry:     false,
	}

	for flag, expected := range data {
		assert.Equal(t, c.Flag(flag), expected, "Access individual CPU flag")
	}
}
