// Written by pricees, inspired by Great Code Club
package vm

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"
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

	exp := 0
	res := len(vm.stack)

	if res != exp {
		t.Error("vm.stack size is ", res, " expected ", exp)
	}
}

func TestStackPush(t *testing.T) {
	vm := NewVM()

	exp := 3
	vm.StackPush(1)
	vm.StackPush(100)
	res := vm.StackPush(1000)

	if res != exp {
		t.Error("vm.stack size is ", res, " expected ", exp)
	}
}

func TestStackPop(t *testing.T) {
	vm := NewVM()

	vm.stack = []int{0, 1, 2, 3, 4000}
	exp := 4000
	res, err := vm.StackPop()

	if err || res != exp {
		t.Error("vm.StackPop() size is ", res, " expected ", exp)
	}
}

func TestStackError(t *testing.T) {
	vm := NewVM()

	vm.stack = []int{}
	_, err := vm.StackPop()

	if !err {
		t.Error("vm.StackPop() should have return err")
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

func TestProgramCounter(t *testing.T) {
	vm := NewVM()

	if vm.pc != 0x200 {
		t.Error("vm not defaulted with pc 0x200!")
	}
}

func TestLoadProgram(t *testing.T) {
	vm := NewVM()

	test_data := []uint8{1, 2, 3, 4, 5}

	vm.LoadProgram(test_data)

	res := vm.memory[512:517] // 0x200 + i

	if !bytes.Equal(res, test_data) {
		t.Error("vm.memory should have a length of ", test_data,
			" but was ", res)
	}
}

func TestDisplayClear(t *testing.T) {
	vm := NewVM()

	if !vm.DisplayClear() {
		t.Error("Display should have cleared")
	}
}

func TestDrawSprite(t *testing.T) {
	vm := NewVM()

	vm.memory[0x21A] = 0xF0 //  11110000  ****
	//	vm.memory[0x21B] = 0x90 //  11110000  *  *
	//	vm.memory[0x21C] = 0x90 //  11110000  *  *
	//	vm.memory[0x21D] = 0x90 //  11110000  *  *
	//	vm.memory[0x21E] = 0xF0 //  11110000  ****

	res := vm.DrawSprite(2, 3, 0x21B, 5)
	exp := true

	if res != exp {
		t.Error("vm.DrawSprite should have output", exp, " but was ", res)
	}
}

func TestStep(t *testing.T) {
	vm := NewVM()

	vm.Step()
	res := vm.Step()
	exp := 516

	if res != exp {
		t.Error("vm.Step should have returned a pc of ", exp,
			" but was ", res)
	}
}

func TestAll(t *testing.T) {
	vm := NewVM()

	data := make([]uint8, 10000)
	file, err := os.Open("roms/LOGO")

	count, err := file.Read(data)

	if err != nil {
		t.Fatal(err)
	}

	vm.LoadProgram(data[:count])
	fmt.Println("Read ", count, "bytes:\t", vm.memory[0x200:(0x200+count)])
	vm.Run()
	time.Sleep(1 * time.Second)
	vm.display.Draw()

}
