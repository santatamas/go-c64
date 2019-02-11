package main

import (
	"encoding/json"
	"fmt"
	"github.com/santatamas/go-c64/MOS6510"
	"github.com/santatamas/go-c64/RAM"
	"github.com/santatamas/go-c64/VIC2"
	"github.com/santatamas/go-c64/internals"
	n "github.com/santatamas/go-c64/numeric"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Emulator struct {
	CPU        *MOS6510.CPU
	Display    *VIC2.MemoryDisplay
	Delay      time.Duration
	cycleCount int64
	pauseFlag  bool
	Debug      bool
	Test       bool
	hub        *internals.Hub
}

type EmulatorState struct {
	Delay      time.Duration
	CycleCount int64
	PauseFlag  bool
	Debug      bool
	Test       bool
}

func NewEmulator(testMode bool) Emulator {

	memory := RAM.NewMemory(testMode)
	display := VIC2.NewMemoryDisplay(&memory)
	cpu := MOS6510.NewCPU(&memory)

	return Emulator{
		CPU:        &cpu,
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

		go func() {
			for {
				select {
				case message := <-emu.hub.Broadcast:
					var command internals.Command
					command.UnmarshalJSON([]byte(message))

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
							}
							state, _ = json.Marshal(emuState)
						}
					case internals.GetMemoryContent:
						{
							state = emu.CPU.Memory.ReadAll()
						}
					}

					telemetry := internals.Telemetry{
						Command: message,
						Payload: state,
					}
					response, _ := json.Marshal(telemetry)
					emu.hub.Broadcast <- string(response)
				}
			}
		}()
	}

	go func() {
		for {
			if !emu.pauseFlag {
				result := emu.CPU.ExecuteCycle()
				if !result {
					break
				}

				emu.cycleCount++
			}
		}
	}()

	//emu.Display.Start()
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
