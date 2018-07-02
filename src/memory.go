package main

type Memory struct {
	memory []byte
}

func newMemory() Memory {
	return Memory{make([]byte, 65536)}
}

func (m Memory) ReadZeroPage(zeroPageAddress byte) byte {
	return m.memory[zeroPageAddress]
}

func (m Memory) ReadAbsolute(absoluteAddress int16) byte {
	return m.memory[absoluteAddress]
}

func (m *Memory) WriteZeroPage(zeroPageAddress byte, value byte) {
	m.memory[zeroPageAddress] = value
}

func (m *Memory) WriteAbsolute(absoluteAddress int16, value byte) {
	m.memory[absoluteAddress] = value
}
