package main

import (
	"bytes"
	"encoding/binary"
)

type CPU struct {
	memory       Memory
	A            byte
	Y            byte // low
	X            byte // high
	S            byte
	P            byte
	PC           int16
	instructions map[AssemblyInstructionType]func(AddressingMode)
}

func newCPU(mem Memory) CPU {
	cpu := CPU{memory: mem}
	cpu.instructions = make(map[AssemblyInstructionType]func(AddressingMode))

	// Wiring instruction implementation methods
	cpu.instructions[BNE] = cpu.BNE
	cpu.instructions[BRK] = cpu.BRK
	cpu.instructions[CPX] = cpu.CPX
	cpu.instructions[INX] = cpu.INX
	cpu.instructions[INY] = cpu.INY
	cpu.instructions[LDA] = cpu.LDA
	cpu.instructions[LDX] = cpu.LDX
	cpu.instructions[STA] = cpu.STA
	cpu.instructions[TAY] = cpu.TAY
	cpu.instructions[TYA] = cpu.TYA
	return cpu
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

func toInt16(data []byte) (ret int16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func (cpu *CPU) Start(PCH byte, PCL byte) {
	// initialise an instruction type map
	instrTypes := assemblyInstructions()

	// Get the initial value of the program counter
	PC := toInt16([]byte{PCL, PCH})
	for {
		// Fetch first executable instruction code from memory
		instrCode := cpu.memory.ReadAbsolute(PC)

		// Resolve instruction by instruction code
		instruction := instrTypes(instrCode)

		// Run instruction, and passing the addressingmode as param
		cpu.instructions[instruction.Type](instruction.AddressingMode)

		// increase program counter
		PC++
		break
	}
}
