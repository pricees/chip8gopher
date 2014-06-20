// Written by pricees, inspired by Great Code Club
package vm

import (
	"testing"
)

func TestMemorySize(t *testing.T) {
	vm := NewVM()

	exp := 4096
	res := len(vm.memory)

	if res != exp {
		t.Error("vm.memory size is ", res, " expected ", exp)
	}
}

func TestMaxStackFrames(t *testing.T) {
	vm := NewVM()

	exp := 40
	res := len(vm.stack)

	if res != exp {
		t.Error("vm.stack size is ", res, " expected ", exp)
	}
}

func TestRegisters(t *testing.T) {
	vm := NewVM()

	exp := 16
	res := len(vm.V)

	if res != exp {
		t.Error("vm.V is ", res, " expected ", exp)
	}
}

func TestI(t *testing.T) {
	vm := NewVM()

	exp := 0
	res := vm.I

	if res != exp {
		t.Error("vm.I is ", res, " expected ", exp)
	}
}

func TestNewVM(t *testing.T) {
	vm := NewVM()

	if vm.pc != 0x200 {
		t.Error("vm not defaulted with pc 0x200!")
	}
}
