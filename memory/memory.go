package memory

import "github.com/WiredSound/goboy/cartridge"

type Memory struct {
	cart            cartridge.Cartridge // game cartridge
	wram            [2][4096]uint8      // working RAM, two regions of 4 KiB
	hram            [127]uint8          // high RAM
	interruptEnable bool                // interrupt enable register (IE)
}

func New() Memory {
	return Memory{
		interruptEnable: false,
	}
}

func (m Memory) Read8(addr uint16) uint8 {
	if addr < 0x8000 { // Cartridge
		return m.cart.Read8(addr)
	} else if addr < 0xA000 { // Video memory
	} else if addr < 0xC000 { // external RAM from cartridge, switchable bank
		return m.cart.Read8(addr)
	} else if addr < 0xD000 { // working RAM
		return m.wram[0][addr-0xC000]
	} else if addr < 0xE000 { // working RAM (switchable bank 1 - 7 for GameBoy Color)
		return m.wram[1][addr-0xD000]
	} else if addr < 0xFE00 { // mirror/echo of working RAM 0xC000 - 0xDDFF
		return m.Read8(addr - 0x2000)
	} else if addr < 0xFEA0 { // object attribute memory
	} else if addr < 0xFF00 { // not usable - this area of memory shouldn't be accessed
		return 0
	} else if addr < 0xFF80 { // I/O
	} else if addr < 0xFFFF { // high RAM
		return m.hram[addr-0xFF80]
	}
	// else: 0xFFFF - interrupt enable
	if m.interruptEnable {
		return 1
	} else {
		return 0
	}
}
