package main

import (
	"fmt"
)

func main() {
	cpu := newCPU()
	cpu.Start(0x01, 0x01)

	cpu.setStatusCarry(true)
	fmt.Printf("%08b", cpu.S)
	fmt.Println(cpu.getStatusCarry())

	cpu.setStatusCarry(false)
	fmt.Printf("%08b", cpu.S)
	fmt.Println(cpu.getStatusCarry())
}
