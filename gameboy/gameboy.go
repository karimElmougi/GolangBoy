package gameboy

import (
	"github.com/karimElmougi/GolangBoy/cartridge"
	"github.com/karimElmougi/GolangBoy/clock"
	"github.com/karimElmougi/GolangBoy/cpu"
	"github.com/karimElmougi/GolangBoy/gpu"
	"github.com/karimElmougi/GolangBoy/interrupts"
	"github.com/karimElmougi/GolangBoy/joypad"
	"github.com/karimElmougi/GolangBoy/mmu"
)

// Boot initializes all the components of the GameBoy
func Boot(romName string) {
	mmu.Init()
	cartridge.Load(romName)
	clock.Init()
	interrupts.Init()
	joypad.Init()
	cpu.Init()
}

// Run steps the entire system by one second
func Run() {
	cyclesPerSecond := uint64(4194304 / 60)
	cyclesEllapsed := uint64(0)
	for i := uint64(0); i < cyclesPerSecond; i += cyclesEllapsed {
		cyclesEllapsed = cpu.Step()
		clock.Tick(cyclesEllapsed)
		gpu.Step(cyclesEllapsed)
		interrupts.ISR()
	}
}
