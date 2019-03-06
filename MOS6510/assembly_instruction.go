package MOS6510

//import "log"

type AssemblyInstruction struct {
	Type           AssemblyInstructionType
	AddressingMode AddressingMode
	Cycles         byte
}

func assemblyInstructions() func(byte) AssemblyInstruction {

	innerMap := map[byte]AssemblyInstruction{
		0:   AssemblyInstruction{BRK, Implied, 7},          // 00 - BRK
		1:   AssemblyInstruction{ORA, IndexedIndirectX, 6}, // 01 - ORA - (Indirect,X}
		2:   AssemblyInstruction{UNDEF, Implied, 0},        // 02 - Future Expansion
		3:   AssemblyInstruction{UNDEF, Implied, 0},        // 03 - Future Expansion
		4:   AssemblyInstruction{UNDEF, Implied, 0},        // 04 - Future Expansion
		5:   AssemblyInstruction{ORA, ZeroPage, 3},         // 05 - ORA - Zero Page
		6:   AssemblyInstruction{ASL, ZeroPage, 5},         // 06 - ASL - Zero Page
		7:   AssemblyInstruction{UNDEF, Implied, 0},        // 07 - Future Expansion
		8:   AssemblyInstruction{PHP, Implied, 3},          // 08 - PHP
		9:   AssemblyInstruction{ORA, Immidiate, 2},        // 09 - ORA - Immediate
		10:  AssemblyInstruction{ASL, Accumulator, 2},      // 0A - ASL - Accumulator
		11:  AssemblyInstruction{UNDEF, Implied, 0},        // 0B - Future Expansion
		12:  AssemblyInstruction{UNDEF, Implied, 0},        // 0C - Future Expansion
		13:  AssemblyInstruction{ORA, Absolute, 4},         // 0D - ORA - Absolute
		14:  AssemblyInstruction{ASL, Absolute, 6},         // 0E - ASL - Absolute
		15:  AssemblyInstruction{UNDEF, Implied, 0},        // 0F - Future Expansion
		16:  AssemblyInstruction{BPL, Relative, 2},         // 10 - BPL
		17:  AssemblyInstruction{ORA, IndirectIndexedY, 5}, // 11 - ORA - (Indirect,Y}
		18:  AssemblyInstruction{UNDEF, Implied, 0},        // 12 - Future Expansion
		19:  AssemblyInstruction{UNDEF, Implied, 0},        // 13 - Future Expansion
		20:  AssemblyInstruction{UNDEF, Implied, 0},        // 14 - Future Expansion
		21:  AssemblyInstruction{ORA, ZeroPageX, 4},        // 15 - ORA - Zero Page,X
		22:  AssemblyInstruction{ASL, ZeroPageX, 6},        // 16 - ASL - Zero Page,X
		23:  AssemblyInstruction{UNDEF, Implied, 0},        // 17 - Future Expansion
		24:  AssemblyInstruction{CLC, Implied, 2},          // 18 - CLC
		25:  AssemblyInstruction{ORA, AbsoluteY, 4},        // 19 - ORA - Absolute,Y
		26:  AssemblyInstruction{UNDEF, Implied, 0},        // 1A - Future Expansion
		27:  AssemblyInstruction{UNDEF, Implied, 0},        // 1B - Future Expansion
		28:  AssemblyInstruction{UNDEF, Implied, 0},        // 1C - Future Expansion
		29:  AssemblyInstruction{ORA, AbsoluteX, 4},        // 1D - ORA - Absolute,X
		30:  AssemblyInstruction{ASL, AbsoluteX, 7},        // 1E - ASL - Absolute,X
		31:  AssemblyInstruction{UNDEF, Implied, 0},        // 1F - Future Expansion
		32:  AssemblyInstruction{JSR, Absolute, 6},         // 20 - JSR
		33:  AssemblyInstruction{AND, IndexedIndirectX, 6}, // 21 - AND - (Indirect,X}
		34:  AssemblyInstruction{UNDEF, Implied, 0},        // 22 - Future Expansion
		35:  AssemblyInstruction{UNDEF, Implied, 0},        // 23 - Future Expansion
		36:  AssemblyInstruction{BIT, ZeroPage, 3},         // 24 - BIT - Zero Page
		37:  AssemblyInstruction{AND, ZeroPage, 3},         // 25 - AND - Zero Page
		38:  AssemblyInstruction{ROL, ZeroPage, 5},         // 26 - ROL - Zero Page
		39:  AssemblyInstruction{UNDEF, Implied, 0},        // 27 - Future Expansion
		40:  AssemblyInstruction{PLP, Implied, 4},          // 28 - PLP
		41:  AssemblyInstruction{AND, Immidiate, 2},        // 29 - AND - Immediate
		42:  AssemblyInstruction{ROL, Accumulator, 2},      // 2A - ROL - Accumulator
		43:  AssemblyInstruction{UNDEF, Implied, 0},        // 2B - Future Expansion
		44:  AssemblyInstruction{BIT, Absolute, 4},         // 2C - BIT - Absolute
		45:  AssemblyInstruction{AND, Absolute, 4},         // 2D - AND - Absolute
		46:  AssemblyInstruction{ROL, Absolute, 6},         // 2E - ROL - Absolute
		47:  AssemblyInstruction{UNDEF, Implied, 0},        // 2F - Future Expansion
		48:  AssemblyInstruction{BMI, Relative, 2},         // 30 - BMI
		49:  AssemblyInstruction{AND, IndirectIndexedY, 5}, // 31 - AND - (Indirect,Y}
		50:  AssemblyInstruction{UNDEF, Implied, 0},        // 32 - Future Expansion
		51:  AssemblyInstruction{UNDEF, Implied, 0},        // 33 - Future Expansion
		52:  AssemblyInstruction{UNDEF, Implied, 0},        // 34 - Future Expansion
		53:  AssemblyInstruction{AND, ZeroPageX, 4},        // 35 - AND - Zero Page,X
		54:  AssemblyInstruction{ROL, ZeroPageX, 6},        // 36 - ROL - Zero Page,X
		55:  AssemblyInstruction{UNDEF, Implied, 0},        // 37 - Future Expansion
		56:  AssemblyInstruction{SEC, Implied, 2},          // 38 - SEC
		57:  AssemblyInstruction{AND, AbsoluteY, 4},        // 39 - AND - Absolute,Y
		58:  AssemblyInstruction{UNDEF, Implied, 0},        // 3A - Future Expansion
		59:  AssemblyInstruction{UNDEF, Implied, 0},        // 3B - Future Expansion
		60:  AssemblyInstruction{UNDEF, Implied, 0},        // 3C - Future Expansion
		61:  AssemblyInstruction{AND, AbsoluteX, 4},        // 3D - AND - Absolute,X
		62:  AssemblyInstruction{ROL, AbsoluteX, 7},        // 3E - ROL - Absolute,X
		63:  AssemblyInstruction{UNDEF, Implied, 0},        // 3F - Future Expansion
		64:  AssemblyInstruction{RTI, Implied, 6},          // 40 - RTI
		65:  AssemblyInstruction{EOR, IndexedIndirectX, 6}, // 41 - EOR - (Indirect,X}
		66:  AssemblyInstruction{UNDEF, Implied, 0},        // 42 - Future Expansion
		67:  AssemblyInstruction{UNDEF, Implied, 0},        // 43 - Future Expansion
		68:  AssemblyInstruction{UNDEF, Implied, 0},        // 44 - Future Expansion
		69:  AssemblyInstruction{EOR, ZeroPage, 3},         // 45 - EOR - Zero Page
		70:  AssemblyInstruction{LSR, ZeroPage, 5},         // 46 - LSR - Zero Page
		71:  AssemblyInstruction{UNDEF, Implied, 0},        // 47 - Future Expansion
		72:  AssemblyInstruction{PHA, Implied, 3},          // 48 - PHA
		73:  AssemblyInstruction{EOR, Immidiate, 2},        // 49 - EOR - Immediate
		74:  AssemblyInstruction{LSR, Accumulator, 2},      // 4A - LSR - Accumulator
		75:  AssemblyInstruction{UNDEF, Implied, 0},        // 4B - Future Expansion
		76:  AssemblyInstruction{JMP, Absolute, 3},         // 4C - JMP - Absolute
		77:  AssemblyInstruction{EOR, Absolute, 4},         // 4D - EOR - Absolute
		78:  AssemblyInstruction{LSR, Absolute, 6},         // 4E - LSR - Absolute
		79:  AssemblyInstruction{UNDEF, Implied, 0},        // 4F - Future Expansion
		80:  AssemblyInstruction{BVC, Relative, 2},         // 50 - BVC
		81:  AssemblyInstruction{EOR, IndirectIndexedY, 5}, // 51 - EOR - (Indirect,Y}
		82:  AssemblyInstruction{UNDEF, Implied, 0},        // 52 - Future Expansion
		83:  AssemblyInstruction{UNDEF, Implied, 0},        // 53 - Future Expansion
		84:  AssemblyInstruction{UNDEF, Implied, 0},        // 54 - Future Expansion
		85:  AssemblyInstruction{EOR, ZeroPageX, 4},        // 55 - EOR - Zero Page,X
		86:  AssemblyInstruction{LSR, ZeroPageX, 6},        // 56 - LSR - Zero Page,X
		87:  AssemblyInstruction{UNDEF, Implied, 0},        // 57 - Future Expansion
		88:  AssemblyInstruction{CLI, Implied, 2},          // 58 - CLI
		89:  AssemblyInstruction{EOR, AbsoluteY, 4},        // 59 - EOR - Absolute,Y
		90:  AssemblyInstruction{UNDEF, Implied, 0},        // 5A - Future Expansion
		91:  AssemblyInstruction{UNDEF, Implied, 0},        // 5B - Future Expansion
		92:  AssemblyInstruction{UNDEF, Implied, 0},        // 5C - Future Expansion
		93:  AssemblyInstruction{EOR, AbsoluteX, 4},        // 50 - EOR - Absolute,X
		94:  AssemblyInstruction{LSR, AbsoluteX, 7},        // 5E - LSR - Absolute,X
		95:  AssemblyInstruction{UNDEF, Implied, 0},        // 5F - Future Expansion
		96:  AssemblyInstruction{RTS, Implied, 6},          // 60 - RTS
		97:  AssemblyInstruction{ADC, IndexedIndirectX, 6}, // 61 - ADC - (Indirect,X}
		98:  AssemblyInstruction{UNDEF, Implied, 0},        // 62 - Future Expansion
		99:  AssemblyInstruction{UNDEF, Implied, 0},        // 63 - Future Expansion
		100: AssemblyInstruction{UNDEF, Implied, 0},        // 64 - Future Expansion
		101: AssemblyInstruction{ADC, ZeroPage, 3},         // 65 - ADC - Zero Page
		102: AssemblyInstruction{ROR, ZeroPage, 5},         // 66 - ROR - Zero Page
		103: AssemblyInstruction{UNDEF, Implied, 0},        // 67 - Future Expansion
		104: AssemblyInstruction{PLA, Implied, 4},          // 68 - PLA
		105: AssemblyInstruction{ADC, Immidiate, 2},        // 69 - ADC - Immediate
		106: AssemblyInstruction{ROR, Accumulator, 2},      // 6A - ROR - Accumulator
		107: AssemblyInstruction{UNDEF, Implied, 0},        // 6B - Future Expansion
		108: AssemblyInstruction{JMP, Indirect, 5},         // 6C - JMP - Indirect
		109: AssemblyInstruction{ADC, Absolute, 4},         // 6D - ADC - Absolute
		110: AssemblyInstruction{ROR, Absolute, 6},         // 6E - ROR - Absolute
		111: AssemblyInstruction{UNDEF, Implied, 0},        // 6F - Future Expansion
		112: AssemblyInstruction{BVS, Relative, 2},         // 70 - BVS
		113: AssemblyInstruction{ADC, IndirectIndexedY, 5}, // 71 - ADC - (Indirect,Y}
		114: AssemblyInstruction{UNDEF, Implied, 0},        // 72 - Future Expansion
		115: AssemblyInstruction{UNDEF, Implied, 0},        // 73 - Future Expansion
		116: AssemblyInstruction{UNDEF, Implied, 0},        // 74 - Future Expansion
		117: AssemblyInstruction{ADC, ZeroPageX, 4},        // 75 - ADC - Zero Page,X
		118: AssemblyInstruction{ROR, ZeroPageX, 6},        // 76 - ROR - Zero Page,X
		119: AssemblyInstruction{UNDEF, Implied, 0},        // 77 - Future Expansion
		120: AssemblyInstruction{SEI, Implied, 2},          // 78 - SEI
		121: AssemblyInstruction{ADC, AbsoluteY, 4},        // 79 - ADC - Absolute,Y
		122: AssemblyInstruction{UNDEF, Implied, 0},        // 7A - Future Expansion
		123: AssemblyInstruction{UNDEF, Implied, 0},        // 7B - Future Expansion
		124: AssemblyInstruction{UNDEF, Implied, 0},        // 7C - Future Expansion
		125: AssemblyInstruction{ADC, AbsoluteX, 4},        // 7D - ADC - Absolute,X
		126: AssemblyInstruction{ROR, AbsoluteX, 7},        // 7E - ROR - Absolute,X
		127: AssemblyInstruction{UNDEF, Implied, 0},        // 7F - Future Expansion
		128: AssemblyInstruction{UNDEF, Implied, 0},        // 80 - Future Expansion
		129: AssemblyInstruction{STA, IndexedIndirectX, 6}, // 81 - STA - (Indirect,X}
		130: AssemblyInstruction{UNDEF, Implied, 0},        // 82 - Future Expansion
		131: AssemblyInstruction{UNDEF, Implied, 0},        // 83 - Future Expansion
		132: AssemblyInstruction{STY, ZeroPage, 3},         // 84 - STY - Zero Page
		133: AssemblyInstruction{STA, ZeroPage, 3},         // 85 - STA - Zero Page
		134: AssemblyInstruction{STX, ZeroPage, 3},         // 86 - STX - Zero Page
		135: AssemblyInstruction{UNDEF, Implied, 0},        // 87 - Future Expansion
		136: AssemblyInstruction{DEY, Implied, 2},          // 88 - DEY
		137: AssemblyInstruction{UNDEF, Implied, 0},        // 89 - Future Expansion
		138: AssemblyInstruction{TXA, Implied, 2},          // 8A - TXA
		139: AssemblyInstruction{UNDEF, Implied, 0},        // 8B - Future Expansion
		140: AssemblyInstruction{STY, Absolute, 4},         // 8C - STY - Absolute
		141: AssemblyInstruction{STA, Absolute, 4},         // 8D - STA - Absolute
		142: AssemblyInstruction{STX, Absolute, 4},         // 8E - STX - Absolute
		143: AssemblyInstruction{UNDEF, Implied, 0},        // 8F - Future Expansion
		144: AssemblyInstruction{BCC, Relative, 2},         // 90 - BCC
		145: AssemblyInstruction{STA, IndirectIndexedY, 6}, // 91 - STA - (Indirect,Y}
		146: AssemblyInstruction{UNDEF, Implied, 0},        // 92 - Future Expansion
		147: AssemblyInstruction{UNDEF, Implied, 0},        // 93 - Future Expansion
		148: AssemblyInstruction{STY, ZeroPageX, 4},        // 94 - STY - Zero Page,X
		149: AssemblyInstruction{STA, ZeroPageX, 4},        // 95 - STA - Zero Page,X
		150: AssemblyInstruction{STX, ZeroPageY, 4},        // 96 - STX - Zero Page,Y
		151: AssemblyInstruction{UNDEF, Implied, 0},        // 97 - Future Expansion
		152: AssemblyInstruction{TYA, Implied, 2},          // 98 - TYA
		153: AssemblyInstruction{STA, AbsoluteY, 5},        // 99 - STA - Absolute,Y
		154: AssemblyInstruction{TXS, Implied, 2},          // 9A - TXS
		155: AssemblyInstruction{UNDEF, Implied, 0},        // 9B - Future Expansion
		156: AssemblyInstruction{UNDEF, Implied, 0},        // 9C - Future Expansion
		157: AssemblyInstruction{STA, AbsoluteX, 5},        // 90 - STA - Absolute,X
		158: AssemblyInstruction{UNDEF, Implied, 0},        // 9E - Future Expansion
		159: AssemblyInstruction{UNDEF, Implied, 0},        // 9F - Future Expansion
		160: AssemblyInstruction{LDY, Immidiate, 2},        // A0 - LDY - Immediate
		161: AssemblyInstruction{LDA, IndexedIndirectX, 6}, // A1 - LDA - (Indirect,X}
		162: AssemblyInstruction{LDX, Immidiate, 2},        // A2 - LDX - Immediate
		163: AssemblyInstruction{UNDEF, Implied, 0},        // A3 - Future Expansion
		164: AssemblyInstruction{LDY, ZeroPage, 3},         // A4 - LDY - Zero Page
		165: AssemblyInstruction{LDA, ZeroPage, 3},         // A5 - LDA - Zero Page
		166: AssemblyInstruction{LDX, ZeroPage, 3},         // A6 - LDX - Zero Page
		167: AssemblyInstruction{UNDEF, Implied, 0},        // A7 - Future Expansion
		168: AssemblyInstruction{TAY, Implied, 2},          // A8 - TAY
		169: AssemblyInstruction{LDA, Immidiate, 2},        // A9 - LDA - Immediate
		170: AssemblyInstruction{TAX, Implied, 2},          // AA - TAX
		171: AssemblyInstruction{UNDEF, Implied, 0},        // AB - Future Expansion
		172: AssemblyInstruction{LDY, Absolute, 4},         // AC - LDY - Absolute
		173: AssemblyInstruction{LDA, Absolute, 4},         // AD - LDA - Absolute
		174: AssemblyInstruction{LDX, Absolute, 4},         // AE - LDX - Absolute
		175: AssemblyInstruction{UNDEF, Implied, 0},        // AF - Future Expansion
		176: AssemblyInstruction{BCS, Relative, 2},         // B0 - BCS
		177: AssemblyInstruction{LDA, IndirectIndexedY, 5}, // B1 - LDA - (Indirect,Y}
		178: AssemblyInstruction{UNDEF, Implied, 0},        // B2 - Future Expansion
		179: AssemblyInstruction{UNDEF, Implied, 0},        // B3 - Future Expansion
		180: AssemblyInstruction{LDY, ZeroPageX, 4},        // B4 - LDY - Zero Page,X
		181: AssemblyInstruction{LDA, ZeroPageX, 4},        // B5 - LDA - Zero Page,X
		182: AssemblyInstruction{LDX, ZeroPageY, 4},        // B6 - LDX - Zero Page,Y
		183: AssemblyInstruction{UNDEF, Implied, 0},        // B7 - Future Expansion
		184: AssemblyInstruction{CLV, Implied, 2},          // B8 - CLV
		185: AssemblyInstruction{LDA, AbsoluteY, 4},        // B9 - LDA - Absolute,Y
		186: AssemblyInstruction{TSX, Implied, 2},          // BA - TSX
		187: AssemblyInstruction{UNDEF, Implied, 0},        // BB - Future Expansion
		188: AssemblyInstruction{LDY, AbsoluteX, 4},        // BC - LDY - Absolute,X
		189: AssemblyInstruction{LDA, AbsoluteX, 4},        // BD - LDA - Absolute,X
		190: AssemblyInstruction{LDX, AbsoluteY, 4},        // BE - LDX - Absolute,Y
		191: AssemblyInstruction{UNDEF, Implied, 0},        // BF - Future Expansion
		192: AssemblyInstruction{CPY, Immidiate, 2},        // C0 - Cpy - Immediate
		193: AssemblyInstruction{CMP, IndexedIndirectX, 6}, // C1 - CMP - (Indirect,X}
		194: AssemblyInstruction{UNDEF, Implied, 0},        // C2 - Future Expansion
		195: AssemblyInstruction{UNDEF, Implied, 0},        // C3 - Future Expansion
		196: AssemblyInstruction{CPY, ZeroPage, 3},         // C4 - CPY - Zero Page
		197: AssemblyInstruction{CMP, ZeroPage, 3},         // C5 - CMP - Zero Page
		198: AssemblyInstruction{DEC, ZeroPage, 5},         // C6 - DEC - Zero Page
		199: AssemblyInstruction{UNDEF, Implied, 0},        // C7 - Future Expansion
		200: AssemblyInstruction{INY, Implied, 2},          // C8 - INY
		201: AssemblyInstruction{CMP, Immidiate, 2},        // C9 - CMP - Immediate
		202: AssemblyInstruction{DEX, Implied, 2},          // CA - DEX
		203: AssemblyInstruction{UNDEF, Implied, 0},        // CB - Future Expansion
		204: AssemblyInstruction{CPY, Absolute, 4},         // CC - CPY - Absolute
		205: AssemblyInstruction{CMP, Absolute, 4},         // CD - CMP - Absolute
		206: AssemblyInstruction{DEC, Absolute, 6},         // CE - DEC - Absolute
		207: AssemblyInstruction{UNDEF, Implied, 0},        // CF - Future Expansion
		208: AssemblyInstruction{BNE, Relative, 2},         // D0 - BNE
		209: AssemblyInstruction{CMP, IndirectIndexedY, 5}, // D1 - CMP   (Indirect,Y}
		210: AssemblyInstruction{UNDEF, Implied, 0},        // D2 - Future Expansion
		211: AssemblyInstruction{UNDEF, Implied, 0},        // D3 - Future Expansion
		212: AssemblyInstruction{UNDEF, Implied, 0},        // D4 - Future Expansion
		213: AssemblyInstruction{CMP, ZeroPageX, 4},        // D5 - CMP - Zero Page,X
		214: AssemblyInstruction{DEC, ZeroPageX, 6},        // D6 - DEC - Zero Page,X
		215: AssemblyInstruction{UNDEF, Implied, 0},        // D7 - Future Expansion
		216: AssemblyInstruction{CLD, Implied, 2},          // D8 - CLD
		217: AssemblyInstruction{CMP, AbsoluteY, 4},        // D9 - CMP - Absolute,Y
		218: AssemblyInstruction{UNDEF, Implied, 0},        // DA - Future Expansion
		219: AssemblyInstruction{UNDEF, Implied, 0},        // DB - Future Expansion
		220: AssemblyInstruction{UNDEF, Implied, 0},        // DC - Future Expansion
		221: AssemblyInstruction{CMP, AbsoluteX, 4},        // DD - CMP - Absolute,X
		222: AssemblyInstruction{DEC, AbsoluteX, 7},        // DE - DEC - Absolute,X
		223: AssemblyInstruction{UNDEF, Implied, 0},        // DF - Future Expansion
		224: AssemblyInstruction{CPX, Immidiate, 2},        // E0 - CPX - Immediate
		225: AssemblyInstruction{SBC, IndexedIndirectX, 6}, // E1 - SBC - (Indirect,X}
		226: AssemblyInstruction{UNDEF, Implied, 0},        // E2 - Future Expansion
		227: AssemblyInstruction{UNDEF, Implied, 0},        // E3 - Future Expansion
		228: AssemblyInstruction{CPX, ZeroPage, 3},         // E4 - CPX - Zero Page
		229: AssemblyInstruction{SBC, ZeroPage, 3},         // E5 - SBC - Zero Page
		230: AssemblyInstruction{INC, ZeroPage, 5},         // E6 - INC - Zero Page
		231: AssemblyInstruction{UNDEF, Implied, 0},        // E7 - Future Expansion
		232: AssemblyInstruction{INX, Implied, 2},          // E8 - INX
		233: AssemblyInstruction{SBC, Immidiate, 2},        // E9 - SBC - Immediate
		234: AssemblyInstruction{NOP, Implied, 2},          // EA - NOP
		235: AssemblyInstruction{UNDEF, Implied, 0},        // EB - Future Expansion
		236: AssemblyInstruction{CPX, Absolute, 4},         // EC - CPX - Absolute
		237: AssemblyInstruction{SBC, Absolute, 4},         // ED - SBC - Absolute
		238: AssemblyInstruction{INC, Absolute, 6},         // EE - INC - Absolute
		239: AssemblyInstruction{UNDEF, Implied, 0},        // EF - Future Expansion
		240: AssemblyInstruction{BEQ, Relative, 2},         // F0 - BEQ
		241: AssemblyInstruction{SBC, IndirectIndexedY, 5}, // F1 - SBC - (Indirect,Y}
		242: AssemblyInstruction{UNDEF, Implied, 0},        // F2 - Future Expansion
		243: AssemblyInstruction{UNDEF, Implied, 0},        // F3 - Future Expansion
		244: AssemblyInstruction{UNDEF, Implied, 0},        // F4 - Future Expansion
		245: AssemblyInstruction{SBC, ZeroPageX, 4},        // F5 - SBC - Zero Page,X
		246: AssemblyInstruction{INC, ZeroPageX, 6},        // F6 - INC - Zero Page,X
		247: AssemblyInstruction{UNDEF, Implied, 0},        // F7 - Future Expansion
		248: AssemblyInstruction{SED, Implied, 2},          // F8 - SED
		249: AssemblyInstruction{SBC, AbsoluteY, 4},        // F9 - SBC - Absolute,Y
		250: AssemblyInstruction{UNDEF, Implied, 0},        // FA - Future Expansion
		251: AssemblyInstruction{UNDEF, Implied, 0},        // FB - Future Expansion
		252: AssemblyInstruction{UNDEF, Implied, 0},        // FC - Future Expansion
		253: AssemblyInstruction{SBC, AbsoluteX, 4},        // FD - SBC - Absolute,X
		254: AssemblyInstruction{INC, AbsoluteX, 7},        // FE - INC - Absolute,X
		255: AssemblyInstruction{UNDEF, Implied, 0},        // FF - Future
	}

	return func(key byte) AssemblyInstruction {
		//log.Println("Resolving instruction by bytecode: ", key)
		return innerMap[key]
	}
}
