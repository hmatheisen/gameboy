package main

import (
	"testing"
)

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
