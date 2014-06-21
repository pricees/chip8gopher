// Written by pricees, inspired by Great Code Club

package vm

import (
	"math/rand"
)

type Op struct {
	opcode   int
	operands int
}

type VM struct {
	display  Display
	cpuSpeed int   // hertz
	stack    []int // TODO: Find size of Chip-8 stack

	// 4KB of memory
	memory [4096]uint8

	// Registers
	V [16]uint8 // Vx: 16 generate purpose 8-bit registers
	I int       // 16-bit register, for storing addresses

	// Program counter, start at 0x200
	pc int
}

func (vm *VM) LoadProgram(data []uint8) {
	for i, val := range data {
		// Program starts are 0x200 offset
		vm.memory[0x200+i] = val
	}
}

func (vm *VM) Step() int {

	// Read the instructions at `pc`. Instructions are 2 bytes long.
	instruction := (vm.memory[vm.pc] << 8) + vm.memory[vm.pc+1]

	// Move to next instruction
	vm.pc += 2

	// Extract common operands.

	tmp_inst := int(instruction)

	x := (tmp_inst & 0x0F00) >> 8
	y := (tmp_inst & 0x00F0) >> 4
	kk := uint8(tmp_inst & 0x00FF)
	nnn := tmp_inst & 0x0FFF
	n := tmp_inst & 0x000F

	// Inspect the first nibble, the opcode.
	// A case for each type of instruction
	switch tmp_inst & 0xF000 {
	case 0x000:
		{
			switch tmp_inst {
			case 0x00EE: // `00EE` - Return from a subroutine
				// pop stack
				new_pc, err := vm.StackPop()

				if !err {
					vm.pc = new_pc
				}
				break
			case 0x00E0:
				//vm.display //.clear()
				break
			}
			break
		}
	case 0x1000:
		// 1nnn - Jump to location nnn.
		vm.pc = nnn
	case 0x2000:
		vm.StackPush(vm.pc)
		vm.pc = nnn
		break
	case 0x3000:
		// 3xkk - Skip next instruction if Vx = kk.
		if vm.V[x] == kk {
			vm.pc += 2 // an instruction is 2 bytes
		}
		break
	case 0x4000:
		// 4xkk - Skip next instruction if Vx != kk
		if vm.V[x] != kk {
			vm.pc += 2
		}
		break
	case 0x6000:
		// 6xkk - Set Vx = kk.
		vm.V[x] = kk
		break
	case 0x7000:
		// 7xkk - Set Vx = Vx + kk.
		vm.V[x] = vm.V[x] + kk
		break
	case 0xA000:
		// Annn - Set I = nnn.
		vm.I = nnn
		break
	case 0xC000:
		// Cxkk - Set Vx = random byte AND kk.
		vm.V[x] = uint8(rand.Intn(0xFF+1)) & kk
		break
	case 0xD000:
		// Dxyn - Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.
		collision := vm.DrawSprite(vm.V[x], vm.V[y], vm.I, n)
		// VF is set to 1 if there is a collision.
		if collision {
			vm.V[0xF] = 1
		} else {
			vm.V[0xF] = 0
		}
		break
	case 0xF000:
		// `Fx1E` - Set I = I + Vx
		vm.I += int(vm.V[x])
		break
	default:
		//FIXME: put a panic here, or something
		break
	}

	return vm.pc
}

func (vm *VM) StackPush(val int) int {
	vm.stack = append(vm.stack, val)
	return len(vm.stack)
}

func (vm *VM) StackPop() (int, bool) {
	size := len(vm.stack)

	ret_val := 0
	err := false

	if size > 0 {
		ret_val = vm.stack[size-1]
		vm.stack = vm.stack[:size-1]
	} else {
		err = true
	}
	return ret_val, err
}

func (vm *VM) DrawSprite(x uint8, y uint8, i int, n int) bool {
	return false
}

func NewVM() *VM {
	return &VM{pc: 0x200, stack: make([]int, 0, 80)}
}
