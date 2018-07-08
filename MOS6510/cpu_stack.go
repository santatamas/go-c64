package MOS6510

import "log"

func (c *CPU) stackPush(value byte) {
	c.stackPointerDec()
	log.Println("Pushing stack value: ", value)
	log.Println("Pushing stack value to address: ", c.SP_LOW+uint16(c.SP))
	c.memory.WriteAbsolute(c.SP_LOW+uint16(c.SP), value)
}

func (c *CPU) stackPop() byte {
	c.stackPointerInc()
	log.Println("Popping stack value from address: ", c.SP_LOW+uint16(c.SP))
	return c.memory.ReadAbsolute(c.SP_LOW + uint16(c.SP))
}

func (c *CPU) stackPointerInc() {
	//TODO: stackoverflow exception
	c.SP++
}

func (c *CPU) stackPointerDec() {
	//TODO: stackunderflow exception
	c.SP--
}
