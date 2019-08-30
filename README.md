# Commodore 64 emulator written in Go
A functional (albeit not feature complete) Commodore64 emulator, running in your terminal. This software comes as-is, and it's not meant to be production-ready. Use it at your own risk!

### Requirements
For the emulator, all you need is a properly set up [Go](https://golang.org) installation.
To run the debug client, you'll have to install [Angular6+](https://angular.io), then install the dependencies via npm.

### How to build the emulator
Run `make build` in the root directory.

### Loading & running binaries
By default, the emulator will load the official C64 BASIC ROM. If you'd like to specify a custom binary, you can use the `--prg-path` flag (see makefile for examples).

### Running unit tests
Run `make test` in the root directory

### Using the internal debug server
1. Start the emulator with the `-debug=true` flag, or run `make debug`
2. Start the Angular application by running `make debug-client`

### Misc
For debugging, you can load and run a complete 6502 instruction set test suite by running `make debug-test-suite`.
Or alternatively, try one of the compiled PRGs from the `_resources/prg` folder (see Makefile for details).

### References & thanks
I'd like to list a few references that I used to learn more about C64, Assembly, etc. They're all invaluable, and were massive help during development.
[6502 instruction set test suite by Klaus Dormann](https://github.com/Klaus2m5/6502_65C02_functional_tests)
[C64 Wiki](https://www.c64-wiki.com/wiki/Main_Page)
[C64 ROM Disassembly](http://www.zimmers.net/anonftp/pub/cbm/src/c64/c64_rom_disassembly.txt)
[Mapping the C64](http://www.unusedino.de/ec64/technical/project64/mapping_c64.html)
[C64 Assembly programming tutorial](https://github.com/petriw/Commodore64Programming)
[C64 Devkit](https://github.com/cliffordcarnmo/c64-devkit)
[Emudore](https://github.com/marioballano/emudore)
[VirtualC64](https://github.com/dirkwhoffmann/virtualc64)

Special thanks to [Balazs Molnar](https://www.linkedin.com/in/balazsm1) for providing the initial test assembly apps & references!

