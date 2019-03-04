package RAM

import (
	"fmt"
	"github.com/santatamas/go-c64/CIA"
	"io/ioutil"
	"log"
	"os"
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

const CIA_ADDR_RANGE_LO = 0xDC00
const CIA_ADDR_RANGE_HI = 0xDC0F

func NewMemory(testMode bool, cia *CIA.CIA) Memory {
	mem := Memory{make([]byte, MemSize),
		make([]byte, MemSize),
		make([]BankConfiguration, 7),
		cia}

	for i := 0; i < 6; i++ {
		mem.banks[i] = RAM
	}

	if testMode {
		log.Println("[MEM] Loading TEST ROM...")
		mem.LoadROM("./_resources/tests/6502_functional_test.bin", ROM_TEST_ADDR)
	} else {
		log.Println("[MEM] Loading BASIC ROM...")
		mem.LoadROM("./_resources/roms/basic.901226-01.bin", ROM_BASIC_ADDR)

		log.Println("[MEM] Loading CHAR ROM...")
		mem.LoadROM("./_resources/roms/characters.901225-01.bin", ROM_CHAR_ADDR)

		log.Println("[MEM] Loading KERNAL ROM...")
		mem.LoadROM("./_resources/roms/kernal.901227-03.bin", ROM_KERNAL_ADDR)
	}

	return mem
}

func (m *Memory) ReadAll() []byte {
	return m.memory_ram
}

func (m *Memory) ReadZeroPage(zeroPageAddress byte) byte {
	return m.memory_ram[zeroPageAddress]
}

func (m *Memory) ReadAbsolute(absoluteAddress uint16) byte {

	/*
	   // VIC-II DMA or Character ROM
	   if (page >= kAddrVicFirstPage && page <= kAddrVicLastPage)
	   {
	     if(banks_[kBankCharen] == kIO)
	       retval = vic_->read_register(addr&0x7f);
	     else if(banks_[kBankCharen] == kROM)
	       retval = mem_rom_[addr];
	     else
	       retval = mem_ram_[addr];
	   }
	   // CIA1
	   else if (page == kAddrCIA1Page)
	   {
	     if(banks_[kBankCharen] == kIO)
	       retval = cia1_->read_register(addr&0x0f);
	     else
	       retval = mem_ram_[addr];
	   }
	   // CIA2
	   else if (page == kAddrCIA2Page)
	   {
	     if(banks_[kBankCharen] == kIO)
	       retval = cia2_->read_register(addr&0x0f);
	     else
	       retval = mem_ram_[addr];
	   }
	   // BASIC or RAM
	   else if (page >= kAddrBasicFirstPage && page <= kAddrBasicLastPage)
	   {
	     if (banks_[kBankBasic] == kROM)
	       retval = mem_rom_[addr];
	     else
	       retval = mem_ram_[addr];
	   }
	   // KERNAL
	   else if (page >= kAddrKernalFirstPage && page <= kAddrKernalLastPage)
	   {
	     if (banks_[kBankKernal] == kROM)
	       retval = mem_rom_[addr];
	     else
	       retval = mem_ram_[addr];
	   }
	   else
	   {
	     retval = mem_ram_[addr];
	   }
	*/

	if absoluteAddress >= CIA_ADDR_RANGE_LO && absoluteAddress <= CIA_ADDR_RANGE_HI {
		return m.Cia.Read(absoluteAddress)
	}
	return m.memory_ram[absoluteAddress]
}

func (m *Memory) SetupBanks(value byte) {

	// Get latch bites for 3 banks only
	hiram := ((value & HIRAM) != 0)
	loram := ((value & LORAM) != 0)
	charen := ((value & CHAREN) != 0)

	/* kernal */
	if hiram {
		m.banks[KERNAL] = ROM
	}
	/* basic */
	if loram && hiram {
		m.banks[BASIC] = ROM
	}
	/* charen */
	if charen && (loram || hiram) {
		m.banks[CHAREN] = IO
	} else if charen && !loram && !hiram {
		m.banks[CHAREN] = RAM
	} else {
		m.banks[CHAREN] = ROM
	}

	/* write the config to the zero page */
	m.WriteZeroPage(AddrMemoryLayout, value)
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

	/*
		// VIC-II DMA or Character ROM
		else if (page >= kAddrVicFirstPage && page <= kAddrVicLastPage)
		{
		if(banks_[kBankCharen] == kIO)
			vic_->write_register(addr&0x7f,v);
		else
			mem_ram_[addr] = v;
		}
		// CIA1
		else if (page == kAddrCIA1Page)
		{
		if(banks_[kBankCharen] == kIO)
			cia1_->write_register(addr&0x0f,v);
		else
			mem_ram_[addr] = v;
		}
		else if (page == kAddrCIA2Page)
		{
		if(banks_[kBankCharen] == kIO)
			cia2_->write_register(addr&0x0f,v);
		else
			mem_ram_[addr] = v;
		}
		else
		{
		mem_ram_[addr] = v;
		}
	*/

	if absoluteAddress >= CIA_ADDR_RANGE_LO && absoluteAddress <= CIA_ADDR_RANGE_HI {
		m.Cia.Write(absoluteAddress, value)
	}

	m.memory_ram[absoluteAddress] = value
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
