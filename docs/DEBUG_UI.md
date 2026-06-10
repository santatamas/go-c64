# Debug UI Documentation

## Overview

The debug UI is an Angular 19 web application that connects to the emulator via WebSocket to provide real-time debugging capabilities. It allows you to:

- Start, stop, and step through CPU execution
- Inspect CPU registers, CIA state, and emulator state
- View memory contents with the current PC highlighted
- Set breakpoints at specific memory addresses

## Setup

### Prerequisites

- Node.js 18 or later
- npm

### Installation

```bash
cd internals/WebUI
npm install
```

### Development

```bash
npm start
```

The development server runs at `http://localhost:4200` with hot module reloading.

### Production Build

```bash
npm run build
```

Output is generated in `dist/WebUI/browser/`.

## Architecture

The debug UI uses Angular 19 standalone components with Angular Material for the UI.

### Technology Stack

| Technology | Version | Purpose |
|-----------|---------|---------|
| Angular | 19.2 | Application framework |
| Angular Material | 19.2 | UI component library |
| RxJS | 7.8 | Reactive WebSocket handling |
| TypeScript | 5.5 | Type-safe development |
| Zone.js | 0.15 | Change detection |

### Component Structure

```
src/app/
├── app.component.ts         # Root shell (toolbar + layout)
├── controls/
│   └── controls.component.ts  # Toolbar buttons (play/pause/step/breakpoint)
├── cpu/
│   └── cpu.component.ts       # CPU register state table
├── cia/
│   └── cia.component.ts       # CIA chip state table
├── emulator/
│   └── emulator.component.ts  # Emulator state table
├── memory/
│   └── memory.component.ts    # Hex memory viewer
├── models/                    # TypeScript interfaces
│   ├── telemetry.model.ts
│   ├── cpustate.model.ts
│   ├── ciastate.model.ts
│   ├── emulatorstate.model.ts
│   └── tablerow.model.ts
└── services/
    ├── websocket.service.ts    # Low-level WebSocket connection
    ├── telemetry.service.ts    # Command sending/receiving
    └── tablehelper.service.ts  # Object → table row conversion
```

## WebSocket Protocol

### Connection

The UI connects to `ws://localhost:8080/ws` on startup.

### Message Format

All messages use the same JSON structure:

```json
{
  "Command": "CommandName",
  "Parameter": "optional_parameter",
  "Payload": "base64_encoded_json_response"
}
```

### Request/Response Flow

1. UI sends a command (e.g., `{"Command": "GetCPUState"}`)
2. Emulator processes the command
3. Emulator responds with the same command name and a base64-encoded JSON payload

### Commands

#### Start
Resume emulation from paused state.
```json
{"Command": "Start"}
→ {"Command": "Start", "Payload": "Ik9LIg=="}  // "OK" base64
```

#### Stop
Pause emulation.
```json
{"Command": "Stop"}
→ {"Command": "Stop", "Payload": "Ik9LIg=="}
```

#### ExecuteNext
Execute a single CPU instruction (step).
```json
{"Command": "ExecuteNext"}
→ {"Command": "ExecuteNext", "Payload": "Ik9LIg=="}
```

#### GetCPUState
Get current CPU register values.
```json
{"Command": "GetCPUState"}
→ {"Command": "GetCPUState", "Payload": "<base64 CPUState JSON>"}
```

Response payload (decoded):
```json
{
  "A": 0,
  "Y": 0,
  "X": 0,
  "S": 48,
  "P": 0,
  "PC": 64738,
  "SP": 255,
  "SP_LOW": 256,
  "SP_HIGH": 0,
  "InstructionType": 120,
  "InstructionName": "SEI"
}
```

#### GetCIAState
Get CIA chip register values.
```json
{"Command": "GetCIAState"}
→ {"Command": "GetCIAState", "Payload": "<base64 CIAState JSON>"}
```

#### GetEmulatorState
Get emulator-level state.
```json
{"Command": "GetEmulatorState"}
→ {"Command": "GetEmulatorState", "Payload": "<base64 EmulatorState JSON>"}
```

#### GetMemoryContent
Get full 64KB memory dump.
```json
{"Command": "GetMemoryContent"}
→ {"Command": "GetMemoryContent", "Payload": "<base64 raw bytes>"}
```

#### SetBreakpoint
Set a breakpoint at a decimal address.
```json
{"Command": "SetBreakpoint", "Parameter": "1024"}
```

The UI converts hex input to decimal before sending.

## Usage Guide

### Starting a Debug Session

1. Start the emulator in debug mode: `./dist/main -debug=true`
2. Start the debug UI: `cd internals/WebUI && npm start`
3. Open `http://localhost:4200` in your browser

The emulator starts paused. You'll see the toolbar with control buttons.

### Toolbar Controls

| Icon | Action | Description |
|------|--------|-------------|
| Play | Start | Resume continuous execution |
| Pause | Stop | Pause execution |
| Skip Next | Step | Execute one instruction |
| Memory | CPU State | Fetch current CPU registers |
| Computer | Emulator State | Fetch emulator state |
| Keyboard | CIA State | Fetch CIA chip state |
| SD Card | Memory | Fetch memory contents |

### Setting Breakpoints

1. Enter a hex address in the breakpoint input (e.g., `FCE2`)
2. Click "Set BP"
3. Start execution - it will pause when PC reaches that address

### Memory Viewer

The memory viewer shows a window of ~200 bytes around the current PC:
- Each line shows 16 bytes in hexadecimal
- The address prefix shows the offset (e.g., `0000fce0 |`)
- The current PC instruction is highlighted in yellow with red text

### Refreshing State

When you click Step or Start/Stop, the UI automatically refreshes all state panels (CPU, CIA, Emulator, Memory) with staggered timing to avoid WebSocket congestion.
