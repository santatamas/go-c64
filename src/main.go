package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(file)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Bye!")
		os.Exit(0)
	}()

	memory, startPCH, startPCL := loadFile("1_first.prg")
	cpu := newCPU(&memory)
	display := newMemoryDisplay(&memory)
	display.Init()

	go cpu.Start(startPCH, startPCL)
	display.Start()

}
