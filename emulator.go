package main

import (
	"encoding/json"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/santatamas/go-c64/CIA"
	"github.com/santatamas/go-c64/MOS6510"
	"github.com/santatamas/go-c64/RAM"
	"github.com/santatamas/go-c64/VIC2"
	"github.com/santatamas/go-c64/internals"
	n "github.com/santatamas/go-c64/numeric"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Emulator struct {
	CPU        *MOS6510.CPU
	CIA        *CIA.CIA
	Display    *VIC2.MemoryDisplay
	Delay      time.Duration
	cycleCount int64
	pauseFlag  bool
	Debug      bool
	Test       bool
	BreakAddr  uint16
	hub        *internals.Hub
}

type EmulatorState struct {
	Delay      time.Duration
	CycleCount int64
	PauseFlag  bool
	Debug      bool
	Test       bool
	BreakAddr  uint16
}

func NewEmulator(testMode bool) Emulator {

	keyPressChannel := make(chan *tcell.EventKey)
	irqChannel := make(chan bool)

	cia := CIA.NewCIA(irqChannel)
	memory := RAM.NewMemory(testMode, &cia)
	cpu := MOS6510.NewCPU(&memory)
	keyboard := CIA.NewKeyboard(&cia)
	display := VIC2.NewMemoryDisplay(&memory, keyPressChannel)

	// Handle async events, like keypresses from terminal, and IRQ events (CIA->CPU)
	go func() {
		for {
			select {
			case key := <-keyPressChannel:
				log.Print("presskey channel new message")
				keyboard.PressKey(key)
			case irq := <-irqChannel:
				log.Print("irq channel new message")
				cpu.SetIRQ(irq)
			}

			//log.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			//time.Sleep(100)
		}
	}()

	return Emulator{
		CPU:        &cpu,
		CIA:        &cia,
		Delay:      0,
		Display:    &display,
		cycleCount: 0,
		pauseFlag:  false,
		Debug:      false,
		Test:       false,
	}
}

func (emu *Emulator) Start() {

	if emu.Debug {
		emu.pauseFlag = true
		go emu.ListenToCommands()
	}

	go func() {
		for {
			// Handle breakpoint
			if !emu.pauseFlag && emu.BreakAddr != 0 && emu.BreakAddr == emu.CPU.PC {
				emu.pauseFlag = true
				log.Printf("[DEBUG] Breakpoint hit on address: %x \n", emu.BreakAddr)
			}

			if !emu.pauseFlag {
				emu.CIA.ExecuteCycle()
				result := emu.CPU.ExecuteCycle()
				if !result {
					break
				}

				emu.cycleCount++
			}

			if emu.pauseFlag {
				time.Sleep(100)
			}
		}
	}()

	emu.Display.Start()
}

func (emu *Emulator) ListenToCommands() {
	for {
		select {
		case request := <-emu.hub.Broadcast:

			var telemetryRequest internals.Telemetry
			json.Unmarshal([]byte(request), &telemetryRequest)

			if err := json.Unmarshal([]byte(request), &telemetryRequest); err != nil {
				panic(err)
			}

			log.Println("[DEBUG] Command received:", telemetryRequest)

			var command internals.Command
			command.UnmarshalJSON(telemetryRequest.Command)

			var state []byte
			switch command {
			case internals.Start:
				{
					emu.pauseFlag = false
					state, _ = json.Marshal("OK")
				}
			case internals.Stop:
				{
					emu.pauseFlag = true
					state, _ = json.Marshal("OK")
				}
			case internals.ExecuteNext:
				{
					emu.CPU.ExecuteCycle()
					emu.cycleCount++
					state, _ = json.Marshal("OK")
				}
			case internals.GetCPUState:
				{
					state, _ = json.Marshal(emu.CPU.GetState())
				}
			case internals.GetEmulatorState:
				{
					emuState := EmulatorState{
						Delay:      emu.Delay,
						CycleCount: emu.cycleCount,
						Debug:      emu.Debug,
						Test:       emu.Test,
						PauseFlag:  emu.pauseFlag,
						BreakAddr:  emu.BreakAddr,
					}
					state, _ = json.Marshal(emuState)
				}
			case internals.SetBreakpoint:
				{
					breakAddr, _ := strconv.ParseUint(telemetryRequest.Parameter, 10, 16)
					emu.BreakAddr = uint16(breakAddr)
					log.Printf("[DEBUG] Breakpoint address set on: %x \n", emu.BreakAddr)
				}
			case internals.GetMemoryContent:
				{
					state = emu.CPU.Memory.ReadAll()
				}
			}

			telemetry := internals.Telemetry{
				Command: telemetryRequest.Command,
				Payload: state,
			}
			response, _ := json.Marshal(telemetry)
			emu.hub.Broadcast <- string(response)
		}
	}
}

func (emu *Emulator) loadFile(path string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error! Can't open file")
	}

	byteContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	startPCH := byteContent[0]
	startPCL := byteContent[1]
	startAddress := n.ToInt16([]byte{startPCH, startPCL})

	currentAddress := startAddress
	for i := 2; i < len(byteContent); i++ {
		emu.CPU.Memory.WriteAbsolute(currentAddress, byteContent[i])
		currentAddress++
	}

	emu.CPU.PC = n.ToInt16_2(startPCH, startPCL)
}
