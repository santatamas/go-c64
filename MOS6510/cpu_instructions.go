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
	case STX:
		cpu.STX(instruction.AddressingMode)
	case TAY:
		cpu.TAY(instruction.AddressingMode)
	case TAX:
		cpu.TAX(instruction.AddressingMode)
	case TXA:
		cpu.TXA(instruction.AddressingMode)
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
	case AND:
		cpu.AND(instruction.AddressingMode)
	case INC:
		cpu.INC(instruction.AddressingMode)
	case JMP:
		cpu.JMP(instruction.AddressingMode)
	case CMP:
		cpu.CMP(instruction.AddressingMode)
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
	log.Println("ADC called -- adr. mode: ", mode.toString())
	if mode == ZeroPage {
		log.Println("Zeropage called")
		zpAdr := cpu.memory.ReadAbsolute(cpu.PC)
		hi := cpu.memory.ReadZeroPage(zpAdr)

		result := hi
		if cpu.getStatusCarry() == true {
			result++
		}

		if uint16(cpu.A)+uint16(result) > 255 {
			cpu.setStatusCarry(true)
		} else {
			cpu.setStatusCarry(false)
		}

		result = cpu.A + result

		isOverflow := (((cpu.A ^ hi) & 0x80) == 0) && ((cpu.A^result)&0x80) > 0
		cpu.setStatusOverflow(isOverflow)
		cpu.setStatusZero(result == 0)

		cpu.A = result
		log.Println("A register: ", cpu.A)
		log.Println("CPU carry flag: ", cpu.getStatusCarry())
		cpu.PC++
	}
}

// Branch on C = 1
func (cpu *CPU) BCS(mode AddressingMode) {
	log.Println("BCS called")
	if cpu.getStatusCarry() == true {
		hi := int8(cpu.memory.ReadAbsolute(cpu.PC))
		log.Println("Moving PC address by: ", hi)

		cpu.PC += uint16(hi) + 1
	} else {
		cpu.PC++
	}
}

func (cpu *CPU) BPL(mode AddressingMode) {
	log.Println("BPL called")
	if cpu.getStatusNegative() == false {
		hi := int8(cpu.memory.ReadAbsolute(cpu.PC))
		log.Println("Moving PC address by: ", hi)

		cpu.PC += uint16(hi) + 1
	} else {
		cpu.PC++
	}
}

func (cpu *CPU) ASL(mode AddressingMode) {
	log.Println("ASL called -- adr. mode: ", mode.toString())
	if mode == Accumulator {

		carryValue := cpu.A & 0x80
		if carryValue == 128 {
			log.Println("Setting status Carry to: true")
			cpu.setStatusCarry(true)
		}

		result := cpu.A << 1

		cpu.setStatusNegative((result & 0x80) == 128)
		cpu.setStatusZero(result == 0)

		cpu.A = result
		log.Println("A register: ", cpu.A)
	}
}

// Operation:  PC + 2 toS, (PC + 1) -> PCL
//                         (PC + 2) -> PCH
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
	log.Println("BNE called -- adr.mode: ", mode.toString())
	if !cpu.getStatusZero() {
		hi := int8(cpu.memory.ReadAbsolute(cpu.PC))
		log.Println("Moving PC address by: ", hi)

		cpu.PC += uint16(hi) + 1
	} else {
		cpu.PC++
	}
}

func (cpu *CPU) BRK(mode AddressingMode) {
	log.Println("BRK called -- do nothing")
}

func (cpu *CPU) CPX(mode AddressingMode) {
	log.Println("CPX called -- adr.mode: ", mode.toString())
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

	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 == 128)
}

func (cpu *CPU) DEY(mode AddressingMode) {
	log.Println("DEY called")
	cpu.Y--

	cpu.setStatusZero(cpu.Y == 0)
	cpu.setStatusNegative(cpu.Y&0x80 == 128)
}

func (cpu *CPU) INX(mode AddressingMode) {
	log.Println("INX called")
	cpu.X++

	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 == 128)
}

func (cpu *CPU) INY(mode AddressingMode) {
	log.Println("INY called")
	cpu.Y++

	cpu.setStatusZero(cpu.Y == 0)
	cpu.setStatusNegative(cpu.Y&0x80 == 128)
}

func (cpu *CPU) LDA(mode AddressingMode) {
	log.Println("LDA called -- adr. mode: ", mode.toString())
	if mode == Immidiate {
		cpu.A = cpu.memory.ReadAbsolute(cpu.PC)
		log.Printf("Value loaded to CPU register A: %x \n", cpu.A)
		cpu.PC++

		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)
	}

	if mode == Absolute {
		hi := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		lo := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cpu.A = cpu.memory.ReadAbsolute(n.ToInt16([]byte{hi, lo}))
		log.Printf("Value loaded to CPU register A: %x \n", cpu.A)

		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)
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
	log.Println("STA called -- adr. mode: ", mode.toString())

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

func (cpu *CPU) STX(mode AddressingMode) {
	log.Println("STX called -- adr. mode: ", mode.toString())

	hi := cpu.memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	// Zeropage only needs a single byte address
	if mode == ZeroPage {
		log.Println("CPU register X value: ", cpu.X)
		log.Println("Setting CPU register X to zero page address: ", hi)
		cpu.memory.WriteZeroPage(hi, cpu.X)
		return
	}

	if mode == ZeroPageY {
		hi += cpu.Y
		log.Println("CPU register X value: ", cpu.X)
		log.Println("Setting CPU register X to address: ", hi)
		cpu.memory.WriteZeroPage(hi, cpu.X)
	}

	if mode == Absolute {
		lo := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		log.Println("CPU register X value: ", cpu.X)
		log.Println("Setting CPU register X to address: ", n.ToInt16([]byte{hi, lo}))
		cpu.memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), cpu.X)
	}
}

func (cpu *CPU) TAX(mode AddressingMode) {
	log.Println("TAX called -- adr. mode: ", mode.toString())
	if mode == Implied {
		log.Println("Setting CPU register X to value: ", cpu.A)

		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)

		cpu.X = cpu.A
	}
}

func (cpu *CPU) TXA(mode AddressingMode) {
	log.Println("TXA called -- adr. mode: ", mode.toString())
	if mode == Implied {
		log.Println("Setting CPU register A to value: ", cpu.X)

		cpu.setStatusZero(cpu.X == 0)
		cpu.setStatusNegative(cpu.X&0x80 == 128)

		cpu.A = cpu.X
	}
}

func (cpu *CPU) TAY(mode AddressingMode) {
	log.Println("TAY called -- adr. mode: ", mode.toString())
	if mode == Implied {
		log.Println("Setting CPU register Y to value: ", cpu.A)

		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)

		cpu.Y = cpu.A
	}
}

func (cpu *CPU) TYA(mode AddressingMode) {
	log.Println("TYA called -- adr. mode: ", mode.toString())
	if mode == Implied {
		log.Println("Setting CPU register A to value: ", cpu.Y)

		cpu.setStatusZero(cpu.Y == 0)
		cpu.setStatusNegative(cpu.Y&0x80 == 128)

		cpu.A = cpu.Y
	}
}

func (cpu *CPU) AND(mode AddressingMode) {
	log.Println("AND called -- adr. mode: ", mode.toString())
	if mode == Immidiate {
		log.Println("Immidiate mode")
		mem := cpu.memory.ReadAbsolute(cpu.PC)
		log.Println("mem value: ", mem)
		cpu.A = mem & cpu.A
		log.Println("CPU.A: ", cpu.A)

		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)

		cpu.PC++
	}

}

func (cpu *CPU) INC(mode AddressingMode) {
	log.Println("INC called -- adr. mode: ", mode.toString())
	if mode == Absolute {
		log.Println("Absolute mode")

		hi := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		lo := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		mem := cpu.memory.ReadAbsolute(n.ToInt16([]byte{hi, lo}))
		cpu.memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), mem+1)

		cpu.setStatusZero(mem+1 == 0)
		cpu.setStatusNegative(mem+1&0x80 == 128)
	}
}

func (cpu *CPU) JMP(mode AddressingMode) {
	log.Println("JMP called -- adr. mode: ", mode.toString())
	if mode == Absolute {
		log.Println("Absolute mode")
		hi := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		lo := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cpu.PC = n.ToInt16([]byte{hi, lo})
	}
}

func (cpu *CPU) CMP(mode AddressingMode) {
	log.Println("CMP called -- adr. mode: ", mode.toString())
	if mode == Immidiate {
		log.Println("Immidiate mode")
		mem := cpu.memory.ReadAbsolute(cpu.PC)
		log.Println("mem value: ", mem)
		log.Println("CPU.A value: ", cpu.A)
		result := cpu.A - mem
		log.Println("result value:", result)

		cpu.setStatusZero(result == 0)
		cpu.setStatusNegative(result&0x80 == 128)
		cpu.setStatusCarry(cpu.A >= mem)

		cpu.PC++
	}

	if mode == Absolute {
		hi := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		lo := cpu.memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		mem := cpu.memory.ReadAbsolute(n.ToInt16([]byte{hi, lo}))
		log.Println("mem value: ", mem)
		log.Println("CPU.A value: ", cpu.A)
		result := cpu.A - mem
		log.Println("result value:", result)

		cpu.setStatusZero(result == 0)
		cpu.setStatusNegative(result&0x80 == 128)
		cpu.setStatusCarry(cpu.A >= mem)
	}
}
