package MOS6510

import (
	"errors"
	"log"
)

func (c *CPU) stackPush(value byte) (err error) {
	err = c.stackPointerDec()
	if err != nil {
		return err
	}

	log.Printf("Pushing stack value: %x", value)
	log.Printf("Pushing stack value to address: %x", c.SP_LOW+uint16(c.SP))
	c.memory.WriteAbsolute(c.SP_LOW+uint16(c.SP), value)
	return
}

func (c *CPU) stackPop() (result byte, err error) {
	log.Printf("Popping stack value from address: %x", c.SP_LOW+uint16(c.SP))
	result = c.memory.ReadAbsolute(c.SP_LOW + uint16(c.SP))
	err = c.stackPointerInc()
	return result, err
}

func (c *CPU) stackPointerInc() (err error) {

	if c.SP == 255 {
		return errors.New("Stackunderflow exception")
	}

	c.SP++

	log.Printf("[INC] Stack pointer value: %x", c.SP)
	return nil
}

func (c *CPU) stackPointerDec() (err error) {

	if c.SP == 0 {
		return errors.New("Stackoverflow error")
	}

	c.SP--

	log.Printf("[DEC] Stack pointer value: %x", c.SP)
	return nil
}
