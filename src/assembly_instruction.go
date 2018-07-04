package main

import "log"

type AssemblyInstruction struct {
	Type           AssemblyInstructionType
	AddressingMode AddressingMode
}

func assemblyInstructions() func(byte) AssemblyInstruction {

	innerMap := map[byte]AssemblyInstruction{
		0:   AssemblyInstruction{BRK, Implied},          // 00 - BRK
		1:   AssemblyInstruction{ORA, IndexedIndirectX}, // 01 - ORA - (Indirect,X}
		2:   AssemblyInstruction{UNDEF, Implied},        // 02 - Future Expansion
		3:   AssemblyInstruction{UNDEF, Implied},        // 03 - Future Expansion
		4:   AssemblyInstruction{UNDEF, Implied},        // 04 - Future Expansion
		5:   AssemblyInstruction{ORA, ZeroPage},         // 05 - ORA - Zero Page
		6:   AssemblyInstruction{ASL, ZeroPage},         // 06 - ASL - Zero Page
		7:   AssemblyInstruction{UNDEF, Implied},        // 07 - Future Expansion
		8:   AssemblyInstruction{PHP, Implied},          // 08 - PHP
		9:   AssemblyInstruction{ORA, Immidiate},        // 09 - ORA - Immediate
		10:  AssemblyInstruction{ASL, Accumulator},      // 0A - ASL - Accumulator
		11:  AssemblyInstruction{UNDEF, Implied},        // 0B - Future Expansion
		12:  AssemblyInstruction{UNDEF, Implied},        // 0C - Future Expansion
		13:  AssemblyInstruction{ORA, Absolute},         // 0D - ORA - Absolute
		14:  AssemblyInstruction{ASL, Absolute},         // 0E - ASL - Absolute
		15:  AssemblyInstruction{UNDEF, Implied},        // 0F - Future Expansion
		16:  AssemblyInstruction{BPL, Relative},         // 10 - BPL
		17:  AssemblyInstruction{ORA, IndirectIndexedY}, // 11 - ORA - (Indirect,Y}
		18:  AssemblyInstruction{UNDEF, Implied},        // 12 - Future Expansion
		19:  AssemblyInstruction{UNDEF, Implied},        // 13 - Future Expansion
		20:  AssemblyInstruction{UNDEF, Implied},        // 14 - Future Expansion
		21:  AssemblyInstruction{ORA, ZeroPageX},        // 15 - ORA - Zero Page,X
		22:  AssemblyInstruction{ASL, ZeroPageX},        // 16 - ASL - Zero Page,X
		23:  AssemblyInstruction{UNDEF, Implied},        // 17 - Future Expansion
		24:  AssemblyInstruction{CLC, Implied},          // 18 - CLC
		25:  AssemblyInstruction{ORA, AbsoluteY},        // 19 - ORA - Absolute,Y
		26:  AssemblyInstruction{UNDEF, Implied},        // 1A - Future Expansion
		27:  AssemblyInstruction{UNDEF, Implied},        // 1B - Future Expansion
		28:  AssemblyInstruction{UNDEF, Implied},        // 1C - Future Expansion
		29:  AssemblyInstruction{ORA, AbsoluteX},        // 1D - ORA - Absolute,X
		30:  AssemblyInstruction{ASL, AbsoluteX},        // 1E - ASL - Absolute,X
		31:  AssemblyInstruction{UNDEF, Implied},        // 1F - Future Expansion
		32:  AssemblyInstruction{JSR, Absolute},         // 20 - JSR
		33:  AssemblyInstruction{AND, IndexedIndirectX}, // 21 - AND - (Indirect,X}
		34:  AssemblyInstruction{UNDEF, Implied},        // 22 - Future Expansion
		35:  AssemblyInstruction{UNDEF, Implied},        // 23 - Future Expansion
		36:  AssemblyInstruction{BIT, ZeroPage},         // 24 - BIT - Zero Page
		37:  AssemblyInstruction{AND, ZeroPage},         // 25 - AND - Zero Page
		38:  AssemblyInstruction{ROL, ZeroPage},         // 26 - ROL - Zero Page
		39:  AssemblyInstruction{UNDEF, Implied},        // 27 - Future Expansion
		40:  AssemblyInstruction{PLP, Implied},          // 28 - PLP
		41:  AssemblyInstruction{AND, Immidiate},        // 29 - AND - Immediate
		42:  AssemblyInstruction{ROL, Accumulator},      // 2A - ROL - Accumulator
		43:  AssemblyInstruction{UNDEF, Implied},        // 2B - Future Expansion
		44:  AssemblyInstruction{BIT, Absolute},         // 2C - BIT - Absolute
		45:  AssemblyInstruction{AND, Absolute},         // 2D - AND - Absolute
		46:  AssemblyInstruction{ROL, Absolute},         // 2E - ROL - Absolute
		47:  AssemblyInstruction{UNDEF, Implied},        // 2F - Future Expansion
		48:  AssemblyInstruction{BMI, Relative},         // 30 - BMI
		49:  AssemblyInstruction{AND, IndirectIndexedY}, // 31 - AND - (Indirect,Y}
		50:  AssemblyInstruction{UNDEF, Implied},        // 32 - Future Expansion
		51:  AssemblyInstruction{UNDEF, Implied},        // 33 - Future Expansion
		52:  AssemblyInstruction{UNDEF, Implied},        // 34 - Future Expansion
		53:  AssemblyInstruction{AND, ZeroPageX},        // 35 - AND - Zero Page,X
		54:  AssemblyInstruction{ROL, ZeroPageX},        // 36 - ROL - Zero Page,X
		55:  AssemblyInstruction{UNDEF, Implied},        // 37 - Future Expansion
		56:  AssemblyInstruction{SEC, Implied},          // 38 - SEC
		57:  AssemblyInstruction{AND, AbsoluteY},        // 39 - AND - Absolute,Y
		58:  AssemblyInstruction{UNDEF, Implied},        // 3A - Future Expansion
		59:  AssemblyInstruction{UNDEF, Implied},        // 3B - Future Expansion
		60:  AssemblyInstruction{UNDEF, Implied},        // 3C - Future Expansion
		61:  AssemblyInstruction{AND, AbsoluteX},        // 3D - AND - Absolute,X
		62:  AssemblyInstruction{ROL, AbsoluteX},        // 3E - ROL - Absolute,X
		63:  AssemblyInstruction{UNDEF, Implied},        // 3F - Future Expansion
		64:  AssemblyInstruction{RTI, Implied},          // 40 - RTI
		65:  AssemblyInstruction{EOR, IndexedIndirectX}, // 41 - EOR - (Indirect,X}
		66:  AssemblyInstruction{UNDEF, Implied},        // 42 - Future Expansion
		67:  AssemblyInstruction{UNDEF, Implied},        // 43 - Future Expansion
		68:  AssemblyInstruction{UNDEF, Implied},        // 44 - Future Expansion
		69:  AssemblyInstruction{EOR, ZeroPage},         // 45 - EOR - Zero Page
		70:  AssemblyInstruction{LSR, ZeroPage},         // 46 - LSR - Zero Page
		71:  AssemblyInstruction{UNDEF, Implied},        // 47 - Future Expansion
		72:  AssemblyInstruction{PHA, Implied},          // 48 - PHA
		73:  AssemblyInstruction{EOR, Immidiate},        // 49 - EOR - Immediate
		74:  AssemblyInstruction{LSR, Accumulator},      // 4A - LSR - Accumulator
		75:  AssemblyInstruction{UNDEF, Implied},        // 4B - Future Expansion
		76:  AssemblyInstruction{JMP, Absolute},         // 4C - JMP - Absolute
		77:  AssemblyInstruction{EOR, Absolute},         // 4D - EOR - Absolute
		78:  AssemblyInstruction{LSR, Absolute},         // 4E - LSR - Absolute
		79:  AssemblyInstruction{UNDEF, Implied},        // 4F - Future Expansion
		80:  AssemblyInstruction{BVC, Relative},         // 50 - BVC
		81:  AssemblyInstruction{EOR, IndirectIndexedY}, // 51 - EOR - (Indirect,Y}
		82:  AssemblyInstruction{UNDEF, Implied},        // 52 - Future Expansion
		83:  AssemblyInstruction{UNDEF, Implied},        // 53 - Future Expansion
		84:  AssemblyInstruction{UNDEF, Implied},        // 54 - Future Expansion
		85:  AssemblyInstruction{EOR, ZeroPageX},        // 55 - EOR - Zero Page,X
		86:  AssemblyInstruction{LSR, ZeroPageX},        // 56 - LSR - Zero Page,X
		87:  AssemblyInstruction{UNDEF, Implied},        // 57 - Future Expansion
		88:  AssemblyInstruction{CLI, Implied},          // 58 - CLI
		89:  AssemblyInstruction{EOR, AbsoluteY},        // 59 - EOR - Absolute,Y
		90:  AssemblyInstruction{UNDEF, Implied},        // 5A - Future Expansion
		91:  AssemblyInstruction{UNDEF, Implied},        // 5B - Future Expansion
		92:  AssemblyInstruction{UNDEF, Implied},        // 5C - Future Expansion
		93:  AssemblyInstruction{EOR, AbsoluteX},        // 50 - EOR - Absolute,X
		94:  AssemblyInstruction{LSR, AbsoluteX},        // 5E - LSR - Absolute,X
		95:  AssemblyInstruction{UNDEF, Implied},        // 5F - Future Expansion
		96:  AssemblyInstruction{RTS, Implied},          // 60 - RTS
		97:  AssemblyInstruction{ADC, IndexedIndirectX}, // 61 - ADC - (Indirect,X}
		98:  AssemblyInstruction{UNDEF, Implied},        // 62 - Future Expansion
		99:  AssemblyInstruction{UNDEF, Implied},        // 63 - Future Expansion
		100: AssemblyInstruction{UNDEF, Implied},        // 64 - Future Expansion
		101: AssemblyInstruction{ADC, ZeroPage},         // 65 - ADC - Zero Page
		102: AssemblyInstruction{ROR, ZeroPage},         // 66 - ROR - Zero Page
		103: AssemblyInstruction{UNDEF, Implied},        // 67 - Future Expansion
		104: AssemblyInstruction{PLA, Implied},          // 68 - PLA
		105: AssemblyInstruction{ADC, Immidiate},        // 69 - ADC - Immediate
		106: AssemblyInstruction{ROR, Accumulator},      // 6A - ROR - Accumulator
		107: AssemblyInstruction{UNDEF, Implied},        // 6B - Future Expansion
		108: AssemblyInstruction{JMP, Indirect},         // 6C - JMP - Indirect
		109: AssemblyInstruction{ADC, Absolute},         // 6D - ADC - Absolute
		110: AssemblyInstruction{ROR, Absolute},         // 6E - ROR - Absolute
		111: AssemblyInstruction{UNDEF, Implied},        // 6F - Future Expansion
		112: AssemblyInstruction{BVS, Relative},         // 70 - BVS
		113: AssemblyInstruction{ADC, IndirectIndexedY}, // 71 - ADC - (Indirect,Y}
		114: AssemblyInstruction{UNDEF, Implied},        // 72 - Future Expansion
		115: AssemblyInstruction{UNDEF, Implied},        // 73 - Future Expansion
		116: AssemblyInstruction{UNDEF, Implied},        // 74 - Future Expansion
		117: AssemblyInstruction{ADC, ZeroPageX},        // 75 - ADC - Zero Page,X
		118: AssemblyInstruction{ROR, ZeroPageX},        // 76 - ROR - Zero Page,X
		119: AssemblyInstruction{UNDEF, Implied},        // 77 - Future Expansion
		120: AssemblyInstruction{SEI, Implied},          // 78 - SEI
		121: AssemblyInstruction{ADC, AbsoluteY},        // 79 - ADC - Absolute,Y
		122: AssemblyInstruction{UNDEF, Implied},        // 7A - Future Expansion
		123: AssemblyInstruction{UNDEF, Implied},        // 7B - Future Expansion
		124: AssemblyInstruction{UNDEF, Implied},        // 7C - Future Expansion
		125: AssemblyInstruction{ADC, AbsoluteX},        // 7D - ADC - Absolute,X
		126: AssemblyInstruction{ROR, AbsoluteX},        // 7E - ROR - Absolute,X
		127: AssemblyInstruction{UNDEF, Implied},        // 7F - Future Expansion
		128: AssemblyInstruction{UNDEF, Implied},        // 80 - Future Expansion
		129: AssemblyInstruction{STA, IndexedIndirectX}, // 81 - STA - (Indirect,X}
		130: AssemblyInstruction{UNDEF, Implied},        // 82 - Future Expansion
		131: AssemblyInstruction{UNDEF, Implied},        // 83 - Future Expansion
		132: AssemblyInstruction{STY, ZeroPage},         // 84 - STY - Zero Page
		133: AssemblyInstruction{STA, ZeroPage},         // 85 - STA - Zero Page
		134: AssemblyInstruction{STX, ZeroPage},         // 86 - STX - Zero Page
		135: AssemblyInstruction{UNDEF, Implied},        // 87 - Future Expansion
		136: AssemblyInstruction{DEY, Implied},          // 88 - DEY
		137: AssemblyInstruction{UNDEF, Implied},        // 89 - Future Expansion
		138: AssemblyInstruction{TXA, Implied},          // 8A - TXA
		139: AssemblyInstruction{UNDEF, Implied},        // 8B - Future Expansion
		140: AssemblyInstruction{STY, Absolute},         // 8C - STY - Absolute
		141: AssemblyInstruction{STA, Absolute},         // 8D - STA - Absolute
		142: AssemblyInstruction{STX, Absolute},         // 8E - STX - Absolute
		143: AssemblyInstruction{UNDEF, Implied},        // 8F - Future Expansion
		144: AssemblyInstruction{BCC, Relative},         // 90 - BCC
		145: AssemblyInstruction{STA, IndirectIndexedY}, // 91 - STA - (Indirect,Y}
		146: AssemblyInstruction{UNDEF, Implied},        // 92 - Future Expansion
		147: AssemblyInstruction{UNDEF, Implied},        // 93 - Future Expansion
		148: AssemblyInstruction{STY, ZeroPageX},        // 94 - STY - Zero Page,X
		149: AssemblyInstruction{STA, ZeroPageX},        // 95 - STA - Zero Page,X
		150: AssemblyInstruction{STX, ZeroPageY},        // 96 - STX - Zero Page,Y
		151: AssemblyInstruction{UNDEF, Implied},        // 97 - Future Expansion
		152: AssemblyInstruction{TYA, Implied},          // 98 - TYA
		153: AssemblyInstruction{STA, AbsoluteY},        // 99 - STA - Absolute,Y
		154: AssemblyInstruction{TXS, Implied},          // 9A - TXS
		155: AssemblyInstruction{UNDEF, Implied},        // 9B - Future Expansion
		156: AssemblyInstruction{UNDEF, Implied},        // 9C - Future Expansion
		157: AssemblyInstruction{STA, AbsoluteX},        // 90 - STA - Absolute,X
		158: AssemblyInstruction{UNDEF, Implied},        // 9E - Future Expansion
		159: AssemblyInstruction{UNDEF, Implied},        // 9F - Future Expansion
		160: AssemblyInstruction{LDY, Immidiate},        // A0 - LDY - Immediate
		161: AssemblyInstruction{LDA, IndexedIndirectX}, // A1 - LDA - (Indirect,X}
		162: AssemblyInstruction{LDX, Immidiate},        // A2 - LDX - Immediate
		163: AssemblyInstruction{UNDEF, Implied},        // A3 - Future Expansion
		164: AssemblyInstruction{LDY, ZeroPage},         // A4 - LDY - Zero Page
		165: AssemblyInstruction{LDA, ZeroPage},         // A5 - LDA - Zero Page
		166: AssemblyInstruction{LDX, ZeroPage},         // A6 - LDX - Zero Page
		167: AssemblyInstruction{UNDEF, Implied},        // A7 - Future Expansion
		168: AssemblyInstruction{TAY, Implied},          // A8 - TAY
		169: AssemblyInstruction{LDA, Immidiate},        // A9 - LDA - Immediate
		170: AssemblyInstruction{TAX, Implied},          // AA - TAX
		171: AssemblyInstruction{UNDEF, Implied},        // AB - Future Expansion
		172: AssemblyInstruction{LDY, Absolute},         // AC - LDY - Absolute
		173: AssemblyInstruction{LDA, Absolute},         // AD - LDA - Absolute
		174: AssemblyInstruction{LDX, Absolute},         // AE - LDX - Absolute
		175: AssemblyInstruction{UNDEF, Implied},        // AF - Future Expansion
		176: AssemblyInstruction{BCS, Relative},         // B0 - BCS
		177: AssemblyInstruction{LDA, IndirectIndexedY}, // B1 - LDA - (Indirect,Y}
		178: AssemblyInstruction{UNDEF, Implied},        // B2 - Future Expansion
		179: AssemblyInstruction{UNDEF, Implied},        // B3 - Future Expansion
		180: AssemblyInstruction{LDY, ZeroPageX},        // B4 - LDY - Zero Page,X
		181: AssemblyInstruction{LDA, ZeroPageX},        // B5 - LDA - Zero Page,X
		182: AssemblyInstruction{LDX, ZeroPageY},        // B6 - LDX - Zero Page,Y
		183: AssemblyInstruction{UNDEF, Implied},        // B7 - Future Expansion
		184: AssemblyInstruction{CLV, Implied},          // B8 - CLV
		185: AssemblyInstruction{LDA, AbsoluteY},        // B9 - LDA - Absolute,Y
		186: AssemblyInstruction{TSX, Implied},          // BA - TSX
		187: AssemblyInstruction{UNDEF, Implied},        // BB - Future Expansion
		188: AssemblyInstruction{LDY, AbsoluteX},        // BC - LDY - Absolute,X
		189: AssemblyInstruction{LDA, AbsoluteX},        // BD - LDA - Absolute,X
		190: AssemblyInstruction{LDX, AbsoluteY},        // BE - LDX - Absolute,Y
		191: AssemblyInstruction{UNDEF, Implied},        // BF - Future Expansion
		192: AssemblyInstruction{CPY, Immidiate},        // C0 - Cpy - Immediate
		193: AssemblyInstruction{CMP, IndexedIndirectX}, // C1 - CMP - (Indirect,X}
		194: AssemblyInstruction{UNDEF, Implied},        // C2 - Future Expansion
		195: AssemblyInstruction{UNDEF, Implied},        // C3 - Future Expansion
		196: AssemblyInstruction{CPY, ZeroPage},         // C4 - CPY - Zero Page
		197: AssemblyInstruction{CMP, ZeroPage},         // C5 - CMP - Zero Page
		198: AssemblyInstruction{DEC, ZeroPage},         // C6 - DEC - Zero Page
		199: AssemblyInstruction{UNDEF, Implied},        // C7 - Future Expansion
		200: AssemblyInstruction{INY, Implied},          // C8 - INY
		201: AssemblyInstruction{CMP, Immidiate},        // C9 - CMP - Immediate
		202: AssemblyInstruction{DEX, Implied},          // CA - DEX
		203: AssemblyInstruction{UNDEF, Implied},        // CB - Future Expansion
		204: AssemblyInstruction{CPY, Absolute},         // CC - CPY - Absolute
		205: AssemblyInstruction{CMP, Absolute},         // CD - CMP - Absolute
		206: AssemblyInstruction{DEC, Absolute},         // CE - DEC - Absolute
		207: AssemblyInstruction{UNDEF, Implied},        // CF - Future Expansion
		208: AssemblyInstruction{BNE, Relative},         // D0 - BNE
		209: AssemblyInstruction{CMP, IndirectIndexedY}, // D1 - CMP   (Indirect,Y}
		210: AssemblyInstruction{UNDEF, Implied},        // D2 - Future Expansion
		211: AssemblyInstruction{UNDEF, Implied},        // D3 - Future Expansion
		212: AssemblyInstruction{UNDEF, Implied},        // D4 - Future Expansion
		213: AssemblyInstruction{CMP, ZeroPageX},        // D5 - CMP - Zero Page,X
		214: AssemblyInstruction{DEC, ZeroPageX},        // D6 - DEC - Zero Page,X
		215: AssemblyInstruction{UNDEF, Implied},        // D7 - Future Expansion
		216: AssemblyInstruction{CLD, Implied},          // D8 - CLD
		217: AssemblyInstruction{CMP, AbsoluteY},        // D9 - CMP - Absolute,Y
		218: AssemblyInstruction{UNDEF, Implied},        // DA - Future Expansion
		219: AssemblyInstruction{UNDEF, Implied},        // DB - Future Expansion
		220: AssemblyInstruction{UNDEF, Implied},        // DC - Future Expansion
		221: AssemblyInstruction{CMP, AbsoluteX},        // DD - CMP - Absolute,X
		222: AssemblyInstruction{DEC, AbsoluteX},        // DE - DEC - Absolute,X
		223: AssemblyInstruction{UNDEF, Implied},        // DF - Future Expansion
		224: AssemblyInstruction{CPX, Immidiate},        // E0 - CPX - Immediate
		225: AssemblyInstruction{SBC, IndexedIndirectX}, // E1 - SBC - (Indirect,X}
		226: AssemblyInstruction{UNDEF, Implied},        // E2 - Future Expansion
		227: AssemblyInstruction{UNDEF, Implied},        // E3 - Future Expansion
		228: AssemblyInstruction{CPX, ZeroPage},         // E4 - CPX - Zero Page
		229: AssemblyInstruction{SBC, ZeroPage},         // E5 - SBC - Zero Page
		230: AssemblyInstruction{INC, ZeroPage},         // E6 - INC - Zero Page
		231: AssemblyInstruction{UNDEF, Implied},        // E7 - Future Expansion
		232: AssemblyInstruction{INX, Implied},          // E8 - INX
		233: AssemblyInstruction{SBC, Immidiate},        // E9 - SBC - Immediate
		234: AssemblyInstruction{NOP, Implied},          // EA - NOP
		235: AssemblyInstruction{UNDEF, Implied},        // EB - Future Expansion
		236: AssemblyInstruction{CPX, Absolute},         // EC - CPX - Absolute
		237: AssemblyInstruction{SBC, Absolute},         // ED - SBC - Absolute
		238: AssemblyInstruction{INC, Absolute},         // EE - INC - Absolute
		239: AssemblyInstruction{UNDEF, Implied},        // EF - Future Expansion
		240: AssemblyInstruction{BEQ, Relative},         // F0 - BEQ
		241: AssemblyInstruction{SBC, IndirectIndexedY}, // F1 - SBC - (Indirect,Y}
		242: AssemblyInstruction{UNDEF, Implied},        // F2 - Future Expansion
		243: AssemblyInstruction{UNDEF, Implied},        // F3 - Future Expansion
		244: AssemblyInstruction{UNDEF, Implied},        // F4 - Future Expansion
		245: AssemblyInstruction{SBC, ZeroPageX},        // F5 - SBC - Zero Page,X
		246: AssemblyInstruction{INC, ZeroPageX},        // F6 - INC - Zero Page,X
		247: AssemblyInstruction{UNDEF, Implied},        // F7 - Future Expansion
		248: AssemblyInstruction{SED, Implied},          // F8 - SED
		249: AssemblyInstruction{SBC, AbsoluteY},        // F9 - SBC - Absolute,Y
		250: AssemblyInstruction{UNDEF, Implied},        // FA - Future Expansion
		251: AssemblyInstruction{UNDEF, Implied},        // FB - Future Expansion
		252: AssemblyInstruction{UNDEF, Implied},        // FC - Future Expansion
		253: AssemblyInstruction{SBC, AbsoluteX},        // FD - SBC - Absolute,X
		254: AssemblyInstruction{INC, AbsoluteX},        // FE - INC - Absolute,X
		255: AssemblyInstruction{UNDEF, Implied},        // FF - Future
	}

	return func(key byte) AssemblyInstruction {
		log.Println("Resolving instruction by bytecode: ", key)
		return innerMap[key]
	}
}
