package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CPU struct {
	memory Memory
	A      byte
	Y      byte // low
	X      byte // high
	S      byte
	P      byte
	PC     uint16
}

func newCPU(mem Memory) CPU {
	return CPU{memory: mem}
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

func toInt16(data []byte) (ret uint16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func (cpu *CPU) Start(PCH byte, PCL byte) {
	// initialise an instruction type map
	instrTypes := assemblyInstructions()

	// Get the initial value of the program counter
	cpu.PC = toInt16([]byte{PCL, PCH})
	fmt.Printf("Start address: %x \n", cpu.PC)
	for {
		fmt.Printf("Current PC address: %x \n", cpu.PC)

		// Fetch first executable instruction code from memory
		instrCode := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		fmt.Printf("Next instruction code: %x \n", instrCode)

		// Resolve instruction by instruction code
		instruction := instrTypes(instrCode)
		if instruction.Type == BRK {
			break
		}

		cpu.callMethod(instruction)
	}
}
