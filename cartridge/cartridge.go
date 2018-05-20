package cartridge

import (
	"fmt"
	"io/ioutil"
	"time"
)

type cartridgeType uint8

const (
	mbc0 cartridgeType = 0
	mbc1 cartridgeType = 1
	mbc2 cartridgeType = 2
	mbc3 cartridgeType = 3
	mbc4 cartridgeType = 4
	mbc5 cartridgeType = 5
)

var (
	rom               []uint8
	ram               []uint8
	romBank           uint16 = 1
	ramBank           uint16
	cartType          cartridgeType
	romBankingEnabled bool
	ramEnabled        bool
	saveFilePath      string
	ramChanged        bool
)

// Load opens the ROM and its accompagying save file (if any) and loads them to cartridge rom and ram
func Load(romName string) {
	data, err := ioutil.ReadFile(romName)
	if err != nil {
		panic(err)
	}
	saveFilePath = romName + ".sav"
	rom = data
	ram = make([]byte, 0x8000)

	mbcFlag := rom[0x147]

	if mbcFlag == 0x00 || (mbcFlag >= 0x08 && mbcFlag <= 0x0d) {
		cartType = mbc0
	} else if mbcFlag >= 0x01 && mbcFlag <= 0x03 {
		cartType = mbc1
		// } else if mbcFlag >= 0x05 && mbcFlag <= 0x06 {
		//     cartType = mbc2
	} else if mbcFlag >= 0x0f && mbcFlag <= 0x13 {
		cartType = mbc3
		// } else if mbcFlag >= 0x15 && mbcFlag <= 0x17 {
		//     cartType = mbc4
		// } else if mbcFlag >= 0x19 && mbcFlag <= 0x1e {
		//     cartType = mbc5
	} else {
		fmt.Printf("Possibly unsupported cartridge type (0x%x)", mbcFlag)
		cartType = mbc0
	}

	switch mbcFlag {
	case 0x3, 0x6, 0x9, 0xD, 0xF, 0x10, 0x13, 0x17, 0x1B, 0x1E:
		saveData, err := ioutil.ReadFile(saveFilePath)
		if err == nil {
			ram = saveData
		}
		ticker := time.NewTicker(time.Second)
		go func() {
			for range ticker.C {
				if ramChanged {
					ramChanged = false
					saveRAM()
				}
			}
		}()
	}
}

// Read fetches the content of the ROM at the given address
func Read(addr uint16) uint8 {
	if addr < 0x4000 {
		return rom[addr]
	}
	newAddr := uint32(addr) - 0x4000
	return rom[newAddr+uint32(romBank)*0x4000]
}

// ReadFromRAM fetches the content of the cartridge's RAM at the given address
func ReadFromRAM(addr uint16) uint8 {
	newAddr := addr - 0xa000
	return ram[newAddr+ramBank*0x2000]
}

// Write writes the given value at the given address in the ROM. The effect depends on the type of cartridge
func Write(addr uint16, value uint8) {
	switch cartType {
	case mbc0:
		return
	case mbc1:
		writembc1(addr, value)
	case mbc2:
		//writembc2(addr, value)
	case mbc3:
		writembc3(addr, value)
	case mbc4:
		return
	case mbc5:
		//writembc5(addr, value)
	}
}

// WriteToRAM writes the given value at the given address in the cartridge RAM
func WriteToRAM(addr uint16, value uint8) {
	if ramEnabled {
		ramChanged = true
		newAddr := addr - 0xa000
		ram[newAddr+ramBank*0x2000] = value
	}
}

func writembc1(addr uint16, value uint8) {
	switch {
	case addr < 0x2000:
		enableRAM(addr, value)
	case addr < 0x4000:
		changeLowerRomBank(value)
	case addr < 0x6000:
		if romBankingEnabled {
			changeUpperRomBank(value)
		} else {
			ramBank = uint16(value & 0x03)
		}
	case addr < 0x8000:
		if value&0x01 == 0 {
			romBankingEnabled = true
			ramBank = 0
		} else {
			romBankingEnabled = false
		}
	}
}

func writembc3(addr uint16, value uint8) {
	switch {
	case addr < 0x2000:
		enableRAM(addr, value)
	case addr < 0x4000:
		romBank = uint16(value & 0x7f)
		if romBank == 0 {
			romBank = 1
		}
	case addr < 0x6000:
		ramBank = uint16(value) & 0x03
	}
}

func enableRAM(addr uint16, value uint8) {
	if cartType == mbc2 && addr&0x08 == 0x10 {
		return
	}

	test := value & 0xF
	if test == 0xA {
		ramEnabled = true
	} else if test == 0x0 {
		ramEnabled = false
	}
}

func changeLowerRomBank(value uint8) {
	if cartType == mbc2 {
		romBank = uint16(value & 0x0f)
	} else {
		romBank &= 0xe0
		romBank |= uint16(value & 0x1f)
	}
	if romBank == 0 {
		romBank++
	}
}

func changeUpperRomBank(value byte) {
	romBank &= 0x1f
	value &= 224
	romBank |= uint16(value & 0xe0)
	if romBank == 0 {
		romBank++
	}
}

func saveRAM() {
	ioutil.WriteFile(saveFilePath, ram, 0644)
}
