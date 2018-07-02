namespace Hackaton
{

    enum AssemblyInstructionType
    {
        UNDEF,
        ADC, //   Add Memory to Accumulator with Carry
        AND, // "AND" Memory with Accumulator
        ASL, // Shift Left One Bit (Memory or Accumulator)
        BCC, // Branch on Carry Clear
        BCS, // Branch on Carry Set
        BEQ, // Branch on Result Zero
        BIT, // Test Bits in Memory with Accumulator
        BMI, // Branch on Result Minus
        BNE, // Branch on Result not Zero
        BPL, // Branch on Result Plus
        BRK, // Force Break
        BVC, // Branch on Overflow Clear
        BVS, // Branch on Overflow Set
        CLC, // Clear Carry Flag
        CLD, // Clear Decimal Mode
        CLI, // Clear interrupt Disable Bit
        CLV, // Clear Overflow Flag
        CMP, // Compare Memory and Accumulator
        CPX, // Compare Memory and Index X
        CPY, // Compare Memory and Index Y
        DEC, // Decrement Memory by One
        DEX, // Decrement Index X by One
        DEY, // Decrement Index Y by One
        EOR, // "Exclusive-Or" Memory with Accumulator
        INC, // Increment Memory by One
        INX, // Increment Index X by One
        INY, // Increment Index Y by One
        JMP, // Jump to New Location
        JSR, // Jump to New Location Saving Return Address
        LDA, // Load Accumulator with Memory
        LDX, // Load Index X with Memory
        LDY, // Load Index Y with Memory
        LSR, // Shift Right One Bit (Memory or Accumulator)
        NOP, // No Operation
        ORA, // "OR" Memory with Accumulator
        PHA, // Push Accumulator on Stack
        PHP, // Push Processor Status on Stack
        PLA, // Pull Accumulator from Stack
        PLP, // Pull Processor Status from Stack
        ROL, // Rotate One Bit Left (Memory or Accumulator)
        ROR, // Rotate One Bit Right (Memory or Accumulator)
        RTI, // Return from Interrupt
        RTS, // Return from Subroutine
        SBC, // Subtract Memory from Accumulator with Borrow
        SEC, // Set Carry Flag
        SED, // Set Decimal Mode
        SEI, // Set Interrupt Disable Status
        STA, // Store Accumulator in Memory
        STX, // Store Index X in Memory
        STY, // Store Index Y in Memory
        TAX, // Transfer Accumulator to Index X
        TAY, // Transfer Accumulator to Index Y
        TSX, // Transfer Stack Pointer to Index X
        TXA, // Transfer Index X to Accumulator
        TXS, // Transfer Index X to Stack Pointer
        TYA, // Transfer Index Y to Accumulator
    }

    enum AddressingMode
    {
        Implied,
        IndexedIndirectX,
        IndirectIndexedY,
        Indirect,
        Absolute,
        AbsoluteX,
        AbsoluteY,
        Immidiate,
        ZeroPage,
        ZeroPageX,
        ZeroPageY,
        Accumulator,
        Relative
    }

    class AssemblyInstruction
    {
        public AssemblyInstructionType InstructionType { get; private set; }
        public AddressingMode AddressingMode { get; private set; }

        public AssemblyInstruction(AssemblyInstructionType instructionType, AddressingMode addressingMode)
        {
            InstructionType = instructionType;
            AddressingMode = addressingMode;
        }
    }

    /// <summary>
    /// Converts machine codes to assembly instructions with addressing mode.
    /// </summary>
    static class AssemblyInstructions
    {
        public static AssemblyInstruction GetInstruction(byte code)
        {
            return Codes[code];
        }

        private static AssemblyInstruction[] Codes = new []
        {
            new AssemblyInstruction(AssemblyInstructionType.BRK, AddressingMode.Implied), // 00 - BRK                        
            new AssemblyInstruction(AssemblyInstructionType.ORA, AddressingMode.IndexedIndirectX), // 01 - ORA - (Indirect,X)         
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 02 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 03 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 04 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.ORA, AddressingMode.ZeroPage), // 05 - ORA - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.ASL, AddressingMode.ZeroPage), // 06 - ASL - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 07 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.PHP, AddressingMode.Implied), // 08 - PHP                        
            new AssemblyInstruction(AssemblyInstructionType.ORA, AddressingMode.Immidiate), // 09 - ORA - Immediate            
            new AssemblyInstruction(AssemblyInstructionType.ASL, AddressingMode.Accumulator), // 0A - ASL - Accumulator          
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 0B - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 0C - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.ORA, AddressingMode.Absolute), // 0D - ORA - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.ASL, AddressingMode.Absolute), // 0E - ASL - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 0F - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.BPL, AddressingMode.Relative), // 10 - BPL                        
            new AssemblyInstruction(AssemblyInstructionType.ORA, AddressingMode.IndirectIndexedY),//11 - ORA - (Indirect),Y         
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 12 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 13 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 14 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.ORA, AddressingMode.ZeroPageX), //15 - ORA - Zero Page,X          
            new AssemblyInstruction(AssemblyInstructionType.ASL, AddressingMode.ZeroPageX), //16 - ASL - Zero Page,X          
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 17 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.CLC, AddressingMode.Implied), //18 - CLC                        
            new AssemblyInstruction(AssemblyInstructionType.ORA, AddressingMode.AbsoluteY), //19 - ORA - Absolute,Y           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 1A - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 1B - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 1C - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.ORA, AddressingMode.AbsoluteX), //1D - ORA - Absolute,X           
            new AssemblyInstruction(AssemblyInstructionType.ASL, AddressingMode.AbsoluteX), //1E - ASL - Absolute,X           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 1F - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.JSR, AddressingMode.Absolute), //20 - JSR
            new AssemblyInstruction(AssemblyInstructionType.AND, AddressingMode.IndexedIndirectX), //21 - AND - (Indirect,X)
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 22 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 23 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.BIT, AddressingMode.ZeroPage), // 24 - BIT - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.AND, AddressingMode.ZeroPage), //25 - AND - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.ROL, AddressingMode.ZeroPage), //26 - ROL - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 27 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.PLP, AddressingMode.Implied), // 28 - PLP
            new AssemblyInstruction(AssemblyInstructionType.AND, AddressingMode.Immidiate), //29 - AND - Immediate
            new AssemblyInstruction(AssemblyInstructionType.ROL, AddressingMode.Accumulator), //2A - ROL - Accumulator
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 2B - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.BIT, AddressingMode.Absolute), //2C - BIT - Absolute
            new AssemblyInstruction(AssemblyInstructionType.AND, AddressingMode.Absolute), //2D - AND - Absolute
            new AssemblyInstruction(AssemblyInstructionType.ROL, AddressingMode.Absolute), //2E - ROL - Absolute
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 2F - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.BMI, AddressingMode.Relative), //30 - BMI
            new AssemblyInstruction(AssemblyInstructionType.AND, AddressingMode.IndirectIndexedY), //31 - AND - (Indirect),Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 32 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 33 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 34 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.AND, AddressingMode.ZeroPageX), // 35 - AND - Zero Page,X
            new AssemblyInstruction(AssemblyInstructionType.ROL, AddressingMode.ZeroPageX), //36 - ROL - Zero Page,X
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 37 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.SEC, AddressingMode.Implied), //38 - SEC
            new AssemblyInstruction(AssemblyInstructionType.AND, AddressingMode.AbsoluteY), //39 - AND - Absolute,Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 3A - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 3B - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 3C - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.AND, AddressingMode.AbsoluteX), //3D - AND - Absolute,X
            new AssemblyInstruction(AssemblyInstructionType.ROL, AddressingMode.AbsoluteX), //3E - ROL - Absolute,X
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 3F - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.RTI, AddressingMode.Implied), //40 - RTI                        
            new AssemblyInstruction(AssemblyInstructionType.EOR, AddressingMode.IndexedIndirectX),//41 - EOR - (Indirect,X)         
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 42 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 43 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 44 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.EOR, AddressingMode.ZeroPage), //45 - EOR - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.LSR, AddressingMode.ZeroPage), //46 - LSR - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 47 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.PHA, AddressingMode.Implied), //48 - PHA                        
            new AssemblyInstruction(AssemblyInstructionType.EOR, AddressingMode.Immidiate), //49 - EOR - Immediate            
            new AssemblyInstruction(AssemblyInstructionType.LSR, AddressingMode.Accumulator), //4A - LSR - Accumulator          
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 4B - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.JMP, AddressingMode.Absolute), //4C - JMP - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.EOR, AddressingMode.Absolute), //4D - EOR - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.LSR, AddressingMode.Absolute), //4E - LSR - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 4F - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.BVC, AddressingMode.Relative), //50 - BVC                        
            new AssemblyInstruction(AssemblyInstructionType.EOR, AddressingMode.IndirectIndexedY),//51 - EOR - (Indirect),Y         
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 52 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 53 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 54 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.EOR, AddressingMode.ZeroPageX), //55 - EOR - Zero Page,X          
            new AssemblyInstruction(AssemblyInstructionType.LSR, AddressingMode.ZeroPageX), //56 - LSR - Zero Page,X          
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 57 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.CLI, AddressingMode.Implied), //58 - CLI                        
            new AssemblyInstruction(AssemblyInstructionType.EOR, AddressingMode.AbsoluteY), //59 - EOR - Absolute,Y           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 5A - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 5B - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 5C - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.EOR, AddressingMode.AbsoluteX), //50 - EOR - Absolute,X           
            new AssemblyInstruction(AssemblyInstructionType.LSR, AddressingMode.AbsoluteX), //5E - LSR - Absolute,X           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 5F - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.RTS, AddressingMode.Implied), //60 - RTS
            new AssemblyInstruction(AssemblyInstructionType.ADC, AddressingMode.IndexedIndirectX), //61 - ADC - (Indirect,X)
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 62 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 63 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 64 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.ADC, AddressingMode.ZeroPage), //65 - ADC - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.ROR, AddressingMode.ZeroPage), //66 - ROR - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 67 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.PLA, AddressingMode.Implied), //68 - PLA
            new AssemblyInstruction(AssemblyInstructionType.ADC, AddressingMode.Immidiate), //69 - ADC - Immediate
            new AssemblyInstruction(AssemblyInstructionType.ROR, AddressingMode.Accumulator), //6A - ROR - Accumulator
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 6B - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.JMP, AddressingMode.Indirect),  //6C - JMP - Indirect
            new AssemblyInstruction(AssemblyInstructionType.ADC, AddressingMode.Absolute), //6D - ADC - Absolute
            new AssemblyInstruction(AssemblyInstructionType.ROR, AddressingMode.Absolute), //6E - ROR - Absolute
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 6F - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.BVS, AddressingMode.Relative), //70 - BVS
            new AssemblyInstruction(AssemblyInstructionType.ADC, AddressingMode.IndirectIndexedY), //71 - ADC - (Indirect),Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 72 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 73 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 74 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.ADC, AddressingMode.ZeroPageX), //75 - ADC - Zero Page,X
            new AssemblyInstruction(AssemblyInstructionType.ROR, AddressingMode.ZeroPageX), //76 - ROR - Zero Page,X
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 77 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.SEI, AddressingMode.Implied), //78 - SEI
            new AssemblyInstruction(AssemblyInstructionType.ADC, AddressingMode.AbsoluteY), //79 - ADC - Absolute,Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 7A - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 7B - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 7C - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.ADC, AddressingMode.AbsoluteX), //7D - ADC - Absolute,X
            new AssemblyInstruction(AssemblyInstructionType.ROR, AddressingMode.AbsoluteX), //7E - ROR - Absolute,X
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 7F - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 80 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.STA, AddressingMode.IndexedIndirectX),//81 - STA - (Indirect,X)         
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 82 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 83 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.STY, AddressingMode.ZeroPage), //84 - STY - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.STA, AddressingMode.ZeroPage), //85 - STA - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.STX, AddressingMode.ZeroPage), //86 - STX - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 87 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.DEY, AddressingMode.Implied), //88 - DEY                        
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 89 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.TXA, AddressingMode.Implied), //8A - TXA                        
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 8B - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.STY, AddressingMode.Absolute), //8C - STY - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.STA, AddressingMode.Absolute), //8D - STA - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.STX, AddressingMode.Absolute), //8E - STX - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 8F - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.BCC, AddressingMode.Relative), //90 - BCC                        
            new AssemblyInstruction(AssemblyInstructionType.STA, AddressingMode.IndirectIndexedY),//91 - STA - (Indirect),Y         
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 92 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 93 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.STY, AddressingMode.ZeroPageX), //94 - STY - Zero Page,X          
            new AssemblyInstruction(AssemblyInstructionType.STA, AddressingMode.ZeroPageX), //95 - STA - Zero Page,X          
            new AssemblyInstruction(AssemblyInstructionType.STX, AddressingMode.ZeroPageY), //96 - STX - Zero Page,Y          
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 97 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.TYA, AddressingMode.Implied), //98 - TYA                        
            new AssemblyInstruction(AssemblyInstructionType.STA, AddressingMode.AbsoluteY), //99 - STA - Absolute,Y           
            new AssemblyInstruction(AssemblyInstructionType.TXS, AddressingMode.Implied), //9A - TXS                        
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 9B - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 9C - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.STA, AddressingMode.AbsoluteX), //90 - STA - Absolute,X           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 9E - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // 9F - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.LDY, AddressingMode.Immidiate), //A0 - LDY - Immediate
            new AssemblyInstruction(AssemblyInstructionType.LDA, AddressingMode.IndexedIndirectX), //A1 - LDA - (Indirect,X)
            new AssemblyInstruction(AssemblyInstructionType.LDX, AddressingMode.Immidiate), //A2 - LDX - Immediate
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // A3 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.LDY, AddressingMode.ZeroPage), //A4 - LDY - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.LDA, AddressingMode.ZeroPage), //A5 - LDA - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.LDX, AddressingMode.ZeroPage), //A6 - LDX - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // A7 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.TAY, AddressingMode.Implied), //A8 - TAY
            new AssemblyInstruction(AssemblyInstructionType.LDA, AddressingMode.Immidiate), //A9 - LDA - Immediate
            new AssemblyInstruction(AssemblyInstructionType.TAX, AddressingMode.Implied), //AA - TAX
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // AB - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.LDY, AddressingMode.Absolute), //AC - LDY - Absolute
            new AssemblyInstruction(AssemblyInstructionType.LDA, AddressingMode.Absolute), //AD - LDA - Absolute
            new AssemblyInstruction(AssemblyInstructionType.LDX, AddressingMode.Absolute), //AE - LDX - Absolute
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // AF - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.BCS, AddressingMode.Relative), //B0 - BCS
            new AssemblyInstruction(AssemblyInstructionType.LDA, AddressingMode.IndirectIndexedY), //B1 - LDA - (Indirect),Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // B2 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // B3 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.LDY, AddressingMode.ZeroPageX), //B4 - LDY - Zero Page,X
            new AssemblyInstruction(AssemblyInstructionType.LDA, AddressingMode.ZeroPageX), //B5 - LDA - Zero Page,X
            new AssemblyInstruction(AssemblyInstructionType.LDX, AddressingMode.ZeroPageY), //B6 - LDX - Zero Page,Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // B7 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.CLV, AddressingMode.Implied), //B8 - CLV
            new AssemblyInstruction(AssemblyInstructionType.LDA, AddressingMode.AbsoluteY), //B9 - LDA - Absolute,Y
            new AssemblyInstruction(AssemblyInstructionType.TSX, AddressingMode.Implied), //BA - TSX
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // BB - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.LDY, AddressingMode.AbsoluteX), //BC - LDY - Absolute,X
            new AssemblyInstruction(AssemblyInstructionType.LDA, AddressingMode.AbsoluteX), //BD - LDA - Absolute,X
            new AssemblyInstruction(AssemblyInstructionType.LDX, AddressingMode.AbsoluteY), //BE - LDX - Absolute,Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // BF - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.CPY, AddressingMode.Immidiate), //C0 - Cpy - Immediate            
            new AssemblyInstruction(AssemblyInstructionType.CMP, AddressingMode.IndexedIndirectX),//C1 - CMP - (Indirect,X)         
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // C2 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // C3 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.CPY, AddressingMode.ZeroPage), // C4 - CPY - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.CMP, AddressingMode.ZeroPage), // C5 - CMP - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.DEC, AddressingMode.ZeroPage), // C6 - DEC - Zero Page            
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // C7 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.INY, AddressingMode.Implied), // C8 - INY                        
            new AssemblyInstruction(AssemblyInstructionType.CMP, AddressingMode.Immidiate), // C9 - CMP - Immediate            
            new AssemblyInstruction(AssemblyInstructionType.DEX, AddressingMode.Implied), // CA - DEX                        
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // CB - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.CPY, AddressingMode.Absolute), // CC - CPY - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.CMP, AddressingMode.Absolute), // CD - CMP - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.DEC, AddressingMode.Absolute), // CE - DEC - Absolute             
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // CF - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.BNE, AddressingMode.Relative), // D0 - BNE                        
            new AssemblyInstruction(AssemblyInstructionType.CMP, AddressingMode.IndirectIndexedY),//D1 - CMP   (Indirect@,Y         
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // D2 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // D3 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // D4 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.CMP, AddressingMode.ZeroPageX), // D5 - CMP - Zero Page,X          
            new AssemblyInstruction(AssemblyInstructionType.DEC, AddressingMode.ZeroPageX), // D6 - DEC - Zero Page,X          
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // D7 - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.CLD, AddressingMode.Implied), //D8 - CLD                        
            new AssemblyInstruction(AssemblyInstructionType.CMP, AddressingMode.AbsoluteY), //D9 - CMP - Absolute,Y           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // DA - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // DB - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // DC - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.CMP, AddressingMode.AbsoluteX), //DD - CMP - Absolute,X           
            new AssemblyInstruction(AssemblyInstructionType.DEC, AddressingMode.AbsoluteX), //DE - DEC - Absolute,X           
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // DF - Future Expansion           
            new AssemblyInstruction(AssemblyInstructionType.CPX, AddressingMode.Immidiate), //E0 - CPX - Immediate
            new AssemblyInstruction(AssemblyInstructionType.SBC, AddressingMode.IndexedIndirectX), //E1 - SBC - (Indirect,X)
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // E2 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // E3 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.CPX, AddressingMode.ZeroPage), //E4 - CPX - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.SBC, AddressingMode.ZeroPage), //E5 - SBC - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.INC, AddressingMode.ZeroPage), //E6 - INC - Zero Page
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // E7 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.INX, AddressingMode.Implied), // E8 - INX
            new AssemblyInstruction(AssemblyInstructionType.SBC, AddressingMode.Immidiate), // E9 - SBC - Immediate
            new AssemblyInstruction(AssemblyInstructionType.NOP, AddressingMode.Implied), // EA - NOP
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // EB - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.CPX, AddressingMode.Absolute), //EC - CPX - Absolute
            new AssemblyInstruction(AssemblyInstructionType.SBC, AddressingMode.Absolute), //ED - SBC - Absolute
            new AssemblyInstruction(AssemblyInstructionType.INC, AddressingMode.Absolute), //EE - INC - Absolute
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // EF - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.BEQ, AddressingMode.Relative), //F0 - BEQ
            new AssemblyInstruction(AssemblyInstructionType.SBC, AddressingMode.IndirectIndexedY), //F1 - SBC - (Indirect),Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // F2 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // F3 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // F4 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.SBC, AddressingMode.ZeroPageX), //F5 - SBC - Zero Page,X
            new AssemblyInstruction(AssemblyInstructionType.INC, AddressingMode.ZeroPageX), // F6 - INC - Zero Page,X
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // F7 - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.SED, AddressingMode.Implied), //F8 - SED
            new AssemblyInstruction(AssemblyInstructionType.SBC, AddressingMode.AbsoluteY), // F9 - SBC - Absolute,Y
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // FA - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // FB - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // FC - Future Expansion
            new AssemblyInstruction(AssemblyInstructionType.SBC, AddressingMode.AbsoluteX), //FD - SBC - Absolute,X
            new AssemblyInstruction(AssemblyInstructionType.INC, AddressingMode.AbsoluteX), // FE - INC - Absolute,X
            new AssemblyInstruction(AssemblyInstructionType.UNDEF, AddressingMode.Implied), // FF - Future Expansion
        };

    }
}