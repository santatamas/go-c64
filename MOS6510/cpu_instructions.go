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
	case TAX:
		cpu.TAX(instruction.AddressingMode)
	case TYA:
		cpu.TYA(instruction.AddressingMode)
	case JSR:
		cpu.JSR(instruction.AddressingMode)
	case ASL:
		cpu.ASL(instruction.AddressingMode)
	case BCS:
		cpu.BCS(instruction.AddressingMode)
	case ADC:
		cpu.ADC(instruction.AddressingMode)
	case RTS:
		cpu.RTS(instruction.AddressingMode)
	case DEX:
		cpu.DEX(instruction.AddressingMode)
	case DEY:
		cpu.DEY(instruction.AddressingMode)
	case BPL:
		cpu.BPL(instruction.AddressingMode)
	}
}

func (cpu *CPU) RTS(mode AddressingMode) {
	log.Println("RTS called")
	lo, _ := cpu.stackPop()
	hi, _ := cpu.stackPop()

	cpu.PC = n.ToInt16_2(lo, hi)
	cpu.PC++
}

// Operation:  A + M + C -> A, C
func (cpu *CPU) ADC(mode AddressingMode) {
	log.Println("ADC called -- adr. mode: ", mode)
	if mode == ZeroPage {
		log.Println("Zeropage called")
		hi := cpu.memory.ReadAbsolute(cpu.PC)

		result := hi
		if cpu.getStatusCarry() == true {
			result++
		}

		result = cpu.A + result

		if result > 255 {
			cpu.setStatusCarry(true)
		} else {
			cpu.setStatusCarry(false)
		}

		//cpu.set

		//setV(!((A ^ op) & 0x80) && ((A ^ sum) & 0x80));

		cpu.A = result
		cpu.PC++
	}
}

func (cpu *CPU) BCS(mode AddressingMode) {
	log.Println("BCS called")
	if cpu.getStatusCarry() == true {
		hi := int8(cpu.memory.ReadAbsolute(cpu.PC))
		log.Println("Moving PC address by: ", hi)

		cpu.PC += uint16(hi) + 1
	}
}

func (cpu *CPU) BPL(mode AddressingMode) {
	log.Println("BPL called")
	if cpu.getStatusNegative() == false {
		hi := int8(cpu.memory.ReadAbsolute(cpu.PC))
		log.Println("Moving PC address by: ", hi)

		cpu.PC += uint16(hi) + 1
	}
}

func (cpu *CPU) ASL(mode AddressingMode) {
	log.Println("ASL called -- adr. mode: ", mode)
	if mode == Accumulator {
		cpu.A = cpu.A << 1
	}
}

func (cpu *CPU) JSR(mode AddressingMode) {
	log.Println("JSR called")
	lo := cpu.memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	cpu.stackPush(n.GetHI(cpu.PC))
	cpu.stackPush(n.GetLO(cpu.PC))

	hi := cpu.memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	cpu.PC = n.ToInt16_2(lo, hi)
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

func (cpu *CPU) DEX(mode AddressingMode) {
	log.Println("DEX called")
	cpu.X--
}

func (cpu *CPU) DEY(mode AddressingMode) {
	log.Println("DEY called")
	cpu.Y--
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

	// Zeropage only needs a single byte address
	if mode == ZeroPage {
		log.Println("CPU register A value: ", cpu.A)
		log.Println("Setting CPU register A to zero page address: ", hi)
		cpu.memory.WriteZeroPage(hi, cpu.A)
		return
	}

	lo := cpu.memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	if mode == Absolute {
		log.Println("CPU register A value: ", cpu.A)
		log.Println("Setting CPU register A to address: ", n.ToInt16([]byte{hi, lo}))
		cpu.memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), cpu.A)
	}

	if mode == AbsoluteX {
		hi += cpu.X
		log.Println("CPU register A value: ", cpu.A)
		log.Println("Setting CPU register A to address: ", n.ToInt16([]byte{hi, lo}))
		cpu.memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), cpu.A)
	}

	if mode == AbsoluteY {
		hi += cpu.Y
		log.Println("CPU register A value: ", cpu.A)
		log.Println("Setting CPU register A to address: ", n.ToInt16([]byte{hi, lo}))
		cpu.memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), cpu.A)
	}

}

func (cpu *CPU) TAX(mode AddressingMode) {
	log.Println("TAX called -- adr. mode: ", mode)
	if mode == Implied {
		log.Println("Setting CPU register X to value: ", cpu.A)
		cpu.X = cpu.A
	}
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
