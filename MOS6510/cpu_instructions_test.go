package MOS6510

import (
	"testing"
)

func TestCPU_RTS(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.stackPush(0)
	cpu.stackPush(100)

	// Act
	cpu.RTS(1)

	// Assert PC == fromS + 1
	if cpu.PC != 101 {
		t.Errorf("Wrong PC value after RTS operation, got: %d, want: %d.", cpu.PC, 101)
	}
}

// Add memory to accumulator with carry
func TestCPU_ADC_zeropage_noCarry(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 0
	cpu.memory.WriteZeroPage(0, 100)

	// Act
	cpu.ADC(9)

	// Assert
	if cpu.A != 100 {
		t.Errorf("Wrong A register value after ADC operation, got: %d, want: %d.", cpu.A, 100)
	}

	if cpu.PC != 1 {
		t.Errorf("Wrong PC value after ADC operation, got: %d, want: %d.", cpu.PC, 1)
	}
}

func TestCPU_ADC_zeropage_initialCarry(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 0
	cpu.setStatusCarry(true)
	cpu.memory.WriteZeroPage(0, 100)

	// Act
	cpu.ADC(9)

	// Assert
	if cpu.A != 101 {
		t.Errorf("Wrong A register value after ADC operation, got: %d, want: %d.", cpu.A, 101)
	}

	if cpu.getStatusCarry() {
		t.Error("Carry flag should be cleared")
	}

	if cpu.PC != 1 {
		t.Errorf("Wrong PC value after ADC operation, got: %d, want: %d.", cpu.PC, 1)
	}
}
