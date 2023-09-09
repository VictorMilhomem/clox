package vm

import (
	"fmt"

	"github.com/VictorMilhomem/golox/cmd/chunk"
	"github.com/VictorMilhomem/golox/cmd/values"
	"github.com/golang-collections/collections/stack"
)

type InterpretResult int

const STACKMAX int = 256

const (
	IntepretOk InterpretResult = iota
	InterpretCompileError
	IntepretRuntimeError
)

type VM struct {
	chunk *chunk.Chunk
	ip    byte
	stack *stack.Stack
}

func NewVm() *VM {
	return &VM{
		chunk: chunk.NewChunk(),
		ip:    0,
		stack: stack.New(),
	}
}

func (vm *VM) InitVM() error {
	vm.resetStack()
	return nil
}

func (vm *VM) FreeVM() error {
	return nil
}

func (vm *VM) resetStack() {
	vm.stack = stack.New()
}

func (vm *VM) Interpret(chunk *chunk.Chunk) (InterpretResult, error) {
	vm.chunk = chunk
	vm.ip = vm.chunk.GetCode()[0]
	return vm.run(), nil
}

func (vm *VM) run() InterpretResult {
	const debug bool = true
	for {
		if debug {
			vm.chunk.DisassembleInstruction(int(vm.ip))
		}
		instruction := vm.readByte()

		switch instruction {
		case byte(chunk.OpConstant):
			{
				constant := vm.readConstant()
				vm.stack.Push(constant)
			}
		case byte(chunk.OpNegate):
			{
				value, _ := vm.stack.Pop().(values.Value)
				vm.stack.Push(-value)
			}
		case byte(chunk.OpReturn):
			value, _ := vm.stack.Pop().(values.Value)
			values.PrintValue(value)
			fmt.Println()
			return IntepretOk
		}
	}
}

func (vm *VM) readByte() byte {
	if vm.ip >= byte(len(vm.chunk.GetCode())) {
		return 0
	}

	byteToRead := vm.chunk.GetCode()[vm.ip]
	vm.ip++
	return byteToRead
}

func (vm *VM) readConstant() values.Value {
	return vm.chunk.GetConstants().GetValues()[vm.readByte()]
}
