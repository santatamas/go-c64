package main

import (
	"github.com/santatamas/go-c64/MOS6510"
	"github.com/santatamas/go-c64/VIC2"
	"log"
	"os"
)

func main() {

	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	memory, startPCH, startPCL := loadFile("./_resources/Prg/4_colors.prg")
	cpu := MOS6510.NewCPU(&memory)
	display := VIC2.NewMemoryDisplay(&memory)

	go cpu.Start(startPCL, startPCH)
	display.Start()

}
