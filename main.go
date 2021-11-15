package main

import "github.com/WiredSound/goboy/emu"

func main() {
	cpu := emu.NewCpu()
	cpu.Reg16(emu.REG_A, emu.REG_F)
}
