package MOS6510

import (
	"testing"
)

func TestCPU_statusCarry(t *testing.T) {
	cpu := getTestCPU()

	cpu.setStatusCarry(true)

	if !cpu.getStatusCarry() {
		t.Error("Carry flag should be set!")
	}

	if cpu.S != 1 {
		t.Error("Possible side effect on Carry set!")
	}
}

func TestCPU_statusNegative(t *testing.T) {
	cpu := getTestCPU()

	cpu.setStatusNegative(true)

	if !cpu.getStatusNegative() {
		t.Error("Negative flag should be set!")
	}

	if cpu.S != 128 {
		t.Error("Possible side effect on Negative set!")
	}
}

func TestCPU_statusZero(t *testing.T) {
	cpu := getTestCPU()

	cpu.setStatusZero(true)

	if !cpu.getStatusZero() {
		t.Error("Zero flag should be set!")
	}

	if cpu.S != 2 {
		t.Error("Possible side effect on Zero set!")
	}
}

func TestCPU_statusIRQ(t *testing.T) {
	cpu := getTestCPU()

	cpu.setStatusIRQ(true)

	if !cpu.getStatusIRQ() {
		t.Error("IRQ flag should be set!")
	}

	if cpu.S != 4 {
		t.Error("Possible side effect on IRQ set!")
	}
}

func TestCPU_statusDecimal(t *testing.T) {
	cpu := getTestCPU()

	cpu.setStatusDecimal(true)

	if !cpu.getStatusDecimal() {
		t.Error("Decimal flag should be set!")
	}

	if cpu.S != 8 {
		t.Error("Possible side effect on Decimal set!")
	}
}

func TestCPU_statusBRK(t *testing.T) {
	cpu := getTestCPU()

	cpu.setStatusBRK(true)

	if !cpu.getStatusBRK() {
		t.Error("BRK flag should be set!")
	}

	if cpu.S != 16 {
		t.Error("Possible side effect on BRK set!")
	}
}

func TestCPU_statusOverflow(t *testing.T) {
	cpu := getTestCPU()

	cpu.setStatusOverflow(true)

	if !cpu.getStatusOverflow() {
		t.Error("Overflow flag should be set!")
	}

	if cpu.S != 64 {
		t.Error("Possible side effect on Overflow set!")
	}
}
