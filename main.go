package main

import (
	//"bufio"
	"flag"
	"github.com/santatamas/go-c64/internals"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	programPathPtr := flag.String("prg-path", "", "The relative path of the program binary to load.")
	delayPtr := flag.Int("delay", 0, "Artificial delay (in milliseconds) between CPU instructions.")
	debugPtr := flag.Bool("debug", false, "Debug mode. If true, you can open http://localhost{:8080} to access the internal debug viewer.")
	testModePtr := flag.Bool("test", false, "Test mode. If true, a test ROM is loaded instead of the standard C64 ROMs.")
	disableLogsPtr := flag.Bool("no-logs", false, "Disable logging. All log output is discarded.")
	flag.Parse()

	emulator := NewEmulator(*testModePtr)
	emulator.Delay = time.Duration(*delayPtr)

	// set normal file logging
	if !*disableLogsPtr {
		file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		log.SetOutput(file)
		log.Println("[DEBUG] Logging to file: log.txt")
	} else {
		log.Println("[DEBUG] Logging is disabled.")
		log.SetOutput(ioutil.Discard)
	}

	if *debugPtr {
		log.Println("[DEBUG] Starting debug server...")
		hub := internals.StartDebugServer()

		// Uncomment to forward logs to websocket
		// [WARNING] - It's prone to buffer overflow in the socket reader/writer
		//debugServerLogger := internals.DebugLog{Hub: hub}
		//log.SetOutput(debugServerLogger)

		emulator.Debug = *debugPtr
		emulator.hub = hub
	}

	if *testModePtr {
		log.Println("[DEBUG] Test mode ON, setting PC to test ROM start address")
		// set program counter to test ROM start address
		emulator.CPU.PC = 0x400
	} else if *programPathPtr == "" {
		log.Println("[DEBUG] Setting PC to hard reset address")
		// set program counter to hard reset address
		emulator.CPU.PC = 0xfce2
	} else {
		log.Println("[DEBUG] Loading PRG...")
		emulator.loadFile(*programPathPtr)
	}

	emulator.Start()

	//bufio.NewReader(os.Stdin).ReadBytes('\n')
}
