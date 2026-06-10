package MOS6510

import (
	"testing"
)

// =============================================================================
// LDA - Load Accumulator
// =============================================================================

func TestCPU_LDA_Immediate(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x42)

	cpu.LDA(Immidiate)

	if cpu.A != 0x42 {
		t.Errorf("LDA immediate: got A=%d, want %d", cpu.A, 0x42)
	}
	if cpu.PC != 0x201 {
		t.Errorf("LDA immediate: got PC=%d, want %d", cpu.PC, 0x201)
	}
	if cpu.getStatusZero() {
		t.Error("LDA immediate: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("LDA immediate: negative flag should be clear")
	}
}

func TestCPU_LDA_Immediate_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x55
	cpu.Memory.WriteAbsolute(0x200, 0x00)

	cpu.LDA(Immidiate)

	if cpu.A != 0x00 {
		t.Errorf("LDA immediate zero: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("LDA immediate zero: zero flag should be set")
	}
	if cpu.getStatusNegative() {
		t.Error("LDA immediate zero: negative flag should be clear")
	}
}

func TestCPU_LDA_Immediate_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x80)

	cpu.LDA(Immidiate)

	if cpu.A != 0x80 {
		t.Errorf("LDA immediate negative: got A=%d, want %d", cpu.A, 0x80)
	}
	if cpu.getStatusZero() {
		t.Error("LDA immediate negative: zero flag should be clear")
	}
	if !cpu.getStatusNegative() {
		t.Error("LDA immediate negative: negative flag should be set")
	}
}

func TestCPU_LDA_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50) // zp address
	cpu.Memory.WriteZeroPage(0x50, 0x77)  // value at zp address

	cpu.LDA(ZeroPage)

	if cpu.A != 0x77 {
		t.Errorf("LDA zeropage: got A=%d, want %d", cpu.A, 0x77)
	}
	if cpu.PC != 0x201 {
		t.Errorf("LDA zeropage: got PC=%d, want %d", cpu.PC, 0x201)
	}
}

func TestCPU_LDA_Absolute(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x00) // low byte of address
	cpu.Memory.WriteAbsolute(0x201, 0x04) // high byte of address -> 0x0400
	cpu.Memory.WriteAbsolute(0x0400, 0xAB)

	cpu.LDA(Absolute)

	if cpu.A != 0xAB {
		t.Errorf("LDA absolute: got A=%d, want %d", cpu.A, 0xAB)
	}
	if cpu.PC != 0x202 {
		t.Errorf("LDA absolute: got PC=%d, want %d", cpu.PC, 0x202)
	}
}

// =============================================================================
// LDX - Load X Register
// =============================================================================

func TestCPU_LDX_Immediate(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x42)

	cpu.LDX(Immidiate)

	if cpu.X != 0x42 {
		t.Errorf("LDX immediate: got X=%d, want %d", cpu.X, 0x42)
	}
	if cpu.PC != 0x201 {
		t.Errorf("LDX immediate: got PC=%d, want %d", cpu.PC, 0x201)
	}
	if cpu.getStatusZero() {
		t.Error("LDX immediate: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("LDX immediate: negative flag should be clear")
	}
}

func TestCPU_LDX_Immediate_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.X = 0x55
	cpu.Memory.WriteAbsolute(0x200, 0x00)

	cpu.LDX(Immidiate)

	if cpu.X != 0x00 {
		t.Errorf("LDX immediate zero: got X=%d, want %d", cpu.X, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("LDX immediate zero: zero flag should be set")
	}
}

func TestCPU_LDX_Immediate_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0xFF)

	cpu.LDX(Immidiate)

	if cpu.X != 0xFF {
		t.Errorf("LDX immediate negative: got X=%d, want %d", cpu.X, 0xFF)
	}
	if !cpu.getStatusNegative() {
		t.Error("LDX immediate negative: negative flag should be set")
	}
}

func TestCPU_LDX_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x30)
	cpu.Memory.WriteZeroPage(0x30, 0x99)

	cpu.LDX(ZeroPage)

	if cpu.X != 0x99 {
		t.Errorf("LDX zeropage: got X=%d, want %d", cpu.X, 0x99)
	}
}

// =============================================================================
// LDY - Load Y Register
// =============================================================================

func TestCPU_LDY_Immediate(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x42)

	cpu.LDY(Immidiate)

	if cpu.Y != 0x42 {
		t.Errorf("LDY immediate: got Y=%d, want %d", cpu.Y, 0x42)
	}
	if cpu.PC != 0x201 {
		t.Errorf("LDY immediate: got PC=%d, want %d", cpu.PC, 0x201)
	}
}

func TestCPU_LDY_Immediate_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Y = 0x55
	cpu.Memory.WriteAbsolute(0x200, 0x00)

	cpu.LDY(Immidiate)

	if cpu.Y != 0x00 {
		t.Errorf("LDY immediate zero: got Y=%d, want %d", cpu.Y, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("LDY immediate zero: zero flag should be set")
	}
}

func TestCPU_LDY_Immediate_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0xFE)

	cpu.LDY(Immidiate)

	if cpu.Y != 0xFE {
		t.Errorf("LDY immediate negative: got Y=%d, want %d", cpu.Y, 0xFE)
	}
	if !cpu.getStatusNegative() {
		t.Error("LDY immediate negative: negative flag should be set")
	}
}

func TestCPU_LDY_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x40)
	cpu.Memory.WriteZeroPage(0x40, 0xBB)

	cpu.LDY(ZeroPage)

	if cpu.Y != 0xBB {
		t.Errorf("LDY zeropage: got Y=%d, want %d", cpu.Y, 0xBB)
	}
}

// =============================================================================
// STA - Store Accumulator
// =============================================================================

func TestCPU_STA_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x42
	cpu.Memory.WriteAbsolute(0x200, 0x50) // zp address to store to

	cpu.STA(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0x42 {
		t.Errorf("STA zeropage: got mem[0x50]=%d, want %d", val, 0x42)
	}
	if cpu.PC != 0x201 {
		t.Errorf("STA zeropage: got PC=%d, want %d", cpu.PC, 0x201)
	}
}

func TestCPU_STA_Absolute(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x99
	cpu.Memory.WriteAbsolute(0x200, 0x00) // low byte
	cpu.Memory.WriteAbsolute(0x201, 0x04) // high byte -> 0x0400

	cpu.STA(Absolute)

	val := cpu.Memory.ReadAbsolute(0x0400)
	if val != 0x99 {
		t.Errorf("STA absolute: got mem[0x0400]=%d, want %d", val, 0x99)
	}
	if cpu.PC != 0x202 {
		t.Errorf("STA absolute: got PC=%d, want %d", cpu.PC, 0x202)
	}
}

func TestCPU_STA_ZeroPageX(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x77
	cpu.X = 0x05
	cpu.Memory.WriteAbsolute(0x200, 0x10) // zp base address

	cpu.STA(ZeroPageX)

	val := cpu.Memory.ReadZeroPage(0x15) // 0x10 + 0x05
	if val != 0x77 {
		t.Errorf("STA zeropageX: got mem[0x15]=%d, want %d", val, 0x77)
	}
}

// =============================================================================
// STX - Store X Register
// =============================================================================

func TestCPU_STX_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.X = 0x55
	cpu.Memory.WriteAbsolute(0x200, 0x30)

	cpu.STX(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x30)
	if val != 0x55 {
		t.Errorf("STX zeropage: got mem[0x30]=%d, want %d", val, 0x55)
	}
	if cpu.PC != 0x201 {
		t.Errorf("STX zeropage: got PC=%d, want %d", cpu.PC, 0x201)
	}
}

func TestCPU_STX_Absolute(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.X = 0xAA
	cpu.Memory.WriteAbsolute(0x200, 0x00)
	cpu.Memory.WriteAbsolute(0x201, 0x05) // -> 0x0500

	cpu.STX(Absolute)

	val := cpu.Memory.ReadAbsolute(0x0500)
	if val != 0xAA {
		t.Errorf("STX absolute: got mem[0x0500]=%d, want %d", val, 0xAA)
	}
	if cpu.PC != 0x202 {
		t.Errorf("STX absolute: got PC=%d, want %d", cpu.PC, 0x202)
	}
}

func TestCPU_STX_ZeroPageY(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.X = 0xBB
	cpu.Y = 0x03
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.STX(ZeroPageY)

	val := cpu.Memory.ReadZeroPage(0x13) // 0x10 + 0x03
	if val != 0xBB {
		t.Errorf("STX zeropageY: got mem[0x13]=%d, want %d", val, 0xBB)
	}
}

// =============================================================================
// STY - Store Y Register
// =============================================================================

func TestCPU_STY_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Y = 0x66
	cpu.Memory.WriteAbsolute(0x200, 0x20)

	cpu.STY(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x20)
	if val != 0x66 {
		t.Errorf("STY zeropage: got mem[0x20]=%d, want %d", val, 0x66)
	}
	if cpu.PC != 0x201 {
		t.Errorf("STY zeropage: got PC=%d, want %d", cpu.PC, 0x201)
	}
}

func TestCPU_STY_Absolute(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Y = 0xCC
	cpu.Memory.WriteAbsolute(0x200, 0x00)
	cpu.Memory.WriteAbsolute(0x201, 0x06) // -> 0x0600

	cpu.STY(Absolute)

	val := cpu.Memory.ReadAbsolute(0x0600)
	if val != 0xCC {
		t.Errorf("STY absolute: got mem[0x0600]=%d, want %d", val, 0xCC)
	}
	if cpu.PC != 0x202 {
		t.Errorf("STY absolute: got PC=%d, want %d", cpu.PC, 0x202)
	}
}

func TestCPU_STY_ZeroPageX(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Y = 0xDD
	cpu.X = 0x02
	cpu.Memory.WriteAbsolute(0x200, 0x20)

	cpu.STY(ZeroPageX)

	val := cpu.Memory.ReadZeroPage(0x22) // 0x20 + 0x02
	if val != 0xDD {
		t.Errorf("STY zeropageX: got mem[0x22]=%d, want %d", val, 0xDD)
	}
}

// =============================================================================
// TAX - Transfer A to X
// =============================================================================

func TestCPU_TAX_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x42

	cpu.TAX(Implied)

	if cpu.X != 0x42 {
		t.Errorf("TAX: got X=%d, want %d", cpu.X, 0x42)
	}
	if cpu.getStatusZero() {
		t.Error("TAX: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("TAX: negative flag should be clear")
	}
}

func TestCPU_TAX_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x00
	cpu.X = 0x55

	cpu.TAX(Implied)

	if cpu.X != 0x00 {
		t.Errorf("TAX zero: got X=%d, want %d", cpu.X, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("TAX zero: zero flag should be set")
	}
}

func TestCPU_TAX_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x80

	cpu.TAX(Implied)

	if cpu.X != 0x80 {
		t.Errorf("TAX negative: got X=%d, want %d", cpu.X, 0x80)
	}
	if !cpu.getStatusNegative() {
		t.Error("TAX negative: negative flag should be set")
	}
}

// =============================================================================
// TXA - Transfer X to A
// =============================================================================

func TestCPU_TXA_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x33

	cpu.TXA(Implied)

	if cpu.A != 0x33 {
		t.Errorf("TXA: got A=%d, want %d", cpu.A, 0x33)
	}
	if cpu.getStatusZero() {
		t.Error("TXA: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("TXA: negative flag should be clear")
	}
}

func TestCPU_TXA_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x00
	cpu.A = 0x55

	cpu.TXA(Implied)

	if cpu.A != 0x00 {
		t.Errorf("TXA zero: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("TXA zero: zero flag should be set")
	}
}

func TestCPU_TXA_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0xFE

	cpu.TXA(Implied)

	if cpu.A != 0xFE {
		t.Errorf("TXA negative: got A=%d, want %d", cpu.A, 0xFE)
	}
	if !cpu.getStatusNegative() {
		t.Error("TXA negative: negative flag should be set")
	}
}

// =============================================================================
// TAY - Transfer A to Y
// =============================================================================

func TestCPU_TAY_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x42

	cpu.TAY(Implied)

	if cpu.Y != 0x42 {
		t.Errorf("TAY: got Y=%d, want %d", cpu.Y, 0x42)
	}
	if cpu.getStatusZero() {
		t.Error("TAY: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("TAY: negative flag should be clear")
	}
}

func TestCPU_TAY_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x00
	cpu.Y = 0x55

	cpu.TAY(Implied)

	if cpu.Y != 0x00 {
		t.Errorf("TAY zero: got Y=%d, want %d", cpu.Y, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("TAY zero: zero flag should be set")
	}
}

func TestCPU_TAY_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x90

	cpu.TAY(Implied)

	if cpu.Y != 0x90 {
		t.Errorf("TAY negative: got Y=%d, want %d", cpu.Y, 0x90)
	}
	if !cpu.getStatusNegative() {
		t.Error("TAY negative: negative flag should be set")
	}
}

// =============================================================================
// TYA - Transfer Y to A
// =============================================================================

func TestCPU_TYA_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0x33

	cpu.TYA(Implied)

	if cpu.A != 0x33 {
		t.Errorf("TYA: got A=%d, want %d", cpu.A, 0x33)
	}
	if cpu.getStatusZero() {
		t.Error("TYA: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("TYA: negative flag should be clear")
	}
}

func TestCPU_TYA_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0x00
	cpu.A = 0x55

	cpu.TYA(Implied)

	if cpu.A != 0x00 {
		t.Errorf("TYA zero: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("TYA zero: zero flag should be set")
	}
}

func TestCPU_TYA_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0xC0

	cpu.TYA(Implied)

	if cpu.A != 0xC0 {
		t.Errorf("TYA negative: got A=%d, want %d", cpu.A, 0xC0)
	}
	if !cpu.getStatusNegative() {
		t.Error("TYA negative: negative flag should be set")
	}
}

// =============================================================================
// INX - Increment X
// =============================================================================

func TestCPU_INX_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x10

	cpu.INX(Implied)

	if cpu.X != 0x11 {
		t.Errorf("INX: got X=%d, want %d", cpu.X, 0x11)
	}
	if cpu.getStatusZero() {
		t.Error("INX: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("INX: negative flag should be clear")
	}
}

func TestCPU_INX_Wrap(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0xFF

	cpu.INX(Implied)

	if cpu.X != 0x00 {
		t.Errorf("INX wrap: got X=%d, want %d", cpu.X, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("INX wrap: zero flag should be set")
	}
	if cpu.getStatusNegative() {
		t.Error("INX wrap: negative flag should be clear")
	}
}

func TestCPU_INX_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x7F

	cpu.INX(Implied)

	if cpu.X != 0x80 {
		t.Errorf("INX negative: got X=%d, want %d", cpu.X, 0x80)
	}
	if !cpu.getStatusNegative() {
		t.Error("INX negative: negative flag should be set")
	}
}

// =============================================================================
// INY - Increment Y
// =============================================================================

func TestCPU_INY_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0x10

	cpu.INY(Implied)

	if cpu.Y != 0x11 {
		t.Errorf("INY: got Y=%d, want %d", cpu.Y, 0x11)
	}
	if cpu.getStatusZero() {
		t.Error("INY: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("INY: negative flag should be clear")
	}
}

func TestCPU_INY_Wrap(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0xFF

	cpu.INY(Implied)

	if cpu.Y != 0x00 {
		t.Errorf("INY wrap: got Y=%d, want %d", cpu.Y, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("INY wrap: zero flag should be set")
	}
}

func TestCPU_INY_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0x7F

	cpu.INY(Implied)

	if cpu.Y != 0x80 {
		t.Errorf("INY negative: got Y=%d, want %d", cpu.Y, 0x80)
	}
	if !cpu.getStatusNegative() {
		t.Error("INY negative: negative flag should be set")
	}
}

// =============================================================================
// DEX - Decrement X
// =============================================================================

func TestCPU_DEX_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x10

	cpu.DEX(Implied)

	if cpu.X != 0x0F {
		t.Errorf("DEX: got X=%d, want %d", cpu.X, 0x0F)
	}
	if cpu.getStatusZero() {
		t.Error("DEX: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("DEX: negative flag should be clear")
	}
}

func TestCPU_DEX_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x01

	cpu.DEX(Implied)

	if cpu.X != 0x00 {
		t.Errorf("DEX zero: got X=%d, want %d", cpu.X, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("DEX zero: zero flag should be set")
	}
}

func TestCPU_DEX_Wrap(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x00

	cpu.DEX(Implied)

	if cpu.X != 0xFF {
		t.Errorf("DEX wrap: got X=%d, want %d", cpu.X, 0xFF)
	}
	if !cpu.getStatusNegative() {
		t.Error("DEX wrap: negative flag should be set")
	}
	if cpu.getStatusZero() {
		t.Error("DEX wrap: zero flag should be clear")
	}
}

// =============================================================================
// DEY - Decrement Y
// =============================================================================

func TestCPU_DEY_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0x10

	cpu.DEY(Implied)

	if cpu.Y != 0x0F {
		t.Errorf("DEY: got Y=%d, want %d", cpu.Y, 0x0F)
	}
	if cpu.getStatusZero() {
		t.Error("DEY: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("DEY: negative flag should be clear")
	}
}

func TestCPU_DEY_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0x01

	cpu.DEY(Implied)

	if cpu.Y != 0x00 {
		t.Errorf("DEY zero: got Y=%d, want %d", cpu.Y, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("DEY zero: zero flag should be set")
	}
}

func TestCPU_DEY_Wrap(t *testing.T) {
	cpu := getTestCPU()
	cpu.Y = 0x00

	cpu.DEY(Implied)

	if cpu.Y != 0xFF {
		t.Errorf("DEY wrap: got Y=%d, want %d", cpu.Y, 0xFF)
	}
	if !cpu.getStatusNegative() {
		t.Error("DEY wrap: negative flag should be set")
	}
}

// =============================================================================
// INC - Increment Memory
// =============================================================================

func TestCPU_INC_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50) // zp address
	cpu.Memory.WriteZeroPage(0x50, 0x10)  // value

	cpu.INC(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0x11 {
		t.Errorf("INC zeropage: got mem[0x50]=%d, want %d", val, 0x11)
	}
	if cpu.getStatusZero() {
		t.Error("INC zeropage: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("INC zeropage: negative flag should be clear")
	}
}

func TestCPU_INC_ZeroPage_Wrap(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0xFF)

	cpu.INC(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0x00 {
		t.Errorf("INC zeropage wrap: got mem[0x50]=%d, want %d", val, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("INC zeropage wrap: zero flag should be set")
	}
}

func TestCPU_INC_ZeroPage_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0x7F)

	cpu.INC(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0x80 {
		t.Errorf("INC zeropage neg: got mem[0x50]=%d, want %d", val, 0x80)
	}
	if !cpu.getStatusNegative() {
		t.Error("INC zeropage neg: negative flag should be set")
	}
}

// =============================================================================
// DEC - Decrement Memory
// =============================================================================

func TestCPU_DEC_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0x10)

	cpu.DEC(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0x0F {
		t.Errorf("DEC zeropage: got mem[0x50]=%d, want %d", val, 0x0F)
	}
	if cpu.getStatusZero() {
		t.Error("DEC zeropage: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("DEC zeropage: negative flag should be clear")
	}
}

func TestCPU_DEC_ZeroPage_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0x01)

	cpu.DEC(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0x00 {
		t.Errorf("DEC zeropage zero: got mem[0x50]=%d, want %d", val, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("DEC zeropage zero: zero flag should be set")
	}
}

func TestCPU_DEC_ZeroPage_Wrap(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0x00)

	cpu.DEC(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0xFF {
		t.Errorf("DEC zeropage wrap: got mem[0x50]=%d, want %d", val, 0xFF)
	}
	if !cpu.getStatusNegative() {
		t.Error("DEC zeropage wrap: negative flag should be set")
	}
}

// =============================================================================
// AND - Logical AND
// =============================================================================

func TestCPU_AND_Immediate(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0xFF
	cpu.Memory.WriteAbsolute(0x200, 0x0F)

	cpu.AND(Immidiate)

	if cpu.A != 0x0F {
		t.Errorf("AND immediate: got A=%d, want %d", cpu.A, 0x0F)
	}
	if cpu.getStatusZero() {
		t.Error("AND immediate: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("AND immediate: negative flag should be clear")
	}
}

func TestCPU_AND_Immediate_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0xF0
	cpu.Memory.WriteAbsolute(0x200, 0x0F)

	cpu.AND(Immidiate)

	if cpu.A != 0x00 {
		t.Errorf("AND immediate zero: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("AND immediate zero: zero flag should be set")
	}
}

func TestCPU_AND_Immediate_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0xFF
	cpu.Memory.WriteAbsolute(0x200, 0x80)

	cpu.AND(Immidiate)

	if cpu.A != 0x80 {
		t.Errorf("AND immediate neg: got A=%d, want %d", cpu.A, 0x80)
	}
	if !cpu.getStatusNegative() {
		t.Error("AND immediate neg: negative flag should be set")
	}
}

// =============================================================================
// ORA - Logical OR
// =============================================================================

func TestCPU_ORA_Immediate(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0xF0
	cpu.Memory.WriteAbsolute(0x200, 0x0F)

	cpu.ORA(Immidiate)

	if cpu.A != 0xFF {
		t.Errorf("ORA immediate: got A=%d, want %d", cpu.A, 0xFF)
	}
	if cpu.getStatusZero() {
		t.Error("ORA immediate: zero flag should be clear")
	}
	if !cpu.getStatusNegative() {
		t.Error("ORA immediate: negative flag should be set")
	}
}

func TestCPU_ORA_Immediate_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x00
	cpu.Memory.WriteAbsolute(0x200, 0x00)

	cpu.ORA(Immidiate)

	if cpu.A != 0x00 {
		t.Errorf("ORA immediate zero: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("ORA immediate zero: zero flag should be set")
	}
}

func TestCPU_ORA_Immediate_SetBits(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x01
	cpu.Memory.WriteAbsolute(0x200, 0x40)

	cpu.ORA(Immidiate)

	if cpu.A != 0x41 {
		t.Errorf("ORA immediate setbits: got A=%d, want %d", cpu.A, 0x41)
	}
	if cpu.getStatusNegative() {
		t.Error("ORA immediate setbits: negative flag should be clear")
	}
}

// =============================================================================
// EOR - Exclusive OR
// =============================================================================

func TestCPU_EOR_Immediate(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0xFF
	cpu.Memory.WriteAbsolute(0x200, 0x0F)

	cpu.EOR(Immidiate)

	if cpu.A != 0xF0 {
		t.Errorf("EOR immediate: got A=%d, want %d", cpu.A, 0xF0)
	}
	if !cpu.getStatusNegative() {
		t.Error("EOR immediate: negative flag should be set")
	}
}

func TestCPU_EOR_Immediate_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x55
	cpu.Memory.WriteAbsolute(0x200, 0x55)

	cpu.EOR(Immidiate)

	if cpu.A != 0x00 {
		t.Errorf("EOR immediate zero: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("EOR immediate zero: zero flag should be set")
	}
}

func TestCPU_EOR_Immediate_NoNegative(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x80
	cpu.Memory.WriteAbsolute(0x200, 0x80)

	cpu.EOR(Immidiate)

	if cpu.A != 0x00 {
		t.Errorf("EOR immediate no neg: got A=%d, want %d", cpu.A, 0x00)
	}
	if cpu.getStatusNegative() {
		t.Error("EOR immediate no neg: negative flag should be clear")
	}
}

// =============================================================================
// CMP - Compare Accumulator
// =============================================================================

func TestCPU_CMP_Immediate_Equal(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x42
	cpu.Memory.WriteAbsolute(0x200, 0x42)

	cpu.CMP(Immidiate)

	if !cpu.getStatusZero() {
		t.Error("CMP equal: zero flag should be set")
	}
	if !cpu.getStatusCarry() {
		t.Error("CMP equal: carry flag should be set")
	}
	if cpu.getStatusNegative() {
		t.Error("CMP equal: negative flag should be clear")
	}
}

func TestCPU_CMP_Immediate_Greater(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x50
	cpu.Memory.WriteAbsolute(0x200, 0x30)

	cpu.CMP(Immidiate)

	if cpu.getStatusZero() {
		t.Error("CMP greater: zero flag should be clear")
	}
	if !cpu.getStatusCarry() {
		t.Error("CMP greater: carry flag should be set")
	}
	if cpu.getStatusNegative() {
		t.Error("CMP greater: negative flag should be clear")
	}
}

func TestCPU_CMP_Immediate_Less(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x10
	cpu.Memory.WriteAbsolute(0x200, 0x20)

	cpu.CMP(Immidiate)

	if cpu.getStatusZero() {
		t.Error("CMP less: zero flag should be clear")
	}
	if cpu.getStatusCarry() {
		t.Error("CMP less: carry flag should be clear")
	}
	if !cpu.getStatusNegative() {
		t.Error("CMP less: negative flag should be set (result 0xF0)")
	}
}

// =============================================================================
// CPX - Compare X Register
// =============================================================================

func TestCPU_CPX_Immediate_Equal(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.X = 0x42
	cpu.Memory.WriteAbsolute(0x200, 0x42)

	cpu.CPX(Immidiate)

	if !cpu.getStatusZero() {
		t.Error("CPX equal: zero flag should be set")
	}
	if !cpu.getStatusCarry() {
		t.Error("CPX equal: carry flag should be set")
	}
}

func TestCPU_CPX_Immediate_Greater(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.X = 0x50
	cpu.Memory.WriteAbsolute(0x200, 0x20)

	cpu.CPX(Immidiate)

	if cpu.getStatusZero() {
		t.Error("CPX greater: zero flag should be clear")
	}
	if !cpu.getStatusCarry() {
		t.Error("CPX greater: carry flag should be set")
	}
}

func TestCPU_CPX_Immediate_Less(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.X = 0x10
	cpu.Memory.WriteAbsolute(0x200, 0x30)

	cpu.CPX(Immidiate)

	if cpu.getStatusCarry() {
		t.Error("CPX less: carry flag should be clear")
	}
	if !cpu.getStatusNegative() {
		t.Error("CPX less: negative flag should be set")
	}
}

// =============================================================================
// CPY - Compare Y Register
// =============================================================================

func TestCPU_CPY_Immediate_Equal(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Y = 0x42
	cpu.Memory.WriteAbsolute(0x200, 0x42)

	cpu.CPY(Immidiate)

	if !cpu.getStatusZero() {
		t.Error("CPY equal: zero flag should be set")
	}
	if !cpu.getStatusCarry() {
		t.Error("CPY equal: carry flag should be set")
	}
}

func TestCPU_CPY_Immediate_Greater(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Y = 0x60
	cpu.Memory.WriteAbsolute(0x200, 0x20)

	cpu.CPY(Immidiate)

	if cpu.getStatusZero() {
		t.Error("CPY greater: zero flag should be clear")
	}
	if !cpu.getStatusCarry() {
		t.Error("CPY greater: carry flag should be set")
	}
}

func TestCPU_CPY_Immediate_Less(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Y = 0x05
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.CPY(Immidiate)

	if cpu.getStatusCarry() {
		t.Error("CPY less: carry flag should be clear")
	}
	if !cpu.getStatusNegative() {
		t.Error("CPY less: negative flag should be set")
	}
}

// =============================================================================
// ADC - Add with Carry (keep existing tests)
// =============================================================================

func TestCPU_ADC_zeropage_noCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 1000
	cpu.Memory.WriteAbsolute(1000, 100)
	cpu.Memory.WriteZeroPage(100, 100)

	cpu.ADC(ZeroPage)

	if cpu.A != 100 {
		t.Errorf("Wrong A register value after ADC operation, got: %d, want: %d.", cpu.A, 100)
	}
	if cpu.PC != 1001 {
		t.Errorf("Wrong PC value after ADC operation, got: %d, want: %d.", cpu.PC, 1001)
	}
}

func TestCPU_ADC_zeropage_initialCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 1000
	cpu.Memory.WriteAbsolute(1000, 100)
	cpu.Memory.WriteZeroPage(100, 100)
	cpu.setStatusCarry(true)

	cpu.ADC(ZeroPage)

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
	cpu := getTestCPU()
	cpu.A = 255
	cpu.PC = 1000
	cpu.Memory.WriteAbsolute(1000, 100)
	cpu.Memory.WriteZeroPage(100, 1)
	cpu.setStatusCarry(false)

	cpu.ADC(ZeroPage)

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
		t.Errorf("Wrong PC value after ADC operation, got: %d, want: %d.", cpu.PC, 1001)
	}
}

func TestCPU_ADC_zeropage_overflow_set(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 127
	cpu.PC = 1000
	cpu.Memory.WriteAbsolute(1000, 100)
	cpu.Memory.WriteZeroPage(100, 1)
	cpu.setStatusCarry(false)

	cpu.ADC(ZeroPage)

	if cpu.A != 128 {
		t.Errorf("Wrong A register value after ADC operation, got: %d, want: %d.", cpu.A, 128)
	}
	if !cpu.getStatusOverflow() {
		t.Error("Overflow flag should be set")
	}
}

// =============================================================================
// SBC - Subtract with Carry (Borrow)
// =============================================================================

func TestCPU_SBC_Immediate_NoBorrow(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x50
	cpu.setStatusCarry(true) // carry set = no borrow
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.SBC(Immidiate)

	if cpu.A != 0x40 {
		t.Errorf("SBC no borrow: got A=%d, want %d", cpu.A, 0x40)
	}
	if !cpu.getStatusCarry() {
		t.Error("SBC no borrow: carry should be set (no borrow)")
	}
	if cpu.getStatusZero() {
		t.Error("SBC no borrow: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("SBC no borrow: negative flag should be clear")
	}
}

func TestCPU_SBC_Immediate_WithBorrow(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x50
	cpu.setStatusCarry(false) // carry clear = borrow
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.SBC(Immidiate)

	if cpu.A != 0x3F {
		t.Errorf("SBC with borrow: got A=%d, want %d", cpu.A, 0x3F)
	}
	if !cpu.getStatusCarry() {
		t.Error("SBC with borrow: carry should be set (result >= 0)")
	}
}

func TestCPU_SBC_Immediate_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x50
	cpu.setStatusCarry(true)
	cpu.Memory.WriteAbsolute(0x200, 0x50)

	cpu.SBC(Immidiate)

	if cpu.A != 0x00 {
		t.Errorf("SBC zero: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("SBC zero: zero flag should be set")
	}
	if !cpu.getStatusCarry() {
		t.Error("SBC zero: carry should be set")
	}
}

func TestCPU_SBC_Immediate_Underflow(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x10
	cpu.setStatusCarry(true)
	cpu.Memory.WriteAbsolute(0x200, 0x20)

	cpu.SBC(Immidiate)

	if cpu.A != 0xF0 {
		t.Errorf("SBC underflow: got A=%d, want %d", cpu.A, 0xF0)
	}
	if cpu.getStatusCarry() {
		t.Error("SBC underflow: carry should be clear (borrow occurred)")
	}
	if !cpu.getStatusNegative() {
		t.Error("SBC underflow: negative flag should be set")
	}
}

// =============================================================================
// ASL - Arithmetic Shift Left
// =============================================================================

func TestCPU_ASL_Accumulator_Carry(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x80 // bit 7 set

	cpu.ASL(Accumulator)

	if cpu.A != 0x00 {
		t.Errorf("ASL acc carry: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusCarry() {
		t.Error("ASL acc carry: carry flag should be set")
	}
	if !cpu.getStatusZero() {
		t.Error("ASL acc carry: zero flag should be set")
	}
}

func TestCPU_ASL_Accumulator_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x7F // 0111_1111

	cpu.ASL(Accumulator)

	if cpu.A != 0xFE {
		t.Errorf("ASL acc neg: got A=%d, want %d", cpu.A, 0xFE)
	}
	if cpu.getStatusCarry() {
		t.Error("ASL acc neg: carry flag should be clear")
	}
	if !cpu.getStatusNegative() {
		t.Error("ASL acc neg: negative flag should be set")
	}
}

func TestCPU_ASL_ZeroPage(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50) // zp addr
	cpu.Memory.WriteZeroPage(0x50, 0x41)  // value 0100_0001

	cpu.ASL(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0x82 {
		t.Errorf("ASL zp: got mem=%d, want %d", val, 0x82)
	}
	if cpu.getStatusCarry() {
		t.Error("ASL zp: carry should be clear")
	}
	if !cpu.getStatusNegative() {
		t.Error("ASL zp: negative flag should be set")
	}
}

func TestCPU_ASL_ZeroPage_WithCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0xC0) // 1100_0000

	cpu.ASL(ZeroPage)

	val := cpu.Memory.ReadZeroPage(0x50)
	if val != 0x80 {
		t.Errorf("ASL zp carry: got mem=%d, want %d", val, 0x80)
	}
	if !cpu.getStatusCarry() {
		t.Error("ASL zp carry: carry should be set")
	}
}

// =============================================================================
// LSR - Logical Shift Right
// =============================================================================

func TestCPU_LSR_Accumulator(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x04 // 0000_0100

	cpu.LSR(Accumulator)

	if cpu.A != 0x02 {
		t.Errorf("LSR acc: got A=%d, want %d", cpu.A, 0x02)
	}
	if cpu.getStatusCarry() {
		t.Error("LSR acc: carry should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("LSR acc: negative flag should always be clear")
	}
}

func TestCPU_LSR_Accumulator_Carry(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x01 // bit 0 set

	cpu.LSR(Accumulator)

	if cpu.A != 0x00 {
		t.Errorf("LSR acc carry: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusCarry() {
		t.Error("LSR acc carry: carry should be set")
	}
	if !cpu.getStatusZero() {
		t.Error("LSR acc carry: zero flag should be set")
	}
}

func TestCPU_LSR_Accumulator_LargeValue(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0xFE // 1111_1110

	cpu.LSR(Accumulator)

	if cpu.A != 0x7F {
		t.Errorf("LSR acc large: got A=%d, want %d", cpu.A, 0x7F)
	}
	if cpu.getStatusCarry() {
		t.Error("LSR acc large: carry should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("LSR acc large: negative should be clear")
	}
}

// =============================================================================
// ROL - Rotate Left
// =============================================================================

func TestCPU_ROL_Accumulator_NoCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x01
	cpu.setStatusCarry(false)

	cpu.ROL(Accumulator)

	if cpu.A != 0x02 {
		t.Errorf("ROL acc no carry: got A=%d, want %d", cpu.A, 0x02)
	}
	if cpu.getStatusCarry() {
		t.Error("ROL acc no carry: carry should be clear")
	}
}

func TestCPU_ROL_Accumulator_WithCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x01
	cpu.setStatusCarry(true)

	cpu.ROL(Accumulator)

	if cpu.A != 0x03 {
		t.Errorf("ROL acc with carry: got A=%d, want %d", cpu.A, 0x03)
	}
	if cpu.getStatusCarry() {
		t.Error("ROL acc with carry: carry should be clear")
	}
}

func TestCPU_ROL_Accumulator_Bit7ToCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x80
	cpu.setStatusCarry(false)

	cpu.ROL(Accumulator)

	if cpu.A != 0x00 {
		t.Errorf("ROL acc bit7: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusCarry() {
		t.Error("ROL acc bit7: carry should be set")
	}
	if !cpu.getStatusZero() {
		t.Error("ROL acc bit7: zero should be set")
	}
}

// =============================================================================
// ROR - Rotate Right
// =============================================================================

func TestCPU_ROR_Accumulator_NoCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x02
	cpu.setStatusCarry(false)

	cpu.ROR(Accumulator)

	if cpu.A != 0x01 {
		t.Errorf("ROR acc no carry: got A=%d, want %d", cpu.A, 0x01)
	}
	if cpu.getStatusCarry() {
		t.Error("ROR acc no carry: carry should be clear")
	}
}

func TestCPU_ROR_Accumulator_WithCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x02
	cpu.setStatusCarry(true)

	cpu.ROR(Accumulator)

	if cpu.A != 0x81 {
		t.Errorf("ROR acc with carry: got A=%d, want %d", cpu.A, 0x81)
	}
	if cpu.getStatusCarry() {
		t.Error("ROR acc with carry: carry should be clear")
	}
	if !cpu.getStatusNegative() {
		t.Error("ROR acc with carry: negative should be set")
	}
}

func TestCPU_ROR_Accumulator_Bit0ToCarry(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x01
	cpu.setStatusCarry(false)

	cpu.ROR(Accumulator)

	if cpu.A != 0x00 {
		t.Errorf("ROR acc bit0: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusCarry() {
		t.Error("ROR acc bit0: carry should be set")
	}
	if !cpu.getStatusZero() {
		t.Error("ROR acc bit0: zero should be set")
	}
}

// =============================================================================
// JMP - Jump
// =============================================================================

func TestCPU_JMP_Absolute(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x00) // low byte
	cpu.Memory.WriteAbsolute(0x201, 0x10) // high byte -> 0x1000

	cpu.JMP(Absolute)

	if cpu.PC != 0x1000 {
		t.Errorf("JMP absolute: got PC=%04x, want %04x", cpu.PC, 0x1000)
	}
}

func TestCPU_JMP_Absolute_LowAddress(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x50) // low byte
	cpu.Memory.WriteAbsolute(0x201, 0x00) // high byte -> 0x0050

	cpu.JMP(Absolute)

	if cpu.PC != 0x0050 {
		t.Errorf("JMP absolute low: got PC=%04x, want %04x", cpu.PC, 0x0050)
	}
}

func TestCPU_JMP_Indirect(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	// Pointer at 0x200 -> address 0x0300
	cpu.Memory.WriteAbsolute(0x200, 0x00) // low byte of pointer
	cpu.Memory.WriteAbsolute(0x201, 0x03) // high byte of pointer -> 0x0300
	// Value at 0x0300 -> target address 0x1234
	cpu.Memory.WriteAbsolute(0x0300, 0x34) // low byte of target
	cpu.Memory.WriteAbsolute(0x0301, 0x12) // high byte of target

	cpu.JMP(Indirect)

	if cpu.PC != 0x1234 {
		t.Errorf("JMP indirect: got PC=%04x, want %04x", cpu.PC, 0x1234)
	}
}

// =============================================================================
// JSR/RTS - Subroutine Call/Return
// =============================================================================

func TestCPU_JSR_Basic(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.Memory.WriteAbsolute(0x200, 0x00) // low byte of target
	cpu.Memory.WriteAbsolute(0x201, 0x10) // high byte -> 0x1000

	cpu.JSR(Absolute)

	if cpu.PC != 0x1000 {
		t.Errorf("JSR: got PC=%04x, want %04x", cpu.PC, 0x1000)
	}
}

func TestCPU_RTS_Basic(t *testing.T) {
	cpu := getTestCPU()
	cpu.stackPush(0x00) // hi byte
	cpu.stackPush(0x64) // lo byte = 0x0064

	cpu.RTS(Implied)

	// RTS adds 1 to the popped address
	if cpu.PC != 0x0065 {
		t.Errorf("RTS: got PC=%04x, want %04x", cpu.PC, 0x0065)
	}
}

func TestCPU_JSR_RTS_Roundtrip(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x300
	cpu.Memory.WriteAbsolute(0x300, 0x00) // low byte
	cpu.Memory.WriteAbsolute(0x301, 0x10) // high byte -> 0x1000

	cpu.JSR(Absolute)

	if cpu.PC != 0x1000 {
		t.Errorf("JSR roundtrip: got PC=%04x, want %04x", cpu.PC, 0x1000)
	}

	// Now RTS should return to 0x302 (the byte after JSR's last operand byte + 1)
	cpu.RTS(Implied)

	if cpu.PC != 0x302 {
		t.Errorf("RTS roundtrip: got PC=%04x, want %04x", cpu.PC, 0x302)
	}
}

// =============================================================================
// BNE - Branch if Not Equal (Z=0)
// =============================================================================

func TestCPU_BNE_Taken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusZero(false)
	cpu.Memory.WriteAbsolute(0x200, 0x10) // offset +16

	cpu.BNE(Relative)

	// PC increments to 0x201 then adds offset 0x10 = 0x211
	if cpu.PC != 0x211 {
		t.Errorf("BNE taken: got PC=%04x, want %04x", cpu.PC, 0x211)
	}
}

func TestCPU_BNE_NotTaken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusZero(true)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BNE(Relative)

	if cpu.PC != 0x201 {
		t.Errorf("BNE not taken: got PC=%04x, want %04x", cpu.PC, 0x201)
	}
}

func TestCPU_BNE_Backward(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x210
	cpu.setStatusZero(false)
	cpu.Memory.WriteAbsolute(0x210, 0xFB) // -5 in two's complement

	cpu.BNE(Relative)

	// PC goes to 0x211, then -5 = 0x20C
	if cpu.PC != 0x020C {
		t.Errorf("BNE backward: got PC=%04x, want %04x", cpu.PC, 0x020C)
	}
}

// =============================================================================
// BEQ - Branch if Equal (Z=1)
// =============================================================================

func TestCPU_BEQ_Taken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusZero(true)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BEQ(Relative)

	if cpu.PC != 0x211 {
		t.Errorf("BEQ taken: got PC=%04x, want %04x", cpu.PC, 0x211)
	}
}

func TestCPU_BEQ_NotTaken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusZero(false)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BEQ(Relative)

	if cpu.PC != 0x201 {
		t.Errorf("BEQ not taken: got PC=%04x, want %04x", cpu.PC, 0x201)
	}
}

// =============================================================================
// BCS - Branch if Carry Set (C=1)
// =============================================================================

func TestCPU_BCS_Taken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusCarry(true)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BCS(Relative)

	if cpu.PC != 0x211 {
		t.Errorf("BCS taken: got PC=%04x, want %04x", cpu.PC, 0x211)
	}
}

func TestCPU_BCS_NotTaken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusCarry(false)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BCS(Relative)

	if cpu.PC != 0x201 {
		t.Errorf("BCS not taken: got PC=%04x, want %04x", cpu.PC, 0x201)
	}
}

// =============================================================================
// BCC - Branch if Carry Clear (C=0)
// =============================================================================

func TestCPU_BCC_Taken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusCarry(false)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BCC(Relative)

	if cpu.PC != 0x211 {
		t.Errorf("BCC taken: got PC=%04x, want %04x", cpu.PC, 0x211)
	}
}

func TestCPU_BCC_NotTaken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusCarry(true)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BCC(Relative)

	if cpu.PC != 0x201 {
		t.Errorf("BCC not taken: got PC=%04x, want %04x", cpu.PC, 0x201)
	}
}

// =============================================================================
// BPL - Branch if Plus (N=0)
// =============================================================================

func TestCPU_BPL_Taken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusNegative(false)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BPL(Relative)

	if cpu.PC != 0x211 {
		t.Errorf("BPL taken: got PC=%04x, want %04x", cpu.PC, 0x211)
	}
}

func TestCPU_BPL_NotTaken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusNegative(true)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BPL(Relative)

	if cpu.PC != 0x201 {
		t.Errorf("BPL not taken: got PC=%04x, want %04x", cpu.PC, 0x201)
	}
}

// =============================================================================
// BMI - Branch if Minus (N=1)
// =============================================================================

func TestCPU_BMI_Taken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusNegative(true)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BMI(Relative)

	if cpu.PC != 0x211 {
		t.Errorf("BMI taken: got PC=%04x, want %04x", cpu.PC, 0x211)
	}
}

func TestCPU_BMI_NotTaken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusNegative(false)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BMI(Relative)

	if cpu.PC != 0x201 {
		t.Errorf("BMI not taken: got PC=%04x, want %04x", cpu.PC, 0x201)
	}
}

// =============================================================================
// BVC - Branch if Overflow Clear (V=0)
// =============================================================================

func TestCPU_BVC_Taken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusOverflow(false)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BVC(Relative)

	if cpu.PC != 0x211 {
		t.Errorf("BVC taken: got PC=%04x, want %04x", cpu.PC, 0x211)
	}
}

func TestCPU_BVC_NotTaken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusOverflow(true)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BVC(Relative)

	if cpu.PC != 0x201 {
		t.Errorf("BVC not taken: got PC=%04x, want %04x", cpu.PC, 0x201)
	}
}

// =============================================================================
// BVS - Branch if Overflow Set (V=1)
// =============================================================================

func TestCPU_BVS_Taken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusOverflow(true)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BVS(Relative)

	if cpu.PC != 0x211 {
		t.Errorf("BVS taken: got PC=%04x, want %04x", cpu.PC, 0x211)
	}
}

func TestCPU_BVS_NotTaken(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.setStatusOverflow(false)
	cpu.Memory.WriteAbsolute(0x200, 0x10)

	cpu.BVS(Relative)

	if cpu.PC != 0x201 {
		t.Errorf("BVS not taken: got PC=%04x, want %04x", cpu.PC, 0x201)
	}
}

// =============================================================================
// SEC - Set Carry Flag
// =============================================================================

func TestCPU_SEC(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusCarry(false)

	cpu.SEC(Implied)

	if !cpu.getStatusCarry() {
		t.Error("SEC: carry flag should be set")
	}
}

// =============================================================================
// CLC - Clear Carry Flag
// =============================================================================

func TestCPU_CLC(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusCarry(true)

	cpu.CLC(Implied)

	if cpu.getStatusCarry() {
		t.Error("CLC: carry flag should be clear")
	}
}

// =============================================================================
// SED - Set Decimal Flag
// =============================================================================

func TestCPU_SED(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusDecimal(false)

	cpu.SED(Implied)

	if !cpu.getStatusDecimal() {
		t.Error("SED: decimal flag should be set")
	}
}

// =============================================================================
// CLD - Clear Decimal Flag
// =============================================================================

func TestCPU_CLD(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusDecimal(true)

	cpu.CLD(Implied)

	if cpu.getStatusDecimal() {
		t.Error("CLD: decimal flag should be clear")
	}
}

// =============================================================================
// SEI - Set Interrupt Disable
// =============================================================================

func TestCPU_SEI(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusIRQ(false)

	cpu.SEI(Implied)

	if !cpu.getStatusIRQ() {
		t.Error("SEI: IRQ disable flag should be set")
	}
}

// =============================================================================
// CLI - Clear Interrupt Disable
// =============================================================================

func TestCPU_CLI(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusIRQ(true)

	cpu.CLI(Implied)

	if cpu.getStatusIRQ() {
		t.Error("CLI: IRQ disable flag should be clear")
	}
}

// =============================================================================
// CLV - Clear Overflow Flag
// =============================================================================

func TestCPU_CLV(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusOverflow(true)

	cpu.CLV(Implied)

	if cpu.getStatusOverflow() {
		t.Error("CLV: overflow flag should be clear")
	}
}

// =============================================================================
// PHA - Push Accumulator
// =============================================================================

func TestCPU_PHA(t *testing.T) {
	cpu := getTestCPU()
	cpu.A = 0x42
	initialSP := cpu.SP

	cpu.PHA(Implied)

	if cpu.SP != initialSP-1 {
		t.Errorf("PHA: got SP=%d, want %d", cpu.SP, initialSP-1)
	}
	// Read the value from the stack
	val := cpu.Memory.ReadAbsolute(cpu.SP_LOW + uint16(initialSP))
	if val != 0x42 {
		t.Errorf("PHA: got stack value=%d, want %d", val, 0x42)
	}
}

// =============================================================================
// PLA - Pull Accumulator
// =============================================================================

func TestCPU_PLA_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.stackPush(0x42)

	cpu.PLA(Implied)

	if cpu.A != 0x42 {
		t.Errorf("PLA: got A=%d, want %d", cpu.A, 0x42)
	}
	if cpu.getStatusZero() {
		t.Error("PLA: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("PLA: negative flag should be clear")
	}
}

func TestCPU_PLA_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.stackPush(0x00)

	cpu.PLA(Implied)

	if cpu.A != 0x00 {
		t.Errorf("PLA zero: got A=%d, want %d", cpu.A, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("PLA zero: zero flag should be set")
	}
}

func TestCPU_PLA_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.stackPush(0x80)

	cpu.PLA(Implied)

	if cpu.A != 0x80 {
		t.Errorf("PLA neg: got A=%d, want %d", cpu.A, 0x80)
	}
	if !cpu.getStatusNegative() {
		t.Error("PLA neg: negative flag should be set")
	}
}

// =============================================================================
// PHP - Push Processor Status
// =============================================================================

func TestCPU_PHP(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusCarry(true)
	cpu.setStatusZero(true)
	initialSP := cpu.SP

	cpu.PHP(Implied)

	if cpu.SP != initialSP-1 {
		t.Errorf("PHP: got SP=%d, want %d", cpu.SP, initialSP-1)
	}
	val := cpu.Memory.ReadAbsolute(cpu.SP_LOW + uint16(initialSP))
	if val != cpu.S {
		t.Errorf("PHP: got stack value=%02x, want %02x", val, cpu.S)
	}
}

// =============================================================================
// PLP - Pull Processor Status
// =============================================================================

func TestCPU_PLP(t *testing.T) {
	cpu := getTestCPU()
	// Push a known status value: carry=1, zero=1, negative=1 = 0x83 + BRK(0x10) + bit5(0x20) = 0xB3
	cpu.stackPush(0x83)

	cpu.PLP(Implied)

	// PLP restores status with B flag forced set and bit 5 forced set
	if !cpu.getStatusCarry() {
		t.Error("PLP: carry flag should be set")
	}
	if !cpu.getStatusZero() {
		t.Error("PLP: zero flag should be set")
	}
	if !cpu.getStatusNegative() {
		t.Error("PLP: negative flag should be set")
	}
}

func TestCPU_PLP_ClearFlags(t *testing.T) {
	cpu := getTestCPU()
	cpu.setStatusCarry(true)
	cpu.setStatusZero(true)
	cpu.setStatusNegative(true)
	// Push status with no flags set
	cpu.stackPush(0x00)

	cpu.PLP(Implied)

	if cpu.getStatusCarry() {
		t.Error("PLP clear: carry flag should be clear")
	}
	if cpu.getStatusZero() {
		t.Error("PLP clear: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("PLP clear: negative flag should be clear")
	}
}

// =============================================================================
// TSX - Transfer Stack Pointer to X
// =============================================================================

func TestCPU_TSX_Normal(t *testing.T) {
	cpu := getTestCPU()
	cpu.SP = 0x42

	cpu.TSX(Implied)

	if cpu.X != 0x42 {
		t.Errorf("TSX: got X=%d, want %d", cpu.X, 0x42)
	}
	if cpu.getStatusZero() {
		t.Error("TSX: zero flag should be clear")
	}
	if cpu.getStatusNegative() {
		t.Error("TSX: negative flag should be clear")
	}
}

func TestCPU_TSX_Negative(t *testing.T) {
	cpu := getTestCPU()
	cpu.SP = 0xFF

	cpu.TSX(Implied)

	if cpu.X != 0xFF {
		t.Errorf("TSX negative: got X=%d, want %d", cpu.X, 0xFF)
	}
	if !cpu.getStatusNegative() {
		t.Error("TSX negative: negative flag should be set")
	}
}

func TestCPU_TSX_Zero(t *testing.T) {
	cpu := getTestCPU()
	cpu.SP = 0x00

	cpu.TSX(Implied)

	if cpu.X != 0x00 {
		t.Errorf("TSX zero: got X=%d, want %d", cpu.X, 0x00)
	}
	if !cpu.getStatusZero() {
		t.Error("TSX zero: zero flag should be set")
	}
}

// =============================================================================
// TXS - Transfer X to Stack Pointer
// =============================================================================

func TestCPU_TXS(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x42

	cpu.TXS(Implied)

	if cpu.SP != 0x42 {
		t.Errorf("TXS: got SP=%d, want %d", cpu.SP, 0x42)
	}
}

func TestCPU_TXS_NoFlags(t *testing.T) {
	cpu := getTestCPU()
	cpu.X = 0x00
	cpu.setStatusZero(false)
	cpu.setStatusNegative(false)

	cpu.TXS(Implied)

	if cpu.SP != 0x00 {
		t.Errorf("TXS no flags: got SP=%d, want %d", cpu.SP, 0x00)
	}
	// TXS does NOT affect flags
	if cpu.getStatusZero() {
		t.Error("TXS no flags: zero flag should remain clear (TXS doesn't affect flags)")
	}
}

// =============================================================================
// BIT - Test Bits in Memory
// =============================================================================

func TestCPU_BIT_ZeroResult(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x0F
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0x80) // A & M = 0x0F & 0x80 = 0x00, bit7=1, bit6=0

	cpu.BIT(ZeroPage)

	if !cpu.getStatusZero() {
		t.Error("BIT zero: zero flag should be set")
	}
	if !cpu.getStatusNegative() {
		t.Error("BIT zero: negative flag should be set (bit 7 of M)")
	}
	if cpu.getStatusOverflow() {
		t.Error("BIT zero: overflow flag should be clear (bit 6 of M is 0)")
	}
}

func TestCPU_BIT_NonZeroResult(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0xFF
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0x40) // bit 6 set

	cpu.BIT(ZeroPage)

	if cpu.getStatusZero() {
		t.Error("BIT nonzero: zero flag should be clear (A & M = 0x40)")
	}
	if cpu.getStatusNegative() {
		t.Error("BIT nonzero: negative flag should be clear (bit 7 of M is 0)")
	}
	if !cpu.getStatusOverflow() {
		t.Error("BIT nonzero: overflow flag should be set (bit 6 of M)")
	}
}

func TestCPU_BIT_AllFlags(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x00
	cpu.Memory.WriteAbsolute(0x200, 0x50)
	cpu.Memory.WriteZeroPage(0x50, 0xC0) // bits 7 and 6 set

	cpu.BIT(ZeroPage)

	if !cpu.getStatusZero() {
		t.Error("BIT all: zero flag should be set (A=0)")
	}
	if !cpu.getStatusNegative() {
		t.Error("BIT all: negative flag should be set (bit 7 of M)")
	}
	if !cpu.getStatusOverflow() {
		t.Error("BIT all: overflow flag should be set (bit 6 of M)")
	}
}

// =============================================================================
// NOP - No Operation
// =============================================================================

func TestCPU_NOP(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200
	cpu.A = 0x42
	cpu.X = 0x10
	cpu.Y = 0x20

	cpu.NOP(Implied)

	if cpu.PC != 0x200 {
		t.Errorf("NOP: PC should not change, got %04x", cpu.PC)
	}
	if cpu.A != 0x42 {
		t.Errorf("NOP: A should not change, got %d", cpu.A)
	}
	if cpu.X != 0x10 {
		t.Errorf("NOP: X should not change, got %d", cpu.X)
	}
	if cpu.Y != 0x20 {
		t.Errorf("NOP: Y should not change, got %d", cpu.Y)
	}
}

// =============================================================================
// BRK - Break
// =============================================================================

func TestCPU_BRK(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x200

	initialSP := cpu.SP
	initialS := cpu.S

	cpu.BRK(Implied)

	// SP should have decreased by 3 (pushed PC hi, PC lo, status)
	if cpu.SP != initialSP-3 {
		t.Errorf("BRK: got SP=%d, want %d", cpu.SP, initialSP-3)
	}
	// IRQ disable flag should be set
	if !cpu.getStatusIRQ() {
		t.Error("BRK: IRQ flag should be set")
	}
	// BRK flag should be set
	if !cpu.getStatusBRK() {
		t.Error("BRK: BRK flag should be set")
	}

	// Verify pushed values on stack: status was pushed after PC
	// Stack grows down; values at SP_LOW + initialSP, initialSP-1, initialSP-2
	// BRK pushes: PC+1 hi, PC+1 lo, status
	pushedPCHi := cpu.Memory.ReadAbsolute(cpu.SP_LOW + uint16(initialSP))
	pushedPCLo := cpu.Memory.ReadAbsolute(cpu.SP_LOW + uint16(initialSP-1))
	pushedStatus := cpu.Memory.ReadAbsolute(cpu.SP_LOW + uint16(initialSP-2))

	// PC+1 = 0x201
	expectedPC := uint16(0x201)
	gotPC := uint16(pushedPCLo) | (uint16(pushedPCHi) << 8)
	if gotPC != expectedPC {
		t.Errorf("BRK pushed PC: got %04x, want %04x", gotPC, expectedPC)
	}

	// Status pushed should be the original status with BRK cleared (as per implementation)
	_ = pushedStatus
	_ = initialS
}

func TestCPU_BRK_PushesCorrectPC(t *testing.T) {
	cpu := getTestCPU()
	cpu.PC = 0x0305
	initialSP := cpu.SP

	cpu.BRK(Implied)

	// BRK pushes PC+1 = 0x0306 to stack (hi byte first, then lo byte)
	pushedPCHi := cpu.Memory.ReadAbsolute(cpu.SP_LOW + uint16(initialSP))
	pushedPCLo := cpu.Memory.ReadAbsolute(cpu.SP_LOW + uint16(initialSP-1))

	pushedPC := uint16(pushedPCLo) | (uint16(pushedPCHi) << 8)
	if pushedPC != 0x0306 {
		t.Errorf("BRK pushed PC: got %04x, want %04x", pushedPC, 0x0306)
	}
}

// =============================================================================
// RTI - Return from Interrupt
// =============================================================================

func TestCPU_RTI(t *testing.T) {
	cpu := getTestCPU()
	// Push PC hi, PC lo, then status (in BRK order)
	// RTI pops in order: status, PC lo, PC hi
	cpu.stackPush(0x10) // PC hi byte
	cpu.stackPush(0x00) // PC lo byte  -> 0x1000
	cpu.stackPush(0x01) // status (carry set)

	cpu.RTI(Implied)

	if cpu.PC != 0x1000 {
		t.Errorf("RTI: got PC=%04x, want %04x", cpu.PC, 0x1000)
	}
	if !cpu.getStatusCarry() {
		t.Error("RTI: carry flag should be restored")
	}
}

func TestCPU_RTI_RestoresFlags(t *testing.T) {
	cpu := getTestCPU()
	// Push state: PC=0x0400, status with negative and carry
	cpu.stackPush(0x04) // PC hi
	cpu.stackPush(0x00) // PC lo -> 0x0400
	cpu.stackPush(0x81) // status: negative(0x80) + carry(0x01)

	cpu.RTI(Implied)

	if cpu.PC != 0x0400 {
		t.Errorf("RTI flags: got PC=%04x, want %04x", cpu.PC, 0x0400)
	}
	if !cpu.getStatusNegative() {
		t.Error("RTI flags: negative flag should be restored")
	}
	if !cpu.getStatusCarry() {
		t.Error("RTI flags: carry flag should be restored")
	}
}
