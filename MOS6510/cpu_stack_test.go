package MOS6510

import (
	"testing"

	"github.com/santatamas/go-c64/RAM"
)

func TestCPU_stackPush_happypath(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory()
	cpu := NewCPU(&memory)

	// Act
	cpu.stackPush(100)

	// Assert
	if cpu.SP != 0xFE {
		t.Errorf("Stack pointer value was incorrect, got: %d, want: %d.", cpu.SP, 0xFE)
	}
}

func TestCPU_stackPush_overflow(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory()
	cpu := NewCPU(&memory)

	// Act
	for i := 0; i < 255; i++ {
		cpu.stackPush(100)
	}

	err := cpu.stackPush(100)

	// Assert
	if err == nil {
		t.Error("Expected a stackoverflow error, got nothing.")
	}
}

func TestCPU_stackPop_underflow(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory()
	cpu := NewCPU(&memory)
	cpu.stackPush(100)

	// Act
	cpu.stackPop()
	_, err := cpu.stackPop()

	// Assert
	if err == nil {
		t.Error("Expected a stack underflow error, got nothing.")
	}
}

func TestCPU_stackPop_happypath(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory()
	cpu := NewCPU(&memory)
	cpu.stackPush(100)

	// Act
	result, _ := cpu.stackPop()

	// Assert
	if result != 100 {
		t.Errorf("Stack value was incorrect, got: %d, want: %d.", result, 100)
	}
}
