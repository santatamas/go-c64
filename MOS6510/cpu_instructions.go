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
	case STY:
		cpu.STY(instruction.AddressingMode)
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
	case BEQ:
		cpu.BEQ(instruction.AddressingMode)
	case SEC:
		cpu.SEC(instruction.AddressingMode)
	case SBC:
		cpu.SBC(instruction.AddressingMode)
	case ROR:
		cpu.ROR(instruction.AddressingMode)
	case EOR:
		cpu.EOR(instruction.AddressingMode)
	case NOP:
		cpu.NOP(instruction.AddressingMode)
	case CPY:
		cpu.CPY(instruction.AddressingMode)
	case BCC:
		cpu.BCC(instruction.AddressingMode)
	case CLC:
		cpu.CLC(instruction.AddressingMode)
	case BMI:
		cpu.BMI(instruction.AddressingMode)
	default:
		log.Println("[WARNING] Unimplemented instruction! ", instruction.Type)
	}
}

func (cpu *CPU) getMemoryValue(mode AddressingMode) byte {
	hi := byte(0)

	switch mode {
	case Immidiate:
		{
			hi = cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
		}
	case ZeroPage:
		{
			hi = cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			hi = cpu.Memory.ReadZeroPage(hi)
		}
	case Absolute:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			hi = cpu.Memory.ReadAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}))
		}
	case AbsoluteX:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			memAdr := n.ToInt16([]byte{adr_hi, adr_lo})
			memAdr += uint16(cpu.X)

			hi = cpu.Memory.ReadAbsolute(memAdr)
		}
	default:
		log.Println("[WARNING] Unsupported addressing mode!")
	}

	return hi
}

func (cpu *CPU) branch() {
	hi := int8(cpu.Memory.ReadAbsolute(cpu.PC))
	log.Println("Moving PC address by: ", hi)

	pcHi := n.GetHI(cpu.PC)

	cpu.PC += uint16(hi)

	if pcHi != n.GetHI(cpu.PC) {
		cpu.PC += 2
	} else {
		cpu.PC++
	}
}

// RTS Return from Subroutine
// Operation:  PC fromS, PC + 1 -> PC
//  N Z C I D V
//  _ _ _ _ _ _
func (cpu *CPU) RTS(mode AddressingMode) {
	log.Println("RTS called -- adr.mode: ", mode.toString())
	lo, _ := cpu.stackPop()
	hi, _ := cpu.stackPop()

	cpu.PC = n.ToInt16_2(lo, hi)
	cpu.PC++
}

// ADC Add memory to accumulator with carry
// Operation:  A + M + C -> A, C
// N Z C I D V
// * * * _ _ *
func (cpu *CPU) ADC(mode AddressingMode) {
	log.Println("ADC called -- adr. mode: ", mode.toString())

	hi := cpu.getMemoryValue(mode)

	if cpu.getStatusCarry() {
		hi++
	}

	cpu.setStatusCarry(uint16(cpu.A)+uint16(hi) > 255)

	result := cpu.A + hi

	isOverflow := (((cpu.A ^ hi) & 0x80) != 128) && ((cpu.A^result)&0x80) == 128
	cpu.setStatusOverflow(isOverflow)
	cpu.setStatusZero(result == 0)

	cpu.A = result
}

// BCS Branch on carry set
// Operation:  Branch on C = 1
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BCS(mode AddressingMode) {
	log.Println("BCS called -- adr.mode: ", mode.toString())
	if cpu.getStatusCarry() {
		cpu.branch()
	} else {
		cpu.PC++
	}
}

// BPL Branch on result plus
// Operation:  Branch on N = 0
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BPL(mode AddressingMode) {
	log.Println("BPL called -- adr.mode: ", mode.toString())
	if !cpu.getStatusNegative() {
		cpu.branch()
	} else {
		cpu.PC++
	}
}

// ASL Shift Left One Bit (Memory or Accumulator)
// 					+-+-+-+-+-+-+-+-+
// Operation:  C <- |7|6|5|4|3|2|1|0| <- 0
//					+-+-+-+-+-+-+-+-+
// N Z C I D V
// * * * _ _ _
func (cpu *CPU) ASL(mode AddressingMode) {
	log.Println("ASL called -- adr. mode: ", mode.toString())
	if mode == Accumulator {

		cpu.setStatusCarry(cpu.A&0x80 == 128)

		result := cpu.A << 1

		cpu.setStatusNegative((result & 0x80) == 128)
		cpu.setStatusZero(result == 0)

		cpu.A = result
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// JSR Jump to new location saving return address
// Operation:  PC + 2 toS, (PC + 1) -> PCL
//                         (PC + 2) -> PCH
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) JSR(mode AddressingMode) {
	log.Println("JSR called -- adr.mode: ", mode.toString())
	lo := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	cpu.stackPush(n.GetHI(cpu.PC))
	cpu.stackPush(n.GetLO(cpu.PC))

	hi := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	cpu.PC = n.ToInt16_2(lo, hi)
}

// BNE Branch on result not zero
// Operation:  Branch on Z = 0
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BNE(mode AddressingMode) {
	log.Println("BNE called -- adr.mode: ", mode.toString())
	if !cpu.getStatusZero() {
		cpu.branch()
	} else {
		cpu.PC++
	}
}

// BEQ Branch on result zero
// Operation:  Branch on Z = 1
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BEQ(mode AddressingMode) {
	log.Println("BEQ called -- adr.mode: ", mode.toString())
	if cpu.getStatusZero() {
		cpu.branch()

	} else {
		cpu.PC++
	}
}

// BRK Force Break
// Operation:  Forced Interrupt PC + 2 toS P toS
// N Z C I D V
// _ _ _ 1 _ _
func (cpu *CPU) BRK(mode AddressingMode) {
	log.Println("BRK called -- do nothing")
}

// CPX Compare Memory and Index X
// Operation:  X - M
// N Z C I D V
// * * * _ _ _
func (cpu *CPU) CPX(mode AddressingMode) {
	log.Println("CPX called -- adr.mode: ", mode.toString())
	if mode == Immidiate {
		aa := cpu.Memory.ReadAbsolute(cpu.PC)
		log.Println("Value to compare with X: ", aa)
		cpu.PC++

		tmp := cpu.X - aa

		cpu.setStatusZero(tmp == 0)
		cpu.setStatusNegative(tmp&0x80 == 128)
		cpu.setStatusCarry(cpu.X >= aa)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// DEX Decrement index X by one
// Operation:  X - 1 -> X
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) DEX(mode AddressingMode) {
	log.Println("DEX called -- adr.mode: ", mode.toString())
	cpu.X--

	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 == 128)
}

// DEY Decrement index Y by one
// Operation:  Y - 1 -> Y
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) DEY(mode AddressingMode) {
	log.Println("DEY called -- adr.mode: ", mode.toString())
	cpu.Y--

	cpu.setStatusZero(cpu.Y == 0)
	cpu.setStatusNegative(cpu.Y&0x80 == 128)
}

// INX Increment Index X by one
// Operation:  X + 1 -> X
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) INX(mode AddressingMode) {
	log.Println("INX called -- adr.mode: ", mode.toString())
	cpu.X++

	log.Println("CPU.X: ", cpu.X)

	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 == 128)
}

// INY Increment Index Y by one
// Operation:  Y + 1 -> Y
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) INY(mode AddressingMode) {
	log.Println("INY called -- adr.mode: ", mode.toString())
	cpu.Y++

	log.Println("CPU.Y: ", cpu.Y)
	cpu.setStatusZero(cpu.Y == 0)
	cpu.setStatusNegative(cpu.Y&0x80 == 128)
}

// LDA Load accumulator with memory
// Operation:  M -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) LDA(mode AddressingMode) {
	log.Println("LDA called -- adr. mode: ", mode.toString())

	hi := cpu.getMemoryValue(mode)

	cpu.A = hi
	log.Printf("Value loaded to CPU register A: %x \n", cpu.A)

	cpu.setStatusZero(cpu.A == 0)
	cpu.setStatusNegative(cpu.A&0x80 == 128)
}

// LDX Load Index X with memory
// Operation:  M -> X
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) LDX(mode AddressingMode) {
	log.Println("LDX called -- adr. mode: ", mode.toString())

	hi := cpu.getMemoryValue(mode)

	cpu.X = hi
	log.Printf("Value loaded to CPU register X: %x \n", cpu.X)
	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 == 128)
}

// LDY Load Index Y with memory
// Operation:  M -> Y
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) LDY(mode AddressingMode) {
	log.Println("LDY called -- adr. mode: ", mode.toString())
	hi := cpu.getMemoryValue(mode)

	cpu.Y = hi
	log.Printf("Value loaded to CPU register Y: %x \n", cpu.Y)
	cpu.setStatusZero(cpu.Y == 0)
	cpu.setStatusNegative(cpu.Y&0x80 == 128)
}

// STA Store accumulator in memory
// Operation:  A -> M
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) STA(mode AddressingMode) {
	log.Println("STA called -- adr. mode: ", mode.toString())

	hi := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	log.Println("CPU register A value: ", cpu.A)

	if mode == ZeroPage {
		log.Println("Setting CPU register A to zero page address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.A)
		return
	}

	if mode == ZeroPageX {
		hi += cpu.X
		log.Println("Setting CPU register A to zero page address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.A)
		return
	}

	if mode == ZeroPageY {
		hi += cpu.Y
		log.Println("Setting CPU register A to zero page address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.A)
		return
	}

	if mode == IndirectIndexedY {
		hiByte := cpu.Memory.ReadZeroPage(hi)
		hiByteNew := cpu.Y + hiByte
		loByte := cpu.Memory.ReadZeroPage(hi + 1)
		if uint16(cpu.Y)+uint16(hiByte) > 255 {
			loByte++
		}
		cpu.Memory.WriteAbsolute(n.ToInt16([]byte{hiByteNew, loByte}), cpu.A)
		return
	}

	if mode == IndexedIndirectX {
		addrHiByteOfLocation := hi + cpu.X
		hiByte := cpu.Memory.ReadZeroPage(addrHiByteOfLocation)
		loByte := cpu.Memory.ReadZeroPage(addrHiByteOfLocation + 1)
		cpu.Memory.WriteAbsolute(n.ToInt16([]byte{hiByte, loByte}), cpu.A)
		return
	}

	lo := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	memLoc := n.ToInt16([]byte{hi, lo})

	if mode == Absolute {
		log.Println("Setting CPU register A to address: ", memLoc)
		cpu.Memory.WriteAbsolute(memLoc, cpu.A)
		return
	}

	if mode == AbsoluteX {
		memLoc += uint16(cpu.X)
		log.Println("Setting CPU register A to address: ", memLoc)
		cpu.Memory.WriteAbsolute(memLoc, cpu.A)
		return
	}

	if mode == AbsoluteY {
		memLoc += uint16(cpu.Y)
		log.Println("Setting CPU register A to address: ", memLoc)
		cpu.Memory.WriteAbsolute(memLoc, cpu.A)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// STX Store Index X in memory
// Operation:  X -> M
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) STX(mode AddressingMode) {
	log.Println("STX called -- adr. mode: ", mode.toString())

	hi := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	// Zeropage only needs a single byte address
	if mode == ZeroPage {
		log.Println("CPU register X value: ", cpu.X)
		log.Println("Setting CPU register X to zero page address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.X)
		return
	}

	if mode == ZeroPageY {
		hi += cpu.Y
		log.Println("CPU register X value: ", cpu.X)
		log.Println("Setting CPU register X to address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.X)
		return
	}

	if mode == Absolute {
		lo := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		log.Println("CPU register X value: ", cpu.X)
		log.Println("Setting CPU register X to address: ", n.ToInt16([]byte{hi, lo}))
		cpu.Memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), cpu.X)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// STY Store Index Y in memory
// Operation:  Y -> M
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) STY(mode AddressingMode) {
	log.Println("STY called -- adr. mode: ", mode.toString())

	hi := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	// Zeropage only needs a single byte address
	if mode == ZeroPage {
		log.Println("CPU register Y value: ", cpu.Y)
		log.Println("Setting CPU register Y to zero page address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.Y)
		return
	}

	if mode == ZeroPageY {
		hi += cpu.Y
		log.Println("CPU register X value: ", cpu.Y)
		log.Println("Setting CPU register X to address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.Y)
		return
	}

	if mode == Absolute {
		lo := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		log.Println("CPU register Y value: ", cpu.Y)
		log.Println("Setting CPU register Y to address: ", n.ToInt16([]byte{hi, lo}))
		cpu.Memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), cpu.Y)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// TAX Transfer accumulator to index X
// Operation:  A -> X
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) TAX(mode AddressingMode) {
	log.Println("TAX called -- adr. mode: ", mode.toString())
	if mode == Implied {
		log.Println("Setting CPU register X to value: ", cpu.A)

		cpu.X = cpu.A

		cpu.setStatusZero(cpu.X == 0)
		cpu.setStatusNegative(cpu.X&0x80 == 128)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// TXA Transfer index X to accumulator
// Operation:  X -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) TXA(mode AddressingMode) {
	log.Println("TXA called -- adr. mode: ", mode.toString())
	if mode == Implied {
		log.Println("Setting CPU register A to value: ", cpu.X)

		cpu.A = cpu.X
		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// TAY Transfer accumulator to index Y
// Operation:  A -> Y
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) TAY(mode AddressingMode) {
	log.Println("TAY called -- adr. mode: ", mode.toString())
	if mode == Implied {
		log.Println("Setting CPU register Y to value: ", cpu.A)

		cpu.Y = cpu.A
		cpu.setStatusZero(cpu.Y == 0)
		cpu.setStatusNegative(cpu.Y&0x80 == 128)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// TYA Transfer index Y to accumulator
// Operation:  Y -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) TYA(mode AddressingMode) {
	log.Println("TYA called -- adr. mode: ", mode.toString())
	if mode == Implied {
		log.Println("Setting CPU register A to value: ", cpu.Y)

		cpu.A = cpu.Y
		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// AND "AND" memory with accumulator
// Operation:  A /\ M -> A
// N Z C I D V
// * * _ _ _
func (cpu *CPU) AND(mode AddressingMode) {
	log.Println("AND called -- adr. mode: ", mode.toString())
	if mode == Immidiate {
		mem := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cpu.A = mem & cpu.A

		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// INC Increment memory by one
// Operation:  M + 1 -> M
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) INC(mode AddressingMode) {
	log.Println("INC called -- adr. mode: ", mode.toString())
	mem := byte(0)
	hi := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	if mode == Absolute {
		lo := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		mem = cpu.Memory.ReadAbsolute(n.ToInt16([]byte{hi, lo}))
		cpu.Memory.WriteAbsolute(n.ToInt16([]byte{hi, lo}), mem+1)
	} else if mode == ZeroPage {
		mem = cpu.Memory.ReadZeroPage(hi)
		cpu.Memory.WriteZeroPage(hi, mem+1)
	} else {
		log.Println("[WARNING] Unsupported addressing mode!")
	}

	cpu.setStatusZero(mem+1 == 0)
	cpu.setStatusNegative(mem+1&0x80 == 128)
}

// JMP Jump to new location
// Operation:  (PC + 1) -> PCL
//             (PC + 2) -> PCH
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) JMP(mode AddressingMode) {
	log.Println("JMP called -- adr. mode: ", mode.toString())

	if mode == Absolute {
		log.Println("Absolute mode")
		hi := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		lo := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cpu.PC = n.ToInt16([]byte{hi, lo})
		return
	}

	if mode == Indirect {
		log.Println("Indirect mode")

		hi := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++
		lo := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		memAdr := n.ToInt16([]byte{hi, lo})
		log.Println("Getting mem value from address: ", memAdr)

		hiByte := cpu.Memory.ReadAbsolute(memAdr)
		loByte := cpu.Memory.ReadAbsolute(memAdr + 1)
		log.Println("Got value: ", loByte)

		cpu.PC = n.ToInt16([]byte{hiByte, loByte})
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// CMP Compare memory and accumulator
// Operation:  A - M
// N Z C I D V
// * * * _ _ _
func (cpu *CPU) CMP(mode AddressingMode) {
	log.Println("CMP called -- adr. mode: ", mode.toString())

	mem := cpu.getMemoryValue(mode)

	log.Println("mem value: ", mem)
	log.Println("CPU.A value: ", cpu.A)
	result := cpu.A - mem
	log.Println("result value:", result)

	cpu.setStatusZero(result == 0)
	cpu.setStatusNegative(result&0x80 == 128)
	cpu.setStatusCarry(cpu.A >= mem)
}

// SEC Set carry flag
// Operation:  1 -> C
// N Z C I D V
// _ _ 1 _ _ _
func (cpu *CPU) SEC(mode AddressingMode) {
	log.Println("SEC called -- adr. mode: ", mode.toString())
	cpu.setStatusCarry(true)
}

// SBC Subtract memory from accumulator with borrow
// Operation:  A - M - C -> A
// N Z C I D V
// * * * _ _ *
// Note:C = Borrow
func (cpu *CPU) SBC(mode AddressingMode) {
	log.Println("SBC called -- adr. mode: ", mode.toString())
	if mode == Immidiate {
		mem := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		result := cpu.A - mem
		if cpu.getStatusCarry() != true {
			result--
		}

		cpu.setStatusCarry(uint16(result) <= 255)

		isOverflow := (((cpu.A ^ result) & 0x80) == 128) && ((cpu.A^mem)&0x80) == 128
		cpu.setStatusOverflow(isOverflow)
		cpu.setStatusZero(result == 0)

		cpu.A = result
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// ROR Rotate one bit right (memory or accumulator)
//               +------------------------------+
//               |                              |
//               |   +-+    +-+-+-+-+-+-+-+-+   |
//  Operation:   +-> |C| -> |7|6|5|4|3|2|1|0| >-+
//                   +-+    +-+-+-+-+-+-+-+-+
// N Z C I D V
// * * * _ _ _
func (cpu *CPU) ROR(mode AddressingMode) {
	log.Println("ROR called -- adr. mode: ", mode.toString())

	previousC := byte(0)
	if cpu.getStatusCarry() {
		previousC = 1
	}

	cpu.setStatusCarry(cpu.A&0x01 == 1)
	cpu.A = cpu.A>>1 | previousC<<7

	cpu.setStatusNegative(cpu.A&0x80 == 128)
	cpu.setStatusZero(cpu.A == 0)
}

// EOR "Exclusive-Or" memory with accumulator
// Operation:  A EOR M -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) EOR(mode AddressingMode) {
	log.Println("EOR called -- adr. mode: ", mode.toString())
	if mode == Immidiate {
		mem := cpu.Memory.ReadAbsolute(cpu.PC)
		cpu.PC++

		cpu.A = mem ^ cpu.A

		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 == 128)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// NOP No operation
// Operation:  No Operation (2 cycles)
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) NOP(mode AddressingMode) {
	log.Println("NOP called -- adr. mode: ", mode.toString())
}

// CPY Compare memory and index Y
// Operation:  Y - M
// N Z C I D V
// * * * _ _ _
func (cpu *CPU) CPY(mode AddressingMode) {
	log.Println("CPY called -- adr. mode: ", mode.toString())
	if mode == Immidiate {
		aa := cpu.Memory.ReadAbsolute(cpu.PC)
		log.Println("Value to compare with Y: ", aa)
		cpu.PC++

		tmp := cpu.Y - aa

		cpu.setStatusZero(tmp == 0)
		cpu.setStatusNegative(tmp&0x80 == 128)
		cpu.setStatusCarry(cpu.Y >= aa)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// BCC Branch on Carry Clear
// Operation:  Branch on C = 0
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BCC(mode AddressingMode) {
	log.Println("BCC called -- adr. mode: ", mode.toString())
	if !cpu.getStatusCarry() {
		cpu.branch()
	} else {
		cpu.PC++
	}
}

// BMI Branch on result minus
// Operation:  Branch on N = 1
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BMI(mode AddressingMode) {
	log.Println("BMI called -- adr. mode: ", mode.toString())
	if cpu.getStatusNegative() {
		cpu.branch()
	} else {
		cpu.PC++
	}
}

// CLC Clear carry flag
// Operation:  0 -> C
// N Z C I D V
// _ _ 0 _ _ _
func (cpu *CPU) CLC(mode AddressingMode) {
	log.Println("CLC called -- adr. mode: ", mode.toString())
	cpu.setStatusCarry(false)
}
