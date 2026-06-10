package MOS6510

import (
	"testing"
)

func TestCPU_statusCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0 // reset S for isolated flag testing

	cpu.setStatusCarry(true)

	if !cpu.getStatusCarry() {
		t.Error("Carry flag should be set!")
	}

	if cpu.S != 1 {
		t.Errorf("Possible side effect on Carry set! got: %d, want: %d", cpu.S, 1)
	}
}

func TestCPU_statusNegative(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0

	cpu.setStatusNegative(true)

	if !cpu.getStatusNegative() {
		t.Error("Negative flag should be set!")
	}

	if cpu.S != 128 {
		t.Errorf("Possible side effect on Negative set! got: %d, want: %d", cpu.S, 128)
	}
}

func TestCPU_statusZero(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0

	cpu.setStatusZero(true)

	if !cpu.getStatusZero() {
		t.Error("Zero flag should be set!")
	}

	if cpu.S != 2 {
		t.Errorf("Possible side effect on Zero set! got: %d, want: %d", cpu.S, 2)
	}
}

func TestCPU_statusIRQ(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0

	cpu.setStatusIRQ(true)

	if !cpu.getStatusIRQ() {
		t.Error("IRQ flag should be set!")
	}

	if cpu.S != 4 {
		t.Errorf("Possible side effect on IRQ set! got: %d, want: %d", cpu.S, 4)
	}
}

func TestCPU_statusDecimal(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0

	cpu.setStatusDecimal(true)

	if !cpu.getStatusDecimal() {
		t.Error("Decimal flag should be set!")
	}

	if cpu.S != 8 {
		t.Errorf("Possible side effect on Decimal set! got: %d, want: %d", cpu.S, 8)
	}
}

func TestCPU_statusBRK(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0

	cpu.setStatusBRK(true)

	if !cpu.getStatusBRK() {
		t.Error("BRK flag should be set!")
	}

	if cpu.S != 16 {
		t.Errorf("Possible side effect on BRK set! got: %d, want: %d", cpu.S, 16)
	}
}

func TestCPU_statusOverflow(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0

	cpu.setStatusOverflow(true)

	if !cpu.getStatusOverflow() {
		t.Error("Overflow flag should be set!")
	}

	if cpu.S != 64 {
		t.Errorf("Possible side effect on Overflow set! got: %d, want: %d", cpu.S, 64)
	}
}

func TestCPU_statusCarry_clear(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0xFF

	cpu.setStatusCarry(false)

	if cpu.getStatusCarry() {
		t.Error("Carry flag should be cleared!")
	}
}

func TestCPU_statusNegative_clear(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0xFF

	cpu.setStatusNegative(false)

	if cpu.getStatusNegative() {
		t.Error("Negative flag should be cleared!")
	}
}

func TestCPU_statusZero_clear(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0xFF

	cpu.setStatusZero(false)

	if cpu.getStatusZero() {
		t.Error("Zero flag should be cleared!")
	}
}

func TestCPU_statusOverflow_clear(t *testing.T) {
	cpu := getTestCPU()
	cpu.S = 0xFF

	cpu.setStatusOverflow(false)

	if cpu.getStatusOverflow() {
		t.Error("Overflow flag should be cleared!")
	}
}
