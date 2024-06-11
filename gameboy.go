package main

type Memory [0xFFFF]uint8

type Gameboy struct {
	CPU     *CPU
	Memory  *Memory
	MCycles int // Machine cycles
}

func NewGameboy() *Gameboy {
	gb := new(Gameboy)

	gb.CPU = NewCPU()
	gb.Memory = new(Memory)

	gb.MCycles = 0

	return gb
}

func (gb *Gameboy) readPC() uint8 {
	value := gb.readMemory(gb.CPU.PC)

	gb.CPU.PC++

	return value
}

func (gb *Gameboy) readSP() uint8 {
	value := gb.readMemory(gb.CPU.SP)

	gb.CPU.SP++

	return value
}

func (gb *Gameboy) readMemory(addr uint16) uint8 {
	value := gb.Memory[addr]

	gb.MCycles++

	return value
}

func (gb *Gameboy) writeMemory(addr uint16, value uint8) {
	gb.Memory[addr] = value

	gb.MCycles++

	return
}
