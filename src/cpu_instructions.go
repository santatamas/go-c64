package main

import (
	"fmt"
)

func (c *CPU) BNE(mode AddressingMode) {
	c.PC++
	fmt.Println("BNE called")
}

func (c *CPU) BRK(mode AddressingMode) {
	fmt.Println("BRK called")
}

func (c *CPU) CPX(mode AddressingMode) {
	fmt.Println("CPX called")
}

func (c *CPU) INX(mode AddressingMode) {
	fmt.Println("INX called")
}

func (c *CPU) INY(mode AddressingMode) {
	fmt.Println("INY called")
}

func (c *CPU) LDA(mode AddressingMode) {
	fmt.Println("LDA called")
}

func (c *CPU) LDX(mode AddressingMode) {
	fmt.Println("LDX called")
}

func (c *CPU) STA(mode AddressingMode) {
	fmt.Println("STA called")
}

func (c *CPU) TAY(mode AddressingMode) {
	fmt.Println("TAY called")
}

func (c *CPU) TYA(mode AddressingMode) {
	fmt.Println("TYA called")
}
