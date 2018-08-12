package RAM

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Memory struct {
	memory []byte
}

const ROM_CHAR_ADDR = 0xD000
const ROM_KERNAL_ADDR = 0xE000
const ROM_BASIC_ADDR = 0xA000

func NewMemory() Memory {
	mem := Memory{make([]byte, 65536)}

	log.Println("[MEM] Loading BASIC ROM...")
	mem.LoadROM("./_resources/roms/basic.901226-01.bin", ROM_BASIC_ADDR)

	log.Println("[MEM] Loading CHAR ROM...")
	mem.LoadROM("./_resources/roms/characters.901225-01.bin", ROM_CHAR_ADDR)

	log.Println("[MEM] Loading KERNAL ROM...")
	mem.LoadROM("./_resources/roms/kernal.901227-03.bin", ROM_KERNAL_ADDR)

	return mem
}

func (m Memory) ReadZeroPage(zeroPageAddress byte) byte {
	return m.memory[zeroPageAddress]
}

func (m Memory) ReadAbsolute(absoluteAddress uint16) byte {
	return m.memory[absoluteAddress]
}

func (m *Memory) WriteZeroPage(zeroPageAddress byte, value byte) {
	m.memory[zeroPageAddress] = value
}

func (m *Memory) WriteAbsolute(absoluteAddress uint16, value byte) {
	m.memory[absoluteAddress] = value
}

func (m *Memory) LoadROM(path string, address uint16) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error! Can't open file %s", path)
	}

	byteContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	currentAddress := address
	for i := 0; i < len(byteContent); i++ {
		m.WriteAbsolute(currentAddress, byteContent[i])
		currentAddress++
	}
}
