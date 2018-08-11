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
	cpu.PC = 1000
	cpu.Memory.WriteAbsolute(1000, 100)
	cpu.Memory.WriteZeroPage(100, 100)

	// Act
	cpu.ADC(9)

	// Assert
	if cpu.A != 100 {
		t.Errorf("Wrong A register value after ADC operation, got: %d, want: %d.", cpu.A, 100)
	}

	if cpu.PC != 1001 {
		t.Errorf("Wrong PC value after ADC operation, got: %d, want: %d.", cpu.PC, 1)
	}
}

func TestCPU_ADC_zeropage_initialCarry(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 1000
	cpu.Memory.WriteAbsolute(1000, 100)
	cpu.Memory.WriteZeroPage(100, 100)
	cpu.setStatusCarry(true)

	// Act
	cpu.ADC(9)

	// Assert
	if cpu.A != 101 {
		t.Errorf("Wrong A register value after ADC operation, got: %d, want: %d.", cpu.A, 101)
	}

	if cpu.getStatusCarry() {
		t.Error("Carry flag should be cleared")
	}

	if cpu.PC != 1001 {
		t.Errorf("Wrong PC value after ADC operation, got: %d, want: %d.", cpu.PC, 1001)
	}

	if cpu.getStatusZero() {
		t.Errorf("Wrong Zero status flag value after ADC operation, got: %t, want: %t.", cpu.getStatusZero(), false)
	}
}

func TestCPU_ADC_zeropage_carryOver(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.A = 255
	cpu.PC = 1000
	cpu.Memory.WriteAbsolute(1000, 100)
	cpu.Memory.WriteZeroPage(100, 1)
	cpu.setStatusCarry(false)

	// Act
	cpu.ADC(9)

	// Assert
	if cpu.A != 0 {
		t.Errorf("Wrong A register value after ADC operation, got: %d, want: %d.", cpu.A, 0)
	}

	if !cpu.getStatusCarry() {
		t.Error("Carry flag should be set")
	}

	if !cpu.getStatusZero() {
		t.Error("Zero flag should be set")
	}

	if cpu.PC != 1001 {
		t.Errorf("Wrong PC value after ADC operation, got: %d, want: %d.", cpu.PC, 1)
	}
}

func TestCPU_ADC_zeropage_overflow_set(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.A = 127
	cpu.PC = 1000
	cpu.Memory.WriteAbsolute(1000, 100)
	cpu.Memory.WriteZeroPage(100, 1)
	cpu.setStatusCarry(false)

	// Act
	cpu.ADC(9)

	// Assert
	if cpu.A != 128 {
		t.Errorf("Wrong A register value after ADC operation, got: %d, want: %d.", cpu.A, 128)
	}

	if !cpu.getStatusOverflow() {
		t.Error("Overflow flag should be set")
	}
}

func TestCPU_BCS_carrySet(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 0
	cpu.setStatusCarry(true)
	cpu.Memory.WriteZeroPage(0, 16)

	// Act
	cpu.BCS(1)

	// Assert
	if cpu.PC != 17 {
		t.Errorf("Wrong PC value after BCS operation, got: %d, want: %d.", cpu.PC, 17)
	}
}

func TestCPU_BCS_carryUnset(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 0
	cpu.Memory.WriteZeroPage(0, 16)

	// Act
	cpu.BCS(1)

	// Assert
	if cpu.PC != 1 {
		t.Errorf("Wrong PC value after BCS operation, got: %d, want: %d.", cpu.PC, 1)
	}
}

func TestCPU_BPL_negativeSet(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 0
	cpu.setStatusNegative(false)
	cpu.Memory.WriteZeroPage(0, 16)

	// Act
	cpu.BPL(1)

	// Assert
	if cpu.PC != 17 {
		t.Errorf("Wrong PC value after BPL operation, got: %d, want: %d.", cpu.PC, 17)
	}
}

func TestCPU_BPL_negativeUnset(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 0
	cpu.setStatusNegative(true)
	cpu.Memory.WriteZeroPage(0, 16)

	// Act
	cpu.BPL(1)

	// Assert
	if cpu.PC != 1 {
		t.Errorf("Wrong PC value after BPL operation, got: %d, want: %d.", cpu.PC, 1)
	}
}

func TestCPU_ASL_carry(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 0
	cpu.A = 128

	// Act
	cpu.ASL(12)

	// Assert
	if !cpu.getStatusCarry() {
		t.Error("Status carry flag should be set!")
	}

	if cpu.getStatusNegative() {
		t.Error("Status Negative flag should be set!")
	}

	if cpu.PC != 0 {
		t.Errorf("Wrong PC value after ASL operation, got: %d, want: %d.", cpu.PC, 0)
	}
}

func TestCPU_ASL_negative(t *testing.T) {
	// Arrange
	cpu := getTestCPU()
	cpu.PC = 0
	cpu.A = 127

	// Act
	cpu.ASL(12)

	// Assert
	if cpu.getStatusCarry() {
		t.Error("Status carry flag should be clear!")
	}

	if !cpu.getStatusNegative() {
		t.Error("Status Negative flag should be set!")
	}

	if cpu.PC != 0 {
		t.Errorf("Wrong PC value after ASL operation, got: %d, want: %d.", cpu.PC, 0)
	}
}

func TestCPU_JSR(t *testing.T) {
	// TODO
}

func TestCPU_BNE(t *testing.T) {
	// TODO
}

func TestCPU_BEQ(t *testing.T) {
	// TODO
}

func TestCPU_CPX(t *testing.T) {
	// TODO
}

func TestCPU_DEX(t *testing.T) {
	// TODO
}

func TestCPU_DEY(t *testing.T) {
	// TODO
}

func TestCPU_INX(t *testing.T) {
	// TODO
}

func TestCPU_INY(t *testing.T) {
	// TODO
}

func TestCPU_LDA(t *testing.T) {
	// TODO
}

func TestCPU_LDX(t *testing.T) {
	// TODO
}

func TestCPU_LDY(t *testing.T) {
	// TODO
}

func TestCPU_STA(t *testing.T) {
	// TODO
}

func TestCPU_STX(t *testing.T) {
	// TODO
}

func TestCPU_STY(t *testing.T) {
	// TODO
}

func TestCPU_TAX(t *testing.T) {
	// TODO
}

func TestCPU_TXA(t *testing.T) {
	// TODO
}

func TestCPU_TAY(t *testing.T) {
	// TODO
}

func TestCPU_TYA(t *testing.T) {
	// TODO
}

func TestCPU_AND(t *testing.T) {
	// TODO
}

func TestCPU_INC(t *testing.T) {
	// TODO
}

func TestCPU_JMP(t *testing.T) {
	// TODO
}

func TestCPU_CMP(t *testing.T) {
	// TODO
}

func TestCPU_SEC(t *testing.T) {
	// TODO
}

func TestCPU_SBC(t *testing.T) {
	// TODO
}

func TestCPU_ROR(t *testing.T) {
	// TODO
}

func TestCPU_EOR(t *testing.T) {
	// TODO
}

func TestCPU_CPY(t *testing.T) {
	// TODO
}

func TestCPU_BCC(t *testing.T) {
	// TODO
}

func TestCPU_BMI(t *testing.T) {
	// TODO
}

func TestCPU_CLC(t *testing.T) {
	// TODO
}

func TestCPU_PHA(t *testing.T) {
	// TODO
}

func TestCPU_PLA(t *testing.T) {
	// TODO
}

func TestCPU_TSX(t *testing.T) {
	// TODO
}
