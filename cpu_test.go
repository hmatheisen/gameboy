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
