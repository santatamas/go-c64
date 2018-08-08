package main

import (
	"fmt"
	"github.com/santatamas/go-c64/MOS6510"
	"github.com/santatamas/go-c64/RAM"
	"github.com/santatamas/go-c64/VIC2"
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
}

func NewEmulator() Emulator {

	memory := RAM.NewMemory()
	display := VIC2.NewMemoryDisplay(&memory)
	cpu := MOS6510.NewCPU(&memory)

	return Emulator{
		CPU:        &cpu,
		Delay:      0,
		Display:    &display,
		cycleCount: 0,
		pauseFlag:  false,
	}
}

func (emu *Emulator) Start() {
	go func() {
		for {
			if !emu.pauseFlag {
				emu.CPU.ExecuteCycle()
				emu.cycleCount++
				time.Sleep(emu.Delay * time.Millisecond)
			}
		}
	}()
	emu.Display.Start()
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
