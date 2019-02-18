package CIA

import ()
import "log"

type CIA struct {
	//56320-56335   $DC00-$DC0F
	//Complex Interface Adapter (CIA) #1 Registers
	Interrupt       bool
	PORT_A          byte
	keyboard_matrix []byte
}

func NewCIA() CIA {
	return CIA{}
}

func (cia *CIA) GetInterrupt() bool {
	if cia.Interrupt {
		cia.Interrupt = false
		return true
	}
	return false
}

func (cia *CIA) SetKey(key byte) {
	cia.Interrupt = true
	log.Printf("SetKey called")

}

func (cia *CIA) ReadRegister() byte {
	if cia.PORT_A == 0x2 {
		return 0x4
	}
	return 0x0
}

func (cia *CIA) WriteRegister(data byte) {
	cia.PORT_A = data
}
