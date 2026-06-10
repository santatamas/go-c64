package numeric

import (
	"testing"
)

func TestToInt16_LittleEndian(t *testing.T) {
	// Little-endian: [lo, hi] -> hi<<8 | lo
	data := []byte{0x34, 0x12}
	result := ToInt16(data)

	if result != 0x1234 {
		t.Errorf("Expected 0x1234, got 0x%04X", result)
	}
}

func TestToInt16_ZeroValue(t *testing.T) {
	data := []byte{0x00, 0x00}
	result := ToInt16(data)

	if result != 0x0000 {
		t.Errorf("Expected 0x0000, got 0x%04X", result)
	}
}

func TestToInt16_MaxValue(t *testing.T) {
	data := []byte{0xFF, 0xFF}
	result := ToInt16(data)

	if result != 0xFFFF {
		t.Errorf("Expected 0xFFFF, got 0x%04X", result)
	}
}

func TestToInt16_OnlyLowByte(t *testing.T) {
	data := []byte{0xAB, 0x00}
	result := ToInt16(data)

	if result != 0x00AB {
		t.Errorf("Expected 0x00AB, got 0x%04X", result)
	}
}

func TestToInt16_OnlyHighByte(t *testing.T) {
	data := []byte{0x00, 0xCD}
	result := ToInt16(data)

	if result != 0xCD00 {
		t.Errorf("Expected 0xCD00, got 0x%04X", result)
	}
}

func TestToInt16_2_HiLoConversion(t *testing.T) {
	// ToInt16_2 takes (hi, lo) but reads as little-endian buffer [hi, lo]
	// So buffer is [hi, lo] read as little-endian = lo<<8 | hi
	result := ToInt16_2(0x34, 0x12)

	// Buffer is [0x34, 0x12], little-endian read = 0x1234
	if result != 0x1234 {
		t.Errorf("Expected 0x1234, got 0x%04X", result)
	}
}

func TestToInt16_2_ZeroValues(t *testing.T) {
	result := ToInt16_2(0x00, 0x00)

	if result != 0x0000 {
		t.Errorf("Expected 0x0000, got 0x%04X", result)
	}
}

func TestToInt16_2_MaxValues(t *testing.T) {
	result := ToInt16_2(0xFF, 0xFF)

	if result != 0xFFFF {
		t.Errorf("Expected 0xFFFF, got 0x%04X", result)
	}
}

func TestGetLO_ExtractsLowByte(t *testing.T) {
	result := GetLO(0x1234)

	if result != 0x34 {
		t.Errorf("Expected 0x34, got 0x%02X", result)
	}
}

func TestGetLO_ZeroValue(t *testing.T) {
	result := GetLO(0x0000)

	if result != 0x00 {
		t.Errorf("Expected 0x00, got 0x%02X", result)
	}
}

func TestGetLO_MaxValue(t *testing.T) {
	result := GetLO(0xFFFF)

	if result != 0xFF {
		t.Errorf("Expected 0xFF, got 0x%02X", result)
	}
}

func TestGetLO_OnlyHighByteSet(t *testing.T) {
	result := GetLO(0xFF00)

	if result != 0x00 {
		t.Errorf("Expected 0x00, got 0x%02X", result)
	}
}

func TestGetLO_OnlyLowByteSet(t *testing.T) {
	result := GetLO(0x00AB)

	if result != 0xAB {
		t.Errorf("Expected 0xAB, got 0x%02X", result)
	}
}

func TestGetHI_ExtractsHighByte(t *testing.T) {
	result := GetHI(0x1234)

	if result != 0x12 {
		t.Errorf("Expected 0x12, got 0x%02X", result)
	}
}

func TestGetHI_ZeroValue(t *testing.T) {
	result := GetHI(0x0000)

	if result != 0x00 {
		t.Errorf("Expected 0x00, got 0x%02X", result)
	}
}

func TestGetHI_MaxValue(t *testing.T) {
	result := GetHI(0xFFFF)

	if result != 0xFF {
		t.Errorf("Expected 0xFF, got 0x%02X", result)
	}
}

func TestGetHI_OnlyHighByteSet(t *testing.T) {
	result := GetHI(0xCD00)

	if result != 0xCD {
		t.Errorf("Expected 0xCD, got 0x%02X", result)
	}
}

func TestGetHI_OnlyLowByteSet(t *testing.T) {
	result := GetHI(0x00FF)

	if result != 0x00 {
		t.Errorf("Expected 0x00, got 0x%02X", result)
	}
}

func TestIsNegative_Bit7Set_ReturnsTrue(t *testing.T) {
	result := IsNegative(128) // 1000 0000

	if !result {
		t.Error("Expected IsNegative to return true for 128")
	}
}

func TestIsNegative_Bit7Clear_ReturnsFalse(t *testing.T) {
	result := IsNegative(127) // 0111 1111

	if result {
		t.Error("Expected IsNegative to return false for 127")
	}
}

func TestIsNegative_AllBitsSet(t *testing.T) {
	result := IsNegative(0xFF) // 1111 1111

	if !result {
		t.Error("Expected IsNegative to return true for 0xFF")
	}
}

func TestIsNegative_Zero(t *testing.T) {
	result := IsNegative(0)

	if result {
		t.Error("Expected IsNegative to return false for 0")
	}
}

func TestIsNegative_JustBit7(t *testing.T) {
	result := IsNegative(0x80)

	if !result {
		t.Error("Expected IsNegative to return true for 0x80")
	}
}

func TestGetBit_Bit0(t *testing.T) {
	result := GetBit(0x01, 0) // bit 0 set

	if !result {
		t.Error("Expected GetBit to return true for bit 0 of 0x01")
	}
}

func TestGetBit_Bit7(t *testing.T) {
	result := GetBit(0x80, 7) // bit 7 set

	if !result {
		t.Error("Expected GetBit to return true for bit 7 of 0x80")
	}
}

func TestGetBit_BitNotSet(t *testing.T) {
	result := GetBit(0x01, 1) // bit 1 not set in 0x01

	if result {
		t.Error("Expected GetBit to return false for bit 1 of 0x01")
	}
}

func TestGetBit_AllBitsSet(t *testing.T) {
	for i := byte(0); i < 8; i++ {
		result := GetBit(0xFF, i)
		if !result {
			t.Errorf("Expected GetBit to return true for bit %d of 0xFF", i)
		}
	}
}

func TestGetBit_NoBitsSet(t *testing.T) {
	for i := byte(0); i < 8; i++ {
		result := GetBit(0x00, i)
		if result {
			t.Errorf("Expected GetBit to return false for bit %d of 0x00", i)
		}
	}
}

func TestGetBit_Bit2_Of_Value20(t *testing.T) {
	// 20 = 0001 0100, bit 2 is set
	result := GetBit(20, 2)

	if !result {
		t.Error("Expected GetBit to return true for bit 2 of 20")
	}
}

func TestGetBit_Bit4_Of_Value20(t *testing.T) {
	// 20 = 0001 0100, bit 4 is set
	result := GetBit(20, 4)

	if !result {
		t.Error("Expected GetBit to return true for bit 4 of 20")
	}
}

func TestGetBit_Bit3_Of_Value20(t *testing.T) {
	// 20 = 0001 0100, bit 3 is not set
	result := GetBit(20, 3)

	if result {
		t.Error("Expected GetBit to return false for bit 3 of 20")
	}
}

func TestGetLO_GetHI_Roundtrip(t *testing.T) {
	original := uint16(0xABCD)
	lo := GetLO(original)
	hi := GetHI(original)

	reconstructed := ToInt16([]byte{lo, hi})
	if reconstructed != original {
		t.Errorf("Expected roundtrip to produce 0x%04X, got 0x%04X", original, reconstructed)
	}
}
