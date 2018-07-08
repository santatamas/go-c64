package main

import (
	"fmt"
	"github.com/santatamas/go-c64/MOS6510"
	"github.com/santatamas/go-c64/VIC2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Bye!")
		os.Exit(0)
	}()

	memory, startPCH, startPCL := loadFile("2_loop.prg")
	cpu := MOS6510.NewCPU(&memory)
	display := VIC2.NewMemoryDisplay(&memory)

	go cpu.Start(startPCH, startPCL)
	display.Start()

}
