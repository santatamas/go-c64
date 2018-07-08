package VIC2

import (
	"fmt"
	"github.com/santatamas/go-c64/RAM"
	"time"
)

type MemoryDisplay struct {
	screenStartAddress uint16
	colorStartAddress  uint16
	memory             *RAM.Memory
	width              int
	height             int
	c64Characters      func(byte) rune
}

func NewMemoryDisplay(memory *RAM.Memory) MemoryDisplay {
	result := MemoryDisplay{0x400, 0xD800, memory, 40, 25, asciiCharacters()}
	return result
}

func (display *MemoryDisplay) ReadCurrentState() [40][25]rune {
	result := [40][25]rune{}
	currentAdr := display.screenStartAddress
	for y := 0; y < display.height; y++ {
		for x := 0; x < display.width; x++ {
			result[x][y] = display.c64Characters(display.memory.ReadAbsolute(currentAdr))
			currentAdr++
		}
	}
	return result
}

func (display *MemoryDisplay) DrawState(state [40][25]rune) {
	for y := 0; y < display.height; y++ {
		fmt.Print("\n")
		for x := 0; x < display.width; x++ {
			fmt.Print(string(state[x][y]))
		}
	}
}

func (display *MemoryDisplay) Start() {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Refresh...")
		display.DrawState(display.ReadCurrentState())
		time.Sleep(1000 * time.Millisecond)
	}
}
