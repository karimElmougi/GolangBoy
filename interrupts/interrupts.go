package interrupts

import (
	"github.com/karimElmougi/GolangBoy/cpu"
	"github.com/karimElmougi/GolangBoy/mmu"
)

const (
	ieFlagAddr = 0xffff
	ifAddr     = 0xff0f
)

// Init registers a SpecialReader with the mmu
func Init() {
	mmu.RegisterSpecialRead(ifAddr, func(addr uint16) uint8 {
		return mmu.RAM[ifAddr] | 0xe0
	})
}

// ISR is the Interrupt Service Routine, which handles the servicing of interrupts
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

// WriteVblankInterrupt signals that a VBLANK interrupt has occured
func WriteVblankInterrupt() {
	writeInterrupt(0x01)
}

// WriteLcdInterrupt signals that a LCD interrupt has occured
func WriteLcdInterrupt() {
	writeInterrupt(0x02)
}

// WriteTimerInterrupt signals that a Timer interrupt has occured
func WriteTimerInterrupt() {
	writeInterrupt(0x04)
}

// WriteSerialInterrupt signals that a Serial has occured
func WriteSerialInterrupt() {
	writeInterrupt(0x08)
}

// WriteJoypadInterrupt signals that a Joypad interrupt has occured
func WriteJoypadInterrupt() {
	writeInterrupt(0x10)
}

func writeInterrupt(interruptSignal uint8) {
	interruptFlags := mmu.Read(ifAddr)
	mmu.Write(ifAddr, interruptFlags|interruptSignal)
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
	cpu.IsHalted = false
	if !cpu.InterruptsEnabled && cpu.IsHalted {
		return
	}
	cpu.InterruptsEnabled = false
	interruptFlags := mmu.Read(ifAddr)
	interruptFlags &= interruptReset
	mmu.Write(ifAddr, interruptFlags)
	cpu.Call(interruptAddress)
}

func isVblankInterruptEnabled() bool {
	return mmu.Read(ieFlagAddr)&0x1 == 0x1
}

func isLcdInterruptEnabled() bool {
	return mmu.Read(ieFlagAddr)&0x2 == 0x2
}

func isTimerInterruptEnabled() bool {
	return mmu.Read(ieFlagAddr)&0x4 == 0x4
}

func isSerialInterruptEnabled() bool {
	return mmu.Read(ieFlagAddr)&0x8 == 0x8
}

func isJoypadInterruptEnabled() bool {
	return mmu.Read(ieFlagAddr)&0x10 == 0x10
}

func vblankInterruptOccured() bool {
	return mmu.Read(ifAddr)&0x1 == 0x1
}

func lcdInterruptOccured() bool {
	return mmu.Read(ifAddr)&0x2 == 0x2
}

func timerInterruptOccured() bool {
	return mmu.Read(ifAddr)&0x4 == 0x4
}

func serialInterruptOccured() bool {
	return mmu.Read(ifAddr)&0x8 == 0x8
}

func joypadInterruptOccured() bool {
	return mmu.Read(ifAddr)&0x10 == 0x10
}

func interruptsReady() bool {
	iflag := mmu.Read(ifAddr)
	ie := mmu.Read(ieFlagAddr)
	return (iflag & ie) != 0x00
}
