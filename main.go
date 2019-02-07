package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	programPathPtr := flag.String("prg-path", "", "The relative path of the program binary to load.")
	delayPtr := flag.Int("delay", 10, "Artificial delay (in milliseconds) between CPU instructions.")
	debugPtr := flag.Bool("debug", false, "Debug mode. If true, you can open http://localhost{:8080} to access the internal debug viewer.")
	testModePtr := flag.Bool("test", false, "Test mode. If true, a test ROM is loaded instead of the standard C64 ROMs.")
	disableLogsPtr := flag.Bool("no-logs", false, "Disable logging. All log output is discarded.")
	flag.Parse()

	emulator := NewEmulator(*testModePtr)
	emulator.Delay = time.Duration(*delayPtr)

	if *debugPtr {
		log.Println("[DEBUG] Starting debug server...")
		hub := startDebugServer()
		debugServerLogger := DebugLog{hub: hub}
		log.SetOutput(debugServerLogger)

		emulator.Debug = *debugPtr
		emulator.hub = hub
	} else {
		// set normal file logging
		if !*disableLogsPtr {
			file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
			log.SetOutput(file)
		} else {
			log.SetOutput(ioutil.Discard)
		}
	}

	if *testModePtr {
		// set program counter to test ROM start address
		emulator.CPU.PC = 0x400
	} else if *programPathPtr == "" {
		// set program counter to hard reset address
		emulator.CPU.PC = 0xfce2

	} else {
		emulator.loadFile(*programPathPtr)
	}

	emulator.Start()
}
