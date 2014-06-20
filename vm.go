// Written by pricees, inspired by Great Code Club

package vm

type Op struct {
	opcode  int
	operand int
}

type VM struct {
	display  Display
	cpuSpeed int     // hertz
	stack    [40]int // TODO: Find size of Chip-8 stack

	// 4KB of memory
	memory [4096]uint8

	// Registers
	V [16]uint8 // Vx: 16 generate purpose 8-bit registers
	I int       // 16-bit register, for storing addresses

	// Program counter, start at 0x200
	pc int
}

func NewVM() *VM {
	return &VM{pc: 0x200}
}
