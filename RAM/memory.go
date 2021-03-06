package RAM

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/santatamas/go-c64/CIA"
)

type Memory struct {
	memory_ram []byte
	memory_rom []byte
	banks      []BankConfiguration
	Cia        *CIA.CIA
}

const MemSize = 0x10000 // 65536

/* memory addresses  */
const BaseAddrBasic = 0xa000
const BaseAddrKernal = 0xe000
const BaseAddrStack = 0x0100
const BaseAddrScreen = 0x0400
const BaseAddrChars = 0xd000
const BaseAddrBitmap = 0x0000
const BaseAddrColorRAM = 0xd800
const AddrResetVector = 0xfffc
const AddrIRQVector = 0xfffe
const AddrNMIVector = 0xfffa
const AddrDataDirection = 0x0000
const AddrMemoryLayout = 0x0001
const AddrColorRAM = 0xd800

/* memory layout */
const AddrZeroPage = 0x0000
const AddrVicFirstPage = 0xd000
const AddrVicLastPage = 0xd300
const AddrCIA1Page = 0xdc00
const AddrCIA2Page = 0xdd00
const AddrBasicFirstPage = 0xa000
const AddrBasicLastPage = 0xbf00
const AddrKernalFirstPage = 0xe000
const AddrKernalLastPage = 0xff00

/* bank switching */
const LORAM = 1 << 0
const HIRAM = 1 << 1
const CHAREN = 1 << 2

/* deprecated constants */
const ROM_CHAR_ADDR = 0xD000
const ROM_KERNAL_ADDR = 0xE000
const ROM_BASIC_ADDR = 0xA000
const ROM_TEST_ADDR = 0x400

const IRQ_VECTOR_ADDR_LO = 0xFFFE
const IRQ_VECTOR_ADDR_HI = 0xFFFF

const CIA_PORT_A = 0xDC00
const CIA_PORT_B = 0xDC01

func NewMemory(cia *CIA.CIA) Memory {
	mem := Memory{make([]byte, MemSize),
		make([]byte, MemSize),
		make([]BankConfiguration, 7),
		cia}

	for i := 0; i < 6; i++ {
		mem.banks[i] = RAM
	}

	mem.SetupBanks(LORAM | HIRAM | CHAREN)

	return mem
}

func (m *Memory) ReadAll() []byte {
	return m.memory_rom
}

func (m *Memory) ReadZeroPage(zeroPageAddress byte) byte {
	return m.ReadAbsolute(uint16(zeroPageAddress)) //m.memory_ram[zeroPageAddress]
}

func (m *Memory) ReadAbsolute(absoluteAddress uint16) byte {

	retval := byte(0x0)
	page := absoluteAddress & 0xff00

	// VIC2
	if page >= AddrVicFirstPage && page <= AddrVicLastPage {
		if m.banks[CHARENG] == IO {
			//log.Println("[MEM] Reading VIC2 IO")
			//retval = vic_->read_register(addr&0x7f);
			retval = m.memory_ram[absoluteAddress]
		} else if m.banks[CHARENG] == ROM {
			//log.Println("[MEM] Reading VIC2 ROM")
			retval = m.memory_rom[absoluteAddress]
		} else {
			//log.Println("[MEM] Reading VIC2 RAM")
			retval = m.memory_ram[absoluteAddress]
		}
		// CIA 1
	} else if page == AddrCIA1Page {

		if m.banks[CHARENG] == IO && absoluteAddress == CIA_PORT_B {
			//log.Println("[MEM] Reading CIA1 IO")
			retval = m.Cia.Read(absoluteAddress)
		} else {
			//log.Println("[MEM] Reading CIA1 RAM")
			retval = m.memory_ram[absoluteAddress]
		}
		// CIA 2
	} else if page == AddrCIA2Page {
		if m.banks[CHARENG] == IO {
			// TODO: implement CIA2
			//retval = m.Cia.Read(absoluteAddress)
			//log.Println("[MEM] Reading CIA2 IO")
			retval = m.memory_ram[absoluteAddress]
		} else {
			//log.Println("[MEM] Reading CIA1 RAM")
			retval = m.memory_ram[absoluteAddress]
		}
		// BASIC
	} else if page >= AddrBasicFirstPage && page <= AddrBasicLastPage {
		if m.banks[BASIC] == ROM {
			//log.Println("[MEM] Reading BASIC ROM")
			retval = m.memory_rom[absoluteAddress]
		} else {
			//log.Println("[MEM] Reading BASIC RAM")
			retval = m.memory_ram[absoluteAddress]
		}
		// KERNAL
	} else if page >= AddrKernalFirstPage && page <= AddrKernalLastPage {
		if m.banks[KERNAL] == ROM {
			//log.Println("[MEM] Reading KERNAL ROM")
			retval = m.memory_rom[absoluteAddress]
		} else {
			//log.Println("[MEM] Reading KERNAL RAM")
			retval = m.memory_ram[absoluteAddress]
		}
		// ELSE return ram content
	} else {
		//log.Println("[MEM] Reading RAM")
		retval = m.memory_ram[absoluteAddress]
	}

	return retval
}

func (m *Memory) SetupBanks(value byte) {

	log.Printf("[MEM] Setting bank configuration to %b", value)

	// Get latch bites for 3 banks only
	hiram := ((value & HIRAM) != 0)
	loram := ((value & LORAM) != 0)
	charen := ((value & CHAREN) != 0)

	/* kernal */
	if hiram {
		log.Println("[MEM] setting KERNAL bank to ROM")
		m.banks[KERNAL] = ROM
	}
	/* basic */
	if loram && hiram {
		log.Println("[MEM] setting BASIC bank to ROM")
		m.banks[BASIC] = ROM
	}
	/* charen */
	if charen && (loram || hiram) {
		log.Println("[MEM] setting CHARENG bank to IO")
		m.banks[CHARENG] = IO
	} else if charen && !loram && !hiram {
		log.Println("[MEM] setting CHARENG bank to RAM")
		m.banks[CHARENG] = RAM
	} else {
		log.Println("[MEM] setting CHARENG bank to ROM")
		m.banks[CHARENG] = ROM
	}

	m.memory_ram[AddrMemoryLayout] = value
}

func (m *Memory) WriteZeroPage(zeroPageAddress byte, value byte) {

	// PLA address - sets bankswitching latches
	if zeroPageAddress == AddrMemoryLayout {
		m.SetupBanks(value)
	} else {
		m.memory_ram[zeroPageAddress] = value
	}
}

func (m *Memory) WriteAbsolute(absoluteAddress uint16, value byte) {

	page := absoluteAddress & 0xff00

	// VIC2
	if page >= AddrVicFirstPage && page <= AddrVicLastPage {
		if m.banks[CHARENG] == IO {
			//retval = vic_->read_register(addr&0x7f);
			m.memory_ram[absoluteAddress] = value
		} else {
			m.memory_ram[absoluteAddress] = value
		}
		// CIA 1
	} else if page == AddrCIA1Page {
		if m.banks[CHARENG] == IO && absoluteAddress == CIA_PORT_A {
			m.Cia.Write(absoluteAddress, value)
		} else {
			m.memory_ram[absoluteAddress] = value
		}
		// CIA 2
	} else if page == AddrCIA2Page {
		if m.banks[CHARENG] == IO {
			// TODO: implement CIA2
			//retval = m.Cia.Read(absoluteAddress)
			m.memory_ram[absoluteAddress] = value
		} else {
			m.memory_ram[absoluteAddress] = value
		}
		// BASIC
	} else {
		m.memory_ram[absoluteAddress] = value
	}
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
		m.memory_rom[currentAddress] = byteContent[i]
		currentAddress++
	}
}
