package MOS6510

import (
	"testing"

	"github.com/santatamas/go-c64/RAM"
)

func TestCPU_stackPush_happypath(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory(nil)
	cpu := NewCPU(&memory)

	// Act
	cpu.stackPush(100)

	// Assert
	if cpu.SP != 0xFE {
		t.Errorf("Stack pointer value was incorrect, got: %d, want: %d.", cpu.SP, 0xFE)
	}
}

func TestCPU_stackPush_multipleValues(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory(nil)
	cpu := NewCPU(&memory)

	// Act
	cpu.stackPush(10)
	cpu.stackPush(20)
	cpu.stackPush(30)

	// Assert
	if cpu.SP != 0xFC {
		t.Errorf("Stack pointer value was incorrect, got: %d, want: %d.", cpu.SP, 0xFC)
	}
}

func TestCPU_stackPush_wraps_around(t *testing.T) {
	// The real 6502 stack wraps around (no error), test that behavior
	memory := RAM.NewMemory(nil)
	cpu := NewCPU(&memory)

	// Push 255 times to wrap SP from 0xFF back to 0xFF
	for i := 0; i < 255; i++ {
		cpu.stackPush(100)
	}

	// SP should have wrapped around to 0x00
	if cpu.SP != 0x00 {
		t.Errorf("Stack pointer value was incorrect after 255 pushes, got: %x, want: %x.", cpu.SP, 0x00)
	}

	// One more push wraps to 0xFF
	cpu.stackPush(100)
	if cpu.SP != 0xFF {
		t.Errorf("Stack pointer value was incorrect after 256 pushes, got: %x, want: %x.", cpu.SP, 0xFF)
	}
}

func TestCPU_stackPop_happypath(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory(nil)
	cpu := NewCPU(&memory)
	cpu.stackPush(100)

	// Act
	result, _ := cpu.stackPop()

	// Assert
	if result != 100 {
		t.Errorf("Stack value was incorrect, got: %d, want: %d.", result, 100)
	}
}

func TestCPU_stackPop_LIFO_order(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory(nil)
	cpu := NewCPU(&memory)
	cpu.stackPush(10)
	cpu.stackPush(20)
	cpu.stackPush(30)

	// Act & Assert - should come back in reverse order
	result, _ := cpu.stackPop()
	if result != 30 {
		t.Errorf("First pop value was incorrect, got: %d, want: %d.", result, 30)
	}

	result, _ = cpu.stackPop()
	if result != 20 {
		t.Errorf("Second pop value was incorrect, got: %d, want: %d.", result, 20)
	}

	result, _ = cpu.stackPop()
	if result != 10 {
		t.Errorf("Third pop value was incorrect, got: %d, want: %d.", result, 10)
	}
}

func TestCPU_stackPop_restores_SP(t *testing.T) {
	// Arrange
	memory := RAM.NewMemory(nil)
	cpu := NewCPU(&memory)
	cpu.stackPush(100)

	// Act
	cpu.stackPop()

	// Assert
	if cpu.SP != 0xFF {
		t.Errorf("Stack pointer not restored, got: %x, want: %x.", cpu.SP, 0xFF)
	}
}

func TestCPU_stack_uses_correct_address_range(t *testing.T) {
	// The 6502 stack lives at 0x0100-0x01FF
	memory := RAM.NewMemory(nil)
	cpu := NewCPU(&memory)

	// Push a value - should go to 0x01FF (SP_LOW + SP)
	cpu.stackPush(42)

	// Read from the expected memory location
	val := memory.ReadAbsolute(0x01FF)
	if val != 42 {
		t.Errorf("Stack value not at correct memory address, got: %d, want: %d.", val, 42)
	}
}
