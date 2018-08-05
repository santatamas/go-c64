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

func (mode AddressingMode) toString() string {
	switch mode {
	case Implied:
		return "Implied"
	case IndexedIndirectX:
		return "IndexedIndirectX"
	case IndirectIndexedY:
		return "IndirectIndexedY"
	case Indirect:
		return "Indirect"
	case Absolute:
		return "Absolute"
	case AbsoluteX:
		return "AbsoluteX"
	case AbsoluteY:
		return "AbsoluteY"
	case Immidiate:
		return "Immidiate"
	case ZeroPage:
		return "ZeroPage"
	case ZeroPageX:
		return "ZeroPageX"
	case ZeroPageY:
		return "ZeroPageY"
	case Accumulator:
		return "Accumulator"
	case Relative:
		return "Relative"
	}

	return "Unknown!"
}
