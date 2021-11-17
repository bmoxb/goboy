package cartridge

type Cartridge interface {
	Read8(addr uint16) uint8
}
