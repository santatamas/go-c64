package CIA

import (
	n "github.com/santatamas/go-c64/numeric"
	"log"
)

type CIA struct {
	//56320-56335   $DC00-$DC0F
	//Complex Interface Adapter (CIA) #1 Registers
	Interrupt           bool
	PORT_A              byte
	keyboard_matrix_row []byte
	keyboard_matrix_col []byte
}

func NewCIA() CIA {

	return CIA{
		Interrupt:           false,
		PORT_A:              0x0,
		keyboard_matrix_row: make([]byte, 8),
		keyboard_matrix_col: make([]byte, 8),
	}
}

func (cia *CIA) GetInterrupt() bool {
	if cia.Interrupt {
		cia.Interrupt = false
		return true
	}
	return false
}

func (cia *CIA) SetKey(row byte, col byte) {
	cia.Interrupt = true
	cia.keyboard_matrix_row[row] &= ^(1 << col)
	cia.keyboard_matrix_col[col] &= ^(1 << row)
	log.Printf("SetKey called")

}

func (cia *CIA) ReadRegister() byte {
	log.Printf("[CIA] ReadRegister called with PORT_A: %x", cia.PORT_A)

	result := byte(0xff)

	for i := 0; i < 8; i++ {
		if n.GetBit(cia.PORT_A, byte(i)) {
			result &= cia.keyboard_matrix_row[byte(i)]
		}
	}

	return result
}

func (cia *CIA) WriteRegister(data byte) {
	log.Printf("[CIA] WriteRegister called with value: %x", data)
	cia.PORT_A = data
}
