package main

type OPCode uint8

type CPU struct {
	A, F uint8
	B, C uint8
	D, E uint8
	H, L uint8
	PC   uint16
}

func NewCPU() *CPU {
	cpu := new(CPU)

	// FIXME: Probably false
	cpu.PC = 0x0000

	return cpu
}

// Register values as 16 bits group
func (c CPU) AF() uint16 { return uint16(c.A)<<8 | uint16(c.F) }
func (c CPU) BC() uint16 { return uint16(c.B)<<8 | uint16(c.C) }
func (c CPU) DE() uint16 { return uint16(c.D)<<8 | uint16(c.E) }
func (c CPU) HL() uint16 { return uint16(c.H)<<8 | uint16(c.L) }

func uint16ToHiLo(value uint16) (Hi uint8, Lo uint8) {
	Hi = uint8(value >> 8)
	Lo = uint8(value & 0x00FF)
	return
}

// Setters for 16 bits registers
func (c *CPU) SetAF(value uint16) { c.A, c.F = uint16ToHiLo(value) }
func (c *CPU) SetBC(value uint16) { c.B, c.C = uint16ToHiLo(value) }
func (c *CPU) SetDE(value uint16) { c.D, c.E = uint16ToHiLo(value) }
func (c *CPU) SetHL(value uint16) { c.H, c.L = uint16ToHiLo(value) }

func (gb *Gameboy) readMemory(addr uint16) uint8 {
	value := gb.Memory[addr]
	gb.CPU.PC++
	return value
}

func (gb *Gameboy) writeMemory(addr uint16, value uint8) {
	gb.Memory[addr] = value
	gb.CPU.PC++
	return
}

func (gb *Gameboy) Fetch() OPCode {
	code := gb.readMemory(gb.CPU.PC)

	return OPCode(code)
}

func (gb *Gameboy) Execute(opCode OPCode) {
	switch opCode {
	case 0x40:
		// LD B, B
		B := gb.CPU.B
		gb.CPU.B = B
	case 0x41:
		// LD B, C
		gb.CPU.B = gb.CPU.C
	case 0x42:
		// LD B, D
		gb.CPU.B = gb.CPU.D
	case 0x43:
		// LD B, E
		gb.CPU.B = gb.CPU.E
	case 0x44:
		// LD B, H
		gb.CPU.B = gb.CPU.H
	case 0x45:
		// LD B, L
		gb.CPU.B = gb.CPU.L
	case 0x46:
		// LD B, [HL]
		gb.CPU.B = gb.readMemory(gb.CPU.HL())
	case 0x47:
		// LD B, A
		gb.CPU.B = gb.CPU.A
	case 0x48:
		// LD C, B
		gb.CPU.C = gb.CPU.B
	case 0x49:
		// LD C, C
		C := gb.CPU.C
		gb.CPU.C = C
	case 0x4A:
		// LD C, D
		gb.CPU.C = gb.CPU.D
	case 0x4B:
		// LD C, E
		gb.CPU.C = gb.CPU.E
	case 0x4C:
		// LD C, H
		gb.CPU.C = gb.CPU.H
	case 0x4D:
		// LD C, L
		gb.CPU.C = gb.CPU.L
	case 0x4E:
		// LD C, [HL]
		gb.CPU.C = gb.readMemory(gb.CPU.HL())
	case 0x4F:
		// LD C, A
		gb.CPU.C = gb.CPU.A
	case 0x50:
		// LD D, B
		gb.CPU.D = gb.CPU.B
	case 0x51:
		// LD D, C
		gb.CPU.D = gb.CPU.C
	case 0x52:
		// LD D, D
		D := gb.CPU.D
		gb.CPU.D = D
	case 0x53:
		// LD D, E
		gb.CPU.D = gb.CPU.E
	case 0x54:
		// LD D, H
		gb.CPU.D = gb.CPU.H
	case 0x55:
		// LD D, L
		gb.CPU.D = gb.CPU.L
	case 0x56:
		// LD D, [HL]
		gb.CPU.D = gb.readMemory(gb.CPU.HL())
	case 0x57:
		// LD D, A
		gb.CPU.D = gb.CPU.A
	case 0x58:
		// LD E, B
		gb.CPU.E = gb.CPU.B
	case 0x59:
		// LD E, C
		gb.CPU.E = gb.CPU.C
	case 0x5A:
		// LD E, D
		gb.CPU.E = gb.CPU.D
	case 0x5B:
		// LD E, E
		E := gb.CPU.E
		gb.CPU.E = E
	case 0x5C:
		// LD E, H
		gb.CPU.E = gb.CPU.H
	case 0x5D:
		// LD E, L
		gb.CPU.E = gb.CPU.L
	case 0x5E:
		// LD E, [HL]
		gb.CPU.E = gb.readMemory(gb.CPU.HL())
	case 0x5F:
		// LD E, A
		gb.CPU.E = gb.CPU.A
	case 0x60:
		// LD H, B
		gb.CPU.H = gb.CPU.B
	case 0x61:
		// LD H, C
		gb.CPU.H = gb.CPU.C
	case 0x62:
		// LD H, D
		gb.CPU.H = gb.CPU.D
	case 0x63:
		// LD H, E
		gb.CPU.H = gb.CPU.E
	case 0x64:
		// LD H, H
		H := gb.CPU.H
		gb.CPU.H = H
	case 0x65:
		// LD H, L
		gb.CPU.H = gb.CPU.L
	case 0x66:
		// LD H, [HL]
		gb.CPU.H = gb.readMemory(gb.CPU.HL())
	case 0x67:
		// LD H, A
		gb.CPU.H = gb.CPU.A
	case 0x68:
		// LD L, B
		gb.CPU.L = gb.CPU.B
	case 0x69:
		// LD L, C
		gb.CPU.L = gb.CPU.C
	case 0x6A:
		// LD L, D
		gb.CPU.L = gb.CPU.D
	case 0x6B:
		// LD L, E
		gb.CPU.L = gb.CPU.E
	case 0x6C:
		// LD L, H
		gb.CPU.L = gb.CPU.H
	case 0x6D:
		// LD L, L
		L := gb.CPU.L
		gb.CPU.L = L
	case 0x6E:
		// LD L, [HL]
		gb.CPU.L = gb.readMemory(gb.CPU.HL())
	case 0x6F:
		// LD L, A
		gb.CPU.L = gb.CPU.A
	case 0x70:
		// LD [HL], B
		gb.writeMemory(gb.CPU.HL(), gb.CPU.B)
	case 0x71:
		// LD [HL], C
		gb.writeMemory(gb.CPU.HL(), gb.CPU.C)
	case 0x72:
		// LD [HL], D
		gb.writeMemory(gb.CPU.HL(), gb.CPU.D)
	case 0x73:
		// LD [HL], E
		gb.writeMemory(gb.CPU.HL(), gb.CPU.E)
	case 0x74:
		// LD [HL], H
		gb.writeMemory(gb.CPU.HL(), gb.CPU.H)
	case 0x75:
		// LD [HL], L
		gb.writeMemory(gb.CPU.HL(), gb.CPU.L)
	case 0x77:
		// LD [HL], A
		gb.writeMemory(gb.CPU.HL(), gb.CPU.A)
	case 0x78:
		// LD A, B
		gb.CPU.A = gb.CPU.B
	case 0x79:
		// LD A, C
		gb.CPU.A = gb.CPU.C
	case 0x7A:
		// LD A, D
		gb.CPU.A = gb.CPU.D
	case 0x7B:
		// LD A, E
		gb.CPU.A = gb.CPU.E
	case 0x7C:
		// LD A, H
		gb.CPU.A = gb.CPU.H
	case 0x7D:
		// LD A, L
		gb.CPU.A = gb.CPU.L
	case 0x7E:
		// LD A, [HL]
		gb.CPU.A = gb.readMemory(gb.CPU.HL())
	case 0x7F:
		// LD A, A
		A := gb.CPU.A
		gb.CPU.A = A
	case 0x02:
		// LD [BC], A
		gb.writeMemory(gb.CPU.BC(), gb.CPU.A)
	case 0x12:
		// LD [DE], A
		gb.writeMemory(gb.CPU.DE(), gb.CPU.A)
	case 0x22:
		// LD [HL+], A
		gb.writeMemory(gb.CPU.HL(), gb.CPU.A)
		gb.CPU.SetHL(gb.CPU.HL() + 1)
	case 0x32:
		// LD [HL-], A
		gb.writeMemory(gb.CPU.HL(), gb.CPU.A)
		gb.CPU.SetHL(gb.CPU.HL() - 1)
	case 0x06:
		// LD B, n8
		gb.CPU.B = gb.readMemory(gb.CPU.PC)
	case 0x16:
		// LD D, n8
		gb.CPU.D = gb.readMemory(gb.CPU.PC)
	case 0x26:
		// LD H, n8
		gb.CPU.H = gb.readMemory(gb.CPU.PC)
	case 0x36:
		// LD [HL], n8
		value := gb.readMemory(gb.CPU.PC)
		gb.writeMemory(gb.CPU.HL(), value)
	case 0x0A:
		// LD A, [BC]
		gb.CPU.A = gb.readMemory(gb.CPU.BC())
	case 0x1A:
		// LD A, [DE]
		gb.CPU.A = gb.readMemory(gb.CPU.DE())
	case 0x2A:
		// LD A, [HL+]
		gb.CPU.A = gb.readMemory(gb.CPU.HL())
		gb.CPU.SetHL(gb.CPU.HL() + 1)
	case 0x3A:
		// LD A, [HL-]
		gb.CPU.A = gb.readMemory(gb.CPU.HL())
		gb.CPU.SetHL(gb.CPU.HL() - 1)
	case 0x0E:
		// LD C, n8
		gb.CPU.C = gb.readMemory(gb.CPU.PC)
	case 0x1E:
		// LD E, n8
		gb.CPU.E = gb.readMemory(gb.CPU.PC)
	case 0x2E:
		// LD L, n8
		gb.CPU.L = gb.readMemory(gb.CPU.PC)
	case 0x3E:
		// LD A, n8
		gb.CPU.A = gb.readMemory(gb.CPU.PC)
	case 0xE0:
		// LDH [a8], A
		value := gb.readMemory(gb.CPU.PC)
		gb.writeMemory(0xFF00+uint16(value), gb.CPU.A)
	case 0xF0:
		// LDH A, [a8]
		value := gb.readMemory(gb.CPU.PC)
		gb.CPU.A = gb.readMemory(0xFF00 + uint16(value))
	case 0xE2:
		// LD [C], A
		gb.writeMemory(0xFF00+uint16(gb.CPU.C), gb.CPU.A)
	case 0xF2:
		// LD A, [C]
		gb.CPU.A = gb.readMemory(0xFF00 + uint16(gb.CPU.C))
	case 0xEA:
		// LD [a16], A
		lsb := uint16(gb.readMemory(gb.CPU.PC))
		msb := uint16(gb.readMemory(gb.CPU.PC))
		addr := lsb<<8 | msb

		gb.writeMemory(addr, gb.CPU.A)
	case 0xFA:
		// LD A, [a16]
		lsb := uint16(gb.readMemory(gb.CPU.PC))
		msb := uint16(gb.readMemory(gb.CPU.PC))
		addr := lsb<<8 | msb

		gb.CPU.A = gb.readMemory(addr)
	}
}
