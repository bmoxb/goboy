package cartridge

import "fmt"

type rom struct {
	data [32768]uint8 // 32 KiB ROM (mapped to 0x0000 - 0x7FFF)
}

func (r rom) Read8(addr uint16) uint8 {
	if addr > 0x7FFF {
		panic(fmt.Sprintf("Cannot read cartridge ROM at address: %d", addr))
	}
	return r.data[addr]
}
