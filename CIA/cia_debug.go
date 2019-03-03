package CIA

type CIAState struct {
	//56320-56335   $DC00-$DC0F
	//Complex Interface Adapter (CIA) #1 Registers
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

func (cia *CIA) GetState() CIAState {

	return CIAState{
		Interrupt:             cia.Interrupt,
		Previous_cpu_cycles:   cia.Previous_cpu_cycles,
		TIMER_A:               cia.TIMER_A,
		TIMER_A_INPUT:         cia.TIMER_A_INPUT,
		TIMER_A_ENABLED:       cia.TIMER_A_ENABLED,
		TIMER_A_LATCH:         cia.TIMER_A_LATCH,
		TIMER_A_IRQ:           cia.TIMER_A_IRQ,
		TIMER_A_IRQ_TRIGGERED: cia.TIMER_A_IRQ_TRIGGERED,
		TIMER_B:               cia.TIMER_B,
		TIMER_B_INPUT:         cia.TIMER_B_INPUT,
		TIMER_B_ENABLED:       cia.TIMER_B_ENABLED,
		TIMER_B_LATCH:         cia.TIMER_B_LATCH,
		TIMER_B_IRQ:           cia.TIMER_B_IRQ,
		TIMER_B_IRQ_TRIGGERED: cia.TIMER_B_IRQ_TRIGGERED,
		PORT_A:                cia.PORT_A,
		Keyboard_matrix:       cia.Keyboard_matrix,
	}
}
