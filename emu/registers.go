package emu

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

type Flag uint8

const (
	FLAG_ZERO Flag = iota
	FLAG_SUBTRACT
	FLAG_HALF_CARRY
	FLAG_CARRY
)

func (f Flag) String() string {
	switch f {
	case FLAG_ZERO:
		return "Z"
	case FLAG_SUBTRACT:
		return "N"
	case FLAG_HALF_CARRY:
		return "H"
	case FLAG_CARRY:
		return "C"
	}
	return ""
}
