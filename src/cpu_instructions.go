package main

import (
	"log"
)

func (cpu *CPU) callMethod(instruction AssemblyInstruction) {
	switch instruction.Type {
	case BNE:
		cpu.BNE(instruction.AddressingMode)
	case BRK:
		cpu.BRK(instruction.AddressingMode)
	case CPX:
		cpu.CPX(instruction.AddressingMode)
	case INX:
		cpu.INX(instruction.AddressingMode)
	case INY:
		cpu.INY(instruction.AddressingMode)
	case LDA:
		cpu.LDA(instruction.AddressingMode)
	case LDX:
		cpu.LDX(instruction.AddressingMode)
	case STA:
		cpu.STA(instruction.AddressingMode)
	case TAY:
		cpu.TAY(instruction.AddressingMode)
	case TYA:
		cpu.TYA(instruction.AddressingMode)
	}
}

func (cpu *CPU) BNE(mode AddressingMode) {
	log.Println("BNE called")
	cpu.PC++
}

func (cpu *CPU) BRK(mode AddressingMode) {
	// Do nothing apparently?
	log.Println("BRK called")
}

func (cpu *CPU) CPX(mode AddressingMode) {
	log.Println("CPX called")
	if mode == Immidiate {
		aa := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		bb := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cc := cpu.memory.ReadAbsolute(toInt16([]byte{aa, bb}))
		cpu.setStatusZero(cc == cpu.X)
		cpu.PC++
	}
}

func (cpu *CPU) INX(mode AddressingMode) {
	log.Println("INX called")
	cpu.X++
	cpu.PC++
}

func (cpu *CPU) INY(mode AddressingMode) {
	log.Println("INY called")
	cpu.Y++
	cpu.PC++
}

func (cpu *CPU) LDA(mode AddressingMode) {
	log.Println("LDA called")
	if mode == Immidiate {
		cpu.A = cpu.memory.ReadAbsolute(cpu.PC)
		log.Printf("Value loaded to CPU register A: %x \n", cpu.A)
		cpu.PC++
	}
}

func (cpu *CPU) LDX(mode AddressingMode) {
	log.Println("LDX called")
}

func (cpu *CPU) STA(mode AddressingMode) {
	log.Println("STA called")
	if mode == Absolute {
		cpu.Y = cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cpu.X = cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		log.Println("CPU register A value: ", cpu.A)
		log.Println("Setting CPU register A to address: ", toInt16([]byte{cpu.Y, cpu.X}))
		cpu.memory.WriteAbsolute(toInt16([]byte{cpu.Y, cpu.X}), cpu.A)
	}
}

func (cpu *CPU) TAY(mode AddressingMode) {
	log.Println("TAY called")
	if mode == Absolute {
		cpu.Y = cpu.A
		cpu.PC++
	}
}

func (cpu *CPU) TYA(mode AddressingMode) {
	log.Println("TYA called")
	if mode == Absolute {
		cpu.A = cpu.Y
		cpu.PC++
	}
}
