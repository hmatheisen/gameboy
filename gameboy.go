package main

type Memory [0xFFFF]uint8

type Gameboy struct {
	CPU    *CPU
	Memory *Memory
}
