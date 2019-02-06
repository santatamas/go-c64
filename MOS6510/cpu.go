package MOS6510

import (
	"github.com/santatamas/go-c64/RAM"
	"log"
	"reflect"
)

type CPU struct {
	Memory     *RAM.Memory
	A          byte
	Y          byte // low
	X          byte // high
	S          byte
	P          byte
	PC         uint16
	SP         byte
	SP_LOW     uint16
	SP_HIGH    uint16
	instrTypes func(byte) AssemblyInstruction
}

func NewCPU(mem *RAM.Memory) CPU {

	return CPU{
		Memory:     mem,
		SP_LOW:     0x0100,
		SP_HIGH:    0x01FF,
		SP:         0xFF,
		instrTypes: assemblyInstructions(),
	}
}

func getTestCPU() (result CPU) {
	memory := RAM.NewMemory(false)
	return NewCPU(&memory)
}

func (cpu *CPU) ExecuteCycle() bool {
	log.Printf("Current PC address: %x \n", cpu.PC)

	// Fetch first executable instruction code from memory
	instrCode := cpu.Memory.ReadAbsolute(cpu.PC)
	cpu.PC++

	//log.Printf("Next instruction code: %x \n", instrCode)

	// Resolve instruction by instruction code
	instruction := cpu.instrTypes(instrCode)

	cpu.callMethod(instruction)
	return instruction.Type != BRK
}

// slow
func (cpu *CPU) callMethodReflection(instruction AssemblyInstruction) {
	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(instruction.AddressingMode)
	reflect.ValueOf(cpu).MethodByName(instruction.Type.String()).Call(inputs)
}

// faster(?)
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
	case PHA:
		cpu.PHA(instruction.AddressingMode)
	case PHP:
		cpu.PHP(instruction.AddressingMode)
	case PLA:
		cpu.PLA(instruction.AddressingMode)
	case TSX:
		cpu.TSX(instruction.AddressingMode)
	case TXS:
		cpu.TXS(instruction.AddressingMode)
	case BIT:
		cpu.BIT(instruction.AddressingMode)
	case BVC:
		cpu.BVC(instruction.AddressingMode)
	case BVS:
		cpu.BVS(instruction.AddressingMode)
	case PLP:
		cpu.PLP(instruction.AddressingMode)
	case SEI:
		cpu.SEI(instruction.AddressingMode)
	case CLD:
		cpu.CLD(instruction.AddressingMode)
	case ORA:
		cpu.ORA(instruction.AddressingMode)
	case ROL:
		cpu.ROL(instruction.AddressingMode)
	case CLI:
		cpu.CLI(instruction.AddressingMode)
	case DEC:
		cpu.DEC(instruction.AddressingMode)
	case LSR:
		cpu.LSR(instruction.AddressingMode)
	default:
		log.Println("[WARNING] Unimplemented instruction! ", instruction.Type.String())
	}
}
