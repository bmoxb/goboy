package cpu

type Register8 uint8

const (
	RegA Register8 = iota
	RegB
	RegC
	RegD
	RegE
	RegF
	RegH
	RegL
)

func (r Register8) String() string {
	switch r {
	case RegA:
		return "A"
	case RegB:
		return "B"
	case RegC:
		return "C"
	case RegD:
		return "D"
	case RegE:
		return "E"
	case RegF:
		return "F"
	case RegH:
		return "H"
	case RegL:
		return "L"
	}
	return ""
}

type Register16 interface {
	String() string
	MostSigComponent() Register8
	LeastSigComponent() Register8
}

type RegStackPointer struct{}

func (r RegStackPointer) String() string               { return "SP" }
func (r RegStackPointer) MostSigComponent() Register8  { panic("") }
func (r RegStackPointer) LeastSigComponent() Register8 { panic("") }

type RegProgramCounter struct{}

func (r RegProgramCounter) String() string               { return "PC" }
func (r RegProgramCounter) MostSigComponent() Register8  { panic("") }
func (r RegProgramCounter) LeastSigComponent() Register8 { panic("") }

type RegAF struct{}

func (r RegAF) String() string               { return "AL" }
func (r RegAF) MostSigComponent() Register8  { return RegA }
func (r RegAF) LeastSigComponent() Register8 { return RegF }

type RegBC struct{}

func (r RegBC) String() string               { return "BC" }
func (r RegBC) MostSigComponent() Register8  { return RegB }
func (r RegBC) LeastSigComponent() Register8 { return RegC }

type RegDE struct{}

func (r RegDE) String() string               { return "DB" }
func (r RegDE) MostSigComponent() Register8  { return RegD }
func (r RegDE) LeastSigComponent() Register8 { return RegE }

type RegHL struct{}

func (r RegHL) String() string               { return "HL" }
func (r RegHL) MostSigComponent() Register8  { return RegH }
func (r RegHL) LeastSigComponent() Register8 { return RegL }

type Flag uint8

const (
	FlagZero Flag = iota
	FlagSubtract
	FlagHalfCarry
	FlagCarry
)

func (f Flag) String() string {
	switch f {
	case FlagZero:
		return "Z"
	case FlagSubtract:
		return "N"
	case FlagHalfCarry:
		return "H"
	case FlagCarry:
		return "C"
	}
	return ""
}
