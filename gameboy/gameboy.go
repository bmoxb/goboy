package gameboy

import (
	"github.com/WiredSound/goboy/cpu"
	"github.com/WiredSound/goboy/media"
	"github.com/WiredSound/goboy/memory"
)

type gameboy struct {
	cpu cpu.Cpu
	mem memory.Memory
}

func New() gameboy {
	return gameboy{
		cpu: cpu.New(),
		mem: memory.New(),
	}
}

func (g gameboy) Update(context media.Context, buttons map[media.Button]bool) {}
