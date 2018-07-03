package main

import (
	"fmt"
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
	fmt.Println("BNE called")
	cpu.PC++
}

func (cpu *CPU) BRK(mode AddressingMode) {
	// Do nothing apparently?
	fmt.Println("BRK called")
}

func (cpu *CPU) CPX(mode AddressingMode) {
	fmt.Println("CPX called")
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
	fmt.Println("INX called")
	cpu.X++
	cpu.PC++
}

func (cpu *CPU) INY(mode AddressingMode) {
	fmt.Println("INY called")
	cpu.Y++
	cpu.PC++
}

func (cpu *CPU) LDA(mode AddressingMode) {
	fmt.Println("LDA called")
	if mode == Immidiate {
		cpu.A = cpu.memory.ReadAbsolute(cpu.PC)
		fmt.Printf("Value loaded to CPU register A: %x \n", cpu.A)
		cpu.PC++
	}
}

func (cpu *CPU) LDX(mode AddressingMode) {
	fmt.Println("LDX called")
}

func (cpu *CPU) STA(mode AddressingMode) {
	fmt.Println("STA called")
	if mode == Absolute {
		cpu.Y = cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cpu.X = cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cpu.memory.WriteAbsolute(toInt16([]byte{cpu.X, cpu.Y}), cpu.A)
	}
}

func (cpu *CPU) TAY(mode AddressingMode) {
	fmt.Println("TAY called")
	if mode == Absolute {
		cpu.Y = cpu.A
		cpu.PC++
	}
}

func (cpu *CPU) TYA(mode AddressingMode) {
	fmt.Println("TYA called")
	if mode == Absolute {
		cpu.A = cpu.Y
		cpu.PC++
	}
}
