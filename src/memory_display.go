package main

import (
	"fmt"
	"time"
)

type MemoryDisplay struct {
	screenStartAddress uint16
	colorStartAddress  uint16
	memory             *Memory
	width              int
	height             int
}

func newMemoryDisplay(memory *Memory) MemoryDisplay {
	result := MemoryDisplay{0x400, 0xD800, memory, 40, 25}
	return result
}

func (display *MemoryDisplay) ReadCurrentState() [40][25]rune {
	result := [40][25]rune{}
	currentAdr := display.screenStartAddress
	for x := 0; x < display.width; x++ {
		for y := 0; y < display.height; y++ {
			display.memory.ReadAbsolute(currentAdr)
			currentAdr++
			result[x][y] = 1 //C64CharConverter.ConvertToAscii(_memory.ReadAbsolute(currentAdr++));
		}
	}
	return result
}

func (display *MemoryDisplay) DrawState(state [40][25]rune) {
	for x := 0; x < display.width; x++ {
		fmt.Print("\n")
		for y := 0; y < display.height; y++ {
			fmt.Print(state[x][y])
		}
	}
}

func (display *MemoryDisplay) Init() {
	//TODO: do initialization
}

func (display *MemoryDisplay) Start() {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Refresh...")
		display.DrawState(display.ReadCurrentState())
		time.Sleep(1000 * time.Millisecond)
	}
}
