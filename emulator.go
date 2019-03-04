package main

import (
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
	"time"
)

type Emulator struct {
	CPU        *MOS6510.CPU
	CIA        *CIA.CIA
	Display    *VIC2.MemoryDisplay
	Delay      time.Duration
	cycleCount uint64
	pauseFlag  bool
	Debug      bool
	Test       bool
	BreakAddr  uint16
	hub        *internals.Hub
}

func NewEmulator(testMode bool) Emulator {

	keyPressChannel := make(chan *tcell.EventKey)
	irqChannel := make(chan bool)

	cia := CIA.NewCIA(irqChannel)
	memory := RAM.NewMemory(testMode, &cia)
	cia = CIA.NewCIA(irqChannel)
	memory.Cia = &cia
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
				emu.CIA.ExecuteCycle(emu.cycleCount * 2)
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
