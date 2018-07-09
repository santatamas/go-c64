package MOS6510

import "log"

func (c *CPU) stackPush(value byte) {
	c.stackPointerDec()
	log.Printf("Pushing stack value: %x", value)
	log.Printf("Pushing stack value to address: %x", c.SP_LOW+uint16(c.SP))
	c.memory.WriteAbsolute(c.SP_LOW+uint16(c.SP), value)
}

func (c *CPU) stackPop() byte {
	log.Printf("Popping stack value from address: %x", c.SP_LOW+uint16(c.SP))
	result := c.memory.ReadAbsolute(c.SP_LOW + uint16(c.SP))
	c.stackPointerInc()
	return result
}

func (c *CPU) stackPointerInc() {
	//TODO: stackoverflow exception
	c.SP++
	log.Printf("[INC] Stack pointer value: %x", c.SP)
}

func (c *CPU) stackPointerDec() {
	//TODO: stackunderflow exception
	c.SP--
	log.Printf("[DEC] Stack pointer value: %x", c.SP)
}
