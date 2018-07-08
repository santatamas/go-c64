package MOS6510

import (
	n "github.com/santatamas/go-c64/numeric"
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
	case LDY:
		cpu.LDY(instruction.AddressingMode)
	case STA:
		cpu.STA(instruction.AddressingMode)
	case TAY:
		cpu.TAY(instruction.AddressingMode)
	case TYA:
		cpu.TYA(instruction.AddressingMode)
	case JSR:
		cpu.JSR(instruction.AddressingMode)
	}
}

func (cpu *CPU) JSR(mode AddressingMode) {
	log.Println("JSR called")
	log.Printf("PC address: %x", cpu.PC)
	lo := cpu.memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	pch := n.GetHI(cpu.PC)
	pcl := n.GetLO(cpu.PC)

	log.Printf("Current PCH: %x", pch)
	log.Printf("Current PCL: %x", pcl)

	cpu.stackPush(pch)
	cpu.stackPush(pcl)

	hi := cpu.memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	log.Printf("PC address: %x", cpu.PC)

	cpu.PC = n.ToInt16_2(lo, hi)
	log.Printf("Jump PC address: %x", cpu.PC)
}

func (cpu *CPU) BNE(mode AddressingMode) {
	log.Println("BNE called")
	if !cpu.getStatusZero() {
		hi := int8(cpu.memory.ReadAbsolute(cpu.PC))
		log.Println("Moving PC address by: ", hi)

		cpu.PC += uint16(hi) + 1
	}
}

func (cpu *CPU) BRK(mode AddressingMode) {
	log.Println("BRK called -- do nothing")
}

func (cpu *CPU) CPX(mode AddressingMode) {
	log.Println("CPX called -- adr.mode: ", mode)
	if mode == Immidiate {
		aa := cpu.memory.ReadAbsolute(cpu.PC)
		log.Println("Value to compare with X: ", aa)
		cpu.PC++

		tmp := cpu.X - aa

		cpu.setStatusZero(tmp == 0)
		cpu.setStatusNegative(tmp < 0)
		cpu.setStatusCarry(cpu.X >= aa)
	}
}

func (cpu *CPU) INX(mode AddressingMode) {
	log.Println("INX called")
	cpu.X++
}

func (cpu *CPU) INY(mode AddressingMode) {
	log.Println("INY called")
	cpu.Y++
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
	if mode == Immidiate {
		cpu.X = cpu.memory.ReadAbsolute(cpu.PC)
		log.Printf("Value loaded to CPU register X: %x \n", cpu.X)
		cpu.PC++
	}
}

func (cpu *CPU) LDY(mode AddressingMode) {
	log.Println("LDY called")
	if mode == Immidiate {
		cpu.Y = cpu.memory.ReadAbsolute(cpu.PC)
		log.Printf("Value loaded to CPU register Y: %x \n", cpu.Y)
		cpu.PC++
	}
}

func (cpu *CPU) STA(mode AddressingMode) {
	log.Println("STA called -- adr. mode: ", mode)

	hi := cpu.memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	lo := cpu.memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	if mode == AbsoluteX {
		hi += cpu.X
	}

	if mode == AbsoluteY {
		hi += cpu.Y
	}

	// if mode == Absolute
	log.Println("CPU register A value: ", cpu.A)
	log.Println("Setting CPU register A to address: ", n.ToInt16([]byte{hi, lo}))
	cpu.memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), cpu.A)
}

func (cpu *CPU) TAY(mode AddressingMode) {
	log.Println("TAY called -- adr. mode: ", mode)
	if mode == Implied {
		log.Println("Setting CPU register Y to value: ", cpu.A)
		cpu.Y = cpu.A
	}
}

func (cpu *CPU) TYA(mode AddressingMode) {
	log.Println("TYA called -- adr. mode: ", mode)
	if mode == Implied {
		log.Println("Setting CPU register A to value: ", cpu.Y)
		cpu.A = cpu.Y
	}
}
