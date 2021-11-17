package main

import "github.com/WiredSound/goboy/cpu"

func main() {
	c := cpu.New()
	c.Reg16(cpu.REG_A, cpu.REG_F)
}
