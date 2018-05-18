package interrupts

import (
	"github.com/karimElmougi/GolangBoy/cpu"
	"github.com/karimElmougi/GolangBoy/mmu"
)

const (
	IE_FLAG = 0xffff
	IF_ADDR = 0xff0f
)

func Init() {
	mmu.RegisterSpecialRead(IF_ADDR, func(addr uint16) uint8 {
		return mmu.RAM[IF_ADDR] | 0xe0
	})
}

func ISR() {
	if cpu.EnablingInterrupts {
		cpu.InterruptsEnabled = true
		cpu.EnablingInterrupts = false
		return
	}
	if !cpu.InterruptsEnabled && !cpu.IsHalted {
		return
	}
	if !interruptsReady() {
		return
	}
	switch {
	case isVblankInterruptEnabled() && vblankInterruptOccured():
		serviceVblankInterrupt()
	case isLcdInterruptEnabled() && lcdInterruptOccured():
		serviceLcdInterrupt()
	case isTimerInterruptEnabled() && timerInterruptOccured():
		serviceTimerInterrupt()
	case isSerialInterruptEnabled() && serialInterruptOccured():
		serviceSerialInterrupt()
	case isJoypadInterruptEnabled() && joypadInterruptOccured():
		serviceJoypadInterrupt()
	}
}

func WriteVblankInterrupt() {
	writeInterrupt(0x01)
}

func WriteLcdInterrupt() {
	writeInterrupt(0x02)
}

func WriteTimerInterrupt() {
	writeInterrupt(0x04)
}

func WriteSerialInterrupt() {
	writeInterrupt(0x08)
}

func WriteJoypadInterrupt() {
	writeInterrupt(0x10)
}

func writeInterrupt(interruptSignal uint8) {
	interruptFlags := mmu.Read(IF_ADDR)
	mmu.Write(IF_ADDR, interruptFlags|interruptSignal)
}

func serviceVblankInterrupt() {
	serviceInterrupt(0x40, 0xfe)
}

func serviceLcdInterrupt() {
	serviceInterrupt(0x48, 0xfd)
}

func serviceTimerInterrupt() {
	serviceInterrupt(0x50, 0xfb)
}

func serviceSerialInterrupt() {
	serviceInterrupt(0x58, 0xf7)
}

func serviceJoypadInterrupt() {
	serviceInterrupt(0x60, 0xef)
}

func serviceInterrupt(interruptAddress uint16, interruptReset uint8) {
	if !cpu.InterruptsEnabled && cpu.IsHalted {
		cpu.IsHalted = false
		return
	}
	cpu.InterruptsEnabled = false
	interruptFlags := mmu.Read(IF_ADDR)
	interruptFlags &= interruptReset
	mmu.Write(IF_ADDR, interruptFlags)
	cpu.Call(interruptAddress)
}

func isVblankInterruptEnabled() bool {
	return mmu.Read(IE_FLAG)&0x1 == 0x1
}

func isLcdInterruptEnabled() bool {
	return mmu.Read(IE_FLAG)&0x2 == 0x2
}

func isTimerInterruptEnabled() bool {
	return mmu.Read(IE_FLAG)&0x4 == 0x4
}

func isSerialInterruptEnabled() bool {
	return mmu.Read(IE_FLAG)&0x8 == 0x8
}

func isJoypadInterruptEnabled() bool {
	return mmu.Read(IE_FLAG)&0x10 == 0x10
}

func vblankInterruptOccured() bool {
	return mmu.Read(IF_ADDR)&0x1 == 0x1
}

func lcdInterruptOccured() bool {
	return mmu.Read(IF_ADDR)&0x2 == 0x2
}

func timerInterruptOccured() bool {
	return mmu.Read(IF_ADDR)&0x4 == 0x4
}

func serialInterruptOccured() bool {
	return mmu.Read(IF_ADDR)&0x8 == 0x8
}

func joypadInterruptOccured() bool {
	return mmu.Read(IF_ADDR)&0x10 == 0x10
}

func interruptsReady() bool {
	iflag := mmu.Read(IF_ADDR)
	ie := mmu.Read(IE_FLAG)
	return (iflag & ie) != 0x00
}
