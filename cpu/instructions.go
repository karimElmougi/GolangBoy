package cpu

import (
	"fmt"

	"github.com/karimElmougi/GolangBoy/mmu"
)

type instruction struct {
	name       string
	nbOperands uint16
	function   func(...*uint8) uint64
	nbCycles   uint64
	registers  []*uint8
}

func executeInstruction(opCode uint8) uint64 {
	inst := instructions[opCode]
	extraCycles := inst.function(inst.registers...)
	pc += inst.nbOperands + 1
	return inst.nbCycles + extraCycles
}

var (
	instructions = [256]instruction{
		{"NOP", 0, nop, 4, []*uint8{}},                     // 0x00
		{"LD BC, NN", 2, ld_rr_nn, 12, []*uint8{&b, &c}},   // 0x01
		{"LD (BC), A", 0, ld_rr_a, 8, []*uint8{&b, &c}},    // 0x02
		{"INC BC", 0, inc_rr, 8, []*uint8{&b, &c}},         // 0x03
		{"INC B", 0, inc_r, 4, []*uint8{&b}},               // 0x04
		{"DEC B", 0, dec_r, 4, []*uint8{&b}},               // 0x05
		{"LD B, N", 1, ld_r_n, 8, []*uint8{&b}},            // 0x06
		{"RLCA", 0, rlca, 4, []*uint8{}},                   // 0x07
		{"LD (NN), SP", 2, ld_nn_sp, 20, []*uint8{}},       // 0x08
		{"ADD HL, BC", 0, add_hl_rr, 8, []*uint8{&b, &c}},  // 0x09
		{"LD A, (BC)", 0, ld_a_rr, 8, []*uint8{&b, &c}},    // 0x0a
		{"DEC BC", 0, dec_rr, 8, []*uint8{&b, &c}},         // 0x0b
		{"INC C", 0, inc_r, 4, []*uint8{&c}},               // 0x0c
		{"DEC C", 0, dec_r, 4, []*uint8{&c}},               // 0x0d
		{"LD C, N", 1, ld_r_n, 8, []*uint8{&c}},            // 0x0e
		{"RRCA", 0, rrca, 4, []*uint8{}},                   // 0x0f
		{"STOP", 0, unimplemented, 4, []*uint8{}},          // 0x10
		{"LD DE, NN", 2, ld_rr_nn, 12, []*uint8{&d, &e}},   // 0x11
		{"LD (DE), A", 0, ld_rr_a, 8, []*uint8{&d, &e}},    // 0x12
		{"INC DE", 0, inc_rr, 8, []*uint8{&d, &e}},         // 0x13
		{"INC D", 0, inc_r, 4, []*uint8{&d}},               // 0x14
		{"DEC D", 0, dec_r, 4, []*uint8{&d}},               // 0x15
		{"LD D, N", 1, ld_r_n, 8, []*uint8{&d}},            // 0x16
		{"RLA", 0, rla, 4, []*uint8{}},                     // 0x17
		{"JR N", 1, jr_n, 8, []*uint8{}},                   // 0x18
		{"ADD HL, DE", 0, add_hl_rr, 8, []*uint8{&d, &e}},  // 0x19
		{"LD A, (DE)", 0, ld_a_rr, 8, []*uint8{&d, &e}},    // 0x1a
		{"DEC DE", 0, dec_rr, 8, []*uint8{&d, &e}},         // 0x1b
		{"INC E", 0, inc_r, 4, []*uint8{&e}},               // 0x1c
		{"DEC E", 0, dec_r, 4, []*uint8{&e}},               // 0x1d
		{"LD E, N", 1, ld_r_n, 8, []*uint8{&e}},            // 0x1e
		{"RRA", 0, rra, 4, []*uint8{}},                     // 0x1f
		{"JR NZ, N", 1, jrnz, 8, []*uint8{}},               // 0x20
		{"LD HL, NN", 2, ld_rr_nn, 12, []*uint8{&h, &l}},   // 0x21
		{"LDI (HL), A", 0, ldi_hl_a, 8, []*uint8{}},        // 0x22
		{"INC HL", 0, inc_rr, 8, []*uint8{&h, &l}},         // 0x23
		{"INC H", 0, inc_r, 4, []*uint8{&h}},               // 0x24
		{"DEC H", 0, dec_r, 4, []*uint8{&h}},               // 0x25
		{"LD H, N", 1, ld_r_n, 8, []*uint8{&h}},            // 0x26
		{"DAA", 0, daa, 4, []*uint8{}},                     // 0x27
		{"JR Z, N", 1, jrz, 8, []*uint8{}},                 // 0x28
		{"ADD HL, HL", 0, add_hl_rr, 8, []*uint8{&h, &l}},  // 0x29
		{"LDI A, (HL)", 0, ldi_a_hl, 8, []*uint8{}},        // 0x2a
		{"DEC HL", 0, dec_rr, 8, []*uint8{&h, &l}},         // 0x2b
		{"INC L", 0, inc_r, 4, []*uint8{&l}},               // 0x2c
		{"DEC L", 0, dec_r, 4, []*uint8{&l}},               // 0x2d
		{"LD L, N", 1, ld_r_n, 8, []*uint8{&l}},            // 0x2e
		{"CPL", 0, cpl, 4, []*uint8{}},                     // 0x2f
		{"JR NC, N", 1, jrnc, 8, []*uint8{}},               // 0x30
		{"LD SP, NN", 2, ld_sp_nn, 12, []*uint8{}},         // 0x31
		{"LDD (HL), A", 0, ldd_hl_a, 8, []*uint8{}},        // 0x32
		{"INC SP", 0, inc_sp, 8, []*uint8{}},               // 0x33
		{"INC (HL)", 0, inc_at_hl, 12, []*uint8{}},         // 0x34
		{"DEC (HL)", 0, dec_at_hl, 12, []*uint8{}},         // 0x35
		{"LD (HL), N", 1, ld_hl_nn, 12, []*uint8{}},        // 0x36
		{"SCF", 0, scf, 4, []*uint8{}},                     // 0x37
		{"JR C, N", 1, jrc, 8, []*uint8{}},                 // 0x38
		{"ADD HL, SP", 0, add_hl_sp, 8, []*uint8{}},        // 0x39
		{"LDD A, (HL)", 0, ldd_a_hl, 8, []*uint8{}},        // 0x3a
		{"DEC SP", 0, dec_sp, 8, []*uint8{}},               // 0x3b
		{"INC A", 0, inc_r, 4, []*uint8{&a}},               // 0x3c
		{"DEC A", 0, dec_r, 4, []*uint8{&a}},               // 0x3d
		{"LD A, N", 1, ld_r_n, 8, []*uint8{&a}},            // 0x3e
		{"CCF", 0, ccf, 4, []*uint8{}},                     // 0x3f
		{"LD B, B", 0, ld_r_r, 4, []*uint8{&b, &b}},        // 0x40
		{"LD B, C", 0, ld_r_r, 4, []*uint8{&b, &c}},        // 0x41
		{"LD B, D", 0, ld_r_r, 4, []*uint8{&b, &d}},        // 0x42
		{"LD B, E", 0, ld_r_r, 4, []*uint8{&b, &e}},        // 0x43
		{"LD B, H", 0, ld_r_r, 4, []*uint8{&b, &h}},        // 0x44
		{"LD B, L", 0, ld_r_r, 4, []*uint8{&b, &l}},        // 0x45
		{"LD B, (HL)", 0, ld_r_hl, 8, []*uint8{&b}},        // 0x46
		{"LD B, A", 0, ld_r_r, 4, []*uint8{&b, &a}},        // 0x47
		{"LD C, B", 0, ld_r_r, 4, []*uint8{&c, &b}},        // 0x48
		{"LD C, C", 0, ld_r_r, 4, []*uint8{&c, &c}},        // 0x49
		{"LD C, D", 0, ld_r_r, 4, []*uint8{&c, &d}},        // 0x4a
		{"LD C, E", 0, ld_r_r, 4, []*uint8{&c, &e}},        // 0x4b
		{"LD C, H", 0, ld_r_r, 4, []*uint8{&c, &h}},        // 0x4c
		{"LD C, L", 0, ld_r_r, 4, []*uint8{&c, &l}},        // 0x4d
		{"LD C, (HL)", 0, ld_r_hl, 8, []*uint8{&c}},        // 0x4e
		{"LD C, A", 0, ld_r_r, 4, []*uint8{&c, &a}},        // 0x4f
		{"LD D, B", 0, ld_r_r, 4, []*uint8{&d, &b}},        // 0x50
		{"LD D, C", 0, ld_r_r, 4, []*uint8{&d, &c}},        // 0x51
		{"LD D, D", 0, ld_r_r, 4, []*uint8{&d, &d}},        // 0x52
		{"LD D, E", 0, ld_r_r, 4, []*uint8{&d, &e}},        // 0x53
		{"LD D, H", 0, ld_r_r, 4, []*uint8{&d, &h}},        // 0x54
		{"LD D, L", 0, ld_r_r, 4, []*uint8{&d, &l}},        // 0x55
		{"LD D, (HL)", 0, ld_r_hl, 8, []*uint8{&d}},        // 0x56
		{"LD D, A", 0, ld_r_r, 4, []*uint8{&d, &a}},        // 0x57
		{"LD E, B", 0, ld_r_r, 4, []*uint8{&e, &b}},        // 0x58
		{"LD E, C", 0, ld_r_r, 4, []*uint8{&e, &c}},        // 0x59
		{"LD E, D", 0, ld_r_r, 4, []*uint8{&e, &d}},        // 0x5a
		{"LD E, E", 0, ld_r_r, 4, []*uint8{&e, &e}},        // 0x5b
		{"LD E, H", 0, ld_r_r, 4, []*uint8{&e, &h}},        // 0x5c
		{"LD E, L", 0, ld_r_r, 4, []*uint8{&e, &l}},        // 0x5d
		{"LD E, (HL)", 0, ld_r_hl, 8, []*uint8{&e}},        // 0x5e
		{"LD E, A", 0, ld_r_r, 4, []*uint8{&e, &a}},        // 0x5f
		{"LD H, B", 0, ld_r_r, 4, []*uint8{&h, &b}},        // 0x60
		{"LD H, C", 0, ld_r_r, 4, []*uint8{&h, &c}},        // 0x61
		{"LD H, D", 0, ld_r_r, 4, []*uint8{&h, &d}},        // 0x62
		{"LD H, E", 0, ld_r_r, 4, []*uint8{&h, &e}},        // 0x63
		{"LD H, H", 0, ld_r_r, 4, []*uint8{&h, &h}},        // 0x64
		{"LD H, L", 0, ld_r_r, 4, []*uint8{&h, &l}},        // 0x65
		{"LD H, (HL)", 0, ld_r_hl, 8, []*uint8{&h}},        // 0x66
		{"LD H, A", 0, ld_r_r, 4, []*uint8{&h, &a}},        // 0x67
		{"LD L, B", 0, ld_r_r, 4, []*uint8{&l, &b}},        // 0x68
		{"LD L, C", 0, ld_r_r, 4, []*uint8{&l, &c}},        // 0x69
		{"LD L, D", 0, ld_r_r, 4, []*uint8{&l, &d}},        // 0x6a
		{"LD L, E", 0, ld_r_r, 4, []*uint8{&l, &e}},        // 0x6b
		{"LD L, H", 0, ld_r_r, 4, []*uint8{&l, &h}},        // 0x6c
		{"LD L, L", 0, ld_r_r, 4, []*uint8{&l, &l}},        // 0x6d
		{"LD L, (HL)", 0, ld_r_hl, 8, []*uint8{&l}},        // 0x6e
		{"LD L, A", 0, ld_r_r, 4, []*uint8{&l, &a}},        // 0x6f
		{"LD (HL), B", 0, ld_hl_r, 8, []*uint8{&b}},        // 0x70
		{"LD (HL), C", 0, ld_hl_r, 8, []*uint8{&c}},        // 0x71
		{"LD (HL), D", 0, ld_hl_r, 8, []*uint8{&d}},        // 0x72
		{"LD (HL), E", 0, ld_hl_r, 8, []*uint8{&e}},        // 0x73
		{"LD (HL), H", 0, ld_hl_r, 8, []*uint8{&h}},        // 0x74
		{"LD (HL), L", 0, ld_hl_r, 8, []*uint8{&l}},        // 0x75
		{"HALT", 0, halt, 0, []*uint8{}},                   // 0x76
		{"LD (HL), A", 0, ld_hl_r, 8, []*uint8{&a}},        // 0x77
		{"LD A, B", 0, ld_r_r, 4, []*uint8{&a, &b}},        // 0x78
		{"LD A, C", 0, ld_r_r, 4, []*uint8{&a, &c}},        // 0x79
		{"LD A, D", 0, ld_r_r, 4, []*uint8{&a, &d}},        // 0x7a
		{"LD A, E", 0, ld_r_r, 4, []*uint8{&a, &e}},        // 0x7b
		{"LD A, H", 0, ld_r_r, 4, []*uint8{&a, &h}},        // 0x7c
		{"LD A, L", 0, ld_r_r, 4, []*uint8{&a, &l}},        // 0x7d
		{"LD A, (HL)", 0, ld_r_hl, 8, []*uint8{&a}},        // 0x7e
		{"LD A, A", 0, ld_r_r, 4, []*uint8{&a, &a}},        // 0x7f
		{"ADD A, B", 0, add_a_r, 4, []*uint8{&b}},          // 0x80
		{"ADD A, C", 0, add_a_r, 4, []*uint8{&c}},          // 0x81
		{"ADD A, D", 0, add_a_r, 4, []*uint8{&d}},          // 0x82
		{"ADD A, E", 0, add_a_r, 4, []*uint8{&e}},          // 0x83
		{"ADD A, H", 0, add_a_r, 4, []*uint8{&h}},          // 0x84
		{"ADD A, L", 0, add_a_r, 4, []*uint8{&l}},          // 0x85
		{"ADD A, (HL)", 0, add_a_hl, 8, []*uint8{}},        // 0x86
		{"ADD A", 0, add_a_r, 4, []*uint8{&a}},             // 0x87
		{"ADC B", 0, adc_a_r, 4, []*uint8{&b}},             // 0x88
		{"ADC C", 0, adc_a_r, 4, []*uint8{&c}},             // 0x89
		{"ADC D", 0, adc_a_r, 4, []*uint8{&d}},             // 0x8a
		{"ADC E", 0, adc_a_r, 4, []*uint8{&e}},             // 0x8b
		{"ADC H", 0, adc_a_r, 4, []*uint8{&h}},             // 0x8c
		{"ADC L", 0, adc_a_r, 4, []*uint8{&l}},             // 0x8d
		{"ADC (HL)", 0, adc_a_hl, 8, []*uint8{}},           // 0x8e
		{"ADC A", 0, adc_a_r, 4, []*uint8{&a}},             // 0x8f
		{"SUB B", 0, sub_a_r, 4, []*uint8{&b}},             // 0x90
		{"SUB C", 0, sub_a_r, 4, []*uint8{&c}},             // 0x91
		{"SUB D", 0, sub_a_r, 4, []*uint8{&d}},             // 0x92
		{"SUB E", 0, sub_a_r, 4, []*uint8{&e}},             // 0x93
		{"SUB H", 0, sub_a_r, 4, []*uint8{&h}},             // 0x94
		{"SUB L", 0, sub_a_r, 4, []*uint8{&l}},             // 0x95
		{"SUB (HL)", 0, sub_a_hl, 8, []*uint8{}},           // 0x96
		{"SUB A", 0, sub_a_r, 4, []*uint8{&a}},             // 0x97
		{"SBC B", 0, sbc_a_r, 4, []*uint8{&b}},             // 0x98
		{"SBC C", 0, sbc_a_r, 4, []*uint8{&c}},             // 0x99
		{"SBC D", 0, sbc_a_r, 4, []*uint8{&d}},             // 0x9a
		{"SBC E", 0, sbc_a_r, 4, []*uint8{&e}},             // 0x9b
		{"SBC H", 0, sbc_a_r, 4, []*uint8{&h}},             // 0x9c
		{"SBC L", 0, sbc_a_r, 4, []*uint8{&l}},             // 0x9d
		{"SBC (HL)", 0, sbc_a_hl, 8, []*uint8{}},           // 0x9e
		{"SBC A", 0, sbc_a_r, 4, []*uint8{&a}},             // 0x9f
		{"AND B", 0, and_a_r, 4, []*uint8{&b}},             // 0xa0
		{"AND C", 0, and_a_r, 4, []*uint8{&c}},             // 0xa1
		{"AND D", 0, and_a_r, 4, []*uint8{&d}},             // 0xa2
		{"AND E", 0, and_a_r, 4, []*uint8{&e}},             // 0xa3
		{"AND H", 0, and_a_r, 4, []*uint8{&h}},             // 0xa4
		{"AND L", 0, and_a_r, 4, []*uint8{&l}},             // 0xa5
		{"AND (HL)", 0, and_a_hl, 8, []*uint8{}},           // 0xa6
		{"AND A", 0, and_a_r, 4, []*uint8{&a}},             // 0xa7
		{"XOR B", 0, xor_a_r, 4, []*uint8{&b}},             // 0xa8
		{"XOR C", 0, xor_a_r, 4, []*uint8{&c}},             // 0xa9
		{"XOR D", 0, xor_a_r, 4, []*uint8{&d}},             // 0xaa
		{"XOR E", 0, xor_a_r, 4, []*uint8{&e}},             // 0xab
		{"XOR H", 0, xor_a_r, 4, []*uint8{&h}},             // 0xac
		{"XOR L", 0, xor_a_r, 4, []*uint8{&l}},             // 0xad
		{"XOR (HL)", 0, xor_a_hl, 8, []*uint8{}},           // 0xae
		{"XOR A", 0, xor_a_r, 4, []*uint8{&a}},             // 0xaf
		{"OR B", 0, or_a_r, 4, []*uint8{&b}},               // 0xb0
		{"OR C", 0, or_a_r, 4, []*uint8{&c}},               // 0xb1
		{"OR D", 0, or_a_r, 4, []*uint8{&d}},               // 0xb2
		{"OR E", 0, or_a_r, 4, []*uint8{&e}},               // 0xb3
		{"OR H", 0, or_a_r, 4, []*uint8{&h}},               // 0xb4
		{"OR L", 0, or_a_r, 4, []*uint8{&l}},               // 0xb5
		{"OR (HL)", 0, or_a_hl, 8, []*uint8{}},             // 0xb6
		{"OR A", 0, or_a_r, 4, []*uint8{&a}},               // 0xb7
		{"CP B", 0, cp_a_r, 4, []*uint8{&b}},               // 0xb8
		{"CP C", 0, cp_a_r, 4, []*uint8{&c}},               // 0xb9
		{"CP D", 0, cp_a_r, 4, []*uint8{&d}},               // 0xba
		{"CP E", 0, cp_a_r, 4, []*uint8{&e}},               // 0xbb
		{"CP H", 0, cp_a_r, 4, []*uint8{&h}},               // 0xbc
		{"CP L", 0, cp_a_r, 4, []*uint8{&l}},               // 0xbd
		{"CP (HL)", 0, cp_a_hl, 8, []*uint8{}},             // 0xbe
		{"CP A", 0, cp_a_r, 4, []*uint8{&a}},               // 0xbf
		{"RET NZ", 0, retnz, 8, []*uint8{}},                // 0xc0
		{"POP BC", 0, pop, 12, []*uint8{&b, &c}},           // 0xc1
		{"JP NZ, NN", 2, jpnz, 12, []*uint8{}},             // 0xc2
		{"JP NN", 2, jpnn, 12, []*uint8{}},                 // 0xc3
		{"CALL NZ, NN", 2, callnz, 12, []*uint8{}},         // 0xc4
		{"PUSH BC", 0, push, 16, []*uint8{&b, &c}},         // 0xc5
		{"ADD A, N", 1, add_a_n, 8, []*uint8{}},            // 0xc6
		{"RST 0x00", 0, rst0x00, 32, []*uint8{}},           // 0xc7
		{"RET Z", 0, retz, 8, []*uint8{}},                  // 0xc8
		{"RET", 0, ret, 8, []*uint8{}},                     // 0xc9
		{"JP Z, NN", 2, jpz, 12, []*uint8{}},               // 0xca
		{"CB N", 1, cb, 0, []*uint8{}},                     // 0xcb
		{"CALL Z, NN", 2, callz, 12, []*uint8{}},           // 0xcc
		{"CALL NN", 2, call, 12, []*uint8{}},               // 0xcd
		{"ADC N", 1, adc_a_n, 8, []*uint8{}},               // 0xce
		{"RST 0x08", 0, rst0x08, 32, []*uint8{}},           // 0xcf
		{"RET NC", 0, retnc, 8, []*uint8{}},                // 0xd0
		{"POP DE", 0, pop, 12, []*uint8{&d, &e}},           // 0xd1
		{"JP NC, NN", 2, jpnc, 12, []*uint8{}},             // 0xd2
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xd3
		{"CALL NC, NN", 2, callnc, 12, []*uint8{}},         // 0xd4
		{"PUSH DE", 0, push, 16, []*uint8{&d, &e}},         // 0xd5
		{"SUB N", 1, sub_a_n, 8, []*uint8{}},               // 0xd6
		{"RST 0x10", 0, rst0x10, 32, []*uint8{}},           // 0xd7
		{"RET C", 0, retc, 8, []*uint8{}},                  // 0xd8
		{"RETI", 0, reti, 8, []*uint8{}},                   // 0xd9
		{"JP C, NN", 2, jpc, 12, []*uint8{}},               // 0xda
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xdb
		{"CALL C, NN", 2, callc, 12, []*uint8{}},           // 0xdc
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xdd
		{"SBC N", 1, sbc_a_n, 8, []*uint8{}},               // 0xde
		{"RST 0x18", 0, rst0x18, 32, []*uint8{}},           // 0xdf
		{"LD (0xFF00 + N), A", 1, ld_n_a, 12, []*uint8{}},  // 0xe0
		{"POP HL", 0, pop, 12, []*uint8{&h, &l}},           // 0xe1
		{"LD (0xFF00 + C), A", 0, ld_c_a, 8, []*uint8{}},   // 0xe2
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xe3
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xe4
		{"PUSH HL", 0, push, 16, []*uint8{&h, &l}},         // 0xe5
		{"AND N", 1, and_a_n, 8, []*uint8{}},               // 0xe6
		{"RST 0x20", 0, rst0x20, 32, []*uint8{}},           // 0xe7
		{"ADD SP,N", 1, add_sp_n, 16, []*uint8{}},          // 0xe8
		{"JP HL", 0, jp_hl, 4, []*uint8{}},                 // 0xe9
		{"LD (NN), A", 2, ld_nn_a, 16, []*uint8{}},         // 0xea
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xeb
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xec
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xed
		{"XOR N", 1, xor_n, 8, []*uint8{}},                 // 0xee
		{"RST 0x28", 0, rst0x28, 32, []*uint8{}},           // 0xef
		{"LD A, (0xFF00 + N)", 1, ld_a_n, 12, []*uint8{}},  // 0xf0
		{"POP AF", 0, pop_af, 12, []*uint8{}},              // 0xf1
		{"LD A, (0xFF00 + C)", 0, ld_a_c, 8, []*uint8{}},   // 0xf2
		{"DI", 0, di, 4, []*uint8{}},                       // 0xf3
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xf4
		{"PUSH AF", 0, push, 16, []*uint8{&a, &f}},         // 0xf5
		{"OR N", 1, or_n, 8, []*uint8{}},                   // 0xf6
		{"RST 0x30", 0, rst0x30, 32, []*uint8{}},           // 0xf7
		{"LD HL, SP+N", 1, ld_hl_sp_n, 12, []*uint8{}},     // 0xf8
		{"LD SP, HL", 0, ld_sp_hl, 8, []*uint8{}},          // 0xf9
		{"LD A, (NN)", 2, ld_a_nn, 16, []*uint8{}},         // 0xfa
		{"EI", 0, ei, 4, []*uint8{}},                       // 0xfb
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xfc
		{"unimplemented", 0, unimplemented, 4, []*uint8{}}, // 0xfd
		{"CP N", 1, cp_a_n, 8, []*uint8{}},                 // 0xfe
		{"RST 0x38", 0, rst0x38, 32, []*uint8{}},           // 0xff
	}
)

func cb(registers ...*uint8) uint64 {
	value := mmu.Read(pc + 1)
	return executeExtendedInstruction(value)
}

func unimplemented(registers ...*uint8) uint64 {
	fmt.Printf("Unimplemented instructions: 0x%x\n", mmu.Read(pc))
	return 0
}

func nop(registers ...*uint8) uint64 {
	return 0
}

func halt(registers ...*uint8) uint64 {
	IsHalted = true
	return 0
}

func ld_rr_nn(registers ...*uint8) uint64 {
	value := mmu.ReadWord(pc + 1)
	*registers[0] = uint8(value >> 8)
	*registers[1] = uint8(value & 0xff)
	return 0
}

func ld_rr_a(registers ...*uint8) uint64 {
	addr := uint16(*registers[0])<<8 | uint16(*registers[1])
	mmu.Write(addr, a)
	return 0
}

func inc_rr(registers ...*uint8) uint64 {
	RR := (uint16(*registers[0])<<8 | uint16(*registers[1])) + 1
	*registers[0] = uint8(RR >> 8)
	*registers[1] = uint8(RR & 0xff)
	return 0
}

func dec_rr(registers ...*uint8) uint64 {
	RR := (uint16(*registers[0])<<8 | uint16(*registers[1])) - 1
	*registers[0] = uint8(RR >> 8)
	*registers[1] = uint8(RR & 0xff)
	return 0
}

func inc_r(registers ...*uint8) uint64 {
	oldCarry := f & 0x10
	*registers[0] = add(*registers[0], 1)
	if oldCarry == 0x10 {
		f |= oldCarry
	} else {
		f &= 0xef
	}
	return 0
}

func dec_r(registers ...*uint8) uint64 {
	*registers[0]--
	f |= 0x40
	if *registers[0] == 0 {
		f |= 0x80
	} else {
		f &= 0x7f
	}
	if (*registers[0] & 0xf) == 0xf {
		f |= 0x20
	} else {
		f &= 0xdf
	}
	return 0
}

func ld_r_n(registers ...*uint8) uint64 {
	*registers[0] = mmu.Read(pc + 1)
	return 0
}

func rlca(registers ...*uint8) uint64 {
	bit7 := a >> 7
	a = (a << 1) | bit7
	if bit7 == 0 {
		f = 0
	} else {
		f = 0x10
	}
	return 0
}

func ld_nn_sp(registers ...*uint8) uint64 {
	addr := mmu.ReadWord(pc + 1)
	mmu.WriteWord(addr, sp)
	return 0
}

func add_hl_rr(registers ...*uint8) uint64 {
	HL := uint16(h)<<8 | uint16(l)
	RR := uint16(*registers[0])<<8 | uint16(*registers[1])
	result := uint32(HL) + uint32(RR)
	f |= 0x40
	if result > 0xffff {
		f |= 0x10
	} else {
		f &= 0xef
	}
	if ((RR & 0xfff) + (HL & 0xfff)) > 0xfff {
		f |= 0x20
	} else {
		f &= 0xdf
	}
	f &= 0xbf
	result16 := uint16(result)
	h = uint8(result16 >> 8)
	l = uint8(result16 & 0xff)
	return 0
}

func ld_a_rr(registers ...*uint8) uint64 {
	addr := uint16(*registers[0])<<8 | uint16(*registers[1])
	a = mmu.Read(addr)
	return 0
}

func rrca(registers ...*uint8) uint64 {
	bit0 := a & 0x1
	a = (a >> 1) | (bit0 << 7)
	if bit0 == 0 {
		f = 0
	} else {
		f = 0x10
	}
	return 0
}

func rla(registers ...*uint8) uint64 {
	bit7 := a >> 7
	a <<= 1
	if (f & 0x10) == 0x10 {
		a |= 0x1
	}
	if bit7 != 0 {
		f = 0x10
	} else {
		f = 0
	}
	return 0
}

func rra(registers ...*uint8) uint64 {
	oldCarry := (f & 0x10) << 3
	if (a & 0x1) == 0x1 {
		f = 0x10
	} else {
		f = 0
	}
	a >>= 1
	a |= oldCarry
	return 0
}

func jr_n(registers ...*uint8) uint64 {
	return jrcc(true)
}

func jrnz(registers ...*uint8) uint64 {
	return jrcc(f&0x80 != 0x80)
}

func jrz(registers ...*uint8) uint64 {
	return jrcc(f&0x80 == 0x80)
}

func jrnc(registers ...*uint8) uint64 {
	return jrcc(f&0x10 != 0x10)
}

func jrc(registers ...*uint8) uint64 {
	return jrcc(f&0x10 == 0x10)
}

func jrcc(condition bool) uint64 {
	if condition {
		pc += uint16(int8(mmu.Read(pc + 1)))
		return 4
	}
	return 0
}

func ldi_hl_a(registers ...*uint8) uint64 {
	HL := uint16(h)<<8 | uint16(l)
	mmu.Write(HL, a)
	HL++
	h = uint8(HL >> 8)
	l = uint8(HL & 0xff)
	return 0
}

func daa(registers ...*uint8) uint64 {
	if f&0x40 == 0x40 {
		if f&0x20 == 0x20 {
			f &= 0xdf
			if f&0x10 == 0x10 {
				a += 0x9a
			} else {
				a += 0xfa
			}
		} else if f&0x10 == 0x10 {
			a += 0xa0
		}
	} else {
		if a > 0x99 || f&0x10 == 0x10 {
			a += 0x60
			f |= 0x10
		}
		if (a&0xf) > 0x09 || f&0x20 == 0x20 {
			a += 0x06
			f &= 0xdf
		}
	}
	if a == 0 {
		f |= 0x80
	} else {
		f &= 0x7f
	}
	return 0
}

func ldi_a_hl(registers ...*uint8) uint64 {
	HL := uint16(h)<<8 | uint16(l)
	a = mmu.Read(HL)
	HL++
	h = uint8(HL >> 8)
	l = uint8(HL & 0xff)
	return 0
}

func cpl(registers ...*uint8) uint64 {
	a = ^a
	f |= 0x40
	f |= 0x20
	return 0
}

func ld_sp_nn(registers ...*uint8) uint64 {
	sp = mmu.ReadWord(pc + 1)
	return 0
}

func ldd_hl_a(registers ...*uint8) uint64 {
	HL := uint16(h)<<8 | uint16(l)
	mmu.Write(HL, a)
	HL--
	h = uint8(HL >> 8)
	l = uint8(HL & 0xff)
	return 0
}

func inc_sp(registers ...*uint8) uint64 {
	sp++
	return 0
}

func inc_at_hl(registers ...*uint8) uint64 {
	oldCarry := f & 0x10
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	value = add(value, 1)
	if oldCarry == 0x10 {
		f |= oldCarry
	} else {
		f &= 0xef
	}
	mmu.Write(addr, value)
	return 0
}

func dec_at_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	value--
	f |= 0x40
	if value == 0 {
		f |= 0x80
	} else {
		f &= 0x7f
	}
	if (value & 0xf) == 0xf {
		f |= 0x20
	} else {
		f &= 0xdf
	}
	mmu.Write(addr, value)
	return 0
}

func ld_hl_nn(registers ...*uint8) uint64 {
	value := mmu.Read(pc + 1)
	addr := uint16(h)<<8 | uint16(l)
	mmu.Write(addr, value)
	return 0
}

func scf(registers ...*uint8) uint64 {
	oldZero := f & 0x80
	f = 0x10
	if oldZero == 0x80 {
		f |= oldZero
	} else {
		f &= 0x7f
	}
	return 0
}

func add_hl_sp(registers ...*uint8) uint64 {
	HL := uint16(h)<<8 | uint16(l)
	SP := sp
	result := uint32(HL) + uint32(SP)
	f |= 0x40
	if result > 0xffff {
		f |= 0x10
	} else {
		f &= 0xef
	}
	if ((SP & 0xfff) + (HL & 0xfff)) > 0xfff {
		f |= 0x20
	} else {
		f &= 0xdf
	}
	f &= 0xbf
	result16 := uint16(result)
	h = uint8(result16 >> 8)
	l = uint8(result16 & 0xff)
	return 0
}

func ldd_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	a = mmu.Read(addr)
	addr--
	h = uint8(addr >> 8)
	l = uint8(addr & 0xff)
	return 0
}

func dec_sp(registers ...*uint8) uint64 {
	sp--
	return 0
}

func ccf(registers ...*uint8) uint64 {
	f ^= 0x10
	f &= 0xdf
	f &= 0xbf
	return 0
}

func ld_r_r(registers ...*uint8) uint64 {
	*registers[0] = *registers[1]
	return 0
}

func ld_r_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	*registers[0] = mmu.Read(addr)
	return 0
}

func ld_hl_r(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	mmu.Write(addr, *registers[0])
	return 0
}

func add_a_r(registers ...*uint8) uint64 {
	a = add(a, *registers[0])
	return 0
}

func adc_a_r(registers ...*uint8) uint64 {
	a = addc(a, *registers[0])
	return 0
}

func add_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	a = add(a, value)
	return 0
}

func adc_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	a = addc(a, value)
	return 0
}

func sub_a_r(registers ...*uint8) uint64 {
	a = sub(a, *registers[0])
	return 0
}

func sbc_a_r(registers ...*uint8) uint64 {
	a = subc(a, *registers[0])
	return 0
}

func sub_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	a = sub(a, value)
	return 0
}

func sbc_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	a = subc(a, value)
	return 0
}

func and_a_r(registers ...*uint8) uint64 {
	a &= *registers[0]
	f = 0x20
	if a == 0 {
		f |= 0x80
	}
	return 0
}

func and_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	a &= value
	f = 0x20
	if a == 0 {
		f |= 0x80
	}
	return 0
}

func xor_a_r(registers ...*uint8) uint64 {
	a ^= *registers[0]
	if a == 0 {
		f = 0x80
	} else {
		f = 0x00
	}
	return 0
}

func xor_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	a ^= value
	if a == 0 {
		f = 0x80
	} else {
		f = 0x00
	}
	return 0
}

func or_a_r(registers ...*uint8) uint64 {
	a |= *registers[0]
	if a == 0 {
		f = 0x80
	} else {
		f = 0x00
	}
	return 0
}

func or_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	a |= value
	if a == 0 {
		f = 0x80
	} else {
		f = 0x00
	}
	return 0
}

func cp_a_r(registers ...*uint8) uint64 {
	sub(a, *registers[0])
	return 0
}

func cp_a_hl(registers ...*uint8) uint64 {
	addr := uint16(h)<<8 | uint16(l)
	value := mmu.Read(addr)
	sub(a, value)
	return 0
}

func ret(registers ...*uint8) uint64 {
	retcc(true)
	return 8
}

func retnz(registers ...*uint8) uint64 {
	return retcc(f&0x80 != 0x80)
}

func retz(registers ...*uint8) uint64 {
	return retcc(f&0x80 == 0x80)
}

func retnc(registers ...*uint8) uint64 {
	return retcc(f&0x10 != 0x10)
}

func retc(registers ...*uint8) uint64 {
	return retcc(f&0x10 == 0x10)
}

func retcc(condition bool) uint64 {
	if condition {
		pc = mmu.ReadWord(sp) - 1
		sp += 2
		return 12
	}
	return 0
}

func pop(registers ...*uint8) uint64 {
	value := mmu.ReadWord(sp)
	*registers[0] = uint8(value >> 8)
	*registers[1] = uint8(value & 0xff)
	sp += 2
	return 0
}

func pop_af(registers ...*uint8) uint64 {
	value := mmu.ReadWord(sp)
	a = uint8(value >> 8)
	f = uint8(value & 0xf0)
	sp += 2
	return 0
}

func jpnn(registers ...*uint8) uint64 {
	return jpcc(true)
}

func jpnz(registers ...*uint8) uint64 {
	return jpcc(f&0x80 != 0x80)
}

func jpz(registers ...*uint8) uint64 {
	return jpcc(f&0x80 == 0x80)
}

func jpnc(registers ...*uint8) uint64 {
	return jpcc(f&0x10 != 0x10)
}

func jpc(registers ...*uint8) uint64 {
	return jpcc(f&0x10 == 0x10)
}

func jpcc(condition bool) uint64 {
	if condition {
		pc = mmu.ReadWord(pc + 1)
		pc -= 3
		return 4
	}
	return 0
}

func call(registers ...*uint8) uint64 {
	return callcc(true)
}

func callnz(registers ...*uint8) uint64 {
	return callcc(f&0x80 != 0x80)
}

func callz(registers ...*uint8) uint64 {
	return callcc(f&0x80 == 0x80)
}

func callnc(registers ...*uint8) uint64 {
	return callcc(f&0x10 != 0x10)
}

func callc(registers ...*uint8) uint64 {
	return callcc(f&0x10 == 0x10)
}

func callcc(condition bool) uint64 {
	if condition {
		sp -= 2
		mmu.WriteWord(sp, pc+3)
		pc = mmu.ReadWord(pc+1) - 3
		return 12
	}
	return 0
}

func push(registers ...*uint8) uint64 {
	sp -= 2
	mmu.WriteWord(sp, uint16(*registers[0])<<8|uint16(*registers[1]))
	return 0
}

func add_a_n(registers ...*uint8) uint64 {
	value := mmu.Read(pc + 1)
	a = add(a, value)
	return 0
}

func rst0x00(registers ...*uint8) uint64 {
	return rst(0x00)
}

func rst0x08(registers ...*uint8) uint64 {
	return rst(0x08)
}

func rst0x10(registers ...*uint8) uint64 {
	return rst(0x10)
}

func rst0x18(registers ...*uint8) uint64 {
	return rst(0x18)
}

func rst0x20(registers ...*uint8) uint64 {
	return rst(0x20)
}

func rst0x28(registers ...*uint8) uint64 {
	return rst(0x28)
}

func rst0x30(registers ...*uint8) uint64 {
	return rst(0x30)
}

func rst0x38(registers ...*uint8) uint64 {
	return rst(0x38)
}

func rst(addr uint16) uint64 {
	sp -= 2
	mmu.WriteWord(sp, pc+1)
	pc = addr - 1
	return 0
}

func adc_a_n(registers ...*uint8) uint64 {
	value := mmu.Read(pc + 1)
	a = addc(a, value)
	return 0
}

func sub_a_n(registers ...*uint8) uint64 {
	value := mmu.Read(pc + 1)
	a = sub(a, value)
	return 0
}

func reti(registers ...*uint8) uint64 {
	extraCycles := ret(registers...)
	EnablingInterrupts = true
	return extraCycles
}

func sbc_a_n(registers ...*uint8) uint64 {
	value := mmu.Read(pc + 1)
	a = subc(a, value)
	return 0
}

func ld_n_a(registers ...*uint8) uint64 {
	addr := 0xff00 + uint16(mmu.Read(pc+1))
	mmu.Write(addr, a)
	return 0
}

func ld_c_a(registers ...*uint8) uint64 {
	addr := 0xff00 + uint16(c)
	mmu.Write(addr, a)
	return 0
}

func and_a_n(registers ...*uint8) uint64 {
	value := mmu.Read(pc + 1)
	a &= value
	f = 0x20
	if a == 0 {
		f |= 0x80
	}
	return 0
}

func add_sp_n(registers ...*uint8) uint64 {
	n := mmu.Read(pc + 1)
	spLow := uint8(sp & 0xff)
	add(n, spLow)
	f &= 0x7f
	f &= 0xbf
	sp += uint16(int8(n))
	return 0
}

func jp_hl(registers ...*uint8) uint64 {
	pc = uint16(h)<<8 | uint16(l) - 1
	return 0
}

func ld_nn_a(registers ...*uint8) uint64 {
	addr := mmu.ReadWord(pc + 1)
	mmu.Write(addr, a)
	return 0
}

func xor_n(registers ...*uint8) uint64 {
	a ^= mmu.Read(pc + 1)
	if a == 0 {
		f = 0x80
	} else {
		f = 0
	}
	return 0
}

func ld_a_n(registers ...*uint8) uint64 {
	addr := 0xff00 + uint16(mmu.Read(pc+1))
	a = mmu.Read(addr)
	return 0
}

func ld_a_c(registers ...*uint8) uint64 {
	addr := 0xff00 + uint16(c)
	a = mmu.Read(addr)
	return 0
}

func di(registers ...*uint8) uint64 {
	InterruptsEnabled = false
	return 0
}

func or_n(registers ...*uint8) uint64 {
	a |= mmu.Read(pc + 1)
	if a == 0 {
		f = 0x80
	} else {
		f = 0
	}
	return 0
}

func ld_hl_sp_n(registers ...*uint8) uint64 {
	n := mmu.Read(pc + 1)
	spLow := uint8(sp & 0xff)
	add(n, spLow)
	f &= 0x7f
	f &= 0xbf
	result := sp + uint16(int8(n))
	h = uint8(result >> 8)
	l = uint8(result & 0xff)
	return 0
}

func ld_sp_hl(registers ...*uint8) uint64 {
	sp = uint16(h)<<8 | uint16(l)
	return 0
}

func ld_a_nn(registers ...*uint8) uint64 {
	addr := uint16(mmu.ReadWord(pc + 1))
	a = mmu.Read(addr)
	return 0
}

func ei(registers ...*uint8) uint64 {
	EnablingInterrupts = true
	return 0
}

func cp_a_n(registers ...*uint8) uint64 {
	value := mmu.Read(pc + 1)
	sub(a, value)
	return 0
}
