# Development Guide

## Prerequisites

- Go 1.21 or later
- Node.js 18+ and npm (for debug UI development)
- A terminal with color support (for running the emulator)

## Setting Up the Development Environment

```bash
# Clone the repository
git clone https://github.com/santatamas/go-c64.git
cd go-c64

# Download Go dependencies
go mod tidy

# Build the project
go build -o ./dist/main

# Run tests
go test ./...
```

## Project Structure

```
go-c64/
├── main.go                  # Application entry point
├── emulator.go              # Core emulator logic
├── emulator_debug.go        # Debug mode handler
├── Makefile                 # Build and run targets
├── go.mod                   # Go module definition
├── MOS6510/                 # CPU package
├── CIA/                     # CIA chip package
├── RAM/                     # Memory package
├── VIC2/                    # Display package
├── numeric/                 # Utility functions
├── internals/               # Debug server
│   └── WebUI/              # Angular debug frontend
├── _resources/              # ROMs, test programs, docs
├── docs/                    # Project documentation
└── dist/                    # Build output
```

## Running the Emulator

### With C64 BASIC ROM (default)

```bash
make run
```

This loads the BASIC, KERNAL, and character ROMs and starts execution at the hard reset vector ($FCE2).

### With a PRG file

```bash
./dist/main -prg-path=_resources/prg/7_Snake.prg -no-logs=true
```

PRG files contain a 2-byte load address header followed by the program data.

### In Debug Mode

```bash
./dist/main -debug=true
```

The emulator starts paused. Use the debug UI or WebSocket commands to step through execution.

## Debug UI Development

The debug frontend is an Angular 19 standalone application.

```bash
cd internals/WebUI

# Install dependencies
npm install

# Development server (hot reload)
npm start
# Access at http://localhost:4200

# Production build
npm run build
# Output at dist/WebUI/
```

### Serving the built UI

After building, you can serve the static files with any HTTP server:

```bash
cd internals/WebUI/dist/WebUI/browser
npx http-server . -p 4200
```

## Adding New CPU Instructions

1. Add the instruction type to `MOS6510/assembly_instruction_type.go`
2. Add the opcode mapping in `MOS6510/assembly_instruction.go`
3. Implement the instruction method in `MOS6510/cpu_instructions.go`
4. Add the dispatch case in the `callMethod` switch in `MOS6510/cpu.go`
5. Write tests in `MOS6510/cpu_instructions_test.go`
6. Regenerate stringer: `go generate ./MOS6510/`

### Instruction Implementation Pattern

```go
// EXAMPLE Instruction description
// Operation:  description of what it does
// N Z C I D V    (which flags are affected)
// * * _ _ _ _    (* = affected, _ = unchanged)
func (cpu *CPU) EXAMPLE(mode AddressingMode) {
    log.Println("EXAMPLE called -- adr. mode: ", mode.String())
    
    // Read operand based on addressing mode
    mem := cpu.getMemoryValue(mode)
    
    // Perform operation
    cpu.A = mem
    
    // Update status flags
    cpu.setStatusZero(cpu.A == 0)
    cpu.setStatusNegative(cpu.A & 0x80 != 0)
}
```

## Testing

### Running All Tests

```bash
go test ./...
```

### Running with Coverage

```bash
go test -cover ./...
```

### Running a Specific Test

```bash
go test -v -run TestCPU_LDA_Immediate ./MOS6510/
```

### Test Organization

Each package has its own `_test.go` files:

| Package | Test File | Coverage |
|---------|-----------|----------|
| MOS6510 | `cpu_instructions_test.go` | All 56 instructions, multiple addressing modes |
| MOS6510 | `cpu_status_test.go` | Status flag set/clear operations |
| MOS6510 | `cpu_stack_test.go` | Stack push/pop, LIFO order, address range |
| CIA | `cia_test.go` | Keyboard matrix, interrupt, register I/O |
| RAM | `memory_test.go` | Bank switching, CIA routing, ROM/RAM access |
| VIC2 | `c64_char_converter_test.go` | Screen code → character mapping |
| VIC2 | `c64_color_converter_test.go` | Color index → terminal color |
| numeric | `numeric_test.go` | Byte operations, bit testing |
| internals | `command_test.go` | Command serialization/deserialization |

## Debugging Tips

### Enable File Logging

By default (without `-no-logs`), the emulator logs extensively to `log.txt`:

```bash
./dist/main &
tail -f log.txt
```

### Using the Profiler

```bash
# CPU profile
go tool pprof http://localhost:8888/debug/pprof/profile?seconds=10

# Memory profile
go tool pprof http://localhost:8888/debug/pprof/heap

# Goroutine dump
curl http://localhost:8888/debug/pprof/goroutine?debug=1
```

### Running the 6502 Test Suite

Klaus Dormann's comprehensive test suite validates all CPU instructions:

```bash
./dist/main -debug=true -test=true
```

This loads the test ROM at $0400 and begins execution. Use the debug UI to verify the test progresses through all instruction tests (the test loops infinitely on failure at the failing test's address).

## Code Generation

The project uses Go's `stringer` tool for enum string methods:

```bash
# Regenerate string methods
cd MOS6510
go generate
```

This produces:
- `addressingmode_string.go` - AddressingMode.String()
- `assemblyinstructiontype_string.go` - AssemblyInstructionType.String()

## ROM Files

The emulator requires C64 ROM files in `_resources/roms/`:

| File | Size | Description |
|------|------|-------------|
| `basic.901226-01.bin` | 8KB | BASIC interpreter |
| `kernal.901227-03.bin` | 8KB | KERNAL operating system |
| `characters.901225-01.bin` | 4KB | Character set |

These are loaded at their respective addresses during initialization:
- BASIC → $A000
- Characters → $D000
- KERNAL → $E000
