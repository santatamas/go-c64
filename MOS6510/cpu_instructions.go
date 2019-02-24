package MOS6510

import (
	"github.com/santatamas/go-c64/RAM"
	n "github.com/santatamas/go-c64/numeric"
	"log"
)

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
			log.Println("Reading memory from zeropage address: ", hi)
			hi = cpu.Memory.ReadZeroPage(hi)
		}
	case ZeroPageX:
		{
			hi = cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			hi += cpu.X
			hi = cpu.Memory.ReadZeroPage(hi)
		}
	case ZeroPageY:
		{
			hi = cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			hi += cpu.Y
			hi = cpu.Memory.ReadZeroPage(hi)
		}
	case Absolute:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			log.Println("Reading memory from absolute address: ", n.ToInt16([]byte{adr_hi, adr_lo}))

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
	case AbsoluteY:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			memAdr := n.ToInt16([]byte{adr_hi, adr_lo})
			memAdr += uint16(cpu.Y)

			hi = cpu.Memory.ReadAbsolute(memAdr)
		}
	case IndirectIndexedY:
		{
			hi = cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			hiByte := cpu.Memory.ReadZeroPage(hi)
			hiByteNew := cpu.Y + hiByte
			loByte := cpu.Memory.ReadZeroPage(hi + 1)
			if uint16(cpu.Y)+uint16(hiByte) > 255 {
				loByte++
			}
			hi = cpu.Memory.ReadAbsolute(n.ToInt16([]byte{hiByteNew, loByte}))
		}
	case IndexedIndirectX:
		{
			hi = cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			addrHiByteOfLocation := hi + cpu.X
			hiByte := cpu.Memory.ReadZeroPage(addrHiByteOfLocation)
			loByte := cpu.Memory.ReadZeroPage(addrHiByteOfLocation + 1)
			hi = cpu.Memory.ReadAbsolute(n.ToInt16([]byte{hiByte, loByte}))
		}
	default:
		log.Println("[WARNING] Unsupported addressing mode!")
	}

	return hi
}

func (cpu *CPU) branch() {
	hi := int8(cpu.Memory.ReadAbsolute(cpu.PC))
	cpu.PC++

	log.Println("Moving PC address by: ", hi)

	cpu.PC += uint16(hi)
}

// RTS Return from Subroutine
// Operation:  PC fromS, PC + 1 -> PC
//  N Z C I D V
//  _ _ _ _ _ _
func (cpu *CPU) RTS(mode AddressingMode) {
	log.Println("RTS called -- adr.mode: ", mode.String())
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
	log.Println("ADC called -- adr. mode: ", mode.String())

	hi := cpu.getMemoryValue(mode)

	result := cpu.A + hi

	if cpu.getStatusCarry() {
		result++
	}

	cpu.setStatusCarry(uint16(cpu.A)+uint16(hi) > 255)

	isOverflow := (((cpu.A ^ hi) & 0x80) == 0) && (((cpu.A ^ result) & 0x80) != 0)
	cpu.setStatusOverflow(isOverflow)
	cpu.setStatusZero(result == 0)

	cpu.A = result
	cpu.setStatusNegative((cpu.A & 0x80) != 0)
}

// BCS Branch on carry set
// Operation:  Branch on C = 1
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BCS(mode AddressingMode) {
	log.Println("BCS called -- adr.mode: ", mode.String())
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
	log.Println("BPL called -- adr.mode: ", mode.String())
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
	log.Println("ASL called -- adr. mode: ", mode.String())

	switch mode {
	case Accumulator:
		{
			cpu.setStatusCarry(cpu.A&0x80 != 0)

			cpu.A = cpu.A << 1

			cpu.setStatusNegative((cpu.A & 0x80) != 0)
			cpu.setStatusZero(cpu.A == 0)
		}
	case ZeroPage:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			mem := cpu.Memory.ReadZeroPage(memAdr)

			cpu.setStatusCarry(mem&0x80 != 0)
			result := mem << 1

			cpu.setStatusNegative((result & 0x80) != 0)
			cpu.setStatusZero(result == 0)

			cpu.Memory.WriteZeroPage(memAdr, result)
		}
	case ZeroPageX:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			memAdr += cpu.X
			mem := cpu.Memory.ReadZeroPage(memAdr)

			cpu.setStatusCarry(mem&0x80 != 0)
			result := mem << 1

			cpu.setStatusNegative((result & 0x80) != 0)
			cpu.setStatusZero(result == 0)

			cpu.Memory.WriteZeroPage(memAdr, result)
		}
	case Absolute:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			mem := cpu.Memory.ReadAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}))
			cpu.setStatusCarry(mem&0x80 != 0)

			result := mem << 1

			cpu.setStatusNegative((result & 0x80) != 0)
			cpu.setStatusZero(result == 0)

			cpu.Memory.WriteAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}), result)
		}
	case AbsoluteX:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			memAdr := n.ToInt16([]byte{adr_hi, adr_lo})
			memAdr += uint16(cpu.X)

			mem := cpu.Memory.ReadAbsolute(memAdr)
			cpu.setStatusCarry(mem&0x80 != 0)

			result := mem << 1

			cpu.setStatusNegative((result & 0x80) != 0)
			cpu.setStatusZero(result == 0)

			cpu.Memory.WriteAbsolute(memAdr, result)
		}
	default:
		log.Println("[WARNING] Unsupported addressing mode: ", mode.String())
	}
}

// JSR Jump to new location saving return address
// Operation:  PC + 2 toS, (PC + 1) -> PCL
//                         (PC + 2) -> PCH
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) JSR(mode AddressingMode) {
	log.Println("JSR called -- adr.mode: ", mode.String())
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
	log.Println("BNE called -- adr.mode: ", mode.String())
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
	log.Println("BEQ called -- adr.mode: ", mode.String())
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
	log.Println("BRK called")
	cpu.stackPush(n.GetHI(cpu.PC + 1))
	cpu.stackPush(n.GetLO(cpu.PC + 1))

	cpu.stackPush(cpu.S) // TODO: clear B flag before pushing to stack

	lo := cpu.Memory.ReadAbsolute(RAM.IRQ_VECTOR_ADDR_LO)
	hi := cpu.Memory.ReadAbsolute(RAM.IRQ_VECTOR_ADDR_HI)

	cpu.PC = n.ToInt16([]byte{lo, hi})
	cpu.setStatusIRQ(true)
	cpu.setStatusBRK(true)
}

// SED Set decimal mode
// Operation:  1 -> D
// N Z C I D V
// _ _ _ _ 1 _
func (cpu *CPU) SED(mode AddressingMode) {
	log.Println("SED called")

	cpu.setStatusDecimal(true)
}

// CLV Clear overflow flag
// Operation: 0 -> V
// N Z C I D V
// _ _ _ _ _ 0
func (cpu *CPU) CLV(mode AddressingMode) {
	log.Println("CLV called")

	cpu.setStatusOverflow(false)
}

// RTI Return from interrupt
// Operation:  P fromS PC fromS
// N Z C I D V
// * * * * * *
func (cpu *CPU) RTI(mode AddressingMode) {
	log.Println("RTI called")

	cpu.S, _ = cpu.stackPop()
	lo, _ := cpu.stackPop()
	hi, _ := cpu.stackPop()

	cpu.PC = n.ToInt16([]byte{lo, hi})
}

// CPX Compare Memory and Index X
// Operation:  X - M
// N Z C I D V
// * * * _ _ _
func (cpu *CPU) CPX(mode AddressingMode) {
	log.Println("CPX called -- adr.mode: ", mode.String())

	mem := cpu.getMemoryValue(mode)

	tmp := cpu.X - mem

	cpu.setStatusZero(tmp == 0)
	cpu.setStatusNegative(tmp&0x80 != 0)
	cpu.setStatusCarry(cpu.X >= mem)
}

// DEX Decrement index X by one
// Operation:  X - 1 -> X
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) DEX(mode AddressingMode) {
	log.Println("DEX called -- adr.mode: ", mode.String())
	cpu.X--

	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 != 0)
}

// DEX Decrement Memory by one
// Operation:  M - 1 -> M
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) DEC(mode AddressingMode) {
	log.Println("DEC called -- adr.mode: ", mode.String())

	hi := byte(0)

	switch mode {
	case ZeroPage:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			hi = cpu.Memory.ReadZeroPage(memAdr)
			hi--
			cpu.Memory.WriteZeroPage(memAdr, hi)
		}
	case ZeroPageX:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			memAdr += cpu.X
			hi = cpu.Memory.ReadZeroPage(memAdr)
			hi--
			cpu.Memory.WriteZeroPage(memAdr, hi)
		}
	case Absolute:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			hi = cpu.Memory.ReadAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}))
			hi--
			cpu.Memory.WriteAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}), hi)
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
			hi--
			cpu.Memory.WriteAbsolute(memAdr, hi)
		}
	default:
		log.Println("[WARNING] Unsupported addressing mode: ", mode.String())
	}

	cpu.setStatusZero(hi == 0)
	cpu.setStatusNegative(hi&0x80 != 0)
}

// DEY Decrement index Y by one
// Operation:  Y - 1 -> Y
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) DEY(mode AddressingMode) {
	log.Println("DEY called -- adr.mode: ", mode.String())
	cpu.Y--

	cpu.setStatusZero(cpu.Y == 0)
	cpu.setStatusNegative(cpu.Y&0x80 != 0)
}

// INX Increment Index X by one
// Operation:  X + 1 -> X
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) INX(mode AddressingMode) {
	log.Println("INX called -- adr.mode: ", mode.String())
	cpu.X++

	log.Println("CPU.X: ", cpu.X)

	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 != 0)
}

// INY Increment Index Y by one
// Operation:  Y + 1 -> Y
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) INY(mode AddressingMode) {
	log.Println("INY called -- adr.mode: ", mode.String())
	cpu.Y++

	log.Println("CPU.Y: ", cpu.Y)
	cpu.setStatusZero(cpu.Y == 0)
	cpu.setStatusNegative(cpu.Y&0x80 != 0)
}

// LDA Load accumulator with memory
// Operation:  M -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) LDA(mode AddressingMode) {
	log.Println("LDA called -- adr. mode: ", mode.String())

	hi := cpu.getMemoryValue(mode)

	cpu.A = hi
	log.Printf("Value loaded to CPU register A: %x \n", cpu.A)

	cpu.setStatusZero(cpu.A == 0)
	cpu.setStatusNegative(cpu.A&0x80 != 0)
}

// LDX Load Index X with memory
// Operation:  M -> X
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) LDX(mode AddressingMode) {
	log.Println("LDX called -- adr. mode: ", mode.String())

	hi := cpu.getMemoryValue(mode)

	cpu.X = hi
	log.Printf("Value loaded to CPU register X: %x \n", cpu.X)
	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 != 0)
}

// LDY Load Index Y with memory
// Operation:  M -> Y
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) LDY(mode AddressingMode) {
	log.Println("LDY called -- adr. mode: ", mode.String())
	hi := cpu.getMemoryValue(mode)

	cpu.Y = hi
	log.Printf("Value loaded to CPU register Y: %x \n", cpu.Y)
	cpu.setStatusZero(cpu.Y == 0)
	cpu.setStatusNegative(cpu.Y&0x80 != 0)
}

// STA Store accumulator in memory
// Operation:  A -> M
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) STA(mode AddressingMode) {
	log.Println("STA called -- adr. mode: ", mode.String())

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
	log.Println("STX called -- adr. mode: ", mode.String())

	hi := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	// Zeropage only needs a single byte address
	if mode == ZeroPage {
		log.Println("CPU register X value: ", cpu.X)
		log.Println("Setting CPU register X to zero page address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.X)
		return
	}

	if mode == ZeroPageX {
		hi += cpu.X
		log.Println("CPU register X value: ", cpu.X)
		log.Println("Setting CPU register X to address: ", hi)
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
	log.Println("STY called -- adr. mode: ", mode.String())

	hi := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	// Zeropage only needs a single byte address
	if mode == ZeroPage {
		log.Println("CPU register Y value: ", cpu.Y)
		log.Println("Setting CPU register Y to zero page address: ", hi)
		cpu.Memory.WriteZeroPage(hi, cpu.Y)
		return
	}

	if mode == ZeroPageX {
		hi += cpu.X
		log.Println("CPU register Y value: ", cpu.Y)
		log.Println("Setting CPU register Y to address: ", hi)
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
	log.Println("TAX called -- adr. mode: ", mode.String())
	if mode == Implied {
		log.Println("Setting CPU register X to value: ", cpu.A)

		cpu.X = cpu.A

		cpu.setStatusZero(cpu.X == 0)
		cpu.setStatusNegative(cpu.X&0x80 != 0)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// TXA Transfer index X to accumulator
// Operation:  X -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) TXA(mode AddressingMode) {
	log.Println("TXA called -- adr. mode: ", mode.String())
	if mode == Implied {
		log.Println("Setting CPU register A to value: ", cpu.X)

		cpu.A = cpu.X
		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 != 0)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// TAY Transfer accumulator to index Y
// Operation:  A -> Y
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) TAY(mode AddressingMode) {
	log.Println("TAY called -- adr. mode: ", mode.String())
	if mode == Implied {
		log.Println("Setting CPU register Y to value: ", cpu.A)

		cpu.Y = cpu.A
		cpu.setStatusZero(cpu.Y == 0)
		cpu.setStatusNegative(cpu.Y&0x80 != 0)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// TYA Transfer index Y to accumulator
// Operation:  Y -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) TYA(mode AddressingMode) {
	log.Println("TYA called -- adr. mode: ", mode.String())
	if mode == Implied {
		log.Println("Setting CPU register A to value: ", cpu.Y)

		cpu.A = cpu.Y
		cpu.setStatusZero(cpu.A == 0)
		cpu.setStatusNegative(cpu.A&0x80 != 0)
		return
	}

	log.Println("[WARNING] Unsupported addressing mode!")
}

// AND "AND" memory with accumulator
// Operation:  A /\ M -> A
// N Z C I D V
// * * _ _ _
func (cpu *CPU) AND(mode AddressingMode) {
	log.Println("AND called -- adr. mode: ", mode.String())

	mem := cpu.getMemoryValue(mode)

	cpu.A = mem & cpu.A

	cpu.setStatusZero(cpu.A == 0)
	cpu.setStatusNegative(cpu.A&0x80 != 0)
}

// INC Increment memory by one
// Operation:  M + 1 -> M
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) INC(mode AddressingMode) {
	log.Println("INC called -- adr. mode: ", mode.String())
	hi := byte(0)

	switch mode {
	case ZeroPage:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			hi = cpu.Memory.ReadZeroPage(memAdr)
			hi++
			cpu.Memory.WriteZeroPage(memAdr, hi)
		}
	case ZeroPageX:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			memAdr += cpu.X
			hi = cpu.Memory.ReadZeroPage(memAdr)
			hi++
			cpu.Memory.WriteZeroPage(memAdr, hi)
		}
	case Absolute:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			hi = cpu.Memory.ReadAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}))
			hi++
			cpu.Memory.WriteAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}), hi)
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
			hi++
			cpu.Memory.WriteAbsolute(memAdr, hi)
		}
	default:
		log.Println("[WARNING] Unsupported addressing mode: ", mode.String())
	}

	cpu.setStatusZero(hi == 0)
	cpu.setStatusNegative((hi)&0x80 != 0)
}

// JMP Jump to new location
// Operation:  (PC + 1) -> PCL
//             (PC + 2) -> PCH
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) JMP(mode AddressingMode) {
	log.Println("JMP called -- adr. mode: ", mode.String())

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
	log.Println("CMP called -- adr. mode: ", mode.String())

	mem := cpu.getMemoryValue(mode)

	log.Printf("mem value: %x", mem)
	log.Printf("CPU.A value: %x", cpu.A)
	result := cpu.A - mem
	log.Printf("result value: %x", result)

	cpu.setStatusZero(result == 0)
	cpu.setStatusNegative(result&0x80 != 0)
	cpu.setStatusCarry(cpu.A >= mem)
}

// SEC Set carry flag
// Operation:  1 -> C
// N Z C I D V
// _ _ 1 _ _ _
func (cpu *CPU) SEC(mode AddressingMode) {
	log.Println("SEC called -- adr. mode: ", mode.String())
	cpu.setStatusCarry(true)
}

// SBC Subtract memory from accumulator with borrow
// Operation:  A - M - C -> A
// N Z C I D V
// * * * _ _ *
// Note:C = Borrow
func (cpu *CPU) SBC(mode AddressingMode) {
	log.Println("SBC called -- adr. mode: ", mode.String())

	mem := cpu.getMemoryValue(mode)
	log.Println("SBC - memory value: ", mem)
	log.Println("SBC - register A value: ", cpu.A)

	result := cpu.A - mem

	memInt := int(mem)
	cpuAInt := int(cpu.A)
	intResult := cpuAInt - memInt
	log.Println("SBC - intResult: ", intResult)

	if cpu.getStatusCarry() {
		result--
		intResult--
	}

	cpu.setStatusCarry(intResult < 0)
	log.Println("SBC - carry value: ", intResult < 0)

	isOverflow := (((cpu.A ^ result) & 0x80) != 0) && (((cpu.A ^ mem) & 0x80) != 0)
	log.Println("SBC - isOverFlow: ", isOverflow)
	cpu.setStatusOverflow(isOverflow)
	cpu.setStatusZero(result == 0)

	log.Println("Setting CPU register A to: ", result)
	cpu.A = result
	cpu.setStatusNegative((cpu.A & 0x80) != 0)
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
	log.Println("ROR called -- adr. mode: ", mode.String())

	previousC := byte(0)
	if cpu.getStatusCarry() {
		previousC = 1
	}

	switch mode {
	case Accumulator:
		{
			cpu.setStatusCarry(cpu.A&0x01 != 0)
			cpu.A = cpu.A>>1 | previousC<<7

			cpu.setStatusNegative(cpu.A&0x80 != 0)
			cpu.setStatusZero(cpu.A == 0)
		}
	case ZeroPage:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			mem := cpu.Memory.ReadZeroPage(memAdr)
			cpu.setStatusCarry(mem&0x01 != 0)
			mem = mem>>1 | previousC<<7
			cpu.setStatusZero(mem == 0)
			cpu.setStatusNegative(mem&0x80 != 0)
			cpu.Memory.WriteZeroPage(memAdr, mem)
		}
	case ZeroPageX:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			memAdr += cpu.X
			mem := cpu.Memory.ReadZeroPage(memAdr)
			cpu.setStatusCarry(mem&0x01 != 0)
			mem = mem>>1 | previousC<<7
			cpu.setStatusZero(mem == 0)
			cpu.setStatusNegative(mem&0x80 != 0)
			cpu.Memory.WriteZeroPage(memAdr, mem)
		}
	case Absolute:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			mem := cpu.Memory.ReadAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}))
			cpu.setStatusCarry(mem&0x01 != 0)
			mem = mem>>1 | previousC<<7
			cpu.setStatusZero(mem == 0)
			cpu.setStatusNegative(mem&0x80 != 0)
			cpu.Memory.WriteAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}), mem)
		}
	case AbsoluteX:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			memAdr := n.ToInt16([]byte{adr_hi, adr_lo})
			memAdr += uint16(cpu.X)

			mem := cpu.Memory.ReadAbsolute(memAdr)
			cpu.setStatusCarry(mem&0x01 != 0)
			mem = mem>>1 | previousC<<7
			cpu.setStatusZero(mem == 0)
			cpu.setStatusNegative(mem&0x80 != 0)
			cpu.Memory.WriteAbsolute(memAdr, mem)
		}
	default:
		log.Println("[WARNING] Unsupported addressing mode: ", mode.String())
	}
}

// ROL Rotate one bit left (memory or accumulator)
//              +------------------------------+
//              |         M or A               |
//              |   +-+-+-+-+-+-+-+-+    +-+   |
// Operation:   +-< |7|6|5|4|3|2|1|0| <- |C| <-+
// 	                +-+-+-+-+-+-+-+-+    +-+
// N Z C I D V
// * * * _ _ _
func (cpu *CPU) ROL(mode AddressingMode) {
	log.Println("ROL called -- adr. mode: ", mode.String())

	previousC := byte(0)
	if cpu.getStatusCarry() {
		previousC = 1
	}

	switch mode {
	case Accumulator:
		{
			cpu.setStatusCarry(cpu.A&0x80 != 0)
			cpu.A = cpu.A<<1 | previousC

			cpu.setStatusNegative(cpu.A&0x80 != 0)
			cpu.setStatusZero(cpu.A == 0)
		}
	case ZeroPage:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			mem := cpu.Memory.ReadZeroPage(memAdr)
			cpu.setStatusCarry(mem&0x80 != 0)
			mem = mem<<1 | previousC
			cpu.setStatusZero(mem == 0)
			cpu.setStatusNegative(mem&0x80 != 0)
			cpu.Memory.WriteZeroPage(memAdr, mem)
		}
	case ZeroPageX:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			memAdr += cpu.X
			mem := cpu.Memory.ReadZeroPage(memAdr)
			cpu.setStatusCarry(mem&0x80 != 0)
			mem = mem<<1 | previousC
			cpu.setStatusZero(mem == 0)
			cpu.setStatusNegative(mem&0x80 != 0)
			cpu.Memory.WriteZeroPage(memAdr, mem)
		}
	case Absolute:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			mem := cpu.Memory.ReadAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}))
			cpu.setStatusCarry(mem&0x80 != 0)
			mem = mem<<1 | previousC
			cpu.setStatusZero(mem == 0)
			cpu.setStatusNegative(mem&0x80 != 0)
			cpu.Memory.WriteAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}), mem)
		}
	case AbsoluteX:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			memAdr := n.ToInt16([]byte{adr_hi, adr_lo})
			memAdr += uint16(cpu.X)

			mem := cpu.Memory.ReadAbsolute(memAdr)
			cpu.setStatusCarry(mem&0x80 != 0)
			mem = mem<<1 | previousC
			cpu.setStatusZero(mem == 0)
			cpu.setStatusNegative(mem&0x80 != 0)
			cpu.Memory.WriteAbsolute(memAdr, mem)
		}
	default:
		log.Println("[WARNING] Unsupported addressing mode: ", mode.String())
	}
}

// EOR "Exclusive-Or" memory with accumulator
// Operation:  A EOR M -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) EOR(mode AddressingMode) {
	log.Println("EOR called -- adr. mode: ", mode.String())

	mem := cpu.getMemoryValue(mode)

	cpu.A = mem ^ cpu.A

	cpu.setStatusZero(cpu.A == 0)
	cpu.setStatusNegative(cpu.A&0x80 != 0)
}

// NOP No operation
// Operation:  No Operation (2 cycles)
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) NOP(mode AddressingMode) {
	log.Println("NOP called -- adr. mode: ", mode.String())
}

// CPY Compare memory and index Y
// Operation:  Y - M
// N Z C I D V
// * * * _ _ _
func (cpu *CPU) CPY(mode AddressingMode) {
	log.Println("CPY called -- adr. mode: ", mode.String())

	mem := cpu.getMemoryValue(mode)

	tmp := cpu.Y - mem

	cpu.setStatusZero(tmp == 0)
	cpu.setStatusNegative(tmp&0x80 != 0)
	cpu.setStatusCarry(cpu.Y >= mem)
}

// BCC Branch on Carry Clear
// Operation:  Branch on C = 0
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BCC(mode AddressingMode) {
	log.Println("BCC called -- adr. mode: ", mode.String())
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
	log.Println("BMI called -- adr. mode: ", mode.String())
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
	log.Println("CLC called -- adr. mode: ", mode.String())
	cpu.setStatusCarry(false)
}

// PHA Push accumulator on stack
// Operation:  A toS
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) PHA(mode AddressingMode) {
	log.Println("PHA called -- adr. mode: ", mode.String())
	cpu.stackPush(cpu.A)
}

// PLA Pull accumulator from stack
// Operation:  A fromS
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) PLA(mode AddressingMode) {
	log.Println("PLA called -- adr. mode: ", mode.String())
	cpu.A, _ = cpu.stackPop()
	cpu.setStatusNegative(n.IsNegative(cpu.A))
	cpu.setStatusZero(cpu.A == 0x0)
}

// PHP Push processor status on stack
// Operation:  P toS
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) PHP(mode AddressingMode) {
	log.Println("PHP called -- adr. mode: ", mode.String())
	cpu.stackPush(cpu.S)
}

// PLP Pull processor status from stack
// Operation:  P fromS
// From Stack
// _ _ _ _ _ _
func (cpu *CPU) PLP(mode AddressingMode) {
	log.Println("PLP called -- adr. mode: ", mode.String())
	stackValue, _ := cpu.stackPop()
	cpu.setStatusWithoutB(stackValue)
}

// TSX Transfer stack pointer to index X
// Operation:  S -> X
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) TSX(mode AddressingMode) {
	log.Println("TSX called -- adr. mode: ", mode.String())
	cpu.X = cpu.SP

	cpu.setStatusZero(cpu.X == 0)
	cpu.setStatusNegative(cpu.X&0x80 != 0)
}

// TXS Transfer index X to stack pointer
// Operation:  X -> S
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) TXS(mode AddressingMode) {
	log.Println("TXS called -- adr. mode: ", mode.String())
	cpu.SP = cpu.X
}

// BIT Test bits in memory with accumulator
// Operation:  A /\ M, M7 -> N, M6 -> V
// Bit 6 and 7 are transferred to the status register.
// If the result of A /\ M is zero then Z = 1, otherwise Z = 0
// N  Z C I D V
// M7 * _ _ _ M6
func (cpu *CPU) BIT(mode AddressingMode) {
	log.Println("BIT called -- adr. mode: ", mode.String())
	mem := cpu.getMemoryValue(mode)

	cpu.setStatusNegative(mem&0x80 != 0)
	cpu.setStatusOverflow(mem&0x40 != 0)
	cpu.setStatusZero(mem&cpu.A == 0)
}

// BVC Branch on overflow clear
// Operation:  Branch on V = 0
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BVC(mode AddressingMode) {
	log.Println("BVC called -- adr.mode: ", mode.String())
	if !cpu.getStatusOverflow() {
		cpu.branch()

	} else {
		cpu.PC++
	}
}

// BVS Branch on overflow set
// Operation:  Branch on V = 1
// N Z C I D V
// _ _ _ _ _ _
func (cpu *CPU) BVS(mode AddressingMode) {
	log.Println("BVS called -- adr.mode: ", mode.String())
	if cpu.getStatusOverflow() {
		cpu.branch()

	} else {
		cpu.PC++
	}
}

// SEI Set interrupt disable status
// Operation:  1 -> I
// N Z C I D V
// _ _ _ 1 _ _
func (cpu *CPU) SEI(mode AddressingMode) {
	log.Println("SEI called -- adr.mode: ", mode.String())
	cpu.setStatusIRQ(true)
}

// CLD Clear decimal mode
// Operation:  0 -> D
// N A C I D V
// _ _ _ _ 0 _
func (cpu *CPU) CLD(mode AddressingMode) {
	log.Println("CLD called -- adr.mode: ", mode.String())
	cpu.setStatusDecimal(false)
}

// CLI Clear interrupt disable bit
// Operation: 0 -> I
// N Z C I D V
// _ _ _ 0 _ _
func (cpu *CPU) CLI(mode AddressingMode) {
	log.Println("CLI called -- adr.mode: ", mode.String())
	cpu.setStatusIRQ(false)
}

// ORA "OR" memory with accumulator
// Operation: A V M -> A
// N Z C I D V
// * * _ _ _ _
func (cpu *CPU) ORA(mode AddressingMode) {
	log.Println("ORA called -- adr. mode: ", mode.String())

	mem := cpu.getMemoryValue(mode)

	cpu.A = cpu.A | mem

	cpu.setStatusZero(cpu.A == 0)
	cpu.setStatusNegative(cpu.A&0x80 != 0)
}

// LSR Shift right one bit (memory or accumulator)
//                   +-+-+-+-+-+-+-+-+
//  Operation:  0 -> |7|6|5|4|3|2|1|0| -> C
//                   +-+-+-+-+-+-+-+-+
//  N Z C I D V
//  0 * * _ _ _
func (cpu *CPU) LSR(mode AddressingMode) {
	log.Println("LSR called -- adr. mode: ", mode.String())

	switch mode {
	case Accumulator:
		{
			cpu.setStatusCarry(cpu.A&0x01 != 0)
			cpu.A = cpu.A >> 1

			cpu.setStatusZero(cpu.A == 0)
		}
	case ZeroPage:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			mem := cpu.Memory.ReadZeroPage(memAdr)
			cpu.setStatusCarry(mem&0x01 != 0)
			mem = mem >> 1
			cpu.setStatusZero(mem == 0)
			cpu.Memory.WriteZeroPage(memAdr, mem)
		}
	case ZeroPageX:
		{
			memAdr := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			memAdr += cpu.X
			mem := cpu.Memory.ReadZeroPage(memAdr)
			cpu.setStatusCarry(mem&0x01 != 0)
			mem = mem >> 1
			cpu.setStatusZero(mem == 0)
			cpu.Memory.WriteZeroPage(memAdr, mem)
		}
	case Absolute:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			mem := cpu.Memory.ReadAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}))
			cpu.setStatusCarry(mem&0x01 != 0)
			mem = mem >> 1
			cpu.setStatusZero(mem == 0)
			cpu.Memory.WriteAbsolute(n.ToInt16([]byte{adr_hi, adr_lo}), mem)
		}
	case AbsoluteX:
		{
			adr_hi := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++
			adr_lo := cpu.Memory.ReadAbsolute(cpu.PC)
			cpu.PC++

			memAdr := n.ToInt16([]byte{adr_hi, adr_lo})
			memAdr += uint16(cpu.X)

			mem := cpu.Memory.ReadAbsolute(memAdr)
			cpu.setStatusCarry(mem&0x01 != 0)
			mem = mem >> 1
			cpu.setStatusZero(mem == 0)
			cpu.Memory.WriteAbsolute(memAdr, mem)
		}
	default:
		log.Println("[WARNING] Unsupported addressing mode: ", mode.String())
	}

	cpu.setStatusNegative(false)
}
