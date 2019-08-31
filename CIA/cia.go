package CIA

import (
	"log"
	"strconv"

	n "github.com/santatamas/go-c64/numeric"
)

const CIA_PORT_A = 0xDC00    // data port A (write register)
const CIA_PORT_B = 0xDC01    // data port B (read register)
const CIA_DDR_A = 0xDC02     // Data Direction Port A
const CIA_DDR_B = 0xDC03     // Data Direction Port B
const CIA_TA_LO = 0xDC04     // Timer A Low Byte
const CIA_TA_HI = 0xDC05     // Timer A High Byte
const CIA_TB_LO = 0xDC06     // Timer B Low Byte
const CIA_TB_HI = 0xDC07     // Timer B High Byte
const CIA_TOD_10THS = 0xDC08 // Real Time Clock 1/10s
const CIA_TOD_SEC = 0xDC09   // Real Time Clock Seconds
const CIA_TOD_MIN = 0xDC0A   // Real Time Clock Minutes
const CIA_TOD_HR = 0xDC0B    // Real Time Clock Hours
const CIA_SDR = 0xDC0C       // Serial shift register
const CIA_ICR = 0xDC0D       // Interrupt Control and status
const CIA_CRA = 0xDC0E       // Control Timer A
const CIA_CRB = 0xDC0F       // Control Timer B

type CIA struct {
	//56320-56335   $DC00-$DC0F
	//Complex Interface Adapter (CIA) #1 Registers
	IrqChannel            chan bool
	Interrupt             bool
	Previous_cpu_cycles   uint64
	TIMER_A               uint16
	TIMER_A_INPUT         byte
	TIMER_A_ENABLED       bool
	TIMER_A_LATCH         uint16
	TIMER_A_IRQ           bool
	TIMER_A_IRQ_TRIGGERED bool
	TIMER_B               uint16
	TIMER_B_INPUT         byte
	TIMER_B_ENABLED       bool
	TIMER_B_LATCH         uint16
	TIMER_B_IRQ           bool
	TIMER_B_IRQ_TRIGGERED bool
	PORT_A                byte
	Keyboard_matrix       []byte
}

func NewCIA(irqChannel chan bool) CIA {

	return CIA{
		IrqChannel:      irqChannel,
		Interrupt:       false,
		PORT_A:          0x0,
		TIMER_A:         0x0,
		TIMER_A_ENABLED: false,
		TIMER_A_IRQ:     false,
		TIMER_B:         0x0,
		TIMER_B_ENABLED: false,
		TIMER_B_IRQ:     false,
		Keyboard_matrix: make([]byte, 8),
	}
}

func (cia *CIA) ExecuteCycle(currentCpuCycleCount uint64) {

	cia.Previous_cpu_cycles = currentCpuCycleCount
}

func (cia *CIA) SendInterrupt() {
	cia.Interrupt = true
}

func (cia *CIA) SetKey(row byte, col byte) {
	log.Printf("[CIA] Setkey called")
	cia.Keyboard_matrix[row] |= (1 << col)

	log.Printf("[CIA] SetKey called with row " + strconv.Itoa(int(row)) + " and col " + strconv.Itoa(int(col)))
	log.Println("[CIA] Keyboard matrix current row:" + strconv.FormatInt(int64(cia.Keyboard_matrix[row]), 2))
}

func (cia *CIA) UnsetKey(row byte, col byte) {
	log.Printf("Unsetkey called")
	cia.Keyboard_matrix[row] &= (0 << col)

	log.Printf("[CIA] UnsetKey called with row " + strconv.Itoa(int(row)) + " and col " + strconv.Itoa(int(col)))
	log.Println("[CIA] Keyboard matrix current row:" + strconv.FormatInt(int64(cia.Keyboard_matrix[row]), 2))
}

func (cia *CIA) ResetKeyboardMatrix() {
	cia.Keyboard_matrix = make([]byte, 8)
}

func (cia *CIA) Read(address uint16) byte {

	log.Printf("[CIA] Read called with memory address %x", address)

	return cia.ReadRegister()
}

func (cia *CIA) Write(address uint16, data byte) {

	log.Printf("[CIA] Read called with memory address %x with data %x", address, data)
	cia.WriteRegister(data)
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
