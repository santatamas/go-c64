package MOS6510

/*
+-+-+-+-+-+-+-+-+
|N|V| |B|D|I|Z|C|  PROCESSOR STATUS REG "P"
+-+-+-+-+-+-+-+-+
 | |   | | | | |
 | |   | | | | +>  CARRY         1=TRUE
 | |   | | | +-->  ZERO          1=RESULT ZERO
 | |   | | +---->  IRQ DISABLE   1=DISABLE
 | |   | +------>  DECIMAL MODE  1=TRUE
 | |   +-------->  BRK COMMAND
 | |
 | +------------>  OVERFLOW      1=TRUE
 +-------------->  NEGATIVE      1=NEG
*/

func (c *CPU) setStatusCarry(flag bool) {
	if flag {
		c.S |= 0x01
	} else {
		c.S &^= 0x01
	}
}

func (c *CPU) getStatusCarry() bool {
	return c.S&0x01 == 0x01
}

func (c *CPU) setStatusZero(flag bool) {
	if flag {
		c.S |= 0x02
	} else {
		c.S &^= 0x02
	}
}

func (c *CPU) getStatusZero() bool {
	return c.S&0x02 == 0x02
}

func (c *CPU) setStatusIRQ(flag bool) {
	if flag {
		c.S |= 0x04
	} else {
		c.S &^= 0x04
	}
}

func (c *CPU) getStatusIRQ() bool {
	return c.S&0x04 == 0x04
}

func (c *CPU) setStatusDecimate(flag bool) {
	if flag {
		c.S |= 0x08
	} else {
		c.S &^= 0x08
	}
}

func (c *CPU) getStatusDecimate() bool {
	return c.S&0x08 == 0x08
}

func (c *CPU) setStatusBRK(flag bool) {
	if flag {
		c.S |= 0x10
	} else {
		c.S &^= 0x10
	}
}

func (c *CPU) getStatusBRK() bool {
	return c.S&0x10 == 0x10
}

func (c *CPU) setStatusOverflow(flag bool) {
	if flag {
		c.S |= 0x40
	} else {
		c.S &^= 0x40
	}
}

func (c *CPU) getStatusOverflow() bool {
	return c.S&0x40 == 0x40
}

func (c *CPU) setStatusNegative(flag bool) {
	if flag {
		c.S |= 0x80
	} else {
		c.S &^= 0x80
	}
}

func (c *CPU) getStatusNegative() bool {
	return c.S&0x80 == 0x80
}
