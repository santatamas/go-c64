package VIC2

import (
	"testing"
)

func TestAsciiCharacters_Space_Code32(t *testing.T) {
	converter := asciiCharacters()

	result := converter(32)
	if result != ' ' {
		t.Errorf("Expected space for code 32, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_LetterA_Code1(t *testing.T) {
	converter := asciiCharacters()

	result := converter(1)
	// Code 1 maps to 'A' (uppercase in the charList)
	if result != 'A' {
		t.Errorf("Expected 'A' for code 1, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_LetterB_Code2(t *testing.T) {
	converter := asciiCharacters()

	result := converter(2)
	if result != 'B' {
		t.Errorf("Expected 'B' for code 2, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_LetterC_Code3(t *testing.T) {
	converter := asciiCharacters()

	result := converter(3)
	if result != 'C' {
		t.Errorf("Expected 'C' for code 3, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_Numbers_Codes48to57(t *testing.T) {
	converter := asciiCharacters()

	for i := byte(48); i <= 57; i++ {
		expected := rune('0' + (i - 48))
		result := converter(i)
		if result != expected {
			t.Errorf("For code %d: expected '%c', got '%c'", i, expected, result)
		}
	}
}

func TestAsciiCharacters_HighBitMaskedOff(t *testing.T) {
	converter := asciiCharacters()

	// Code 1 with high bit set (1 + 128 = 129) should still map like code 1
	result := converter(129) // 129 & 127 = 1
	expected := converter(1)
	if result != expected {
		t.Errorf("Expected high-bit masked code 129 to equal code 1, got '%c' vs '%c'", result, expected)
	}
}

func TestAsciiCharacters_HighBitMaskedOff_Space(t *testing.T) {
	converter := asciiCharacters()

	// Code 32 with high bit set (32 + 128 = 160) should still map like code 32
	result := converter(160) // 160 & 127 = 32
	expected := converter(32)
	if result != expected {
		t.Errorf("Expected high-bit masked code 160 to equal code 32 (space), got '%c' vs '%c'", result, expected)
	}
}

func TestAsciiCharacters_UnmappedCode_ReturnsSpace(t *testing.T) {
	converter := asciiCharacters()

	// Code 96 maps to "SPACE" in the charList
	result := converter(96)
	if result != ' ' {
		t.Errorf("Expected space for code 96, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_LetterZ_Code26(t *testing.T) {
	converter := asciiCharacters()

	result := converter(26)
	if result != 'Z' {
		t.Errorf("Expected 'Z' for code 26, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_Exclamation_Code33(t *testing.T) {
	converter := asciiCharacters()

	result := converter(33)
	if result != '!' {
		t.Errorf("Expected '!' for code 33, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_Colon_Code58(t *testing.T) {
	converter := asciiCharacters()

	result := converter(58)
	if result != ':' {
		t.Errorf("Expected ':' for code 58, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_Semicolon_Code59(t *testing.T) {
	converter := asciiCharacters()

	result := converter(59)
	if result != ';' {
		t.Errorf("Expected ';' for code 59, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_Slash_Code47(t *testing.T) {
	converter := asciiCharacters()

	result := converter(47)
	if result != '/' {
		t.Errorf("Expected '/' for code 47, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_Plus_Code43(t *testing.T) {
	converter := asciiCharacters()

	result := converter(43)
	if result != '+' {
		t.Errorf("Expected '+' for code 43, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_Minus_Code45(t *testing.T) {
	converter := asciiCharacters()

	result := converter(45)
	if result != '-' {
		t.Errorf("Expected '-' for code 45, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_Zero_Code0(t *testing.T) {
	converter := asciiCharacters()

	// Code 0 maps to ' ' (space character) in the charList
	result := converter(0)
	if result != ' ' {
		t.Errorf("Expected space for code 0, got '%c' (%d)", result, result)
	}
}

func TestAsciiCharacters_ConsistentResults(t *testing.T) {
	converter := asciiCharacters()

	// Calling with same value should always return same result
	result1 := converter(5)
	result2 := converter(5)
	if result1 != result2 {
		t.Errorf("Inconsistent results for same code: '%c' vs '%c'", result1, result2)
	}
}
