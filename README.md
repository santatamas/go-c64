# Commodore 64 Emulator in Go

A functional Commodore 64 emulator written in Go that runs in your terminal. This project emulates the core hardware of the C64 including the MOS 6510 CPU, CIA chip, VIC-II display (text mode), and full 64KB memory with bank switching.

## Features

- **MOS 6510 CPU** - Full instruction set emulation with all 13 addressing modes
- **CIA (Complex Interface Adapter)** - Timer support, keyboard matrix scanning, and interrupt generation
- **Memory Bank Switching** - Accurate C64 memory map with ROM/RAM/IO bank configurations
- **VIC-II Text Display** - Terminal-based 40x25 character display using tcell
- **Keyboard Input** - Real-time keyboard handling mapped to the C64 keyboard matrix
- **PRG File Loading** - Load and run compiled 6502/6510 assembly binaries
- **Debug Mode** - WebSocket-based debug server with an Angular UI for stepping through code
- **Built-in Profiler** - Go pprof accessible via HTTP for performance analysis
- **6502 Test Suite** - Support for Klaus Dormann's functional test suite

## Requirements

- **Go 1.21+** (developed with Go 1.25)
- **Node.js 18+** (for the debug UI, optional)
- A terminal emulator with color support

## Quick Start

### Build

```bash
make build
```

### Run with C64 BASIC ROM

```bash
make run
```

### Run a PRG file

```bash
./dist/main -prg-path=_resources/prg/7_Snake.prg -no-logs=true
```

### Run with debug mode

```bash
./dist/main -debug=true
```

Then open the debug UI at `http://localhost:8080` (or serve the Angular build).

## Command-Line Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-prg-path` | `""` | Path to a `.prg` binary to load and execute |
| `-delay` | `0` | Artificial delay (ms) between CPU instructions |
| `-debug` | `false` | Enable debug server on port 8080 |
| `-test` | `false` | Load 6502 functional test ROM instead of C64 ROMs |
| `-no-logs` | `false` | Disable file logging (logs go to `log.txt` by default) |

## Architecture

```
go-c64/
├── main.go              # Entry point, CLI flag parsing
├── emulator.go          # Emulator struct, main loop, ROM loading
├── emulator_debug.go    # Debug command handler (WebSocket)
├── MOS6510/             # CPU emulation
│   ├── cpu.go           # CPU struct, cycle execution, instruction dispatch
│   ├── cpu_instructions.go  # All 56 instruction implementations
│   ├── cpu_status.go    # Processor status register (N,V,B,D,I,Z,C flags)
│   ├── cpu_stack.go     # Stack push/pop operations
│   ├── addressing_mode.go   # 13 addressing mode definitions
│   ├── assembly_instruction.go      # Opcode → instruction mapping (256 entries)
│   └── assembly_instruction_type.go # Instruction type enum
├── CIA/                  # Complex Interface Adapter
│   ├── cia.go           # CIA registers, timers, keyboard matrix I/O
│   ├── cia_debug.go     # CIA state serialization for debug UI
│   └── keyboard.go      # Key-to-matrix mapping, keyboard handling
├── RAM/                  # Memory subsystem
│   ├── memory.go        # 64KB memory with bank switching logic
│   ├── bank.go          # Bank type definitions (BASIC, CHARENG, KERNAL)
│   └── bank_configuration.go  # Bank configuration modes (ROM, RAM, IO)
├── VIC2/                 # Video chip (text mode)
│   ├── memory_display.go      # Terminal rendering with tcell
│   ├── c64_char_converter.go  # Screen code → ASCII character mapping
│   └── c64_color_converter.go # C64 color index → terminal color mapping
├── numeric/              # Utility functions
│   └── numeric.go       # Byte/uint16 conversions, bit operations
├── internals/            # Debug infrastructure
│   ├── hub.go           # WebSocket hub (client management)
│   ├── client.go        # WebSocket client (read/write pumps)
│   ├── server.go        # HTTP + WebSocket debug server
│   ├── command.go       # Debug command enum and JSON serialization
│   ├── telemetry.go     # Telemetry message struct
│   └── WebUI/           # Angular 19 debug frontend
└── _resources/           # ROMs, test programs, documentation
    ├── roms/            # C64 BASIC, KERNAL, and character ROMs
    ├── prg/             # Compiled assembly programs
    ├── bas/             # BASIC programs
    ├── tests/           # 6502 functional test suite
    └── Doc/             # Hardware documentation
```

## Hardware Emulation Details

### MOS 6510 CPU

The CPU emulation covers all official 6502/6510 instructions:

| Category | Instructions |
|----------|-------------|
| Load/Store | LDA, LDX, LDY, STA, STX, STY |
| Transfer | TAX, TXA, TAY, TYA, TSX, TXS |
| Arithmetic | ADC, SBC, INC, DEC, INX, DEX, INY, DEY |
| Logic | AND, ORA, EOR, BIT |
| Shift/Rotate | ASL, LSR, ROL, ROR |
| Compare | CMP, CPX, CPY |
| Branch | BCC, BCS, BEQ, BNE, BMI, BPL, BVC, BVS |
| Jump | JMP, JSR, RTS, RTI |
| Stack | PHA, PLA, PHP, PLP |
| Flags | SEC, CLC, SED, CLD, SEI, CLI, CLV |
| System | BRK, NOP |

**Addressing Modes:** Implied, Immediate, Zero Page, Zero Page X/Y, Absolute, Absolute X/Y, Indirect, Indexed Indirect (X), Indirect Indexed (Y), Relative, Accumulator

### Memory Map

The emulator implements the C64's bank switching mechanism controlled by the PLA (Programmable Logic Array) at address $0001:

| Address Range | Default Configuration |
|---------------|----------------------|
| $0000-$00FF | Zero Page (RAM) |
| $0100-$01FF | Stack (RAM) |
| $0200-$9FFF | RAM |
| $A000-$BFFF | BASIC ROM (switchable to RAM) |
| $C000-$CFFF | RAM |
| $D000-$D3FF | VIC-II I/O (switchable to Character ROM or RAM) |
| $D400-$DBFF | I/O area |
| $DC00-$DC0F | CIA 1 (keyboard, joystick) |
| $DD00-$DD0F | CIA 2 |
| $E000-$FFFF | KERNAL ROM (switchable to RAM) |

### CIA Chip

- Keyboard matrix scanning (8x8 matrix)
- Timer A and Timer B registers
- Interrupt generation for the CPU
- Port A/B read and write

### VIC-II (Text Mode)

- 40x25 character display
- 16 C64 colors mapped to terminal colors
- Background and border color registers
- Screen memory at $0400, color RAM at $D800

## Debug UI

The debug frontend is an Angular 19 application that communicates with the emulator via WebSocket.

### Building the Debug UI

```bash
cd internals/WebUI
npm install
npm run build
```

### Debug Commands

| Command | Description |
|---------|-------------|
| `Start` | Resume emulation |
| `Stop` | Pause emulation |
| `ExecuteNext` | Execute single instruction |
| `GetCPUState` | Get CPU register values |
| `GetCIAState` | Get CIA register values |
| `GetEmulatorState` | Get emulator state (cycle count, flags) |
| `GetMemoryContent` | Get full 64KB memory dump |
| `SetBreakpoint` | Set breakpoint at address |

### Debug UI Features

- Play/Pause/Step execution controls
- CPU register state table (A, X, Y, SP, PC, Status)
- CIA chip state (timers, interrupts, PORT_A)
- Emulator state (cycle count, pause flag, breakpoint address)
- Memory hex viewer with PC highlight (scrolling window around current instruction)

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test -v ./MOS6510/
go test -v ./CIA/
go test -v ./RAM/
go test -v ./VIC2/
go test -v ./numeric/
go test -v ./internals/
```

The test suite includes 285 tests covering:
- All CPU instructions with multiple addressing modes
- CPU status flag operations
- Stack push/pop behavior
- Memory bank switching
- CIA keyboard matrix operations
- VIC-II character and color conversion
- Numeric utility functions
- Debug command serialization

## Makefile Targets

| Target | Description |
|--------|-------------|
| `make build` | Build the emulator binary |
| `make run` | Build and run with BASIC ROM (no logs) |
| `make run-log` | Build and run with logging enabled |
| `make run-snake` | Run the Snake game |
| `make run-colors` | Run the color demo |
| `make debug` | Run in debug mode |
| `make debug-test-suite` | Run with 6502 test suite in debug mode |

## Profiling

The emulator exposes a Go pprof HTTP server on port 8888:

```bash
# View available profiles
open http://localhost:8888/debug/pprof

# Generate a CPU profile
go tool pprof http://localhost:8888/debug/pprof/profile?seconds=30

# Generate a flame graph
docker run uber/go-torch -u http://<host>:8888/debug/pprof -p -t=30 > torch.svg
```

## Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| github.com/gdamore/tcell/v2 | v2.7.4 | Terminal UI rendering |
| github.com/gorilla/websocket | v1.5.3 | WebSocket for debug server |

## Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests for new functionality
4. Ensure all tests pass: `go test ./...`
5. Submit a pull request

## Known Limitations

- VIC-II only supports text mode (no bitmap/sprite graphics)
- CIA timers are simplified (no full cycle-accurate timing)
- No SID (sound) chip emulation
- No joystick support
- Illegal/undocumented 6502 opcodes are not implemented

## References

- [6502 Instruction Set Test Suite by Klaus Dormann](https://github.com/Klaus2m5/6502_65C02_functional_tests)
- [C64 Wiki](https://www.c64-wiki.com/wiki/Main_Page)
- [C64 ROM Disassembly](http://www.zimmers.net/anonftp/pub/cbm/src/c64/c64_rom_disassembly.txt)
- [Mapping the C64](http://www.unusedino.de/ec64/technical/project64/mapping_c64.html)
- [C64 Assembly Programming Tutorial](https://github.com/petriw/Commodore64Programming)
- [C64 Devkit](https://github.com/cliffordcarnmo/c64-devkit)
- [Emudore](https://github.com/marioballano/emudore)
- [VirtualC64](https://github.com/dirkwhoffmann/virtualc64)

## License

This software comes as-is and is not meant to be production-ready. Use it at your own risk.

## Acknowledgments

Special thanks to [Balazs Molnar](https://www.linkedin.com/in/balazsm1) for providing the initial test assembly apps and references.
