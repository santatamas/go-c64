package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func main() {

	programPathPtr := flag.String("programPath", "./_resources/Phase1/Prg/4_Colors.prg", "The relative path of the program binary to load.")
	delayPtr := flag.Int("delay", 0, "Artificial delay (in milliseconds) between CPU instructions.")
	debugPtr := flag.Bool("debug", false, "Debug mode. If true, you can open http://localhost{:8080} to access the internal debug viewer.")
	flag.Parse()

	emulator := NewEmulator()
	emulator.Delay = time.Duration(*delayPtr)

	if *debugPtr {
		hub := startDebugServer()
		debugServerLogger := DebugLog{hub: hub}
		log.SetOutput(debugServerLogger)

		emulator.Debug = *debugPtr
		emulator.hub = hub
	} else {
		// set normal file logging
		file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		log.SetOutput(file)
	}

	emulator.loadFile(*programPathPtr)
	emulator.Start()
}