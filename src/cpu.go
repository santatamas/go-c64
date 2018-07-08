package main

import (
	"log"
)

type CPU struct {
	memory  *Memory
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

func newCPU(mem *Memory) CPU {
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

func (cpu *CPU) Start(PCH byte, PCL byte) {
	// initialise an instruction type map
	instrTypes := assemblyInstructions()

	// Get the initial value of the program counter
	cpu.PC = toInt16([]byte{PCL, PCH})
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
