package CIA

import (
	"github.com/santatamas/go-c64/MOS6510"
)

type CIA struct {
	//56320-56335   $DC00-$DC0F
	//Complex Interface Adapter (CIA) #1 Registers
	Cpu *MOS6510.CPU
}

func NewCIA(cpu *MOS6510.CPU) CIA {
	return CIA{
		Cpu: cpu,
	}
}
