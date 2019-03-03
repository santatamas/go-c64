package MOS6510

type CPUState struct {
	A               byte
	Y               byte // low
	X               byte // high
	S               byte
	P               byte
	PC              uint16
	SP              byte
	SP_LOW          uint16
	SP_HIGH         uint16
	InstructionType byte
	InstructionName string
}

func (cpu *CPU) GetState() CPUState {

	instrCode := cpu.Memory.ReadAbsolute(cpu.PC)
	instruction := cpu.instrTypes(instrCode)

	return CPUState{
		A:               cpu.A,
		Y:               cpu.Y,
		X:               cpu.X,
		S:               cpu.S,
		P:               cpu.P,
		PC:              cpu.PC,
		SP:              cpu.SP,
		SP_LOW:          cpu.SP_LOW,
		InstructionType: instrCode,
		InstructionName: instruction.Type.String(),
	}
}
