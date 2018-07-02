package MOS6510

type AddressingMode int

const (
	Implied AddressingMode = iota + 1
	IndexedIndirectX
	IndirectIndexedY
	Indirect
	Absolute
	AbsoluteX
	AbsoluteY
	Immidiate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Accumulator
	Relative
)
