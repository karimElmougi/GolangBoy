package cpu

import (
	"github.com/karimElmougi/GolangBoy/mmu"
)

type extendedInstruction struct {
	name     string
	function func(*uint8, uint8)
	nbCycles uint64
	register *uint8
}

func executeExtendedInstruction(opCode uint8) uint64 {
	inst := extendedInstructions[opCode]
	inst.function(inst.register, ((opCode/8)-8)%8)
	return inst.nbCycles
}

var (
	extendedInstructions = [256]extendedInstruction{
		{"RLC B", rlc_r, 8, &b},          // 0x00
		{"RLC C", rlc_r, 8, &c},          // 0x01
		{"RLC D", rlc_r, 8, &d},          // 0x02
		{"RLC E", rlc_r, 8, &e},          // 0x03
		{"RLC H", rlc_r, 8, &h},          // 0x04
		{"RLC L", rlc_r, 8, &l},          // 0x05
		{"RLC (HL)", rlc_hl, 16, nil},    // 0x06
		{"RLC A", rlc_r, 8, &a},          // 0x07
		{"RRC B", rrc_r, 8, &b},          // 0x08
		{"RRC C", rrc_r, 8, &c},          // 0x09
		{"RRC D", rrc_r, 8, &d},          // 0x0a
		{"RRC E", rrc_r, 8, &e},          // 0x0b
		{"RRC H", rrc_r, 8, &h},          // 0x0c
		{"RRC L", rrc_r, 8, &l},          // 0x0d
		{"RRC (HL)", rrc_hl, 16, nil},    // 0x0e
		{"RRC A", rrc_r, 8, &a},          // 0x0f
		{"RL B", rl_r, 8, &b},            // 0x10
		{"RL C", rl_r, 8, &c},            // 0x11
		{"RL D", rl_r, 8, &d},            // 0x12
		{"RL E", rl_r, 8, &e},            // 0x13
		{"RL H", rl_r, 8, &h},            // 0x14
		{"RL L", rl_r, 8, &l},            // 0x15
		{"RL (HL)", rl_hl, 16, nil},      // 0x16
		{"RL A", rl_r, 8, &a},            // 0x17
		{"RR B", rr_r, 8, &b},            // 0x18
		{"RR C", rr_r, 8, &c},            // 0x19
		{"RR D", rr_r, 8, &d},            // 0x1a
		{"RR E", rr_r, 8, &e},            // 0x1b
		{"RR H", rr_r, 8, &h},            // 0x1c
		{"RR L", rr_r, 8, &l},            // 0x1d
		{"RR (HL)", rr_hl, 16, nil},      // 0x1e
		{"RR A", rr_r, 8, &a},            // 0x1f
		{"SLA B", sla_r, 8, &b},          // 0x20
		{"SLA C", sla_r, 8, &c},          // 0x21
		{"SLA D", sla_r, 8, &d},          // 0x22
		{"SLA E", sla_r, 8, &e},          // 0x23
		{"SLA H", sla_r, 8, &h},          // 0x24
		{"SLA L", sla_r, 8, &l},          // 0x25
		{"SLA (HL)", sla_hl, 16, nil},    // 0x26
		{"SLA A", sla_r, 8, &a},          // 0x27
		{"SRA B", sra_r, 8, &b},          // 0x28
		{"SRA C", sra_r, 8, &c},          // 0x29
		{"SRA D", sra_r, 8, &d},          // 0x2a
		{"SRA E", sra_r, 8, &e},          // 0x2b
		{"SRA H", sra_r, 8, &h},          // 0x2c
		{"SRA L", sra_r, 8, &l},          // 0x2d
		{"SRA (HL)", sra_hl, 16, nil},    // 0x2e
		{"SRA A", sra_r, 8, &a},          // 0x2f
		{"SWAP B", swap_r, 8, &b},        // 0x30
		{"SWAP C", swap_r, 8, &c},        // 0x31
		{"SWAP D", swap_r, 8, &d},        // 0x32
		{"SWAP E", swap_r, 8, &e},        // 0x33
		{"SWAP H", swap_r, 8, &h},        // 0x34
		{"SWAP L", swap_r, 8, &l},        // 0x35
		{"SWAP (HL)", swap_hl, 16, nil},  // 0x36
		{"SWAP A", swap_r, 8, &a},        // 0x37
		{"SRL B", srl_r, 8, &b},          // 0x38
		{"SRL C", srl_r, 8, &c},          // 0x39
		{"SRL D", srl_r, 8, &d},          // 0x3a
		{"SRL E", srl_r, 8, &e},          // 0x3b
		{"SRL H", srl_r, 8, &h},          // 0x3c
		{"SRL L", srl_r, 8, &l},          // 0x3d
		{"SRL (HL)", srl_hl, 16, nil},    // 0x3e
		{"SRL A", srl_r, 8, &a},          // 0x3f
		{"BIT 0, B", bit, 8, &b},         // 0x40
		{"BIT 0, C", bit, 8, &c},         // 0x41
		{"BIT 0, D", bit, 8, &d},         // 0x42
		{"BIT 0, E", bit, 8, &e},         // 0x43
		{"BIT 0, H", bit, 8, &h},         // 0x44
		{"BIT 0, L", bit, 8, &l},         // 0x45
		{"BIT 0, (HL)", bit_hl, 12, nil}, // 0x46
		{"BIT 0, A", bit, 8, &a},         // 0x47
		{"BIT 1, B", bit, 8, &b},         // 0x48
		{"BIT 1, C", bit, 8, &c},         // 0x49
		{"BIT 1, D", bit, 8, &d},         // 0x4a
		{"BIT 1, E", bit, 8, &e},         // 0x4b
		{"BIT 1, H", bit, 8, &h},         // 0x4c
		{"BIT 1, L", bit, 8, &l},         // 0x4d
		{"BIT 1, (HL)", bit_hl, 12, nil}, // 0x4e
		{"BIT 1, A", bit, 8, &a},         // 0x4f
		{"BIT 2, B", bit, 8, &b},         // 0x50
		{"BIT 2, C", bit, 8, &c},         // 0x51
		{"BIT 2, D", bit, 8, &d},         // 0x52
		{"BIT 2, E", bit, 8, &e},         // 0x53
		{"BIT 2, H", bit, 8, &h},         // 0x54
		{"BIT 2, L", bit, 8, &l},         // 0x55
		{"BIT 2, (HL)", bit_hl, 12, nil}, // 0x56
		{"BIT 2, A", bit, 8, &a},         // 0x57
		{"BIT 3, B", bit, 8, &b},         // 0x58
		{"BIT 3, C", bit, 8, &c},         // 0x59
		{"BIT 3, D", bit, 8, &d},         // 0x5a
		{"BIT 3, E", bit, 8, &e},         // 0x5b
		{"BIT 3, H", bit, 8, &h},         // 0x5c
		{"BIT 3, L", bit, 8, &l},         // 0x5d
		{"BIT 3, (HL)", bit_hl, 12, nil}, // 0x5e
		{"BIT 3, A", bit, 8, &a},         // 0x5f
		{"BIT 4, B", bit, 8, &b},         // 0x60
		{"BIT 4, C", bit, 8, &c},         // 0x61
		{"BIT 4, D", bit, 8, &d},         // 0x62
		{"BIT 4, E", bit, 8, &e},         // 0x63
		{"BIT 4, H", bit, 8, &h},         // 0x64
		{"BIT 4, L", bit, 8, &l},         // 0x65
		{"BIT 4, (HL)", bit_hl, 12, nil}, // 0x66
		{"BIT 4, A", bit, 8, &a},         // 0x67
		{"BIT 5, B", bit, 8, &b},         // 0x68
		{"BIT 5, C", bit, 8, &c},         // 0x69
		{"BIT 5, D", bit, 8, &d},         // 0x6a
		{"BIT 5, E", bit, 8, &e},         // 0x6b
		{"BIT 6, H", bit, 8, &h},         // 0x6c
		{"BIT 6, L", bit, 8, &l},         // 0x6d
		{"BIT 5, (HL)", bit_hl, 12, nil}, // 0x6e
		{"BIT 5, A", bit, 8, &a},         // 0x6f
		{"BIT 6, B", bit, 8, &b},         // 0x70
		{"BIT 6, C", bit, 8, &c},         // 0x71
		{"BIT 6, D", bit, 8, &d},         // 0x72
		{"BIT 6, E", bit, 8, &e},         // 0x73
		{"BIT 6, H", bit, 8, &h},         // 0x74
		{"BIT 6, L", bit, 8, &l},         // 0x75
		{"BIT 6, (HL)", bit_hl, 12, nil}, // 0x76
		{"BIT 6, A", bit, 8, &a},         // 0x77
		{"BIT 7, B", bit, 8, &b},         // 0x78
		{"BIT 7, C", bit, 8, &c},         // 0x79
		{"BIT 7, D", bit, 8, &d},         // 0x7a
		{"BIT 7, E", bit, 8, &e},         // 0x7b
		{"BIT 7, H", bit, 8, &h},         // 0x7c
		{"BIT 7, L", bit, 8, &l},         // 0x7d
		{"BIT 7, (HL)", bit_hl, 12, nil}, // 0x7e
		{"BIT 7, A", bit, 8, &a},         // 0x7f
		{"RES 0, B", res, 8, &b},         // 0x80
		{"RES 0, C", res, 8, &c},         // 0x81
		{"RES 0, D", res, 8, &d},         // 0x82
		{"RES 0, E", res, 8, &e},         // 0x83
		{"RES 0, H", res, 8, &h},         // 0x84
		{"RES 0, L", res, 8, &l},         // 0x85
		{"RES 0, (HL)", res_hl, 16, nil}, // 0x86
		{"RES 0, A", res, 8, &a},         // 0x87
		{"RES 1, B", res, 8, &b},         // 0x88
		{"RES 1, C", res, 8, &c},         // 0x89
		{"RES 1, D", res, 8, &d},         // 0x8a
		{"RES 1, E", res, 8, &e},         // 0x8b
		{"RES 1, H", res, 8, &h},         // 0x8c
		{"RES 1, L", res, 8, &l},         // 0x8d
		{"RES 1, (HL)", res_hl, 16, nil}, // 0x8e
		{"RES 1, A", res, 8, &a},         // 0x8f
		{"RES 2, B", res, 8, &b},         // 0x90
		{"RES 2, C", res, 8, &c},         // 0x91
		{"RES 2, D", res, 8, &d},         // 0x92
		{"RES 2, E", res, 8, &e},         // 0x93
		{"RES 2, H", res, 8, &h},         // 0x94
		{"RES 2, L", res, 8, &l},         // 0x95
		{"RES 2, (HL)", res_hl, 16, nil}, // 0x96
		{"RES 2, A", res, 8, &a},         // 0x97
		{"RES 3, B", res, 8, &b},         // 0x98
		{"RES 3, C", res, 8, &c},         // 0x99
		{"RES 3, D", res, 8, &d},         // 0x9a
		{"RES 3, E", res, 8, &e},         // 0x9b
		{"RES 3, H", res, 8, &h},         // 0x9c
		{"RES 3, L", res, 8, &l},         // 0x9d
		{"RES 3, (HL)", res_hl, 16, nil}, // 0x9e
		{"RES 3, A", res, 8, &a},         // 0x9f
		{"RES 4, B", res, 8, &b},         // 0xa0
		{"RES 4, C", res, 8, &c},         // 0xa1
		{"RES 4, D", res, 8, &d},         // 0xa2
		{"RES 4, E", res, 8, &e},         // 0xa3
		{"RES 4, H", res, 8, &h},         // 0xa4
		{"RES 4, L", res, 8, &l},         // 0xa5
		{"RES 4, (HL)", res_hl, 16, nil}, // 0xa6
		{"RES 4, A", res, 8, &a},         // 0xa7
		{"RES 5, B", res, 8, &b},         // 0xa8
		{"RES 5, C", res, 8, &c},         // 0xa9
		{"RES 5, D", res, 8, &d},         // 0xaa
		{"RES 5, E", res, 8, &e},         // 0xab
		{"RES 5, H", res, 8, &h},         // 0xac
		{"RES 5, L", res, 8, &l},         // 0xad
		{"RES 5, (HL)", res_hl, 16, nil}, // 0xae
		{"RES 5, A", res, 8, &a},         // 0xaf
		{"RES 6, B", res, 8, &b},         // 0xb0
		{"RES 6, C", res, 8, &c},         // 0xb1
		{"RES 6, D", res, 8, &d},         // 0xb2
		{"RES 6, E", res, 8, &e},         // 0xb3
		{"RES 6, H", res, 8, &h},         // 0xb4
		{"RES 6, L", res, 8, &l},         // 0xb5
		{"RES 6, (HL)", res_hl, 16, nil}, // 0xb6
		{"RES 6, A", res, 8, &a},         // 0xb7
		{"RES 7, B", res, 8, &b},         // 0xb8
		{"RES 7, C", res, 8, &c},         // 0xb9
		{"RES 7, D", res, 8, &d},         // 0xba
		{"RES 7, E", res, 8, &e},         // 0xbb
		{"RES 7, H", res, 8, &h},         // 0xbc
		{"RES 7, L", res, 8, &l},         // 0xbd
		{"RES 7, (HL)", res_hl, 16, nil}, // 0xbe
		{"RES 7, A", res, 8, &a},         // 0xbf
		{"SET 0, B", set, 8, &b},         // 0xc0
		{"SET 0, C", set, 8, &c},         // 0xc1
		{"SET 0, D", set, 8, &d},         // 0xc2
		{"SET 0, E", set, 8, &e},         // 0xc3
		{"SET 0, H", set, 8, &h},         // 0xc4
		{"SET 0, L", set, 8, &l},         // 0xc5
		{"SET 0, (HL)", set_hl, 16, nil}, // 0xc6
		{"SET 0, A", set, 8, &a},         // 0xc7
		{"SET 1, B", set, 8, &b},         // 0xc8
		{"SET 1, C", set, 8, &c},         // 0xc9
		{"SET 1, D", set, 8, &d},         // 0xca
		{"SET 1, E", set, 8, &e},         // 0xcb
		{"SET 1, H", set, 8, &h},         // 0xcc
		{"SET 1, L", set, 8, &l},         // 0xcd
		{"SET 1, (HL)", set_hl, 16, nil}, // 0xce
		{"SET 1, A", set, 8, &a},         // 0xcf
		{"SET 2, B", set, 8, &b},         // 0xd0
		{"SET 2, C", set, 8, &c},         // 0xd1
		{"SET 2, D", set, 8, &d},         // 0xd2
		{"SET 2, E", set, 8, &e},         // 0xd3
		{"SET 2, H", set, 8, &h},         // 0xd4
		{"SET 2, L", set, 8, &l},         // 0xd5
		{"SET 2, (HL)", set_hl, 16, nil}, // 0xd6
		{"SET 2, A", set, 8, &a},         // 0xd7
		{"SET 3, B", set, 8, &b},         // 0xd8
		{"SET 3, C", set, 8, &c},         // 0xd9
		{"SET 3, D", set, 8, &d},         // 0xda
		{"SET 3, E", set, 8, &e},         // 0xdb
		{"SET 3, H", set, 8, &h},         // 0xdc
		{"SET 3, L", set, 8, &l},         // 0xdd
		{"SET 3, (HL)", set_hl, 16, nil}, // 0xde
		{"SET 3, A", set, 8, &a},         // 0xdf
		{"SET 4, B", set, 8, &b},         // 0xe0
		{"SET 4, C", set, 8, &c},         // 0xe1
		{"SET 4, D", set, 8, &d},         // 0xe2
		{"SET 4, E", set, 8, &e},         // 0xe3
		{"SET 4, H", set, 8, &h},         // 0xe4
		{"SET 4, L", set, 8, &l},         // 0xe5
		{"SET 4, (HL)", set_hl, 16, nil}, // 0xe6
		{"SET 4, A", set, 8, &a},         // 0xe7
		{"SET 5, B", set, 8, &b},         // 0xe8
		{"SET 5, C", set, 8, &c},         // 0xe9
		{"SET 5, D", set, 8, &d},         // 0xea
		{"SET 5, E", set, 8, &e},         // 0xeb
		{"SET 5, H", set, 8, &h},         // 0xec
		{"SET 5, L", set, 8, &l},         // 0xed
		{"SET 5, (HL)", set_hl, 16, nil}, // 0xee
		{"SET 5, A", set, 8, &a},         // 0xef
		{"SET 6, B", set, 8, &b},         // 0xf0
		{"SET 6, C", set, 8, &c},         // 0xf1
		{"SET 6, D", set, 8, &d},         // 0xf2
		{"SET 6, E", set, 8, &e},         // 0xf3
		{"SET 6, H", set, 8, &h},         // 0xf4
		{"SET 6, L", set, 8, &l},         // 0xf5
		{"SET 6, (HL)", set_hl, 16, nil}, // 0xf6
		{"SET 6, A", set, 8, &a},         // 0xf7
		{"SET 7, B", set, 8, &b},         // 0xf8
		{"SET 7, C", set, 8, &c},         // 0xf9
		{"SET 7, D", set, 8, &d},         // 0xfa
		{"SET 7, E", set, 8, &e},         // 0xfb
		{"SET 7, H", set, 8, &h},         // 0xfc
		{"SET 7, L", set, 8, &l},         // 0xfd
		{"SET 7, (HL)", set_hl, 16, nil}, // 0xfe
		{"SET 7, A", set, 8, &a},         // 0xff
	}
)

func fn_hl(fn func(*uint8, uint8), n uint8) {
	HL := uint16(h)<<8 | uint16(l)
	value := mmu.Read(HL)
	fn(&value, n)
	mmu.Write(HL, value)
}

func rlc_r(register *uint8, n uint8) {
	bit7 := *register >> 7
	*register = (*register << 1) | bit7
	if bit7 == 0 {
		f = 0
	} else {
		f = 0x10
	}
	if *register == 0 {
		f |= 0x80
	}
}

func rlc_hl(register *uint8, n uint8) {
	fn_hl(rlc_r, n)
}

func rrc_r(register *uint8, n uint8) {
	bit0 := *register & 0x1
	*register = (*register >> 1) | (bit0 << 7)
	if bit0 == 0 {
		f = 0
	} else {
		f = 0x10
	}
	if *register == 0 {
		f |= 0x80
	}
}

func rrc_hl(register *uint8, n uint8) {
	fn_hl(rrc_r, n)
}

func rl_r(register *uint8, n uint8) {
	bit7 := *register >> 7
	*register <<= 1
	if (f & 0x10) == 0x10 {
		*register |= 0x1
	}
	if bit7 == 0 {
		f = 0
	} else {
		f = 0x10
	}
	if *register == 0 {
		f |= 0x80
	}
}

func rl_hl(register *uint8, n uint8) {
	fn_hl(rl_r, n)
}

func rr_r(register *uint8, n uint8) {
	oldCarry := f & 0x10
	if (*register & 0x1) == 0x1 {
		f = 0x10
	} else {
		f = 0
	}
	*register >>= 1
	if oldCarry == 0x10 {
		*register |= 0x80
	}
	if *register == 0 {
		f |= 0x80
	}
}

func rr_hl(register *uint8, n uint8) {
	fn_hl(rr_r, n)
}

func sla_r(register *uint8, n uint8) {
	bit7 := *register >> 7
	*register <<= 1
	if bit7 == 0 {
		f = 0
	} else {
		f = 0x10
	}
	if *register == 0 {
		f |= 0x80
	}
}

func sla_hl(register *uint8, n uint8) {
	fn_hl(sla_r, n)
}

func sra_r(register *uint8, n uint8) {
	bit0 := *register & 0x1
	bit7 := *register & 0x80
	*register = (*register >> 1) | bit7
	if bit0 == 0 {
		f = 0
	} else {
		f = 0x10
	}
	if *register == 0 {
		f |= 0x80
	}
}

func sra_hl(register *uint8, n uint8) {
	fn_hl(sra_r, n)
}

func swap_r(register *uint8, n uint8) {
	low := *register & 0xf
	*register = (*register >> 4) | (low << 4)
	if *register == 0 {
		f = 0x80
	} else {
		f = 0
	}
}

func swap_hl(register *uint8, n uint8) {
	fn_hl(swap_r, n)
}

func srl_r(register *uint8, n uint8) {
	if *register&0x1 == 0x1 {
		f = 0x10
	} else {
		f = 0
	}
	*register >>= 1
	if *register == 0 {
		f |= 0x80
	}
}

func srl_hl(register *uint8, n uint8) {
	fn_hl(srl_r, n)
}

func bit(register *uint8, n uint8) {
	if *register>>n&0x1 == 0 {
		f |= 0x80
	} else {
		f &= 0x7f
	}
	f |= 0x20
	f &= 0xbf
}

func bit_hl(register *uint8, n uint8) {
	fn_hl(bit, n)
}

func res(register *uint8, n uint8) {
	*register &= ^(0x1 << n)
}

func res_hl(register *uint8, n uint8) {
	fn_hl(res, n)
}

func set(register *uint8, n uint8) {
	*register |= (0x1 << n)
}

func set_hl(register *uint8, n uint8) {
	fn_hl(set, n)
}
