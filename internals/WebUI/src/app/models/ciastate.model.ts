export class CIAState {
    Interrupt: boolean;
    Previous_cpu_cycles: number;
    TIMER_A:               number;
    TIMER_A_INPUT:         number;
    TIMER_A_ENABLED:       boolean;
    TIMER_A_LATCH:         number;
    TIMER_A_IRQ:           boolean;
    TIMER_A_IRQ_TRIGGERED: boolean;
    TIMER_B:               number;
    TIMER_B_INPUT:         number;
    TIMER_B_ENABLED:       boolean;
    TIMER_B_LATCH:         number;
    TIMER_B_IRQ:           boolean;
    TIMER_B_IRQ_TRIGGERED: boolean;
    PORT_A:                number;
    Keyboard_matrix:       number[];
}