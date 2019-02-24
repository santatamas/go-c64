package CIA

import (
	n "github.com/santatamas/go-c64/numeric"
	"log"
	"strconv"
)

type CIA struct {
	//56320-56335   $DC00-$DC0F
	//Complex Interface Adapter (CIA) #1 Registers
	Interrupt           bool
	PORT_A              byte
	keyboard_matrix_row []byte
	keyboard_matrix_col []byte
	Keyboard_matrix     []byte
}

func NewCIA() CIA {

	return CIA{
		Interrupt:       false,
		PORT_A:          0x0,
		Keyboard_matrix: make([]byte, 8),
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
	cia.Keyboard_matrix[row] |= (1 << col)

	log.Printf("[CIA] SetKey called with row " + strconv.Itoa(int(row)) + " and col " + strconv.Itoa(int(col)))
	log.Println("[CIA] Keyboard matrix current row:" + strconv.FormatInt(int64(cia.Keyboard_matrix[row]), 2))
}

func (cia *CIA) ReadRegister() byte {
	log.Println("[CIA] ReadRegister called with PORT_A:" + strconv.FormatInt(int64(cia.PORT_A), 2))

	result := byte(0x00)

	for i := 0; i < 8; i++ {
		if !n.GetBit(cia.PORT_A, byte(i)) {
			result |= cia.Keyboard_matrix[byte(i)]
		}
	}

	// invert bits
	result ^= 0xFF

	log.Println("[CIA] Returning read register:" + strconv.FormatInt(int64(result), 2))
	return result
}

func (cia *CIA) WriteRegister(data byte) {
	log.Println("[CIA] WriteRegister called with value:", strconv.FormatInt(int64(data), 2))
	cia.PORT_A = data
}
