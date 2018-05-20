package cpu

import (
	"github.com/karimElmougi/GolangBoy/mmu"
)

var (
	// InterruptsEnabled indicates whether interrupts are enabled (IME)
	InterruptsEnabled bool
	// EnablingInterrupts indicates whether interrupts should be turned on during the ISR
	EnablingInterrupts bool
	// IsHalted indicates that the CPU is in a halted state
	IsHalted               bool
	a, b, c, d, e, h, l, f uint8
	pc, sp                 uint16
)

// Init initializes the CPU with the post-boot ROM register values
func Init() {
	a = 1
	b = 0
	c = 19
	d = 0
	e = 216
	h = 1
	l = 77
	f = 176
	pc = 0x100
	sp = 0xfffe
}

// Step advances the CPU by one instruction and returns the corresponding number of hardware cycles
func Step() uint64 {
	if IsHalted {
		return 4
	}
	opCode := mmu.Read(pc)
	return executeInstruction(opCode)
}

// Call pushes the current Program Counter to the stack and sets it to the given address
func Call(address uint16) {
	sp -= 2
	mmu.WriteWord(sp, pc)
	pc = address
}

func add(op1 uint8, op2 uint8) uint8 {
	return addImpl(op1, op2, false)
}

func addc(op1 uint8, op2 uint8) uint8 {
	return addImpl(op1, op2, true)
}

func addImpl(op1 uint8, op2 uint8, carrying bool) uint8 {
	carry := uint16(0)
	if carrying {
		carry = uint16((f & 0x10) >> 4)
	}
	result16 := uint16(op1) + uint16(op2) + carry
	result := uint8(result16)
	f = 0
	if result == 0 {
		f |= 0x80
	}
	if result16 > 0xff {
		f |= 0x10
	}
	if (op1&0xf)+(op2&0xf)+uint8(carry) > 0xf {
		f |= 0x20
	}
	return result
}

func sub(op1 uint8, op2 uint8) uint8 {
	return subImpl(op1, op2, false)
}

func subc(op1 uint8, op2 uint8) uint8 {
	return subImpl(op1, op2, true)
}

func subImpl(op1 uint8, op2 uint8, borrowing bool) uint8 {
	borrow := uint16(0)
	if borrowing {
		borrow = uint16((f & 0x10) >> 4)
	}
	result16 := uint16(op1) - uint16(op2) - borrow
	result := uint8(result16)
	f = 0x40
	if result == 0 {
		f |= 0x80
	}
	if result16 > 0xff {
		f |= 0x10
	}
	if (op1&0xf - op2&0xf - uint8(borrow)) > 0xf {
		f |= 0x20
	}
	return result
}
