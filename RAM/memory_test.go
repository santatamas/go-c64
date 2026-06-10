package RAM

import (
	"testing"

	"github.com/santatamas/go-c64/CIA"
)

func TestNewMemory_Creates64KBMemory(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	if len(mem.memory_ram) != MemSize {
		t.Errorf("Expected RAM size %d, got %d", MemSize, len(mem.memory_ram))
	}
	if len(mem.memory_rom) != MemSize {
		t.Errorf("Expected ROM size %d, got %d", MemSize, len(mem.memory_rom))
	}
}

func TestWriteAbsolute_ReadAbsolute_Roundtrip(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Write to a RAM-only area (below 0xA000)
	mem.WriteAbsolute(0x0800, 0x42)
	result := mem.ReadAbsolute(0x0800)

	if result != 0x42 {
		t.Errorf("Expected 0x42, got 0x%02X", result)
	}
}

func TestWriteAbsolute_ReadAbsolute_MultipleAddresses(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	addresses := []uint16{0x0000, 0x00FF, 0x0100, 0x0800, 0x4000, 0x8000}
	for _, addr := range addresses {
		mem.WriteAbsolute(addr, byte(addr&0xFF))
		result := mem.ReadAbsolute(addr)
		if result != byte(addr&0xFF) {
			t.Errorf("At address 0x%04X: expected 0x%02X, got 0x%02X", addr, byte(addr&0xFF), result)
		}
	}
}

func TestWriteZeroPage_ReadZeroPage_Roundtrip(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.WriteZeroPage(0x10, 0xAB)
	result := mem.ReadZeroPage(0x10)

	if result != 0xAB {
		t.Errorf("Expected 0xAB, got 0x%02X", result)
	}
}

func TestWriteZeroPage_ReadZeroPage_MultipleAddresses(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	for i := byte(2); i < 0x10; i++ {
		mem.WriteZeroPage(i, i*2)
		result := mem.ReadZeroPage(i)
		if result != i*2 {
			t.Errorf("At zero page address 0x%02X: expected 0x%02X, got 0x%02X", i, i*2, result)
		}
	}
}

func TestWriteZeroPage_Address0x01_TriggersBankSetup(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Writing to address 0x01 should call SetupBanks
	// Reset CHARENG bank to RAM manually to test that SetupBanks changes it
	mem.banks[CHARENG] = RAM

	// Write CHAREN | LORAM to address 0x01 - should trigger SetupBanks
	// and set CHARENG to IO
	mem.WriteZeroPage(AddrMemoryLayout, CHAREN|LORAM)

	if mem.banks[CHARENG] != IO {
		t.Errorf("Expected CHARENG bank to be IO after writing CHAREN|LORAM to 0x01, got %d", mem.banks[CHARENG])
	}

	// Verify the memory layout byte was updated
	if mem.memory_ram[AddrMemoryLayout] != CHAREN|LORAM {
		t.Errorf("Expected memory[0x0001] to be %d, got %d", CHAREN|LORAM, mem.memory_ram[AddrMemoryLayout])
	}
}

func TestSetupBanks_HIRAM_SetsKernalToROM(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Reset banks
	for i := 0; i < len(mem.banks); i++ {
		mem.banks[i] = RAM
	}

	mem.SetupBanks(HIRAM)

	if mem.banks[KERNAL] != ROM {
		t.Errorf("Expected KERNAL bank to be ROM when HIRAM set, got %d", mem.banks[KERNAL])
	}
}

func TestSetupBanks_LORAM_HIRAM_SetsBasicToROM(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Reset banks
	for i := 0; i < len(mem.banks); i++ {
		mem.banks[i] = RAM
	}

	mem.SetupBanks(LORAM | HIRAM)

	if mem.banks[BASIC] != ROM {
		t.Errorf("Expected BASIC bank to be ROM when LORAM|HIRAM set, got %d", mem.banks[BASIC])
	}
}

func TestSetupBanks_LORAM_Only_DoesNotSetBasicToROM(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Reset banks
	for i := 0; i < len(mem.banks); i++ {
		mem.banks[i] = RAM
	}

	mem.SetupBanks(LORAM)

	if mem.banks[BASIC] == ROM {
		t.Errorf("BASIC bank should not be ROM with only LORAM set")
	}
}

func TestSetupBanks_CHAREN_WithLORAM_SetsIOMode(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.SetupBanks(CHAREN | LORAM)

	if mem.banks[CHARENG] != IO {
		t.Errorf("Expected CHARENG bank to be IO when CHAREN and LORAM set, got %d", mem.banks[CHARENG])
	}
}

func TestSetupBanks_CHAREN_WithHIRAM_SetsIOMode(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.SetupBanks(CHAREN | HIRAM)

	if mem.banks[CHARENG] != IO {
		t.Errorf("Expected CHARENG bank to be IO when CHAREN and HIRAM set, got %d", mem.banks[CHARENG])
	}
}

func TestSetupBanks_NoCharenWithLORAM_SetsROMMode(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.SetupBanks(LORAM)

	if mem.banks[CHARENG] != ROM {
		t.Errorf("Expected CHARENG bank to be ROM when CHAREN not set and LORAM set, got %d", mem.banks[CHARENG])
	}
}

func TestSetupBanks_CHAREN_NoLORAM_NoHIRAM_SetsRAMMode(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.SetupBanks(CHAREN)

	if mem.banks[CHARENG] != RAM {
		t.Errorf("Expected CHARENG bank to be RAM when CHAREN set but LORAM/HIRAM clear, got %d", mem.banks[CHARENG])
	}
}

func TestReadAbsolute_KernalROM_WhenHIRAMSet(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Write different values to ROM and RAM at kernal address
	mem.memory_rom[0xE000] = 0xAA
	mem.memory_ram[0xE000] = 0xBB

	// With HIRAM set, KERNAL should be ROM
	mem.SetupBanks(HIRAM | CHAREN)

	result := mem.ReadAbsolute(0xE000)
	if result != 0xAA {
		t.Errorf("Expected ROM value 0xAA, got 0x%02X (should read from ROM when HIRAM set)", result)
	}
}

func TestReadAbsolute_KernalRAM_WhenHIRAMClear(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Reset all banks to RAM
	for i := 0; i < len(mem.banks); i++ {
		mem.banks[i] = RAM
	}

	mem.memory_rom[0xE000] = 0xAA
	mem.memory_ram[0xE000] = 0xBB

	result := mem.ReadAbsolute(0xE000)
	if result != 0xBB {
		t.Errorf("Expected RAM value 0xBB, got 0x%02X (should read from RAM when HIRAM clear)", result)
	}
}

func TestReadAbsolute_BasicROM_WhenLORAMAndHIRAMSet(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.memory_rom[0xA000] = 0xCC
	mem.memory_ram[0xA000] = 0xDD

	mem.SetupBanks(LORAM | HIRAM | CHAREN)

	result := mem.ReadAbsolute(0xA000)
	if result != 0xCC {
		t.Errorf("Expected ROM value 0xCC, got 0x%02X (should read from BASIC ROM when LORAM|HIRAM set)", result)
	}
}

func TestReadAbsolute_CharacterROM_WhenIONotSet(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.memory_rom[0xD000] = 0xEE
	mem.memory_ram[0xD000] = 0xFF

	// CHAREN clear, LORAM set -> CHARENG bank = ROM
	mem.SetupBanks(LORAM)

	result := mem.ReadAbsolute(0xD000)
	if result != 0xEE {
		t.Errorf("Expected ROM value 0xEE, got 0x%02X (should read character ROM when CHAREN clear)", result)
	}
}

func TestReadAbsolute_IOArea_WhenCharenSet(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.memory_ram[0xD000] = 0x77

	// CHAREN + LORAM/HIRAM -> IO mode for char area
	mem.SetupBanks(CHAREN | LORAM | HIRAM)

	result := mem.ReadAbsolute(0xD000)
	if result != 0x77 {
		t.Errorf("Expected RAM/IO value 0x77, got 0x%02X", result)
	}
}

func TestReadAbsolute_CIA_PortB_RoutesToCIA(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Set up for IO mode (default bank config)
	mem.SetupBanks(LORAM | HIRAM | CHAREN)

	// Set a key in CIA so ReadRegister returns non-zero
	cia.SetKey(0, 0)
	cia.WriteRegister(0xFE) // Enable row 0

	result := mem.ReadAbsolute(CIA_PORT_B)

	// When reading CIA_PORT_B in IO mode, it should route to CIA.Read
	// CIA.Read calls ReadRegister which returns keyboard scan result
	if result == 0 {
		// With key set at row 0 col 0, and PORT_A = 0xFE (bit 0 clear),
		// the result should be non-zero
		t.Errorf("Expected non-zero CIA read result, got 0x%02X", result)
	}
}

func TestWriteAbsolute_CIA_PortA_RoutesToCIA(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Set up for IO mode
	mem.SetupBanks(LORAM | HIRAM | CHAREN)

	// Write to CIA_PORT_A should route to CIA.Write
	mem.WriteAbsolute(CIA_PORT_A, 0xFD)

	// CIA.Write calls WriteRegister which sets PORT_A
	if cia.PORT_A != 0xFD {
		t.Errorf("Expected CIA PORT_A to be 0xFD, got 0x%02X", cia.PORT_A)
	}
}

func TestReadAll_ReturnsROMContent(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	mem.memory_rom[0x0000] = 0x11
	mem.memory_rom[0xFFFF] = 0x22

	romContent := mem.ReadAll()
	if len(romContent) != MemSize {
		t.Errorf("Expected ROM content length %d, got %d", MemSize, len(romContent))
	}
	if romContent[0x0000] != 0x11 {
		t.Errorf("Expected ROM[0x0000] = 0x11, got 0x%02X", romContent[0x0000])
	}
	if romContent[0xFFFF] != 0x22 {
		t.Errorf("Expected ROM[0xFFFF] = 0x22, got 0x%02X", romContent[0xFFFF])
	}
}

func TestNewMemory_DefaultBankConfig(t *testing.T) {
	cia := CIA.NewCIA(nil)
	mem := NewMemory(&cia)

	// Default setup is LORAM|HIRAM|CHAREN
	// KERNAL should be ROM
	if mem.banks[KERNAL] != ROM {
		t.Errorf("Expected default KERNAL bank to be ROM, got %d", mem.banks[KERNAL])
	}
	// BASIC should be ROM
	if mem.banks[BASIC] != ROM {
		t.Errorf("Expected default BASIC bank to be ROM, got %d", mem.banks[BASIC])
	}
	// CHARENG should be IO
	if mem.banks[CHARENG] != IO {
		t.Errorf("Expected default CHARENG bank to be IO, got %d", mem.banks[CHARENG])
	}
}
