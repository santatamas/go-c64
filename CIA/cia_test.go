package CIA

import (
	"testing"
)

func TestCIA_SetKey_SetsMatrixProperly_SingleKey(t *testing.T) {
	// Arrange
	cia := NewCIA()

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
	cia := NewCIA()

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
	cia := NewCIA()

	// Act
	cia.SetKey(1, 2)
	cia.WriteRegister(253) // 1111 1101 - read 2nd row values

	// Assert
	rowValue := cia.ReadRegister()

	// 0000 0100
	if rowValue != 4 {
		t.Errorf("Keyboard matrix value was incorrect, got: %d, want: %d.", rowValue, 4)
	}
}

func TestCIA_ReadRegister_ReturnsZeroValue_SingleKey(t *testing.T) {
	// Arrange
	cia := NewCIA()

	// Act
	cia.SetKey(1, 2)
	cia.WriteRegister(0xFF) // 1111 1111 - ignore all row values

	// Assert
	rowValue := cia.ReadRegister()

	// 0000 0000
	if rowValue != 0 {
		t.Errorf("Keyboard matrix value was incorrect, got: %d, want: %d.", rowValue, 0)
	}
}

func TestCIA_ReadRegister_ReturnsCorrectValue_MultipleKey(t *testing.T) {
	// Arrange
	cia := NewCIA()

	// Act
	cia.SetKey(1, 2)
	cia.SetKey(1, 4)
	cia.WriteRegister(253) // 1111 1001 - read 2nd row values

	// Assert
	rowValue := cia.ReadRegister()

	// 0001 0100
	if rowValue != 20 {
		t.Errorf("Keyboard matrix value was incorrect, got: %d, want: %d.", rowValue, 20)
	}
}
