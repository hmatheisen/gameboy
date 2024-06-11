package main

import "math/bits"

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

func (c *CPU) rl(val uint8, carry bool) uint8 {
	result := val<<1 | uint8(c.F&0x10)>>4

	c.SetZFlag(result == 0)
	c.SetNFlag(false)
	c.SetHFlag(false)
	c.SetCFlag(val&0x80 == 0x80)

	return result
}

func (c *CPU) rr(val uint8, carry bool) uint8 {
	result := val>>1 | uint8(c.F&0x10)<<3

	c.SetZFlag(result == 0)
	c.SetNFlag(false)
	c.SetHFlag(false)
	c.SetCFlag(val&0x01 == 0x01)

	return result
}
