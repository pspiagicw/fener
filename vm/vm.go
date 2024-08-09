package vm

import (
	"encoding/binary"
	"fmt"

	"github.com/pspiagicw/fener/code"
	"github.com/pspiagicw/fener/object"
)

const StackSize = 2048

type VM struct {
	stack        []object.Object
	stackPointer int
	frames       []*Frame
	framePointer int

	constants []object.Object
}

func New(bytes []byte, constants []object.Object) *VM {
	mainFrame := NewFrame(bytes)
	return &VM{
		stack:        make([]object.Object, StackSize),
		stackPointer: 0,

		constants:    constants,
		frames:       []*Frame{mainFrame},
		framePointer: 1,
	}
}
func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.framePointer-1]
}

func (vm *VM) push(obj object.Object) error {
	if vm.stackPointer >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.stackPointer] = obj
	vm.stackPointer++
	return nil
}
func (vm *VM) Run() error {
	var err error
	var ip int
	var ins []byte
	for vm.currentFrame().IP < len(vm.currentFrame().Instructions) {
		ip = vm.currentFrame().IP
		ins = vm.currentFrame().Instructions

		instr := ins[ip]

		switch code.OpCode(instr) {
		case code.PUSH:
			operand := binary.BigEndian.Uint16(ins[ip+1:])
			constant := vm.constants[operand]

			vm.currentFrame().IP += 3

			err = vm.push(constant)

		default:
			err = fmt.Errorf("unknown opcode %d", instr)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (vm *VM) StackTop() object.Object {
	return vm.stack[vm.stackPointer-1]
}
