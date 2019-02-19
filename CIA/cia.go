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
	keyboard_matrix     []byte
}

func NewCIA() CIA {

	return CIA{
		Interrupt:           false,
		PORT_A:              0x0,
		keyboard_matrix_row: make([]byte, 8),
		keyboard_matrix_col: make([]byte, 8),
		keyboard_matrix:     make([]byte, 8),
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
	cia.keyboard_matrix[row] = (1 << col)
	//cia.keyboard_matrix_col[col] = (1 << row)

	//cia.keyboard_matrix[row] ^= 0xF
	//cia.keyboard_matrix_row[row] ^= 0xF
	//cia.keyboard_matrix_col[col] ^= 0xF

	log.Printf("[CIA] SetKey called with row " + strconv.Itoa(int(row)) + " and col " + strconv.Itoa(int(col)))
	log.Println("[CIA] Keyboard matrix current row:" + strconv.FormatInt(int64(cia.keyboard_matrix_row[row]), 2))
	log.Println("[CIA] Keyboard matrix current col:" + strconv.FormatInt(int64(cia.keyboard_matrix_col[col]), 2))

}

func (cia *CIA) ReadRegister() byte {
	log.Println("[CIA] ReadRegister called with PORT_A:" + strconv.FormatInt(int64(cia.PORT_A), 2))

	result := byte(0xFF)

	for i := 0; i < 8; i++ {
		if !n.GetBit(cia.PORT_A, byte(i)) {
			result &= cia.keyboard_matrix[byte(i)]
		}
	}

	result ^= 0xF

	log.Println("[CIA] Returning read register:" + strconv.FormatInt(int64(result), 2))
	return result
}

func (cia *CIA) WriteRegister(data byte) {
	log.Println("[CIA] WriteRegister called with value:", strconv.FormatInt(int64(data), 2))
	cia.PORT_A = data
}
