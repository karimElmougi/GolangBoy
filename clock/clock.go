package clock

import (
	"github.com/karimElmougi/GolangBoy/interrupts"
	"github.com/karimElmougi/GolangBoy/mmu"
)

const (
	dividerAddr      = 0xff04
	counterAddr      = 0xff05
	moduloAddr       = 0xff06
	timerControlAddr = 0xff07
)

var (
	timerCounter   = 1024
	dividerCounter int
)

// Init initializes the clock with SpecialWriters in the mmu
func Init() {
	mmu.RegisterSpecialWrite(dividerAddr,
		func(addr uint16, value uint8) {
			mmu.RAM[dividerAddr] = 0
			dividerCounter = 0
			timerCounter = getTimerFrequency()
		})
	mmu.RegisterSpecialWrite(counterAddr,
		func(addr uint16, value uint8) {
			mmu.RAM[counterAddr] = value
			timerCounter = getTimerFrequency()
		})
	mmu.RegisterSpecialWrite(timerControlAddr,
		func(addr uint16, value uint8) {
			old := mmu.RAM[timerControlAddr]
			mmu.RAM[timerControlAddr] = value
			if old != value {
				timerCounter = getTimerFrequency()
			}
		})
}

// Tick steps the clock through time based on the number of cycles provided in argument
func Tick(cycles uint64) {
	dividerCounter += int(cycles)
	if dividerCounter >= 255 {
		dividerCounter -= 255
		mmu.RAM[dividerAddr]++
	}
	if isTimerRunning() {
		timerCounter -= int(cycles)
		if timerCounter <= 0 {
			timerCounter += getTimerFrequency()
			incrementTimer()
		}
	}
}

func isTimerRunning() bool {
	return mmu.Read(timerControlAddr)&0x04 == 0x04
}

func getTimerFrequency() int {
	switch mmu.Read(timerControlAddr) & 0x03 {
	case 0x0:
		return 1024
	case 0x1:
		return 16
	case 0x2:
		return 64
	case 0x3:
		fallthrough
	default:
		return 256
	}
}

func incrementTimer() {
	timerVal := mmu.Read(counterAddr)
	if timerVal == 255 {
		modulo := mmu.Read(moduloAddr)
		mmu.Write(counterAddr, modulo)
		interrupts.WriteTimerInterrupt()
	} else {
		mmu.Write(counterAddr, timerVal+1)
	}
}
