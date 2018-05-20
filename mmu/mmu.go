package mmu

import (
	"github.com/karimElmougi/GolangBoy/cartridge"
)

type specialReader struct {
	address uint16
	read    func(uint16) uint8
}

type specialWriter struct {
	address uint16
	write   func(uint16, uint8)
}

var (
	RAM          [65536]uint8
	specialRead  []specialReader
	specialWrite []specialWriter
)

func Init() {
	specialRead = make([]specialReader, 0)
	specialWrite = make([]specialWriter, 0)

	RAM[0xff0f] = 0xe1
	RAM[0xff10] = 0x80
	RAM[0xff11] = 0xbf
	RAM[0xff12] = 0xf3
	RAM[0xff14] = 0xbf
	RAM[0xff16] = 0x3f
	RAM[0xff19] = 0xbf
	RAM[0xff1a] = 0x7f
	RAM[0xff1b] = 0xff
	RAM[0xff1c] = 0x9f
	RAM[0xff1e] = 0xbf
	RAM[0xff20] = 0xff
	RAM[0xff23] = 0xbf
	RAM[0xff24] = 0x77
	RAM[0xff25] = 0xf3
	RAM[0xff26] = 0xf1
	RAM[0xff40] = 0x91
	RAM[0xff41] = 0x85
	RAM[0xff47] = 0xfc
	RAM[0xff48] = 0xff
	RAM[0xff49] = 0xff
	RAM[0xff70] = 0x01
}

func RegisterSpecialRead(addr uint16, f func(uint16) uint8) {
	specialRead = append(specialRead, specialReader{addr, f})
}

func RegisterSpecialWrite(addr uint16, f func(uint16, uint8)) {
	specialWrite = append(specialWrite, specialWriter{addr, f})
}

func Read(addr uint16) uint8 {
	for _, r := range specialRead {
		if addr == r.address {
			return r.read(addr)
		}
	}
	if addr >= 0x0000 && addr <= 0x7fff {
		return cartridge.Read(addr)
	} else if addr >= 0xa000 && addr <= 0xbfff {
		return cartridge.ReadFromRAM(addr)
	}
	return RAM[addr]
}

func Write(addr uint16, value uint8) {
	for _, r := range specialWrite {
		if addr == r.address {
			r.write(addr, value)
			return
		}
	}
	if addr >= 0x0000 && addr <= 0x7fff {
		cartridge.Write(addr, value)
	} else if addr >= 0xa000 && addr <= 0xbfff {
		cartridge.WriteToRAM(addr, value)
	} else if addr == 0xff02 {
		return
	} else if addr == 0xff44 {
		RAM[addr] = 0
	} else if addr == 0xff46 {
		RAM[addr] = value
		dmaTransfer(value)
	} else if addr == 0xff4d {
		RAM[addr] = 0
	} else if addr == 0xff55 {
		RAM[addr] = value
	} else if addr == 0xff70 {
		return
	} else {
		RAM[addr] = value
	}
}

func ReadWord(addr uint16) uint16 {
	return uint16(Read(addr)) | uint16(Read(addr+1))<<8
}

func WriteWord(addr uint16, value uint16) {
	Write(addr, uint8(value&0xff))
	Write(addr+1, uint8(value>>8))
}

func dmaTransfer(value uint8) {
	address := uint16(value) << 8
	for i := uint16(0); i < 0xa0; i++ {
		Write(0xfe00+i, Read(address+i))
	}
}
