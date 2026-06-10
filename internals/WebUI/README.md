# C64 Emulator Debug UI

An Angular 19 web application for debugging the go-c64 emulator in real-time via WebSocket.

## Features

- Real-time CPU register inspection
- CIA chip state monitoring
- Emulator state tracking
- Memory hex viewer with PC highlighting
- Execution control (start/stop/step)
- Breakpoint support

## Quick Start

```bash
# Install dependencies
npm install

# Start development server
npm start

# Production build
npm run build
```

## Requirements

- Node.js 18+
- npm 9+
- The go-c64 emulator running with `-debug=true` flag

## Technology Stack

- Angular 19.2
- Angular Material 19.2
- RxJS 7.8
- TypeScript 5.5

## Usage

1. Start the emulator: `../../dist/main -debug=true`
2. Start the dev server: `npm start`
3. Open http://localhost:4200

See [docs/DEBUG_UI.md](../../docs/DEBUG_UI.md) for detailed documentation.
