package MOS6510

import (
	"github.com/santatamas/go-c64/RAM"
	n "github.com/santatamas/go-c64/numeric"
	"log"
)

type CPU struct {
	memory  *RAM.Memory
	A       byte
	Y       byte // low
	X       byte // high
	S       byte
	P       byte
	PC      uint16
	SP      byte
	SP_LOW  uint16
	SP_HIGH uint16
}

func NewCPU(mem *RAM.Memory) CPU {
	return CPU{
		memory:  mem,
		SP_LOW:  0x0100,
		SP_HIGH: 0x01FF,
		SP:      0xFF,
	}
}

func getTestCPU() (result CPU) {
	memory := RAM.NewMemory()
	return NewCPU(&memory)
}

func (cpu *CPU) Start(PCL byte, PCH byte) {
	// initialise an instruction type map
	instrTypes := assemblyInstructions()

	// Get the initial value of the program counter
	cpu.PC = n.ToInt16_2(PCH, PCL)
	log.Printf("Start address: %x \n", cpu.PC)
	for {
		log.Printf("Current PC address: %x \n", cpu.PC)

		// Fetch first executable instruction code from memory
		instrCode := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		log.Printf("Next instruction code: %x \n", instrCode)

		// Resolve instruction by instruction code
		instruction := instrTypes(instrCode)

		cpu.callMethod(instruction)
		if instruction.Type == BRK {
			break
		}
	}
}
