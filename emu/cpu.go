package emu

import "log"

type Cpu struct {}

func (Cpu) Init() {
    log.Println("CPU initialised")
}
