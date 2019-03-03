package main

import (
	"encoding/json"
	"github.com/santatamas/go-c64/internals"
	"log"
	"strconv"
	"time"
)

type EmulatorState struct {
	Delay      time.Duration
	CycleCount uint64
	PauseFlag  bool
	Debug      bool
	Test       bool
	BreakAddr  uint16
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
			case internals.GetCIAState:
				{
					state, _ = json.Marshal(emu.CIA.GetState())
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
