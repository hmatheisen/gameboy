package main

import (
	"math/bits"
)

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

// Getters for registers as 16 bits pair
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

// Flags in F register
func (c *CPU) setFlag(on bool, pos int) {
	if on {
		c.F |= (1 << pos)
	} else {
		c.F &= ^(1 << pos)
	}
}

func (c *CPU) getFlag(pos int) bool {
	return c.F>>pos&1 == 1
}

func (c *CPU) SetZFlag(on bool) { c.setFlag(on, 3) }
func (c *CPU) SetNFlag(on bool) { c.setFlag(on, 2) }
func (c *CPU) SetHFlag(on bool) { c.setFlag(on, 1) }
func (c *CPU) SetCFlag(on bool) { c.setFlag(on, 0) }

func (c *CPU) ZFlag() bool { return c.getFlag(3) }
func (c *CPU) NFlag() bool { return c.getFlag(2) }
func (c *CPU) HFlag() bool { return c.getFlag(1) }
func (c *CPU) CFlag() bool { return c.getFlag(0) }

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

func (c *CPU) inc(val uint8) uint8 {
	result := val + 1

	c.SetZFlag(result == 0)
	c.SetNFlag(false)
	c.SetHFlag((val&0xF)+(1&0xF) > 0xF)

	return result
}

func (c *CPU) dec(val uint8) uint8 {
	result := val - 1

	c.SetZFlag(result == 0)
	c.SetNFlag(false)
	c.SetHFlag(val&0x0F == 0)

	return result
}

func (c *CPU) add(val1, val2 uint8, carry uint) uint8 {
	result, carryOut := bits.Add(uint(val1), uint(val2), carry)

	c.SetZFlag(result == 0)
	c.SetNFlag(false)
	c.SetHFlag((val1&0xF)+(val2&0xF)+uint8(carryOut) > 0xF)
	c.SetCFlag(result > 0xFF)

	return uint8(result)
}

func (c *CPU) sub(val1, val2 uint8, borrow uint) uint8 {
	result, borrowOut := bits.Sub(uint(val1), uint(val2), borrow)

	c.SetZFlag(result == 0)
	c.SetNFlag(true)
	c.SetHFlag((val1&0xF)-(val2&0xF)-uint8(borrowOut) < 0)
	c.SetCFlag(result < 0)

	return uint8(result)
}

func (c *CPU) and(val1, val2 uint8) uint8 {
	result := val1 & val2

	c.SetZFlag(result == 0)
	c.SetNFlag(false)
	c.SetHFlag(true)
	c.SetCFlag(false)

	return result
}

func (c *CPU) xor(val1, val2 uint8) uint8 {
	result := val1 ^ val2

	c.SetZFlag(result == 0)
	c.SetNFlag(false)
	c.SetHFlag(false)
	c.SetCFlag(false)

	return result
}

func (c *CPU) or(val1, val2 uint8) uint8 {
	result := val1 | val2

	c.SetZFlag(result == 0)
	c.SetNFlag(false)
	c.SetHFlag(false)
	c.SetCFlag(false)

	return result
}

func (c *CPU) cp(val1, val2 uint8) {
	result := val2 - val1

	c.SetZFlag(result == 0)
	c.SetNFlag(true)
	c.SetHFlag((val1 & 0x0F) > (val2 & 0x0F))
	c.SetCFlag(val1 > val2)
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
	case 0x04:
		// INC B
		gb.CPU.B = gb.CPU.inc(gb.CPU.B)
	case 0x14:
		// INC D
		gb.CPU.D = gb.CPU.inc(gb.CPU.D)
	case 0x24:
		// INC H
		gb.CPU.H = gb.CPU.inc(gb.CPU.H)
	case 0x34:
		// INC [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.writeMemory(gb.CPU.HL(), gb.CPU.inc(data))
	case 0x05:
		// DEC B
		gb.CPU.B = gb.CPU.dec(gb.CPU.B)
	case 0x15:
		// DEC D
		gb.CPU.D = gb.CPU.dec(gb.CPU.D)
	case 0x25:
		// DEC H
		gb.CPU.H = gb.CPU.dec(gb.CPU.H)
	case 0x35:
		// DEC [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.writeMemory(gb.CPU.HL(), gb.CPU.dec(data))
	case 0x27:
		// DAA
		if !gb.CPU.NFlag() {
			if gb.CPU.CFlag() || gb.CPU.A > 0x99 {
				gb.CPU.A = gb.CPU.A + 0x60
				gb.CPU.SetCFlag(true)
			}
			if gb.CPU.HFlag() || gb.CPU.A&0xF > 0x9 {
				gb.CPU.A = gb.CPU.A + 0x06
				gb.CPU.SetHFlag(false)
			}
		} else if gb.CPU.CFlag() && gb.CPU.HFlag() {
			gb.CPU.A = gb.CPU.A + 0x9A
			gb.CPU.SetHFlag(false)
		} else if gb.CPU.CFlag() {
			gb.CPU.A = gb.CPU.A + 0xA0
		} else if gb.CPU.HFlag() {
			gb.CPU.A = gb.CPU.A + 0xFA
			gb.CPU.SetHFlag(false)
		}
		gb.CPU.SetZFlag(gb.CPU.A == 0)
	case 0x37:
		// SCF
		gb.CPU.SetNFlag(false)
		gb.CPU.SetHFlag(false)
		gb.CPU.SetCFlag(true)
	case 0x0C:
		// INC C
		gb.CPU.C = gb.CPU.inc(gb.CPU.C)
	case 0x1C:
		// INC E
		gb.CPU.E = gb.CPU.inc(gb.CPU.E)
	case 0x2C:
		// INC L
		gb.CPU.L = gb.CPU.inc(gb.CPU.L)
	case 0x3C:
		// INC A
		gb.CPU.A = gb.CPU.inc(gb.CPU.A)
	case 0x0D:
		// DEC C
		gb.CPU.C = gb.CPU.dec(gb.CPU.C)
	case 0x1D:
		// DEC E
		gb.CPU.E = gb.CPU.dec(gb.CPU.E)
	case 0x2D:
		// DEC L
		gb.CPU.L = gb.CPU.dec(gb.CPU.L)
	case 0x3D:
		// DEC A
		gb.CPU.A = gb.CPU.dec(gb.CPU.A)
	case 0x2F:
		// CPL
		gb.CPU.A = 0xFF ^ gb.CPU.A
		gb.CPU.SetNFlag(true)
		gb.CPU.SetHFlag(true)
	case 0x3F:
		// CCF
		gb.CPU.SetNFlag(false)
		gb.CPU.SetHFlag(false)
		gb.CPU.SetCFlag(!gb.CPU.CFlag())
	case 0x80:
		// ADD A, B
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.B, 0)
	case 0x81:
		// ADD A, C
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.C, 0)
	case 0x82:
		// ADD A, D
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.D, 0)
	case 0x83:
		// ADD A, E
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.E, 0)
	case 0x84:
		// ADD A, H
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.H, 0)
	case 0x85:
		// ADD A, L
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.L, 0)
	case 0x86:
		// ADD A, [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.CPU.A = gb.CPU.add(gb.CPU.A, data, 0)
	case 0x87:
		// ADD A, A
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.A, 0)
	case 0x88:
		// ADC A, B
		gb.CPU.A = gb.CPU.add(gb.CPU.B, gb.CPU.L, 1)
	case 0x89:
		// ADC A, C
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.C, 1)
	case 0x8A:
		// ADC A, D
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.D, 1)
	case 0x8B:
		// ADC A, E
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.E, 1)
	case 0x8C:
		// ADC A, H
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.H, 1)
	case 0x8D:
		// ADC A, L
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.L, 1)
	case 0x8E:
		// ADC A, [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.CPU.A = gb.CPU.add(gb.CPU.A, data, 1)
	case 0x8F:
		// ADC A, A
		gb.CPU.A = gb.CPU.add(gb.CPU.A, gb.CPU.L, 1)
	case 0x90:
		// SUB A, B
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.B, 0)
	case 0x91:
		// SUB A, C
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.C, 0)
	case 0x92:
		// SUB A, D
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.D, 0)
	case 0x93:
		// SUB A, E
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.E, 0)
	case 0x94:
		// SUB A, H
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.H, 0)
	case 0x95:
		// SUB A, L
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.B, 0)
	case 0x96:
		// SUB A, [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, data, 0)
	case 0x97:
		// SUB A, A
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.A, 0)
	case 0x98:
		// SBC A, B
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.B, 1)
	case 0x99:
		// SBC A, C
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.C, 1)
	case 0x9A:
		// SBC A, D
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.D, 1)
	case 0x9B:
		// SBC A, E
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.E, 1)
	case 0x9C:
		// SBC A, H
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.H, 1)
	case 0x9D:
		// SBC A, L
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.L, 1)
	case 0x9E:
		// SBC A, [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, data, 1)
	case 0x9F:
		// SBC A, A
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, gb.CPU.A, 1)
	case 0xA0:
		// AND A, B
		gb.CPU.A = gb.CPU.and(gb.CPU.A, gb.CPU.B)
	case 0xA1:
		// AND A, C
		gb.CPU.A = gb.CPU.and(gb.CPU.A, gb.CPU.C)
	case 0xA2:
		// AND A, D
		gb.CPU.A = gb.CPU.and(gb.CPU.A, gb.CPU.D)
	case 0xA3:
		// AND A, E
		gb.CPU.A = gb.CPU.and(gb.CPU.A, gb.CPU.E)
	case 0xA4:
		// AND A, H
		gb.CPU.A = gb.CPU.and(gb.CPU.A, gb.CPU.H)
	case 0xA5:
		// AND A, L
		gb.CPU.A = gb.CPU.and(gb.CPU.A, gb.CPU.L)
	case 0xA6:
		// AND A, [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.CPU.A = gb.CPU.and(gb.CPU.A, data)
	case 0xA7:
		// AND A, A
		gb.CPU.A = gb.CPU.and(gb.CPU.A, gb.CPU.A)
	case 0xA8:
		// XOR A, B
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, gb.CPU.B)
	case 0xA9:
		// XOR A, C
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, gb.CPU.C)
	case 0xAA:
		// XOR A, D
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, gb.CPU.D)
	case 0xAB:
		// XOR A, E
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, gb.CPU.E)
	case 0xAC:
		// XOR A, H
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, gb.CPU.H)
	case 0xAD:
		// XOR A, L
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, gb.CPU.L)
	case 0xAE:
		// XOR A, [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, data)
	case 0xAF:
		// XOR A, A
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, gb.CPU.A)
	case 0xB0:
		// OR A, B
		gb.CPU.A = gb.CPU.or(gb.CPU.A, gb.CPU.B)
	case 0xB1:
		// OR A, C
		gb.CPU.A = gb.CPU.or(gb.CPU.A, gb.CPU.C)
	case 0xB2:
		// OR A, D
		gb.CPU.A = gb.CPU.or(gb.CPU.A, gb.CPU.D)
	case 0xB3:
		// OR A, E
		gb.CPU.A = gb.CPU.or(gb.CPU.A, gb.CPU.E)
	case 0xB4:
		// OR A, H
		gb.CPU.A = gb.CPU.or(gb.CPU.A, gb.CPU.H)
	case 0xB5:
		// OR A, L
		gb.CPU.A = gb.CPU.or(gb.CPU.A, gb.CPU.L)
	case 0xB6:
		// OR A, [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.CPU.A = gb.CPU.or(gb.CPU.A, data)
	case 0xB7:
		// OR A, A
		gb.CPU.A = gb.CPU.or(gb.CPU.A, gb.CPU.A)
	case 0xB8:
		// CP A, B
		gb.CPU.cp(gb.CPU.A, gb.CPU.B)
	case 0xB9:
		// CP A, C
		gb.CPU.cp(gb.CPU.A, gb.CPU.C)
	case 0xBA:
		// CP A, D
		gb.CPU.cp(gb.CPU.A, gb.CPU.D)
	case 0xBB:
		// CP A, E
		gb.CPU.cp(gb.CPU.A, gb.CPU.E)
	case 0xBC:
		// CP A, H
		gb.CPU.cp(gb.CPU.A, gb.CPU.H)
	case 0xBD:
		// CP A, L
		gb.CPU.cp(gb.CPU.A, gb.CPU.L)
	case 0xBE:
		// CP A, [HL]
		data := gb.readMemory(gb.CPU.HL())
		gb.CPU.cp(gb.CPU.A, data)
	case 0xBF:
		// CP A, A
		gb.CPU.cp(gb.CPU.A, gb.CPU.A)
	case 0xC6:
		// ADD A, n8
		n := gb.readMemory(gb.CPU.PC)
		gb.CPU.A = gb.CPU.add(gb.CPU.A, n, 0)
	case 0xD6:
		// SUB A, n8
		n := gb.readMemory(gb.CPU.PC)
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, n, 0)
	case 0xE6:
		// AND A, n8
		n := gb.readMemory(gb.CPU.PC)
		gb.CPU.A = gb.CPU.and(gb.CPU.A, n)
	case 0xF6:
		// OR A, n8
		n := gb.readMemory(gb.CPU.PC)
		gb.CPU.A = gb.CPU.or(gb.CPU.A, n)
	case 0xCE:
		// ADC A, n8
		n := gb.readMemory(gb.CPU.PC)
		gb.CPU.A = gb.CPU.add(gb.CPU.A, n, 1)
	case 0xDE:
		// SBC A, n8
		n := gb.readMemory(gb.CPU.PC)
		gb.CPU.A = gb.CPU.sub(gb.CPU.A, n, 1)
	case 0xEE:
		// XOR A, n8
		n := gb.readMemory(gb.CPU.PC)
		gb.CPU.A = gb.CPU.xor(gb.CPU.A, n)
	case 0xFE:
		// CP A, n8
		n := gb.readMemory(gb.CPU.PC)
		gb.CPU.cp(gb.CPU.A, n)
	}
}
