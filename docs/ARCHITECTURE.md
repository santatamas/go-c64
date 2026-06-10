# Architecture Documentation

## System Overview

The go-c64 emulator follows a modular architecture that mirrors the physical hardware of the Commodore 64. Each major hardware component is represented by its own Go package.

```
┌─────────────────────────────────────────────────────┐
│                    main.go                           │
│              (CLI, initialization)                   │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                 Emulator                             │
│         (Main loop, coordination)                   │
├──────────┬──────────┬──────────┬────────────────────┤
│  MOS6510 │   CIA    │   RAM    │       VIC2         │
│  (CPU)   │  (I/O)   │ (Memory) │    (Display)       │
└──────────┴──────────┴──────────┴────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│               internals                             │
│    (Debug server, WebSocket, Telemetry)             │
└─────────────────────────────────────────────────────┘
```

## Package Details

### `main` (root package)

**Files:** `main.go`, `emulator.go`, `emulator_debug.go`

The main package ties everything together:

- **main.go** - Parses CLI flags, initializes the emulator, starts the pprof server
- **emulator.go** - `Emulator` struct coordinates CPU, CIA, memory, and display. The main emulation loop runs in a goroutine executing CPU cycles while the display runs in the main goroutine
- **emulator_debug.go** - Handles debug commands received via WebSocket (start/stop/step/get state)

**Emulation Loop:**
```
for {
    CIA.ExecuteCycle()          // Update CIA timers
    if CIA.Interrupt → CPU.Interrupt()  // Handle IRQ
    CPU.ExecuteCycle()          // Fetch-decode-execute one instruction
}
```

### `MOS6510` - CPU Emulation

**Key types:**
- `CPU` - Main CPU struct with registers (A, X, Y, SP, PC, S), memory pointer, instruction resolver
- `AssemblyInstruction` - Opcode mapping (type + addressing mode + cycle count)
- `AddressingMode` - Enum of 13 addressing modes
- `AssemblyInstructionType` - Enum of 56 instruction types

**Execution flow:**
1. Read opcode byte from memory at PC
2. Increment PC
3. Look up `AssemblyInstruction` from opcode table
4. Dispatch to instruction handler method
5. Instruction reads operands (incrementing PC as needed)
6. Instruction modifies CPU state (registers, flags, memory)

**Status Register (P):**
```
Bit: 7 6 5 4 3 2 1 0
     N V - B D I Z C
     │ │   │ │ │ │ └─ Carry
     │ │   │ │ │ └─── Zero
     │ │   │ │ └───── IRQ Disable
     │ │   │ └─────── Decimal Mode
     │ │   └───────── Break
     │ └───────────── Overflow
     └─────────────── Negative
```

### `CIA` - Complex Interface Adapter

**Key types:**
- `CIA` - CIA chip state (timers, keyboard matrix, ports, interrupt flag)
- `Keyboard` - Keyboard input handler

**Keyboard scanning:** The C64 uses an 8x8 keyboard matrix. When the CPU writes to PORT_A ($DC00), it selects which rows to scan. Reading PORT_B ($DC01) returns which columns have keys pressed in those rows.

### `RAM` - Memory Subsystem

**Key types:**
- `Memory` - Dual 64KB arrays (RAM + ROM) with bank switching logic
- `Bank` - Bank identifiers (BASIC, CHARENG, KERNAL)
- `BankConfiguration` - Bank modes (ROM, RAM, IO)

**Bank switching:** Writing to address $0001 configures memory banks via three control bits:
- LORAM (bit 0) - BASIC ROM visibility
- HIRAM (bit 1) - KERNAL ROM visibility
- CHAREN (bit 2) - Character ROM vs I/O area

### `VIC2` - Video Display

**Key types:**
- `MemoryDisplay` - Terminal-based display using tcell

**Display loop:** Runs at ~20fps, reading screen memory ($0400-$07E7) and color RAM ($D800-$DBE7) to render the 40x25 character grid in the terminal.

### `numeric` - Utility Package

Helper functions for byte manipulation:
- `ToInt16` / `ToInt16_2` - Little-endian byte-to-uint16 conversion
- `GetLO` / `GetHI` - Extract low/high bytes from uint16
- `IsNegative` - Check bit 7 (sign bit)
- `GetBit` - Test arbitrary bit position

### `internals` - Debug Infrastructure

**Components:**
- `Hub` - WebSocket hub managing connected debug clients
- `Client` - Individual WebSocket connection with read/write goroutines
- `Command` - Enum of debug commands with JSON marshaling
- `Telemetry` - Message format for debug communication

## Concurrency Model

The emulator uses goroutines for:
1. **Main goroutine** - Display rendering loop (tcell event loop)
2. **Emulation goroutine** - CPU/CIA execution loop
3. **Event handler goroutine** - Channels for keyboard events and IRQ signals
4. **Debug server goroutines** - HTTP server, WebSocket hub, per-client read/write

Communication between goroutines uses Go channels:
- `keyPressChannel` - Terminal key events → keyboard handler
- `irqChannel` - CIA interrupt signals → CPU

## Data Flow

```
Terminal Input → tcell → keyPressChannel → CIA.SetKey() → keyboard matrix
                                                              │
CPU reads $DC01 ←── CIA.ReadRegister() ←── keyboard matrix ──┘

CIA timer → interrupt → irqChannel → CPU.SetIRQ() → CPU.Interrupt()
                                                        │
                                          pushes PC & status to stack
                                          jumps to IRQ vector ($FFFE)
```
