package clock

import (
	"github.com/karimElmougi/GolangBoy/interrupts"
	"github.com/karimElmougi/GolangBoy/mmu"
)

const (
	DIVIDER_ADDR       = 0xff04
	COUNTER_ADDR       = 0xff05
	MODULO_ADDR        = 0xff06
	TIMER_CONTROL_ADDR = 0xff07
)

var (
	timerCounter   int = 1024
	dividerCounter int = 0
)

func Init() {
	mmu.RegisterSpecialWrite(DIVIDER_ADDR,
		func(addr uint16, value uint8) {
			mmu.RAM[DIVIDER_ADDR] = 0
			dividerCounter = 0
			timerCounter = getTimerFrequency()
		})
	mmu.RegisterSpecialWrite(COUNTER_ADDR,
		func(addr uint16, value uint8) {
			mmu.RAM[COUNTER_ADDR] = value
			timerCounter = getTimerFrequency()
		})
	mmu.RegisterSpecialWrite(TIMER_CONTROL_ADDR,
		func(addr uint16, value uint8) {
			old := mmu.RAM[TIMER_CONTROL_ADDR]
			mmu.RAM[TIMER_CONTROL_ADDR] = value
			if old != value {
				timerCounter = getTimerFrequency()
			}
		})
}

func Tick(cycles uint64) {
	dividerCounter += int(cycles)
	if dividerCounter >= 255 {
		dividerCounter -= 255
		mmu.RAM[DIVIDER_ADDR]++
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
	return mmu.Read(TIMER_CONTROL_ADDR)&0x04 == 0x04
}

func getTimerFrequency() int {
	switch mmu.Read(TIMER_CONTROL_ADDR) & 0x03 {
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
	timerVal := mmu.Read(COUNTER_ADDR)
	if timerVal == 255 {
		modulo := mmu.Read(MODULO_ADDR)
		mmu.Write(COUNTER_ADDR, modulo)
		interrupts.WriteTimerInterrupt()
	} else {
		mmu.Write(COUNTER_ADDR, timerVal+1)
	}
}
