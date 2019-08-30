package main

import (
	//"bufio"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/santatamas/go-c64/internals"
)

func main() {

	programPathPtr := flag.String("prg-path", "", "The relative path of the program binary to load.")
	delayPtr := flag.Int("delay", 0, "Artificial delay (in milliseconds) between CPU instructions.")
	debugPtr := flag.Bool("debug", false, "Debug mode. If true, you can open http://localhost{:8080} to access the internal debug viewer.")
	testModePtr := flag.Bool("test", false, "Test mode. If true, a test ROM is loaded instead of the standard C64 ROMs.")
	disableLogsPtr := flag.Bool("no-logs", false, "Disable logging. All log output is discarded.")
	flag.Parse()

	if !*disableLogsPtr {
		// set normal file logging
		file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		log.SetOutput(file)
		log.Println("[DEBUG] Logging to file: log.txt")
	} else {
		log.Println("[DEBUG] Logging is disabled.")
		log.SetOutput(ioutil.Discard)
	}

	emulator := NewEmulator(*testModePtr)
	emulator.Delay = time.Duration(*delayPtr)

	if *debugPtr {
		// Start custom debug server (using a Websocket connection to send telemetry)
		log.Println("[DEBUG] Starting debug server...")
		hub := internals.StartDebugServer()
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
		log.Println("[DEBUG] Loading PRG, setting PC to custom address")
		emulator.loadFile(*programPathPtr)
	}

	// Starting up the PPROF debug server
	s := &http.Server{
		Addr: ":8888",
	}
	go s.ListenAndServe()

	// Starting emulation and display/keyboard input listeners
	emulator.Start()
}
