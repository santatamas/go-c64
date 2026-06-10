package VIC2

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestGetColor_Black(t *testing.T) {
	result := getColor(0)
	if result != tcell.ColorBlack {
		t.Errorf("Expected ColorBlack for value 0, got %v", result)
	}
}

func TestGetColor_White(t *testing.T) {
	result := getColor(1)
	if result != tcell.ColorWhite {
		t.Errorf("Expected ColorWhite for value 1, got %v", result)
	}
}

func TestGetColor_DarkRed(t *testing.T) {
	result := getColor(2)
	if result != tcell.ColorDarkRed {
		t.Errorf("Expected ColorDarkRed for value 2, got %v", result)
	}
}

func TestGetColor_LightCyan(t *testing.T) {
	result := getColor(3)
	if result != tcell.ColorLightCyan {
		t.Errorf("Expected ColorLightCyan for value 3, got %v", result)
	}
}

func TestGetColor_Maroon(t *testing.T) {
	result := getColor(4)
	if result != tcell.ColorMaroon {
		t.Errorf("Expected ColorMaroon for value 4, got %v", result)
	}
}

func TestGetColor_DarkGreen(t *testing.T) {
	result := getColor(5)
	if result != tcell.ColorDarkGreen {
		t.Errorf("Expected ColorDarkGreen for value 5, got %v", result)
	}
}

func TestGetColor_DarkBlue(t *testing.T) {
	result := getColor(6)
	if result != tcell.ColorDarkBlue {
		t.Errorf("Expected ColorDarkBlue for value 6, got %v", result)
	}
}

func TestGetColor_Yellow(t *testing.T) {
	result := getColor(7)
	if result != tcell.ColorYellow {
		t.Errorf("Expected ColorYellow for value 7, got %v", result)
	}
}

func TestGetColor_YellowGreen(t *testing.T) {
	result := getColor(8)
	if result != tcell.ColorYellowGreen {
		t.Errorf("Expected ColorYellowGreen for value 8, got %v", result)
	}
}

func TestGetColor_DarkMagenta(t *testing.T) {
	result := getColor(9)
	if result != tcell.ColorDarkMagenta {
		t.Errorf("Expected ColorDarkMagenta for value 9, got %v", result)
	}
}

func TestGetColor_Red(t *testing.T) {
	result := getColor(10)
	if result != tcell.ColorRed {
		t.Errorf("Expected ColorRed for value 10, got %v", result)
	}
}

func TestGetColor_DarkGray(t *testing.T) {
	result := getColor(11)
	if result != tcell.ColorDarkGray {
		t.Errorf("Expected ColorDarkGray for value 11, got %v", result)
	}
}

func TestGetColor_Gray(t *testing.T) {
	result := getColor(12)
	if result != tcell.ColorGray {
		t.Errorf("Expected ColorGray for value 12, got %v", result)
	}
}

func TestGetColor_Green(t *testing.T) {
	result := getColor(13)
	if result != tcell.ColorGreen {
		t.Errorf("Expected ColorGreen for value 13, got %v", result)
	}
}

func TestGetColor_Blue(t *testing.T) {
	result := getColor(14)
	if result != tcell.ColorBlue {
		t.Errorf("Expected ColorBlue for value 14, got %v", result)
	}
}

func TestGetColor_DarkCyan(t *testing.T) {
	result := getColor(15)
	if result != tcell.ColorDarkCyan {
		t.Errorf("Expected ColorDarkCyan for value 15, got %v", result)
	}
}

func TestGetColor_MasksHighBits(t *testing.T) {
	// Value 16 should be masked to 0 (16 & 15 = 0)
	result := getColor(16)
	if result != tcell.ColorBlack {
		t.Errorf("Expected ColorBlack for value 16 (masked to 0), got %v", result)
	}
}

func TestGetColor_MasksHighBits_Value17(t *testing.T) {
	// Value 17 should be masked to 1 (17 & 15 = 1)
	result := getColor(17)
	if result != tcell.ColorWhite {
		t.Errorf("Expected ColorWhite for value 17 (masked to 1), got %v", result)
	}
}

func TestGetColor_MasksHighBits_Value255(t *testing.T) {
	// Value 255 should be masked to 15 (255 & 15 = 15)
	result := getColor(255)
	if result != tcell.ColorDarkCyan {
		t.Errorf("Expected ColorDarkCyan for value 255 (masked to 15), got %v", result)
	}
}

func TestGetColor_MasksHighBits_Value32(t *testing.T) {
	// Value 32 should be masked to 0 (32 & 15 = 0)
	result := getColor(32)
	if result != tcell.ColorBlack {
		t.Errorf("Expected ColorBlack for value 32 (masked to 0), got %v", result)
	}
}

func TestGetColor_AllColors_TableDriven(t *testing.T) {
	expected := map[byte]tcell.Color{
		0:  tcell.ColorBlack,
		1:  tcell.ColorWhite,
		2:  tcell.ColorDarkRed,
		3:  tcell.ColorLightCyan,
		4:  tcell.ColorMaroon,
		5:  tcell.ColorDarkGreen,
		6:  tcell.ColorDarkBlue,
		7:  tcell.ColorYellow,
		8:  tcell.ColorYellowGreen,
		9:  tcell.ColorDarkMagenta,
		10: tcell.ColorRed,
		11: tcell.ColorDarkGray,
		12: tcell.ColorGray,
		13: tcell.ColorGreen,
		14: tcell.ColorBlue,
		15: tcell.ColorDarkCyan,
	}

	for val, expectedColor := range expected {
		result := getColor(val)
		if result != expectedColor {
			t.Errorf("For value %d: expected %v, got %v", val, expectedColor, result)
		}
	}
}
