package CIA

import (
	"testing"
)

func TestCIA_NewCIA_Initialization(t *testing.T) {
	irqChan := make(chan bool, 1)
	cia := NewCIA(irqChan)

	if cia.Interrupt != false {
		t.Errorf("Expected Interrupt to be false, got %v", cia.Interrupt)
	}
	if cia.PORT_A != 0x0 {
		t.Errorf("Expected PORT_A to be 0x00, got 0x%02X", cia.PORT_A)
	}
	if cia.TIMER_A != 0x0 {
		t.Errorf("Expected TIMER_A to be 0, got %d", cia.TIMER_A)
	}
	if cia.TIMER_A_ENABLED != false {
		t.Errorf("Expected TIMER_A_ENABLED to be false, got %v", cia.TIMER_A_ENABLED)
	}
	if cia.TIMER_A_IRQ != false {
		t.Errorf("Expected TIMER_A_IRQ to be false, got %v", cia.TIMER_A_IRQ)
	}
	if cia.TIMER_B != 0x0 {
		t.Errorf("Expected TIMER_B to be 0, got %d", cia.TIMER_B)
	}
	if cia.TIMER_B_ENABLED != false {
		t.Errorf("Expected TIMER_B_ENABLED to be false, got %v", cia.TIMER_B_ENABLED)
	}
	if cia.TIMER_B_IRQ != false {
		t.Errorf("Expected TIMER_B_IRQ to be false, got %v", cia.TIMER_B_IRQ)
	}
	if len(cia.Keyboard_matrix) != 8 {
		t.Errorf("Expected keyboard matrix length 8, got %d", len(cia.Keyboard_matrix))
	}
	if cia.IrqChannel == nil {
		t.Errorf("Expected IrqChannel to be set, got nil")
	}
}

func TestCIA_NewCIA_NilChannel(t *testing.T) {
	cia := NewCIA(nil)

	if cia.IrqChannel != nil {
		t.Errorf("Expected IrqChannel to be nil")
	}
	if len(cia.Keyboard_matrix) != 8 {
		t.Errorf("Expected keyboard matrix length 8, got %d", len(cia.Keyboard_matrix))
	}
}

func TestCIA_ExecuteCycle_UpdatesPreviousCycles(t *testing.T) {
	cia := NewCIA(nil)

	cia.ExecuteCycle(100)
	if cia.Previous_cpu_cycles != 100 {
		t.Errorf("Expected Previous_cpu_cycles to be 100, got %d", cia.Previous_cpu_cycles)
	}

	cia.ExecuteCycle(200)
	if cia.Previous_cpu_cycles != 200 {
		t.Errorf("Expected Previous_cpu_cycles to be 200, got %d", cia.Previous_cpu_cycles)
	}
}

func TestCIA_SendInterrupt_SetsFlag(t *testing.T) {
	cia := NewCIA(nil)

	if cia.Interrupt != false {
		t.Errorf("Expected Interrupt to initially be false")
	}

	cia.SendInterrupt()

	if cia.Interrupt != true {
		t.Errorf("Expected Interrupt to be true after SendInterrupt, got %v", cia.Interrupt)
	}
}

func TestCIA_ResetKeyboardMatrix_ClearsAll(t *testing.T) {
	cia := NewCIA(nil)

	// Set some keys
	cia.SetKey(0, 0)
	cia.SetKey(1, 2)
	cia.SetKey(3, 5)
	cia.SetKey(7, 7)

	// Verify keys are set
	if cia.Keyboard_matrix[0] == 0 {
		t.Errorf("Expected matrix[0] to be non-zero after SetKey")
	}

	// Reset
	cia.ResetKeyboardMatrix()

	// Verify all cleared
	for i := 0; i < 8; i++ {
		if cia.Keyboard_matrix[i] != 0 {
			t.Errorf("Expected keyboard matrix[%d] to be 0 after reset, got %d", i, cia.Keyboard_matrix[i])
		}
	}
}

func TestCIA_Write_SetsPortA(t *testing.T) {
	cia := NewCIA(nil)

	cia.Write(CIA_PORT_A, 0xFD)

	if cia.PORT_A != 0xFD {
		t.Errorf("Expected PORT_A to be 0xFD, got 0x%02X", cia.PORT_A)
	}
}

func TestCIA_Write_SetsPortA_AllValues(t *testing.T) {
	cia := NewCIA(nil)

	testValues := []byte{0x00, 0x01, 0x7F, 0x80, 0xFE, 0xFF}
	for _, val := range testValues {
		cia.Write(CIA_PORT_A, val)
		if cia.PORT_A != val {
			t.Errorf("Expected PORT_A to be 0x%02X, got 0x%02X", val, cia.PORT_A)
		}
	}
}

func TestCIA_Read_ReturnsRegisterValue(t *testing.T) {
	cia := NewCIA(nil)

	// Set a key and configure PORT_A to read that row
	cia.SetKey(0, 3)        // row 0, col 3
	cia.WriteRegister(0xFE) // bit 0 clear = read row 0

	result := cia.Read(CIA_PORT_B)

	// ReadRegister should return the keyboard scan with inversion
	// Row 0 has bit 3 set (from SetKey), PORT_A bit 0 is clear so row 0 is read
	// Result = keyboard_matrix[0] = 0x08, inverted = 0xF7
	expected := byte(0x08) ^ 0xFF
	if result != expected {
		t.Errorf("Expected Read to return 0x%02X, got 0x%02X", expected, result)
	}
}

func TestCIA_UnsetKey(t *testing.T) {
	cia := NewCIA(nil)

	// Set multiple keys in same row
	cia.SetKey(2, 3)
	cia.SetKey(2, 5)

	// Verify both are set: bit 3 and bit 5 = 0010 1000 = 0x28
	if cia.Keyboard_matrix[2] != 0x28 {
		t.Errorf("Expected matrix[2] = 0x28, got 0x%02X", cia.Keyboard_matrix[2])
	}

	// Unset one key
	cia.UnsetKey(2, 3)

	// Only bit 5 should remain: 0010 0000 = 0x20
	if cia.Keyboard_matrix[2] != 0x20 {
		t.Errorf("Expected matrix[2] = 0x20 after UnsetKey, got 0x%02X", cia.Keyboard_matrix[2])
	}
}

func TestCIA_UnsetKey_ClearsAll(t *testing.T) {
	cia := NewCIA(nil)

	cia.SetKey(4, 6)
	cia.UnsetKey(4, 6)

	if cia.Keyboard_matrix[4] != 0 {
		t.Errorf("Expected matrix[4] = 0 after unsetting only key, got 0x%02X", cia.Keyboard_matrix[4])
	}
}

func TestCIA_SetKey_SetsMatrixProperly_SingleKey(t *testing.T) {
	// Arrange
	cia := NewCIA(nil)

	// Act
	cia.SetKey(1, 2)

	// Assert
	rowValue := cia.Keyboard_matrix[1]

	// 0000 0100
	if rowValue != 4 {
		t.Errorf("Keyboard matrix value was incorrect, got: %d, want: %d.", rowValue, 4)
	}
}

func TestCIA_SetKey_SetsMatrixProperly_MultipleKeys(t *testing.T) {
	// Arrange
	cia := NewCIA(nil)

	// Act
	cia.SetKey(1, 2)
	cia.SetKey(1, 5)

	// Assert
	rowValue := cia.Keyboard_matrix[1]

	// 0010 0100
	if rowValue != 36 {
		t.Errorf("Keyboard matrix value was incorrect, got: %d, want: %d.", rowValue, 36)
	}
}

func TestCIA_ReadRegister_ReturnsCorrectValue_SingleKey(t *testing.T) {
	// Arrange
	cia := NewCIA(nil)

	// Act
	cia.SetKey(1, 2)
	cia.WriteRegister(253) // 1111 1101 - read 2nd row values (bit 1 clear)

	// Assert
	rowValue := cia.ReadRegister()

	// matrix[1] = 0x04, result = 0x04 ^ 0xFF = 0xFB (251)
	// ReadRegister inverts the bits (active-low signaling)
	expected := byte(0x04 ^ 0xFF)
	if rowValue != expected {
		t.Errorf("Keyboard matrix value was incorrect, got: %d, want: %d.", rowValue, expected)
	}
}

func TestCIA_ReadRegister_ReturnsZeroValue_SingleKey(t *testing.T) {
	// Arrange
	cia := NewCIA(nil)

	// Act
	cia.SetKey(1, 2)
	cia.WriteRegister(0xFF) // 1111 1111 - all bits set, no rows read

	// Assert
	rowValue := cia.ReadRegister()

	// No rows read (all PORT_A bits high), result = 0x00 ^ 0xFF = 0xFF
	if rowValue != 0xFF {
		t.Errorf("Keyboard matrix value was incorrect, got: %d, want: %d.", rowValue, 0xFF)
	}
}

func TestCIA_ReadRegister_ReturnsCorrectValue_MultipleKey(t *testing.T) {
	// Arrange
	cia := NewCIA(nil)

	// Act
	cia.SetKey(1, 2)
	cia.SetKey(1, 4)
	cia.WriteRegister(253) // 1111 1101 - bit 1 clear, read row 1 values

	// Assert
	rowValue := cia.ReadRegister()

	// matrix[1] = 0x14 (bit 2 + bit 4), result = 0x14 ^ 0xFF = 0xEB (235)
	expected := byte(0x14 ^ 0xFF)
	if rowValue != expected {
		t.Errorf("Keyboard matrix value was incorrect, got: %d, want: %d.", rowValue, expected)
	}
}

func TestCIA_WriteRegister_SetsPortA(t *testing.T) {
	cia := NewCIA(nil)

	cia.WriteRegister(0xAB)

	if cia.PORT_A != 0xAB {
		t.Errorf("Expected PORT_A to be 0xAB, got 0x%02X", cia.PORT_A)
	}
}

func TestCIA_ReadRegister_NoKeysPressed_ReturnsZero(t *testing.T) {
	cia := NewCIA(nil)

	// No keys set, all matrix bytes are 0
	cia.WriteRegister(0x00) // All rows enabled

	result := cia.ReadRegister()

	// No keys pressed, OR of all zeros = 0, XOR 0xFF = 0xFF
	if result != 0xFF {
		t.Errorf("Expected 0xFF when no keys pressed, got 0x%02X", result)
	}
}
