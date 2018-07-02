package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type CPU struct {
	memory Memory
	A      byte
	Y      byte // low
	X      byte // high
	S      byte
	P      byte
	PC     int16
}

func newCPU() CPU {
	return CPU{memory: newMemory()}
}

func (c *CPU) setStatusCarry(flag bool) {
	if flag {
		c.S |= 0x01
	} else {
		c.S &^= 0x01
	}
}

func (c *CPU) getStatusCarry() bool {
	return c.S&0x01 == 0x01
}

func (c *CPU) setStatusZero(flag bool) {
	if flag {
		c.S |= 0x02
	} else {
		c.S &^= 0x02
	}
}

func (c *CPU) getStatusZero() bool {
	return c.S&0x02 == 0x02
}

func (c *CPU) setStatusNegative(flag bool) {
	if flag {
		c.S |= 0x80
	} else {
		c.S &^= 0x80
	}
}

func (c *CPU) getStatusNegative() bool {
	return c.S&0x80 == 0x80
}

func toInt16(data []byte) (ret int16) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func (c *CPU) Start(PCH byte, PCL byte) {
	PC := toInt16([]byte{PCL, PCH})
	for {
		instrCode := c.memory.ReadAbsolute(PC)
		fmt.Println(instrCode)
	}
}
