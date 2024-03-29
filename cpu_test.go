package main

import (
	"testing"
)

func TestGetters(t *testing.T) {
	cpu := NewCPU()

	cpu.A, cpu.F = 0x00, 0x11
	cpu.B, cpu.C = 0x11, 0x00
	cpu.D, cpu.E = 0x10, 0x01
	cpu.H, cpu.L = 0x01, 0x10

	if cpu.AF() != 0x0011 {
		t.Errorf("want: AF = 0x0011; got AF = %x", cpu.AF())
	}
	if cpu.BC() != 0x1100 {
		t.Errorf("want: BC = 0x1100; got BC = %x", cpu.BC())
	}
	if cpu.DE() != 0x1001 {
		t.Errorf("want: DE = 0x1001; got DE = %x", cpu.DE())
	}
	if cpu.HL() != 0x0110 {
		t.Errorf("want: HL = 0x0110; got HL = %x", cpu.HL())
	}
}

func TestSetters(t *testing.T) {
	cpu := NewCPU()

	cpu.SetAF(0x0011)
	cpu.SetBC(0x1100)
	cpu.SetDE(0x1001)
	cpu.SetHL(0x0110)

	if cpu.A != 0x00 {
		t.Errorf("want: A = 0x00; got A = %x", cpu.A)
	}
	if cpu.F != 0x11 {
		t.Errorf("want: F = 0x11; got F = %x", cpu.F)
	}
	if cpu.B != 0x11 {
		t.Errorf("want: B = 0x11; got B = %x", cpu.B)
	}
	if cpu.C != 0x00 {
		t.Errorf("want: C = 0x00; got C = %x", cpu.C)
	}
	if cpu.D != 0x10 {
		t.Errorf("want: D = 0x10; got D = %x", cpu.D)
	}
	if cpu.E != 0x01 {
		t.Errorf("want: E = 0x01; got E = %x", cpu.E)
	}
	if cpu.H != 0x01 {
		t.Errorf("want: H = 0x01; got H = %x", cpu.H)
	}
	if cpu.L != 0x10 {
		t.Errorf("want: L = 0x10; got L = %x", cpu.L)
	}
}

func TestFlagGetters(t *testing.T) {
	cpu := NewCPU()

	cpu.F = 0b00000000
	if cpu.ZFlag() {
		t.Errorf("want: Z = false; got Z = %t", cpu.ZFlag())
	}

	cpu.F = 0b00001000
	if !cpu.ZFlag() {
		t.Errorf("want: Z = true; got Z = %t", cpu.ZFlag())
	}

	cpu.F = 0b00000100
	if !cpu.NFlag() {
		t.Errorf("want: N = true; got N = %t", cpu.NFlag())
	}

	cpu.F = 0b00000010
	if !cpu.HFlag() {
		t.Errorf("want: F = true; got F = %t", cpu.HFlag())
	}

	cpu.F = 0b00000001
	if !cpu.CFlag() {
		t.Errorf("want: C = true; got C = %t", cpu.CFlag())
	}
}

func TestFlagSetters(t *testing.T) {
	cpu := NewCPU()
	cpu.F = 0b00000000

	cpu.SetZFlag(true)
	if cpu.F != 0b00001000 {
		t.Errorf("want F = 0b00001000; got F = %08b", cpu.F)
	}
	cpu.SetZFlag(false)

	cpu.SetNFlag(true)
	if cpu.F != 0b00000100 {
		t.Errorf("want F = 0b00001000; got F = %08b", cpu.F)
	}
	cpu.SetNFlag(false)

	cpu.SetHFlag(true)
	if cpu.F != 0b00000010 {
		t.Errorf("want F = 0b00001000; got F = %08b", cpu.F)
	}
	cpu.SetHFlag(false)

	cpu.SetCFlag(true)
	if cpu.F != 0b00000001 {
		t.Errorf("want F = 0b00001000; got F = %08b", cpu.F)
	}
	cpu.SetCFlag(false)

	if cpu.F != 0b00000000 {
		t.Errorf("want F = 0b00000000; got F = %08b", cpu.F)
	}

	cpu.F = 0b00000000
	cpu.SetZFlag(false)
	if cpu.F != 0b00000000 {
		t.Errorf("want F = 0b00000000; got F = %08b", cpu.F)
	}
}

func TestInc(t *testing.T) {
	cpu := NewCPU()
	cpu.B = 12

	if cpu.inc(cpu.B) != 13 {
		t.Errorf("want: B = 13; got B = %d", cpu.B)
	}

	if cpu.ZFlag() != false {
		t.Errorf("want: Z = false; got Z = %t", cpu.ZFlag())
	}
	if cpu.NFlag() != false {
		t.Errorf("want: N = false; got N = %t", cpu.NFlag())
	}
	if cpu.HFlag() != false {
		t.Errorf("want: H = false; got H = %t", cpu.HFlag())
	}
}

func TestDec(t *testing.T) {
	cpu := NewCPU()
	cpu.B = 12

	if cpu.dec(cpu.B) != 11 {
		t.Errorf("want: B = 13; got B = %d", cpu.B)
	}

	if cpu.ZFlag() != false {
		t.Errorf("want: Z = false; got Z = %t", cpu.ZFlag())
	}
	if cpu.NFlag() != false {
		t.Errorf("want: N = false; got N = %t", cpu.NFlag())
	}
	if cpu.HFlag() != false {
		t.Errorf("want: H = false; got H = %t", cpu.HFlag())
	}
}

func TestAdd(t *testing.T) {
	cpu := NewCPU()
	cpu.A = 10
	cpu.B = 12

	result := cpu.add(cpu.A, cpu.B, 0)

	if result != 22 {
		t.Errorf("want: A+B = 13; got A+B = %d", result)
	}

	if cpu.ZFlag() != false {
		t.Errorf("want: Z = false; got Z = %t", cpu.ZFlag())
	}
	if cpu.NFlag() != false {
		t.Errorf("want: N = false; got N = %t", cpu.NFlag())
	}
	if cpu.HFlag() != true {
		t.Errorf("want: H = true; got H = %t", cpu.HFlag())
	}
	if cpu.CFlag() != false {
		t.Errorf("want: C = false; got C = %t", cpu.CFlag())
	}
}

func TestSub(t *testing.T) {
	cpu := NewCPU()
	cpu.A = 12
	cpu.B = 10

	result := cpu.sub(cpu.A, cpu.B, 0)

	if result != 2 {
		t.Errorf("want: A-B = 13; got A-B = %d", result)
	}

	if cpu.ZFlag() != false {
		t.Errorf("want: Z = false; got Z = %t", cpu.ZFlag())
	}
	if cpu.NFlag() != true {
		t.Errorf("want: N = true; got N = %t", cpu.NFlag())
	}
	if cpu.HFlag() != false {
		t.Errorf("want: H = false; got H = %t", cpu.HFlag())
	}
	if cpu.CFlag() != false {
		t.Errorf("want: C = false; got C = %t", cpu.CFlag())
	}
}

func TestXOR(t *testing.T) {
	cpu := NewCPU()
	cpu.A = 10
	cpu.B = 12

	result := cpu.xor(cpu.A, cpu.B)

	if result != 6 {
		t.Errorf("want: A^B = 12; got A^B = %d", result)
	}

	if cpu.ZFlag() != false {
		t.Errorf("want: Z = false; got Z = %t", cpu.ZFlag())
	}
	if cpu.NFlag() != false {
		t.Errorf("want: N = false; got N = %t", cpu.NFlag())
	}
	if cpu.HFlag() != false {
		t.Errorf("want: H = false; got H = %t", cpu.HFlag())
	}
	if cpu.CFlag() != false {
		t.Errorf("want: C = false; got C = %t", cpu.CFlag())
	}
}

func TestOr(t *testing.T) {
	cpu := NewCPU()
	cpu.A = 10
	cpu.B = 12

	result := cpu.or(cpu.A, cpu.B)

	if result != 14 {
		t.Errorf("want: A|B = 12; got A|B = %d", result)
	}

	if cpu.ZFlag() != false {
		t.Errorf("want: Z = false; got Z = %t", cpu.ZFlag())
	}
	if cpu.NFlag() != false {
		t.Errorf("want: N = false; got N = %t", cpu.NFlag())
	}
	if cpu.HFlag() != false {
		t.Errorf("want: H = false; got H = %t", cpu.HFlag())
	}
	if cpu.CFlag() != false {
		t.Errorf("want: C = false; got C = %t", cpu.CFlag())
	}
}

func TestCP(t *testing.T) {
	cpu := NewCPU()
	cpu.A = 10
	cpu.B = 0

	cpu.cp(cpu.A, cpu.B)
	
	if cpu.ZFlag() != false {
		t.Errorf("want: Z = false; got Z = %t", cpu.ZFlag())
	}
	if cpu.NFlag() != true {
		t.Errorf("want: N = true; got N = %t", cpu.NFlag())
	}
	if cpu.HFlag() != true {
		t.Errorf("want: H = true; got H = %t", cpu.HFlag())
	}
	if cpu.CFlag() != true {
		t.Errorf("want: C = true; got C = %t", cpu.CFlag())
	}
}
