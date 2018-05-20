package joypad

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/karimElmougi/GolangBoy/interrupts"
	"github.com/karimElmougi/GolangBoy/mmu"
)

type keyBinding struct {
	key       ebiten.Key
	binding   uint8
	isPressed bool
}

func (k *keyBinding) wasPressed() bool {
	return ebiten.IsKeyPressed(k.key) && !k.isPressed
}

var (
	inputMask   uint8 = 0xff
	keyBindings       = []keyBinding{
		keyBinding{ebiten.KeyS, 0x01, false},
		keyBinding{ebiten.KeyA, 0x02, false},
		keyBinding{ebiten.KeyQ, 0x04, false},
		keyBinding{ebiten.KeyW, 0x08, false},
		keyBinding{ebiten.KeyRight, 0x10, false},
		keyBinding{ebiten.KeyLeft, 0x20, false},
		keyBinding{ebiten.KeyUp, 0x40, false},
		keyBinding{ebiten.KeyDown, 0x80, false},
	}
)

func Init() {
	mmu.RegisterSpecialRead(0xff00, func(addr uint16) uint8 {
		return getJoypadRegister(mmu.RAM[addr])
	})
}

func UpdateInputs() {
	for _, key := range keyBindings {
		if key.wasPressed() {
			key.isPressed = true
			interrupts.WriteJoypadInterrupt()
			inputMask &= ^key.binding
		}
		if !ebiten.IsKeyPressed(key.key) {
			key.isPressed = false
			inputMask |= key.binding
		}
	}
}

func getJoypadRegister(currentValue uint8) uint8 {
	if currentValue&0x10 == 0x10 {
		return 0x0f&inputMask | currentValue | 0xc0
	} else if currentValue&0x20 == 0x20 {
		return 0x0f&(inputMask>>4) | currentValue | 0xc0
	}
	return 0x0f&currentValue | 0xc0
}
