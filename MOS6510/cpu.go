package MOS6510

import (
	"github.com/santatamas/go-c64/RAM"
	"log"
)

type CPU struct {
	Memory     *RAM.Memory
	A          byte
	Y          byte // low
	X          byte // high
	S          byte
	P          byte
	PC         uint16
	SP         byte
	SP_LOW     uint16
	SP_HIGH    uint16
	instrTypes func(byte) AssemblyInstruction
}

func NewCPU(mem *RAM.Memory) CPU {

	return CPU{
		Memory:     mem,
		SP_LOW:     0x0100,
		SP_HIGH:    0x01FF,
		SP:         0xFF,
		instrTypes: assemblyInstructions(),
	}
}

func getTestCPU() (result CPU) {
	memory := RAM.NewMemory()
	return NewCPU(&memory)
}

func (cpu *CPU) ExecuteCycle() bool {
	log.Printf("Current PC address: %x \n", cpu.PC)

	// Fetch first executable instruction code from memory
	instrCode := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	log.Printf("Next instruction code: %x \n", instrCode)

	// Resolve instruction by instruction code
	instruction := cpu.instrTypes(instrCode)

	cpu.callMethod(instruction)
	return instruction.Type != BRK
}
