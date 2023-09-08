package vm

import (
	"fmt"

	"github.com/VictorMilhomem/golox/cmd/chunk"
	"github.com/VictorMilhomem/golox/cmd/values"
)

type InterpretResult int

const (
	IntepretOk InterpretResult = iota
	InterpretCompileError
	IntepretRuntimeError
)

type VM struct {
	chunk *chunk.Chunk
	ip    []byte
}

func NewVm() *VM {
	return &VM{
		chunk: chunk.NewChunk(),
	}
}

func (vm *VM) InitVM() error {
	return nil
}

func (vm *VM) FreeVM() error {
	return nil
}

func (vm *VM) Interpret(chunk *chunk.Chunk) (InterpretResult, error) {
	vm.chunk = chunk
	vm.ip = vm.chunk.GetCode()
	return vm.run(), nil
}

func (vm *VM) run() InterpretResult {
	for {
		instruction := vm.readByte()

		switch instruction {
		case byte(chunk.OpConstant):
			{
				constant := vm.readConstant()
				values.PrintValue(constant)
				fmt.Println()
			}
		case byte(chunk.OpReturn):
			return IntepretOk
		}
	}
}

func (vm *VM) readByte() byte {
	if len(vm.ip) == 0 {
		return 0
	}

	byteToRead := vm.ip[0]
	vm.ip = vm.ip[1:] // Avançar para o próximo byte
	return byteToRead
}

func (vm *VM) readConstant() values.Value {
	return vm.chunk.GetConstants().GetValues()[vm.readByte()]
}
