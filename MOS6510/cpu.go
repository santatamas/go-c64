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

func (c *CPU) setStatusCarry(flag bool) {
	if flag {
		c.S |= 0x01
	} else {
		c.S &^= 0x01
	}
}

func (c *CPU) getStatusCarry() bool {
	return c.S&0x01 == 0x01
}

func (c *CPU) setStatusZero(flag bool) {
	if flag {
		c.S |= 0x02
	} else {
		c.S &^= 0x02
	}
}

func (c *CPU) getStatusZero() bool {
	return c.S&0x02 == 0x02
}

func (c *CPU) setStatusNegative(flag bool) {
	if flag {
		c.S |= 0x80
	} else {
		c.S &^= 0x80
	}
}

func (c *CPU) getStatusNegative() bool {
	return c.S&0x80 == 0x80
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
